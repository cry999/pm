package web

import (
	"net/http"
	"sync"

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
		logger      Logger
	}

	// ShutdownFunc は server が Close される時に実行されるべき関数
	ShutdownFunc func(Logger)
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
		logger:      NewDefaultLogger(LoggerLevelDebug),
	}
}

// GlobalUse ...
func (s *server) GlobalUse(middlewares ...Middleware) {
	s.middlewares = append(s.middlewares, middlewares...)
}

// Route ...
func (s *server) Route(method, path string, handler HandlerFunc, middlewares ...Middleware) {
	middlewares = append(s.middlewares, middlewares...)
	s.mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		rc := NewRequestContext(w, r, NewRequestLogger(r, LoggerLevelDebug))
		for i := len(middlewares) - 1; i >= 0; i-- {
			middleware := middlewares[i]
			handler = middleware(handler)
		}
		handler(rc)
	}).Methods(method)
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
