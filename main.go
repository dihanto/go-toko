package main

import (
	"fmt"
	"net/http"

	"github.com/dihanto/go-toko/config"
	"github.com/dihanto/go-toko/controller"
	"github.com/dihanto/go-toko/helper"
	"github.com/dihanto/go-toko/repository"
	"github.com/dihanto/go-toko/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := config.NewDb()
	validate := validator.New()
	var timeout int
	router := httprouter.New()

	{
		repository := repository.NewCustomerRepositoryImpl()
		usecase := usecase.NewCustomerUsecaseImpl(repository, db, validate, timeout)
		controller.NewCustomerControllerImpl(usecase, router)
	}
	{
		repository := repository.NewSellerRepositoryImpl()
		usecase := usecase.NewSellerUsecaseImpl(repository, db, validate, timeout)
		controller.NewSellerControllerImpl(usecase, router)
	}
	{
		repository := repository.NewProductRepositoryImpl()
		usecase := usecase.NewProductUsecaseImpl(repository, db, validate, timeout)
		controller.NewProductControllerImpl(usecase, router)
	}
	{
		repository := repository.NewWalletRepository()
		usecase := usecase.NewWalletUsecase(repository, db, validate, timeout)
		controller.NewWalletController(usecase, router)
	}
	{
		repository := repository.NewOrderRepositoryImpl()
		usecase := usecase.NewOrderUsecaseImpl(repository, db, validate, timeout)
		controller.NewOrderControllerImpl(usecase, router)
	}
	server := http.Server{
		Addr:    "localhost:2000",
		Handler: router,
	}
	fmt.Println("server running")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
