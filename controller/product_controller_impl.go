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
	route.POST("/products", middleware.ProductMiddleware(controller.AddProduct))
	route.GET("/products", middleware.MindMiddleware(controller.GetProduct))
	route.GET("/products/:id", middleware.MindMiddleware(controller.FindById))
	route.PUT("/products/:id", middleware.ProductMiddleware(controller.UpdateProduct))
	route.DELETE("/products/:id", middleware.ProductMiddleware(controller.DeleteProduct))
	route.GET("/products/", middleware.ProductMiddleware(controller.FindByName))
}

// addProduct
// @Summary Add product
// @Description Add product
// @Tags Product
// @Param request_body body request.AddProduct true "Add Product"
// @Success 201 {object} response.WebResponse
// @Security JWTAuth
// @Router /product [post]
func (controller *ProductControllerImpl) AddProduct(writer http.ResponseWriter, req *http.Request, param httprouter.Params) {
	request := request.AddProduct{}
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}

	authHeader := req.Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	request.IdSeller, err = helper.GenerateIdFromToken(tokenString)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}

	productResponse, err := controller.Usecase.AddProduct(req.Context(), request)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}
	webResponse := response.WebResponse{
		Code:    http.StatusCreated,
		Message: "Add product success",
		Data:    productResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(webResponse)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}
}

// getProduct
// @Summary Get product
// @Description Get product
// @Tags Product
// @Success 200 {object} response.WebResponse
// @Security JWTAuth
// @Router /product [get]
func (controller *ProductControllerImpl) GetProduct(writer http.ResponseWriter, req *http.Request, param httprouter.Params) {
	productResponses, err := controller.Usecase.GetProduct(req.Context())
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}

	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Message: "Success get all product",
		Data:    productResponses,
	}

	writer.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(webResponse)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}
}

// findProduct
// @Summary Find product
// @Description Find product
// @Tags Product
// @Param id path integer true "Product ID"
// @Success 200 {object} response.WebResponse
// @Security JWTAuth
// @Router /product/{id} [get]
func (controller *ProductControllerImpl) FindById(writer http.ResponseWriter, req *http.Request, param httprouter.Params) {
	idString := param.ByName("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}
	productResponse, err := controller.Usecase.FindById(req.Context(), id)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}

	webReponse := response.WebResponse{
		Code:    http.StatusOK,
		Message: "Success find product by id",
		Data:    productResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(webReponse)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}
}

// updateProduct
// @Summary Update product
// @Description Update product
// @Tags Product
// @Param request_body body request.UpdateProduct true "Update Product"
// @Param id path integer true "Product ID"
// @Success 200 {object} response.WebResponse
// @Security JWTAuth
// @Router /product/{id} [put]
func (controller *ProductControllerImpl) UpdateProduct(writer http.ResponseWriter, req *http.Request, param httprouter.Params) {
	request := request.UpdateProduct{}
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}
	idString := param.ByName("id")
	request.Id, err = strconv.Atoi(idString)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}

	productResponse, err := controller.Usecase.UpdateProduct(req.Context(), request)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}

	webRespones := response.WebResponse{
		Code:    http.StatusOK,
		Message: "Success update product",
		Data:    productResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(webRespones)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}
}

// deleteProduct
// @Summary Delete product
// @Description Delete product
// @Tags Product
// @Param id path integer true "Product ID"
// @Success 200 {object} response.WebResponse
// @Security JWTAuth
// @Router /product/{id} [delete]
func (controller *ProductControllerImpl) DeleteProduct(writer http.ResponseWriter, req *http.Request, param httprouter.Params) {
	idString := param.ByName("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}

	err = controller.Usecase.DeleteProduct(req.Context(), id)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}

	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Message: "Product success deleted",
	}

	writer.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(webResponse)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}
}

func (controller *ProductControllerImpl) FindByName(writer http.ResponseWriter, req *http.Request, param httprouter.Params) {
	search := req.URL.Query().Get("search")
	offset := req.URL.Query().Get("offset")
	limit := req.URL.Query().Get("limit")

	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}

	productsWithPagination, err := controller.Usecase.FindByName(req.Context(), search, offsetInt, limitInt)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}

	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Message: "Success find product",
		Data:    productsWithPagination,
	}

	writer.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(webResponse)
	if err != nil {
		exception.ErrorHandler(writer, req, err)
		return
	}
}
