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
	route.POST("/customer/register", controller.RegisterCustomer)
	route.POST("/customer/login", controller.LoginCustomer)
	route.PUT("/customer", middleware.MindMiddleware(controller.UpdateCustomer))
	route.DELETE("/customer", middleware.MindMiddleware(controller.DeleteCustomer))
}

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
		Data:    tokenString,
	}

	writer.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(writer).Encode(webResponse); err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}
}

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
