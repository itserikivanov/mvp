package main

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Generate recipe for given products
//
//  @Summary      Generate recipe
//  @Description  generate recipe
//  @Tags         recipe
//  @Accept       json
//  @Produce      json
//  @Success      200  {array}   string
//  @Router       /api/recipe [post]
func (h *handler) generateRecipe(c echo.Context) error {
	var ingredients = make([]string, 0)
	json.NewDecoder(c.Request().Body).Decode(&ingredients)
	return c.JSON(http.StatusOK, Response{Data: ingredients})
}
