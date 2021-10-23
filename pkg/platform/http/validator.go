package http

import (
	"net/http"

	v "github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if _, err := v.ValidateStruct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}

func AddValidator(s *Server) {
	s.Validator = &CustomValidator{}
}