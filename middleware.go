package grest

import (
	"log"
	"net/http"
)

type Middleware func(next http.HandlerFunc) http.HandlerFunc

// ChainMiddleware provides syntactic sugar to create a new middleware
// which will be the result of chaining the ones received as parameters.
func ChainMiddleware(middleware ...Middleware) Middleware {
	return func(final http.HandlerFunc) http.HandlerFunc {
		return func(response http.ResponseWriter, request *http.Request) {
			last := final
			for i := len(middleware) - 1; i >= 0; i-- {
				last = middleware[i](last)
			}
			last(response, request)
		}
	}
}

// TODO: Figure out logic
func NewMiddleware(middleware http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		middleware(response, request)
		middleware.ServeHTTP(response, request)
	}
}

var RequestLogger = NewMiddleware(func(response http.ResponseWriter, request *http.Request) {
	log.Printf("%s request to %s", request.Method, request.RequestURI)
})
