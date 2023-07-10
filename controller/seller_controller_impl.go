package controller

import (
	"encoding/json"
	"net/http"

	"github.com/dihanto/go-toko/exception"
	"github.com/dihanto/go-toko/helper"
	"github.com/dihanto/go-toko/middleware"
	"github.com/dihanto/go-toko/model/web/request"
	"github.com/dihanto/go-toko/model/web/response"
	"github.com/dihanto/go-toko/usecase"
	"github.com/julienschmidt/httprouter"
)

type SellerControllerImpl struct {
	Usecase usecase.SellerUsecase
	Route   *httprouter.Router
}

func NewSellerControllerImpl(usecase usecase.SellerUsecase, route *httprouter.Router) SellerController {
	controller := &SellerControllerImpl{
		Usecase: usecase,
		Route:   route,
	}

	controller.router(route)
	return controller
}

func (controller *SellerControllerImpl) router(route *httprouter.Router) {
	route.POST("/seller/register", controller.RegisterSeller)
	route.POST("/seller/login", controller.LoginSeller)
	route.PUT("/seller", middleware.MindMiddleware(controller.UpdateSeller))
	route.DELETE("/seller", middleware.MindMiddleware(controller.DeleteSeller))
}

// registerSeller
// @Summary Register seller
// @Description Register Seller
// @Tags Seller
// @Param request_body body request.SellerRegister true "Register Seller"
// @Success      201  {object}   response.WebResponse
// @Router /seller/register [post]
func (controller *SellerControllerImpl) RegisterSeller(writer http.ResponseWriter, req *http.Request, param httprouter.Params) {
	seller := request.SellerRegister{}
	err := json.NewDecoder(req.Body).Decode(&seller)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}

	sellerResponse, err := controller.Usecase.RegisterSeller(req.Context(), seller)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}

	webResponse := response.WebResponse{
		Code:    http.StatusCreated,
		Message: "Seller successfully registered",
		Data:    sellerResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(webResponse)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}
}

// loginSeller
// @Summary Register seller
// @Description Register seller
// @Tags Seller
// @Param request_body body request.SellerLogin true "Login Seller"
// @Success      200  {object}   response.WebResponse
// @Router /seller/login [post]
func (controller *SellerControllerImpl) LoginSeller(writer http.ResponseWriter, req *http.Request, param httprouter.Params) {
	seller := request.SellerLogin{}
	err := json.NewDecoder(req.Body).Decode(&seller)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}

	id, result, err := controller.Usecase.LoginSeller(req.Context(), seller)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}
	if !result {
		exception.ErrorHandler(writer, req, result)
		return
	}

	tokenString, err := helper.GenerateSellerJWTToken(id)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}

	webRespone := response.WebResponse{
		Code:    http.StatusOK,
		Message: "Login success",
		Data:    tokenString,
	}

	writer.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(webRespone)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}

}

// updateSeller
// @Summary Update seller
// @Description Update seller
// @Tags Seller
// @Param request_body body request.SellerUpdate true "Update Seller"
// @Success      200  {object}   response.WebResponse
// @Security JWTAuth
// @Router /seller [put]
func (controller *SellerControllerImpl) UpdateSeller(writer http.ResponseWriter, req *http.Request, param httprouter.Params) {
	seller := request.SellerUpdate{}
	err := json.NewDecoder(req.Body).Decode(&seller)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}

	sellerResponse, err := controller.Usecase.UpdateSeller(req.Context(), seller)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}

	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Message: "Seller successfully updated",
		Data:    sellerResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(webResponse)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}
}

// deleteSeller
// @Summary Delete seller
// @Description Delete seller
// @Tags Seller
// @Param request_body body request.SellerDelete true "Delete Seller"
// @Success      200  {object}   response.WebResponse
// @Security JWTAuth
// @Router /seller [delete]
func (controller *SellerControllerImpl) DeleteSeller(writer http.ResponseWriter, req *http.Request, param httprouter.Params) {
	seller := request.SellerDelete{}
	err := json.NewDecoder(req.Body).Decode(&seller)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}

	err = controller.Usecase.DeleteSeller(req.Context(), seller)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}

	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Message: "Seller successfully deleted",
	}

	writer.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(webResponse)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}

}
