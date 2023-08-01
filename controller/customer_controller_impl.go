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

type CustomerControllerImpl struct {
	Usecase usecase.CustomerUsecase
	Route   *httprouter.Router
}

func NewCustomerControllerImpl(usecase usecase.CustomerUsecase, route *httprouter.Router) CustomerController {
	controller := &CustomerControllerImpl{
		Usecase: usecase,
		Route:   route,
	}
	controller.router(route)
	return controller
}

func (controller *CustomerControllerImpl) router(route *httprouter.Router) {
	route.POST("/customers/register", controller.RegisterCustomer)
	route.POST("/customers/login", controller.LoginCustomer)
	route.PUT("/customers", middleware.CommonMiddleware(controller.UpdateCustomer))
	route.DELETE("/customers", middleware.CommonMiddleware(controller.DeleteCustomer))
}

// registerCustomer
// @Summary Register customer
// @Description Register Customer
// @Tags Customer
// @Param request_body body request.CustomerRegister true "Register Customer"
// @Success      201  {object}   response.WebResponse
// @Router /customer/register [post]
func (controller *CustomerControllerImpl) RegisterCustomer(writer http.ResponseWriter, req *http.Request, param httprouter.Params) {

	customer := request.CustomerRegister{}
	err := json.NewDecoder(req.Body).Decode(&customer)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}

	customerResponse, err := controller.Usecase.RegisterCustomer(req.Context(), customer)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}

	webResponse := response.WebResponse{
		Code:    http.StatusCreated,
		Message: "Customer successfully registered",
		Data:    customerResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(webResponse)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}

}

// loginCustomer
// @Summary Login customer
// @Description Login customer
// @Tags Customer
// @Param request_body body request.CustomerLogin true "Login Customer"
// @Success 200 {object} response.WebResponse
// @Router /customer/login [post]
func (controller *CustomerControllerImpl) LoginCustomer(writer http.ResponseWriter, req *http.Request, param httprouter.Params) {
	customer := request.CustomerLogin{}
	err := json.NewDecoder(req.Body).Decode(&customer)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}

	id, result, err := controller.Usecase.LoginCustomer(req.Context(), customer)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}
	if !result {
		exception.ErrorHandler(writer, req, result)
		return
	}

	tokenString, err := helper.GenerateCustomerJWTToken(id)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}

	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Message: "Login Success",
		Data:    "Token : " + tokenString,
	}

	writer.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(writer).Encode(webResponse); err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}
}

// updateCustomer
// @Summary Update customer
// @Description Update customer
// @Tags Customer
// @Param request_body body request.CustomerUpdate true "Update Customer"
// @Success 200 {object} response.WebResponse
// @Security JWTAuth
// @Router /customer [put]
func (controller *CustomerControllerImpl) UpdateCustomer(writer http.ResponseWriter, req *http.Request, param httprouter.Params) {
	customer := request.CustomerUpdate{}
	err := json.NewDecoder(req.Body).Decode(&customer)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}

	customerResponse, err := controller.Usecase.UpdateCustomer(req.Context(), customer)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}

	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Message: "Customer successfully updated",
		Data:    customerResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(webResponse)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}
}

// deleteCustomer
// @Summary Delete customer
// @Description Delete customer
// @Tags Customer
// @Param request_body body request.CustomerDelete true "Delete customer"
// @Success 200 {object} response.WebResponse
// @Security JWTAuth
// @Router /customer [delete]
func (controller *CustomerControllerImpl) DeleteCustomer(writer http.ResponseWriter, req *http.Request, param httprouter.Params) {
	customer := request.CustomerDelete{}
	err := json.NewDecoder(req.Body).Decode(&customer)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}

	err = controller.Usecase.DeleteCustomer(req.Context(), customer)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}
	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Message: "Customer successfully deleted",
	}

	writer.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(webResponse)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}
}
