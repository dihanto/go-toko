package controller

import (
	"net/http"

	"github.com/dihanto/go-toko/usecase"
	"github.com/julienschmidt/httprouter"
)

type OrderControllerImpl struct {
	Usecase usecase.OrderUsecase
	Route   httprouter.Router
}

func (controller *OrderControllerImpl) AddOrder(writer http.ResponseWriter, req *http.Request, param httprouter.Params) {
	panic("not implemented") // TODO: Implement
}

func (controller *OrderControllerImpl) FindOrder(writer http.ResponseWriter, req *http.Request, param httprouter.Params) {
	panic("not implemented") // TODO: Implement
}
