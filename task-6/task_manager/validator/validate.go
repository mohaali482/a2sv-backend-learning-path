package validator

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ReturnErrorResponse(err error, c *gin.Context) {
	if err != nil {
		var validationErrors []ErrorResponse
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, ErrorResponse{
				Field:   err.Field(),
				Message: SetValidationResult(err.Tag()),
			})
		}
		c.JSON(http.StatusBadRequest, validationErrors)
		return

	}
}

func SetValidationResult(tag string) string {
	switch tag {
	case "required":
		tag = "This field is required"
	case "email":
		tag = "This field is not a valid email"
	case "min":
		tag = "This field is too short"
	case "max":
		tag = "This field is too long"
	case "eqfield":
		tag = "This field is not equal to the other field"
	case "mongodb":
		tag = "This field is not a valid ObjectID"
	}
	return tag
}

func BindingValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
}
