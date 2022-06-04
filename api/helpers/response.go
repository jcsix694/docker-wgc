package helpers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ResponseData struct {
	Status  int         `json:"status"`
	Meta    interface{} `json:"meta"`
	Data    interface{} `json:"data"`
	Success interface{} `json:"success"`
	Errors  interface{} `json:"errors"`
}

type ApiError struct {
	Field string `json:"field"`
	Msg   string `json:"message"`
}

func RespondJSON(w *gin.Context, status int, payload interface{}, err error, meta interface{}) {
	var res ResponseData
	var validate validator.ValidationErrors

	res.Status = status
	res.Meta = meta
	res.Data = payload

	if errors.As(err, &validate) {
		out := make([]ApiError, len(validate))
		for i, fe := range validate {
			out[i] = ApiError{fe.Field(), ErrorMessage(fe.Tag())}
		}

		res.Errors = out
	} else if err != nil {
		res.Errors = err.Error()
	} else {
		res.Success = SuccessMessage(status)
	}

	w.JSON(status, res)
}

func ErrorMessage(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	}
	return ""
}

func SuccessMessage(status int) string {
	switch status {
	case http.StatusOK:
		return "ok"
	case http.StatusCreated:
		return "created"
	}
	return ""
}
