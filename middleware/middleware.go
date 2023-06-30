package middleware

import (
	"net/http"
	"strings"

	"github.com/dihanto/go-toko/exception"
	"github.com/dihanto/go-toko/helper"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

func MindMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
		logger := logrus.New()
		logger.Infoln(request.Method)
		logger.Infoln(request.RequestURI)

		authHeader := request.Header.Get("Authorization")
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			exception.ErrorHandler(writer, request, tokenString)
			return
		}

		token, err := helper.ParseJWTString(tokenString)
		if err != nil || !token.Valid {
			exception.ErrorHandler(writer, request, err)
			return
		}

		next(writer, request, param)
	}
}

func ProductMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
		authHeader := request.Header.Get("Authorization")
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			exception.ErrorHandler(writer, request, tokenString)
			return
		}

		token, err := helper.ParseJWTString(tokenString)
		if err != nil || !token.Valid {
			exception.ErrorHandler(writer, request, err)
			return
		}

		role, err := helper.GenerateRoleFromToken(token)
		if err != nil {
			exception.ErrorHandler(writer, request, err)
			return
		}

		if role != "seller" {
			exception.ErrorHandler(writer, request, false)
			return
		}

		next(writer, request, param)

	}

}

func OrderMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
		authHeader := request.Header.Get("Authorization")
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			exception.ErrorHandler(writer, request, tokenString)
			return
		}

		token, err := helper.ParseJWTString(tokenString)
		if err != nil || !token.Valid {
			exception.ErrorHandler(writer, request, err)
			return
		}

		role, err := helper.GenerateRoleFromToken(token)
		if err != nil {
			exception.ErrorHandler(writer, request, err)
			return
		}

		if role != "customer" {
			exception.ErrorHandler(writer, request, false)
			return
		}

		next(writer, request, param)

	}

}
