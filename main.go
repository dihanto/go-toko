package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dihanto/go-toko/config"
	"github.com/dihanto/go-toko/controller"
	_ "github.com/dihanto/go-toko/docs"
	"github.com/dihanto/go-toko/helper"
	"github.com/dihanto/go-toko/repository"
	"github.com/dihanto/go-toko/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"github.com/spf13/viper"
)

// @title           Go-Toko API
// @version         1.0
// @description     This is a simple API for marketplace.

// @contact.name   Dihanto
// @contact.email  dihanto2306@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:2000

// @SecurityDefinitions JWTAuth
// @TokenURL /login
// @SecurityScheme bearerAuth

func main() {

	config.InitLoadConfiguration()
	// serverHost := viper.GetString("server.host")
	serverPort := viper.GetString("server.port")
	timeout := viper.GetDuration("usecase.timeout")

	db := config.InitDatabaseConnection()

	validate := validator.New()
	validate.RegisterValidation("email_unique", helper.ValdateEmailUnique)
	validate.RegisterValidation("wallet", helper.ValidateUserOnlyHaveOneWallet)
	validate.RegisterValidation("wishlist", helper.ValidateUserHaveOneWishlistInOneProduct)

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
		Addr:    ":" + serverPort,
		Handler: router,
	}

	fmt.Println("server running")
	err := server.ListenAndServe()
	log.Println(err)
}
