package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/dihanto/go-toko/exception"
	"github.com/dihanto/go-toko/helper"
	"github.com/dihanto/go-toko/middleware"
	"github.com/dihanto/go-toko/model/web/request"
	"github.com/dihanto/go-toko/model/web/response"
	"github.com/dihanto/go-toko/usecase"
	"github.com/julienschmidt/httprouter"
)

type OrderControllerImpl struct {
	Usecase usecase.OrderUsecase
	Route   *httprouter.Router
}

func NewOrderControllerImpl(usecase usecase.OrderUsecase, route *httprouter.Router) OrderController {
	controller := &OrderControllerImpl{
		Usecase: usecase,
		Route:   route,
	}

	controller.router(route)
	return controller
}

func (controller *OrderControllerImpl) router(route *httprouter.Router) {
	route.POST("/orders", middleware.OrderMiddleware(controller.AddOrder))
	route.GET("/orders/:id", middleware.OrderMiddleware(controller.FindOrder))
}

// addOrder
// @Summary Add order
// @Description Add order
// @Tags Order
// @Param request_body body request.AddOrder true "Request Body"
// @Success 201 {object} response.WebResponse
// @Security JWTAuth
// @Router /order [post]
func (controller *OrderControllerImpl) AddOrder(writer http.ResponseWriter, req *http.Request, param httprouter.Params) {
	request := request.AddOrder{}
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

	orderResponse, err := controller.Usecase.AddOrder(req.Context(), request)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}

	webResponse := response.WebResponse{
		Code:    http.StatusCreated,
		Message: "Order successfully created",
		Data:    orderResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(webResponse)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}

}

// findOrder
// @Summary Find order
// @Description Find order
// @Tags Order
// @Param id path integer true "Order ID"
// @Success 200 {object} response.WebResponse
// @Security JWTAuth
// @Router /order/{id} [post]
func (controller *OrderControllerImpl) FindOrder(writer http.ResponseWriter, req *http.Request, param httprouter.Params) {
	idString := param.ByName("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}

	orderResponse, err := controller.Usecase.FindOrder(req.Context(), id)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}

	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Message: "Success find Order",
		Data:    orderResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(webResponse)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}
}
