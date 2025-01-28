package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestEmptyRecipe(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/recipe", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &handler{}

	if assert.NoError(t, h.generateRecipe(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		result := Response{}
		if err := json.NewDecoder(rec.Body).Decode(&result); err != nil {
			log.Fatalln(err)
		}
		assert.EqualValues(t, Response{Data: []interface{}{}, Errors: []string{}}, result)
	}
}

func TestRecipe(t *testing.T) {
	e := echo.New()
	ingredients, _ := json.Marshal([]string{"flour", "milk", "eggs"})
	req := httptest.NewRequest(http.MethodPost, "/recipe", strings.NewReader(string(ingredients)))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &handler{}

	if assert.NoError(t, h.generateRecipe(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		result := Response{}
		if err := json.NewDecoder(rec.Body).Decode(&result); err != nil {
			log.Fatalln(err)
		}
		assert.EqualValues(t, Response{Data: []interface{}{"flour", "milk", "eggs"}, Errors: []string{}}, result)
	}
}
