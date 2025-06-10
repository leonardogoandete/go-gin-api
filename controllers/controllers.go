package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/leonardogoandete/go-gin-api/database"
	"github.com/leonardogoandete/go-gin-api/models"
	"net/http"
)

func ExibePaginaIndex(ctx *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"alunos": alunos,
	})
}
func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// ExibeTodosAlunos godoc
// @Summary Exibe todos os alunos
// @Description Retorna uma lista de todos os alunos cadastrados
// @Tags alunos
// @Accept json
// @Produce json
// @Success 200 {array} models.Aluno
// @Router /api/alunos [get]
func ExibeTodosAlunos(ctx *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	ctx.JSON(http.StatusOK, alunos)
}

// ExibeAlunoPorID godoc
// @Summary Exibe um aluno por ID
// @Description Retorna os detalhes de um aluno específico pelo ID
// @Tags alunos
// @Accept json
// @Produce json
// @Param id path string true "ID do aluno"
// @Success 200 {object} models.Aluno
// @Failure 400 {object} gin.H "ID inválido"
// @Failure 404 {object} gin.H "Aluno não encontrado"
// @Router /api/alunos/{id} [get]
func ExibeAlunoPorID(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ID do aluno não pode ser vazio",
		})
		return
	}
	var aluno models.Aluno
	database.DB.First(&aluno, id)
	if aluno.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Aluno não encontrado",
		})
		return
	}
	ctx.JSON(http.StatusOK, aluno)
}

// ExibeAlunoPorNome godoc
// @Summary Exibe um aluno por nome
// @Description Retorna os detalhes de um aluno específico pelo nome
// @Tags alunos
// @Accept json
// @Produce json
// @Param nome path string true "Nome do aluno"
// @Success 200 {object} models.Aluno
// @Failure 400 {object} gin.H "Nome inválido"
// @Failure 404 {object} gin.H "Aluno não encontrado"
// @Router /api/alunos/nome/{nome} [get]
func ExibeAlunoPorNome(ctx *gin.Context) {
	nome := ctx.Param("nome")
	if nome == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Nome do aluno não pode ser vazio",
		})
		return
	}
	var aluno models.Aluno
	database.DB.Where(&models.Aluno{Nome: nome}).First(&aluno)
	if aluno.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Aluno não encontrado",
		})
		return
	}
	ctx.JSON(http.StatusOK, aluno)
}

// CriaNovoAluno godoc
// @Summary Cria um novo aluno
// @Description Cria um novo aluno com os dados fornecidos
// @Tags alunos
// @Accept json
// @Produce json
// @Param aluno body models.Aluno true "Dados do aluno"
// @Success 201 {object} models.Aluno
// @Failure 400 {object} gin.H "Dados inválidos ou falha de validação"
// @Router /api/alunos [post]
func CriaNovoAluno(ctx *gin.Context) {
	var aluno models.Aluno

	if err := ctx.ShouldBindJSON(&aluno); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Dados inválidos",
		})
		return
	}

	if err := models.ValidaAluno(&aluno); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Validação falhou: " + err.Error(),
		})
		return
	}
	database.DB.Create(&aluno)
	ctx.JSON(http.StatusCreated, aluno)
}

// DeletaAluno godoc
// @Summary Deleta um aluno
// @Description Deleta um aluno específico pelo ID
// @Tags alunos
// @Accept json
// @Produce json
// @Param id path string true "ID do aluno"
// @Success 200 {object} gin.H "OK"
// @Failure 400 {object} gin.H "ID inválido"
// @Failure 404 {object} gin.H "Aluno não encontrado"
// @Router /api/alunos/{id} [delete]
func DeletaAluno(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ID do aluno não pode ser vazio",
		})
		return
	}
	var aluno models.Aluno
	database.DB.Where("id = ?", id).First(&aluno)
	if aluno.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Aluno não encontrado",
		})
		return
	}
	database.DB.Delete(&aluno)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Aluno deletado com sucesso",
	})
}

// AtualizaAluno godoc
// @Summary Atualiza um aluno
// @Description Atualiza os dados de um aluno específico pelo ID
// @Tags alunos
// @Accept json
// @Produce json
// @Param id path string true "ID do aluno"
// @Param aluno body models.Aluno true "Dados do aluno"
// @Success 200 {object} models.Aluno
// @Failure 400 {object} gin.H "ID inválido ou dados inválidos"
// @Failure 404 {object} gin.H "Aluno não encontrado"
// @Router /api/alunos/{id} [patch]
func AtualizaAluno(ctx *gin.Context) {
	var aluno models.Aluno
	id := ctx.Params.ByName("id")
	database.DB.First(&aluno, id)

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ID do aluno não pode ser vazio",
		})
		return
	}
	if err := ctx.ShouldBindJSON(&aluno); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Dados inválidos",
		})
		return
	}
	if err := models.ValidaAluno(&aluno); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DB.Save(&aluno)
	ctx.JSON(http.StatusOK, aluno)
}

func ExibePaginaNotFound(ctx *gin.Context) {
	ctx.HTML(http.StatusNotFound, "404.html", nil)
}
