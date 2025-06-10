package main

import (
	"github.com/gin-gonic/gin"
	"github.com/leonardogoandete/go-gin-api/controllers"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetupRotasDeTeste() *gin.Engine {
	rotas := gin.Default()
	return rotas
}

func TestVerificaRetonoPingComSucesso(t *testing.T) {
	r := SetupRotasDeTeste()
	r.GET("/ping", controllers.Ping)
	req, _ := http.NewRequest("GET", "/ping", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code, "Status code deve ser 200 OK")
}
