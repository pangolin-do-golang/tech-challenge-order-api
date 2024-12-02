package controller_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/pangolin-do-golang/tech-challenge-order-api/internal/adapters/rest/controller"
	"github.com/pangolin-do-golang/tech-challenge-order-api/internal/errutil"
	"github.com/stretchr/testify/assert"
)

func TestBusinessErrorReturnsUnprocessableEntity(t *testing.T) {
	ctrl := &controller.AbstractController{}
	router := gin.Default()
	router.GET("/error", func(c *gin.Context) {
		ctrl.Error(c, errutil.NewBusinessError(errors.New("original error"), "business error occurred"))
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/error", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 422, w.Code)
	assert.JSONEq(t, `{"error":"business error occurred"}`, w.Body.String())
}

func TestInputErrorReturnsBadRequest(t *testing.T) {
	ctrl := &controller.AbstractController{}
	router := gin.Default()
	router.GET("/error", func(c *gin.Context) {
		ctrl.Error(c, errutil.NewInputError(errors.New("original error")))
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/error", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.JSONEq(t, `{"error":"Bad Request"}`, w.Body.String())
}

func TestUnknownErrorTypeReturnsInternalServerError(t *testing.T) {
	ctrl := &controller.AbstractController{}
	router := gin.Default()
	router.GET("/error", func(c *gin.Context) {
		ctrl.Error(c, &errutil.Error{Type: "UNKNOWN", Message: "unknown error"})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/error", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
	assert.JSONEq(t, `{"error":"Internal Server Error"}`, w.Body.String())
}

func TestNonCustomErrorReturnsInternalServerError(t *testing.T) {
	ctrl := &controller.AbstractController{}
	router := gin.Default()
	router.GET("/error", func(c *gin.Context) {
		ctrl.Error(c, errors.New("some error"))
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/error", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
	assert.JSONEq(t, `{"error":"Internal Server Error"}`, w.Body.String())
}
