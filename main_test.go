package main

import (
	"github.com/gin-gonic/gin"
	"github.com/leonardogoandete/go-gin-api/controllers"
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
	if resp.Code != http.StatusOK {
		t.Errorf("Código de resposta esperado: %d, mas obteve: %d", http.StatusOK, resp.Code)
	}
}
