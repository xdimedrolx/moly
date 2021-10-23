package middleware

import (
	"net/http"

	v "github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
)

func Validator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			if ve, ok := err.(*v.Error); ok {
				return echo.NewHTTPError(http.StatusBadRequest, ve.Error())
			} else if ve, ok := err.(v.Errors); ok {
				return echo.NewHTTPError(http.StatusBadRequest, ve.Error())
			} else {
				return err
			}
		}
		return nil
	}
}