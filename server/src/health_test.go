package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/health", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &handler{}

	if assert.NoError(t, h.healthCheck(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		result := Response{}
		if err := json.NewDecoder(rec.Body).Decode(&result); err != nil {
			log.Fatalln(err)
		}
		assert.EqualValues(t, Response{Data: "OK", Errors: []string{}}, result)
	}
}
