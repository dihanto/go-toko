package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CustomerController interface {
	RegisterCustomer(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	LoginCustomer(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	UpdateCustomer(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	DeleteCustomer(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
}
