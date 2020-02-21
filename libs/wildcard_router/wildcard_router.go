package wildcard_router

import (
	"net/http"
	"strings"
)

// WildcardRouter holds registered route handlers
type WildcardRouter struct {
	middlewares     []func(writer http.ResponseWriter, request *http.Request)
	handlers        []http.Handler
	notFoundHandler http.HandlerFunc
}

// New return a new WildcardRouter
func New() *WildcardRouter {
	return &WildcardRouter{}
}

// MountTo mount the service into mux (HTTP request multiplexer) with given path
func (w *WildcardRouter) MountTo(mountTo string, mux *http.ServeMux) {
	mountTo = "/" + strings.Trim(mountTo, "/")

	mux.Handle(mountTo, w)
	mux.Handle(mountTo+"/", w)
}

// AddHandler will append new handler to Handlers
func (w *WildcardRouter) AddHandler(handler http.Handler) {
	w.handlers = append(w.handlers, handler)
}

// NoRoute will set handler to handle 404
func (w *WildcardRouter) NoRoute(handler http.HandlerFunc) {
	w.notFoundHandler = handler
}

// Use will append new middleware
func (w *WildcardRouter) Use(middleware func(writer http.ResponseWriter, request *http.Request)) {
	w.middlewares = append(w.middlewares, middleware)
}

// ServeHTTP serve http for wildcard router
func (w *WildcardRouter) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	wildcardRouterWriter := &WildcardRouterWriter{writer, 0, false}

	for _, middleware := range w.middlewares {
		middleware(writer, req)
	}

	for _, handler := range w.handlers {
		if handler.ServeHTTP(wildcardRouterWriter, req); wildcardRouterWriter.isProcessed() {
			return
		}
		wildcardRouterWriter.reset()
	}

	wildcardRouterWriter.skipNotFoundCheck = true
	if w.notFoundHandler != nil {
		w.notFoundHandler(writer, req)
	} else {
		http.NotFound(wildcardRouterWriter, req)
	}
}

// WildcardRouterWriter will used to capture status
type WildcardRouterWriter struct {
	http.ResponseWriter
	// Keep status code
	status int
	// Used to skip status check
	skipNotFoundCheck bool
}

// Status will return request's status code
func (w WildcardRouterWriter) Status() int {
	return w.status
}

// WriteHeader only set status code when not 404
func (w *WildcardRouterWriter) WriteHeader(statusCode int) {
	if w.skipNotFoundCheck || statusCode != http.StatusNotFound {
		w.ResponseWriter.WriteHeader(statusCode)
	}
	w.status = statusCode
}

// Write only set content when not 404
func (w *WildcardRouterWriter) Write(data []byte) (int, error) {
	if w.skipNotFoundCheck || w.status != http.StatusNotFound {
		w.status = http.StatusOK
		return w.ResponseWriter.Write(data)
	}
	return 0, nil
}

func (w *WildcardRouterWriter) reset() {
	w.skipNotFoundCheck = false
	w.Header().Set("Content-Type", "")
	w.status = 0
}

func (w WildcardRouterWriter) isProcessed() bool {
	return w.status != http.StatusNotFound && w.status != 0
}
