package controller

import (
	"encoding/json"
	"net/http"

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
		http.Error(writer, "Invalid request body", http.StatusBadRequest)
		return
	}

	customerResponse, err := controller.Usecase.RegisterCustomer(req.Context(), customer)
	if err != nil {
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
		return
	}

	webResponse := response.WebResponse{
		Code:    http.StatusCreated,
		Message: "Customer successfully registered",
		Data:    customerResponse,
	}

	err = json.NewEncoder(writer).Encode(webResponse)
	if err != nil {
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
		return
	}

}

func (controller *CustomerControllerImpl) LoginCustomer(writer http.ResponseWriter, req *http.Request, param httprouter.Params) {
	customer := request.CustomerLogin{}
	if err := json.NewDecoder(req.Body).Decode(&customer); err != nil {
		http.Error(writer, "Invalid request body", http.StatusBadRequest)
		return
	}

	email := customer.Email
	password := customer.Password

	id, ok, err := controller.Usecase.LoginCustomer(req.Context(), email, password)
	if !ok {
		http.Error(writer, "Unathorized", http.StatusUnauthorized)
		return
	}
	if err != nil {
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
		return
	}

	tokenString, err := helper.GenerateCustomerJWTToken(id)
	if err != nil {
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
		return
	}

	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Message: "Login Success",
		Data:    tokenString,
	}

	if err := json.NewEncoder(writer).Encode(webResponse); err != nil {
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
		return
	}

}

func (controller *CustomerControllerImpl) UpdateCustomer(writer http.ResponseWriter, req *http.Request, param httprouter.Params) {
	customer := request.CustomerUpdate{}
	err := json.NewDecoder(req.Body).Decode(&customer)
	if err != nil {
		http.Error(writer, "Invalid request body", http.StatusBadRequest)
		return
	}

	customerResponse, err := controller.Usecase.UpdateCustomer(req.Context(), customer)
	if err != nil {
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
		return
	}

	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Message: "Customer successfully updated",
		Data:    customerResponse,
	}
	err = json.NewEncoder(writer).Encode(webResponse)
	if err != nil {
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
		return
	}

}

func (controller *CustomerControllerImpl) DeleteCustomer(writer http.ResponseWriter, req *http.Request, param httprouter.Params) {
	customer := request.CustomerDelete{}
	err := json.NewDecoder(req.Body).Decode(&customer)
	if err != nil {
		http.Error(writer, "Invalid request body", http.StatusBadRequest)
	}

	err = controller.Usecase.DeleteCustomer(req.Context(), customer.Email)
	if err != nil {
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
	}
	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Message: "Customer successfully deleted",
	}

	writer.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(webResponse)
	if err != nil {
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
		return
	}
}
