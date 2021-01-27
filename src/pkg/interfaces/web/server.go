package web

import (
	"net/http"
	"reflect"
	"runtime"
	"sync"
	"time"

	"github.com/cry999/pm-projects/pkg/interfaces/logger"
	"github.com/gorilla/mux"
)

type (
	// HandlerFunc ...
	HandlerFunc func(*RequestContext) error

	// Middleware ...
	Middleware func(HandlerFunc) HandlerFunc

	// Server ...
	Server interface {
		GlobalUse(...Middleware)
		Route(method, path string, h HandlerFunc, middlewares ...Middleware)
		Run(add string) error
		Close() error
	}

	server struct {
		mux         *mux.Router
		middlewares []Middleware
		logger      logger.Logger
	}

	// ShutdownFunc は server が Close される時に実行されるべき関数
	ShutdownFunc func(logger.Logger)
)

var (
	// shutdowns は server が close される時に実行されるべき関数
	shutdowns []ShutdownFunc
)

// RegistShutdown ...
func RegistShutdown(f ShutdownFunc) {
	shutdowns = append(shutdowns, f)
}

// NewServer ...
func NewServer() Server {
	return &server{
		mux:         mux.NewRouter(),
		middlewares: []Middleware{},
		logger:      NewDefaultLogger(logger.LoggerLevelDebug),
	}
}

// GlobalUse ...
func (s *server) GlobalUse(middlewares ...Middleware) {
	s.middlewares = append(s.middlewares, middlewares...)
}

// Route ...
// TODO: CORS 対応を強制で実行している。うまいこと middleware に落とし込めるようにしたい。
func (s *server) Route(method, path string, handler HandlerFunc, middlewares ...Middleware) {
	middlewares = append(s.middlewares, middlewares...)

	handlerName := runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()

	s.mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		rc := NewRequestContext(w, r, NewRequestLogger(r, logger.LoggerLevelDebug))
		rc.Logger().Info("access from %s", r.UserAgent())
		rc.Logger().Debug("handle by '%s'", handlerName)

		start := time.Now()
		defer rc.Logger().Debug("elapsed time: %v", time.Now().Sub(start))

		for key, val := range mux.Vars(r) {
			rc.SetParam("path."+key, val)
		}

		for i := len(middlewares) - 1; i >= 0; i-- {
			middleware := middlewares[i]
			handler = middleware(handler)
		}
		if r.Method == http.MethodOptions {
			// CORS preflight
			if origin := rc.RequestHeader().Get("Origin"); origin != "" && origin != "*" {
				rc.ResponseHeader().Set("Access-Control-Allow-Origin", rc.RequestHeader().Get("Origin"))
				rc.ResponseHeader().Set("Access-Control-Allow-Methods", "*")
				rc.ResponseHeader().Set("Access-Control-Allow-Headers", "*")
				rc.ResponseHeader().Set("Access-Control-Max-Age", "86400")
			}
			rc.NoContent()
			return
		}
		cors(handler)(rc)
	}).Methods(method, http.MethodOptions)
}

// TODO: CORS を適切なファイルに移動する
func cors(handler HandlerFunc) HandlerFunc {
	return func(rc *RequestContext) error {
		rc.Logger().Debug("CORS called")
		if origin := rc.RequestHeader().Get("Origin"); origin != "" && origin != "*" {
			rc.ResponseHeader().Set("Access-Control-Allow-Origin", origin)
		}
		err := handler(rc)
		return err
	}
}

// Run ...
func (s *server) Run(addr string) error {
	return http.ListenAndServe(addr, s.mux)
}

func (s *server) Close() error {
	wg := sync.WaitGroup{}

	for _, shutdown := range shutdowns {
		wg.Add(1)
		go func(shutdown ShutdownFunc) {
			defer wg.Done()
			defer func() {
				if err := recover(); err != nil {
					s.logger.Error("shutdown failed: %v", err)
				}
			}()
			shutdown(s.logger)
		}(shutdown)
	}

	wg.Wait()
	return nil
}
