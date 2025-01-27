package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handler) healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, Response{Status: "OK"})
}
