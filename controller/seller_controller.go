package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type SellerController interface {
	RegisterSeller(writer http.ResponseWriter, req *http.Request, param httprouter.Params)
	LoginSeller(writer http.ResponseWriter, req *http.Request, param httprouter.Params)
	UpdateSeller(writer http.ResponseWriter, req *http.Request, param httprouter.Params)
	DeleteSeller(writer http.ResponseWriter, req *http.Request, param httprouter.Params)
}
