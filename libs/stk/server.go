package stk

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type HandlerFunc func(*Context)

type Middleware func(HandlerFunc) HandlerFunc

type Server struct {
	Router      *httprouter.Router
	Middlewares []Middleware
}

func NewServer() *Server {
	return &Server{
		Router:      httprouter.New(),
		Middlewares: []Middleware{},
	}
}

func (s *Server) Use(middleware Middleware) {
	s.Middlewares = append(s.Middlewares, middleware)
}

func (s *Server) Get(path string, handler HandlerFunc) {
	s.Router.GET(path, wrapHandlerFunc(s.applyMiddleware(handler)))
}

func (s *Server) Post(path string, handler HandlerFunc) {
	s.Router.POST(path, wrapHandlerFunc(s.applyMiddleware(handler)))
}

func (s *Server) Put(path string, handler HandlerFunc) {
	s.Router.PUT(path, wrapHandlerFunc(s.applyMiddleware(handler)))
}

func (s *Server) Delete(path string, handler HandlerFunc) {
	s.Router.DELETE(path, wrapHandlerFunc(s.applyMiddleware(handler)))
}

func (s *Server) Patch(path string, handler HandlerFunc) {
	s.Router.PATCH(path, wrapHandlerFunc(s.applyMiddleware(handler)))
}

func (s *Server) ServeHTTP(port string) error {
	return http.ListenAndServe(port, s.Router)
}

func wrapHandlerFunc(handler HandlerFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		handlerContext := &Context{
			Params:  p,
			Request: r,
			Writer:  w,
		}
		handler(handlerContext)
	}
}

func (s *Server) applyMiddleware(handler HandlerFunc) HandlerFunc {
	updatedHandler := handler
	for i := len(s.Middlewares) - 1; i >= 0; i-- {
		updatedHandler = s.Middlewares[i](updatedHandler)
	}
	return updatedHandler
}
