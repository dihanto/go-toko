package controller

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/dihanto/go-toko/exception"
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
	route.GET("/wallet", middleware.MindMiddleware(controller.GetWallet))
	route.PUT("/wallet", middleware.MindMiddleware(controller.UpdateWallet))
}

// addWallet
// @Summary Add wallet
// @Description Add wallet
// @Tags Wallet
// @Param request_body body request.AddWallet true "Add Wallet"
// @Success 201 {object} response.WebResponse
// @Security JWTAuth
// @Router /wallet [post]
func (controller *WalletControllerImpl) AddWallet(writer http.ResponseWriter, req *http.Request, param httprouter.Params) {
	request := request.AddWallet{}
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}
	authHeader := req.Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	request.IdCustomer, err = helper.GenerateIdFromToken(tokenString)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}

	walletResponse, err := controller.Usecase.AddWallet(req.Context(), request)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}

	webResponse := response.WebResponse{
		Code:    http.StatusCreated,
		Message: "Wallet successfully created",
		Data:    walletResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(webResponse)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}
}

// getWallet
// @Summary Get wallet
// @Description Get wallet
// @Tags Wallet
// @Success 200 {object} response.WebResponse
// @Security JWTAuth
// @Router /wallet [get]
func (controller *WalletControllerImpl) GetWallet(writer http.ResponseWriter, req *http.Request, param httprouter.Params) {
	authHeader := req.Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	idCustomer, err := helper.GenerateIdFromToken(tokenString)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}

	walletResponse, err := controller.Usecase.GetWallet(req.Context(), idCustomer)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}

	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Message: "Success get wallet",
		Data:    walletResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(webResponse)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}
}

// updateWallet
// @Summary Update wallet
// @Description Update wallet
// @Tags Wallet
// @Param request_body body request.UpdateWallet true "Update Wallet"
// @Success 200 {object} response.WebResponse
// @Security JWTAuth
// @Router /wallet [put]
func (controller *WalletControllerImpl) UpdateWallet(writer http.ResponseWriter, req *http.Request, param httprouter.Params) {
	request := request.UpdateWallet{}
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}
	authHeader := req.Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	request.IdCustomer, err = helper.GenerateIdFromToken(tokenString)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}

	walletResponse, err := controller.Usecase.UpdateWallet(req.Context(), request)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}

	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Message: "Success update wallet",
		Data:    walletResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(webResponse)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}

}
