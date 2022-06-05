package helpers

import (
	"encoding/json"
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

type ValidationError struct {
	Field string `json:"field"`
	Msg   string `json:"message"`
}

func RespondJSON(w *gin.Context, status int, payload interface{}, err error, meta interface{}) {
	var res ResponseData
	var validatorErr validator.ValidationErrors
	var jsonErr *json.UnmarshalTypeError

	res.Status = status
	res.Meta = meta
	res.Data = payload

	if errors.As(err, &validatorErr) {
		out := []ValidationError{}

		for _, f := range validatorErr {
			out = append(out, ValidationError{Field: f.Field(), Msg: ValidationErrorMessages(f.Tag())})
		}

		res.Errors = out
	} else if errors.As(err, &jsonErr) {
		// To Do - Validate them all at once instead of one by one

		out := []ValidationError{}
		out = append(out, ValidationError{Field: jsonErr.Field, Msg: ValidationErrorMessages(jsonErr.Type.String())})

		res.Errors = out
	} else if err != nil {
		res.Errors = err.Error()
	} else {
		res.Success = SuccessMessage(status)
	}

	w.JSON(status, res)
}

func SuccessMessage(status int) string {
	switch status {
	case http.StatusOK:
		return "ok"
	case http.StatusCreated:
		return "created"
	}
	return "-"
}
