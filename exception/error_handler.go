package exception

import (
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
	internalServerError(writer, request, err)
}

func validationError(writer http.ResponseWriter, request *http.Request, errs interface{}) bool {
	exception, ok := errs.(validator.ValidationErrors)
	if ok {

		var messages []string
		for _, err := range exception {

			fieldName := strings.ToLower(err.Field())
			tag := err.ActualTag()
			switch tag {
			case "required":
				messages = append(messages, fmt.Sprintf("%s is required", fieldName))
			case "email":
				messages = append(messages, fmt.Sprintf("%s is not a valid email", fieldName))
			case "email_unique":
				messages = append(messages, fmt.Sprintf("%s must be unique", fieldName))
			case "min":
				messages = append(messages, fmt.Sprintf("%s must be at least %s characters long", fieldName, err.Param()))
			default:
				messages = append(messages, fmt.Sprintf("validation error for %s: %s", fieldName, tag))
			}
		}
		webResponse := response.WebResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad Request",
			Data:    messages,
		}
		writer.Header().Add("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(webResponse)
		return true
	} else {
		return false
	}

}

func unauthorized(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	if err == "" {
		webResponse := response.WebResponse{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized",
			Data:    "JWT token cannot be empty",
		}

		writer.Header().Add("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(webResponse)
		return
	}

	if err == false {
		webResponse := response.WebResponse{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized",
			Data:    "Login failed",
		}

		writer.Header().Add("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(webResponse)
		return
	}

	webResponse := response.WebResponse{
		Code:    http.StatusUnauthorized,
		Message: "Unauthorized",
		Data:    err,
	}

	writer.Header().Add("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(webResponse)

}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := response.WebResponse{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error",
		Data:    fmt.Sprintf("%v", err),
	}

	writer.Header().Add("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(webResponse)

}
