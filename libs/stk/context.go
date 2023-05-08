package stk

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Context struct {
	Request *http.Request
	Writer  http.ResponseWriter

	Params httprouter.Params
}
