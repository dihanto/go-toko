package exception

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/dihanto/go-toko/model/web/response"
	"github.com/go-playground/validator/v10"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if err == false || err == "" {
		unauthorized(writer, request, err)
		return
	}

	if validationError(writer, request, err) {
		return
	}

	if err == sql.ErrNoRows {
		badRequestError(writer, request, err)
		return
	}

	internalServerError(writer, request, err)
}

func validationError(writer http.ResponseWriter, request *http.Request, errs interface{}) bool {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusBadRequest)

	exception, ok := errs.(validator.ValidationErrors)
	if ok {

		var messages interface{}
		message := make(map[string]string)

		for _, err := range exception {
			fieldName := strings.ToLower(err.Field())
			tag := err.ActualTag()

			switch tag {
			case "required":
				message[fieldName] = fmt.Sprintf("%s is required", fieldName)
				messages = message
			case "email":
				message[fieldName] = fmt.Sprintf("%s in not a valid email", fieldName)
				messages = message
			case "email_unique":
				message[fieldName] = fmt.Sprintf("%s must be unique", fieldName)
				messages = message
			case "min":
				message[fieldName] = fmt.Sprintf("%s must be at least %s characters long", fieldName, err.Param())
				messages = message
			case "wallet":
				message[fieldName] = "customer cannot have more than one wallet"
				messages = message
			default:
				message[fieldName] = fmt.Sprintf("validation error for %s: %s", fieldName, tag)
				messages = message
			}
		}

		webResponse := response.WebResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad Request",
			Data:    messages,
		}
		json.NewEncoder(writer).Encode(webResponse)

		return true

	} else {
		return false
	}

}

func badRequestError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-type", "application/json")

	webResponse := response.WebResponse{
		Code:    http.StatusBadRequest,
		Message: "Bad Request",
		Data:    fmt.Sprintf("%v", err),
	}

	json.NewEncoder(writer).Encode(webResponse)

}

func unauthorized(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-type", "application/json")
	writer.WriteHeader(http.StatusUnauthorized)

	if err == "" {
		webResponse := response.WebResponse{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized",
			Data:    "JWT token cannot be empty",
		}

		json.NewEncoder(writer).Encode(webResponse)
		return
	}

	if err == false {
		webResponse := response.WebResponse{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized",
			Data:    "Login failed/Password do not match",
		}

		json.NewEncoder(writer).Encode(webResponse)
		return
	}

	webResponse := response.WebResponse{
		Code:    http.StatusUnauthorized,
		Message: "Unauthorized",
		Data:    err,
	}

	json.NewEncoder(writer).Encode(webResponse)

}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-type", "application/json")

	webResponse := response.WebResponse{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error",
		Data:    fmt.Sprintf("%v", err),
	}

	json.NewEncoder(writer).Encode(webResponse)

}
