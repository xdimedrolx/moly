package http

import (
	"github.com/labstack/echo/v4"
	useCase "github.com/xdimedrolx/moly/internal/app/application/use_case"
	"net/http"
)

// @Summary Echo
// @Tags webhook
// @Accept  json
// @Produce json
// @Param body body useCase.EchoRequest true "dto"
// @Success 200 {object} useCase.EchoResponse
// @Failure 400,404 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Failure default {object} echo.HTTPError
// @Router /webhook/echo [get]
func (r *RouteHandlers) GetWebhookEcho(c echo.Context) error {
	cmd := new(useCase.EchoRequest)

	if err := c.Bind(cmd); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	resp, err := r.useCases.Echo(c.Request().Context(), cmd)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}
