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

	if err != nil {
		errorMessage := fmt.Sprintf("%v", err)
		if strings.Contains(errorMessage, "token has invalid claims: token is expired") {
			expiredTokenError(writer, request, err)
			return
		}
	}

	internalServerError(writer, request, err)
}

func expiredTokenError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-type", "application/json")
	writer.WriteHeader(http.StatusUnauthorized)

	errorResponse := response.ErrorResponse{
		Code:    http.StatusUnauthorized,
		Message: "Token is expired",
	}

	json.NewEncoder(writer).Encode(errorResponse)
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

		errorResponse := response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: messages,
		}
		json.NewEncoder(writer).Encode(errorResponse)

		return true

	} else {
		return false
	}

}

func badRequestError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-type", "application/json")

	errorResponse := response.ErrorResponse{
		Code:    http.StatusBadRequest,
		Message: fmt.Sprintf("%v", err),
	}

	json.NewEncoder(writer).Encode(errorResponse)

}

func unauthorized(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-type", "application/json")
	writer.WriteHeader(http.StatusUnauthorized)

	if err == "" {
		errorResponse := response.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: "JWT token cannot be empty",
		}

		json.NewEncoder(writer).Encode(errorResponse)
		return
	}

	if err == false {
		errorResponse := response.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: "Login failed/Password do not match",
		}

		json.NewEncoder(writer).Encode(errorResponse)
		return
	}

	errorResponse := response.ErrorResponse{
		Code:    http.StatusUnauthorized,
		Message: err,
	}
	json.NewEncoder(writer).Encode(errorResponse)

}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-type", "application/json")

	errorResponse := response.ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: fmt.Sprintf("%v", err),
	}

	json.NewEncoder(writer).Encode(errorResponse)

}
