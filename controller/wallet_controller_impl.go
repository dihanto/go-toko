package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/dihanto/go-toko/helper"
	"github.com/dihanto/go-toko/middleware"
	"github.com/dihanto/go-toko/model/web/request"
	"github.com/dihanto/go-toko/model/web/response"
	"github.com/dihanto/go-toko/usecase"
	"github.com/julienschmidt/httprouter"
)

type WalletControllerImpl struct {
	Usecase usecase.WalletUsecase
	Route   *httprouter.Router
}

func NewWalletController(usecase usecase.WalletUsecase, route *httprouter.Router) WalletController {
	controller := &WalletControllerImpl{
		Usecase: usecase,
		Route:   route,
	}
	controller.router(route)
	return controller
}

func (controller *WalletControllerImpl) router(route *httprouter.Router) {
	route.POST("/wallet", middleware.MindMiddleware(controller.AddWallet))
	route.GET("/wallet/:id", middleware.MindMiddleware(controller.GetWallet))
	route.PUT("/wallet/:id", middleware.MindMiddleware(controller.UpdateWallet))
}

func (controller *WalletControllerImpl) AddWallet(writer http.ResponseWriter, req *http.Request, param httprouter.Params) {
	request := request.AddWallet{}
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		log.Println(err)
		http.Error(writer, "invalid request body", http.StatusBadRequest)
		return
	}
	authHeader := req.Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	request.IdCustomer, err = helper.GenerateIdFromToken(tokenString)
	if err != nil {
		log.Println(err)
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
		return
	}

	walletResponse, err := controller.Usecase.AddWallet(req.Context(), request)
	if err != nil {
		log.Println(err)
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
		return
	}

	webResponse := response.WebResponse{
		Code:    http.StatusCreated,
		Message: "Wallet successfully created",
		Data:    walletResponse,
	}

	err = json.NewEncoder(writer).Encode(webResponse)
	if err != nil {
		log.Println(err)
		http.Error(writer, "internal server error", http.StatusInternalServerError)
		return
	}
}

func (controller *WalletControllerImpl) GetWallet(writer http.ResponseWriter, req *http.Request, param httprouter.Params) {
	idString := param.ByName("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Println(err)
		http.Error(writer, "internal server error", http.StatusInternalServerError)
		return
	}

	walletResponse, err := controller.Usecase.GetWallet(req.Context(), id)
	if err != nil {
		log.Println(err)
		http.Error(writer, "internal server error", http.StatusInternalServerError)
		return
	}

	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Message: "Success get wallet",
		Data:    walletResponse,
	}

	err = json.NewEncoder(writer).Encode(webResponse)
	if err != nil {
		log.Println(err)
		http.Error(writer, "internal server error", http.StatusInternalServerError)
		return
	}
}

func (controller *WalletControllerImpl) UpdateWallet(writer http.ResponseWriter, req *http.Request, param httprouter.Params) {
	request := request.UpdateWallet{}
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		log.Println(err)
		http.Error(writer, "invalid response body", http.StatusInternalServerError)
		return
	}

	idString := param.ByName("id")
	request.Id, err = strconv.Atoi(idString)
	if err != nil {
		log.Println(err)
		http.Error(writer, "internal server error", http.StatusInternalServerError)
		return
	}

	walletResponse, err := controller.Usecase.UpdateWallet(req.Context(), request)
	if err != nil {
		log.Println(err)
		http.Error(writer, "internal server error", http.StatusInternalServerError)
		return
	}

	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Message: "Success update wallet",
		Data:    walletResponse,
	}

	err = json.NewEncoder(writer).Encode(webResponse)
	if err != nil {
		log.Println(err)
		http.Error(writer, "internal server error", http.StatusInternalServerError)
		return
	}

}
