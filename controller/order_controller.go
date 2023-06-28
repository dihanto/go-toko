package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type OrderController interface {
	AddOrder(writer http.ResponseWriter, req *http.Request, param httprouter.Params)
	FindOrder(writer http.ResponseWriter, req *http.Request, param httprouter.Params)
}
