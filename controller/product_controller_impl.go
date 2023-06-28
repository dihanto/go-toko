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

type ProductControllerImpl struct {
	Usecase usecase.ProductUsecase
	Route   *httprouter.Router
}

func NewProductControllerImpl(usecase usecase.ProductUsecase, route *httprouter.Router) ProductController {
	controller := &ProductControllerImpl{
		Usecase: usecase,
		Route:   route,
	}
	controller.router(route)
	return controller
}

func (controller *ProductControllerImpl) router(route *httprouter.Router) {
	route.POST("/product", middleware.ProductMiddleware(controller.AddProduct))
	route.GET("/product", middleware.ProductMiddleware(controller.GetProduct))
	route.GET("/product/:id", middleware.ProductMiddleware(controller.FindById))
	route.PUT("/product/:id", middleware.ProductMiddleware(controller.UpdateProduct))
	route.DELETE("/product/:id", middleware.ProductMiddleware(controller.DeleteProduct))
}

func (controller *ProductControllerImpl) AddProduct(writer http.ResponseWriter, req *http.Request, param httprouter.Params) {
	request := request.AddProduct{}
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		log.Println(err)
		http.Error(writer, "Invalid request body", http.StatusBadRequest)
		return
	}

	authHeader := req.Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	request.IdSeller, err = helper.GenerateIdFromToken(tokenString)
	if err != nil {
		log.Println(err)
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
		return
	}

	productResponse, err := controller.Usecase.AddProduct(req.Context(), request)
	if err != nil {
		log.Println(err)
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
		return
	}
	webResponse := response.WebResponse{
		Code:    http.StatusCreated,
		Message: "Add product success",
		Data:    productResponse,
	}

	err = json.NewEncoder(writer).Encode(webResponse)
	if err != nil {
		log.Println(err)
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func (controller *ProductControllerImpl) GetProduct(writer http.ResponseWriter, req *http.Request, param httprouter.Params) {
	productResponses, err := controller.Usecase.GetProduct(req.Context())
	if err != nil {
		log.Println(err)
		http.Error(writer, "internal server error", http.StatusInternalServerError)
		return
	}

	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Message: "Success get all product",
		Data:    productResponses,
	}

	err = json.NewEncoder(writer).Encode(webResponse)
	if err != nil {
		log.Println(err)
		http.Error(writer, "internal server error", http.StatusInternalServerError)
		return
	}
}

func (controller *ProductControllerImpl) FindById(writer http.ResponseWriter, req *http.Request, param httprouter.Params) {
	idString := param.ByName("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Println(err)
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
		return
	}
	productResponse, err := controller.Usecase.FindById(req.Context(), id)
	if err != nil {
		log.Println(err)
		http.Error(writer, "internal server error", http.StatusInternalServerError)
		return
	}

	webReponse := response.WebResponse{
		Code:    http.StatusOK,
		Message: "Success find product by id",
		Data:    productResponse,
	}

	err = json.NewEncoder(writer).Encode(webReponse)
	if err != nil {
		log.Println(err)
		http.Error(writer, "internal server error", http.StatusInternalServerError)
		return
	}
}

func (controller *ProductControllerImpl) UpdateProduct(writer http.ResponseWriter, req *http.Request, param httprouter.Params) {
	request := request.UpdateProduct{}
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		log.Println(err)
		http.Error(writer, "Invalid request body", http.StatusBadRequest)
		return
	}
	idString := param.ByName("id")
	request.Id, err = strconv.Atoi(idString)
	if err != nil {
		log.Println(err)
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
		return
	}

	productResponse, err := controller.Usecase.UpdateProduct(req.Context(), request)
	if err != nil {
		log.Println(err)
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
		return
	}

	webRespones := response.WebResponse{
		Code:    http.StatusOK,
		Message: "Success update product",
		Data:    productResponse,
	}

	err = json.NewEncoder(writer).Encode(webRespones)
	if err != nil {
		log.Println(err)
		http.Error(writer, "internal server error", http.StatusInternalServerError)
	}
}

func (controller *ProductControllerImpl) DeleteProduct(writer http.ResponseWriter, req *http.Request, param httprouter.Params) {
	idString := param.ByName("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Println(err)
		http.Error(writer, "internal server error", http.StatusInternalServerError)
		return
	}

	err = controller.Usecase.DeleteProduct(req.Context(), id)
	if err != nil {
		log.Println(err)
		http.Error(writer, "internal server error", http.StatusInternalServerError)
		return
	}

	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Message: "Product success deleted",
	}

	err = json.NewEncoder(writer).Encode(webResponse)
	if err != nil {
		log.Println(err)
		http.Error(writer, "internal server error", http.StatusInternalServerError)
		return
	}
}
