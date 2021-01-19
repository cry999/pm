package web

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// RequestContext ...
type RequestContext struct {
	w http.ResponseWriter
	r *http.Request
	l Logger
	p map[string]interface{}
}

// NewRequestContext creates a new RequestContext instance
func NewRequestContext(w http.ResponseWriter, r *http.Request, l Logger) *RequestContext {
	return &RequestContext{
		w: w,
		r: r,
		l: l,
		p: make(map[string]interface{}),
	}
}

// Logger ...
func (rc *RequestContext) Logger() Logger { return rc.l }

// Context ...
func (rc *RequestContext) Context() context.Context { return rc.r.Context() }

// RequestHeader ...
// ? Request.Header の変更を許可するか否か。Clone() を渡すべき?
func (rc *RequestContext) RequestHeader() http.Header {
	return rc.r.Header
}

// ResponseHeader ...
// ? http.Header を見せるべきか否か。メソッドを仲介するべき？
func (rc *RequestContext) ResponseHeader() http.Header {
	return rc.w.Header()
}

// JSONResponse ...
func (rc *RequestContext) JSONResponse(code int, v interface{}) {
	rc.w.Header().Set("Content-Type", "application/json")
	rc.w.WriteHeader(code)
	if err := json.NewEncoder(rc.w).Encode(v); err != nil {
		rc.Logger().Error("failed to write response: %v", err)
	}
}

type errorResponse struct {
	Message string `json:"message"`
}

// JSONErrorResponse ...
func (rc *RequestContext) JSONErrorResponse(err error) {
	code := ErrorCode(err)
	rc.Logger().Error("error occurred: %d: %v", code, err)

	rc.RawJSONErrorResponse(code, err.Error())
}

// RawJSONErrorResponse ...
func (rc *RequestContext) RawJSONErrorResponse(code int, f string, a ...interface{}) {
	rc.JSONResponse(code, errorResponse{Message: fmt.Sprintf(f, a...)})
}

// JSONRequest ...
func (rc *RequestContext) JSONRequest(v interface{}) error {
	return json.NewDecoder(rc.r.Body).Decode(v)
}

// SetParam ...
func (rc *RequestContext) SetParam(key string, val interface{}) {
	rc.p[key] = val
}

// GetParam ...
func (rc *RequestContext) GetParam(key string) (interface{}, bool) {
	val, ok := rc.p[key]
	return val, ok
}

// GetParamString ...
func (rc *RequestContext) GetParamString(key string) (string, error) {
	val, ok := rc.p[key]
	if !ok {
		return "", fmt.Errorf("'%s' is not set", key)
	}
	str, ok := val.(string)
	if !ok {
		return "", fmt.Errorf("'%s' is not string value: %v", key, val)
	}
	return str, nil
}
