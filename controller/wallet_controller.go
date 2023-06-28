package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type WalletController interface {
	AddWallet(writer http.ResponseWriter, req *http.Request, param httprouter.Params)
	GetWallet(writer http.ResponseWriter, req *http.Request, param httprouter.Params)
	UpdateWallet(writer http.ResponseWriter, req *http.Request, param httprouter.Params)
}
