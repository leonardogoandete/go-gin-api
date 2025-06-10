package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/leonardogoandete/go-gin-api/controllers"
	"github.com/leonardogoandete/go-gin-api/database"
	"github.com/leonardogoandete/go-gin-api/models"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var ID int

func SetupRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func CriaAlunoMock() {
	aluno := models.Aluno{Nome: "Aluno Mock", CPF: "12345678901", RG: "123456789"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletaAlunoMock() {
	aluno := models.Aluno{}
	database.DB.Delete(&aluno, ID)

}
func TestVerificaRetonoPingComSucesso(t *testing.T) {
	r := SetupRotasDeTeste()
	r.GET("/api/ping", controllers.Ping)
	req, _ := http.NewRequest("GET", "/api/ping", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code, "Status code deve ser 200 OK")
	mockDaResposta := `{"message":"pong"}`
	respBody, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, mockDaResposta, string(respBody), "Resposta deve ser igual a "+mockDaResposta)
}

func TestExibeTodosAlunosComSucesso(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupRotasDeTeste()
	r.GET("/api/alunos", controllers.ExibeTodosAlunos)
	req, _ := http.NewRequest("GET", "/api/alunos", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code, "Status code deve ser 200 OK")
	fmt.Println(resp.Body)
}

func TestExibeAlunoPorNomeComSucesso(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupRotasDeTeste()
	r.GET("/api/alunos/nome/:nome", controllers.ExibeAlunoPorNome)
	req, _ := http.NewRequest("GET", fmt.Sprintf("/api/alunos/nome/%s", "Aluno Mock"), nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code, "Status code deve ser 200 OK")
	fmt.Println(resp.Body)
}

func TestExibeAlunoPorIDComSucesso(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupRotasDeTeste()
	r.GET("/api/alunos/:id", controllers.ExibeAlunoPorID)
	req, _ := http.NewRequest("GET", fmt.Sprintf("/api/alunos/%d", ID), nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code, "Status code deve ser 200 OK")
	var alunoMock models.Aluno
	json.Unmarshal(resp.Body.Bytes(), &alunoMock)
	assert.Equal(t, "Aluno Mock", alunoMock.Nome, "Nome do aluno deve ser 'Aluno Mock'")
	assert.Equal(t, "12345678901", alunoMock.CPF, "CPF do aluno deve ser '12345678901'")
	assert.Equal(t, "123456789", alunoMock.RG, "RG do aluno deve ser '123456789'")
}

func TestDeletarAlunoComSucesso(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	r := SetupRotasDeTeste()
	r.DELETE("/api/alunos/:id", controllers.DeletaAluno)
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/api/alunos/%d", ID), nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code, "Status code deve ser 200 OK")
}

func TestAtualizarAlunoComSucesso(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupRotasDeTeste()
	r.PATCH("/api/alunos/:id", controllers.AtualizaAluno)
	alunoAtualizado := models.Aluno{Nome: "Aluno Atualizado", CPF: "09876543210", RG: "987654321"}
	jsonAluno, _ := json.Marshal(alunoAtualizado)
	req, _ := http.NewRequest("PATCH", fmt.Sprintf("/api/alunos/%d", ID), bytes.NewBuffer(jsonAluno))
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code, "Status code deve ser 200 OK")
	var alunoMock models.Aluno
	json.Unmarshal(resp.Body.Bytes(), &alunoMock)
	assert.Equal(t, "Aluno Atualizado", alunoMock.Nome, "Nome do aluno deve ser 'Aluno Atualizado'")
	assert.Equal(t, "09876543210", alunoMock.CPF, "CPF do aluno deve ser '09876543210'")
	assert.Equal(t, "987654321", alunoMock.RG, "RG do aluno deve ser '987654321'")
}
