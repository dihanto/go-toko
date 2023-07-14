package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ProductController interface {
	AddProduct(writer http.ResponseWriter, req *http.Request, param httprouter.Params)
	GetProduct(writer http.ResponseWriter, req *http.Request, param httprouter.Params)
	FindById(writer http.ResponseWriter, req *http.Request, param httprouter.Params)
	UpdateProduct(writer http.ResponseWriter, req *http.Request, param httprouter.Params)
	DeleteProduct(writer http.ResponseWriter, req *http.Request, param httprouter.Params)
	FindByName(writer http.ResponseWriter, req *http.Request, param httprouter.Params)
}
