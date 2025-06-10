package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/leonardogoandete/go-gin-api/database"
	"github.com/leonardogoandete/go-gin-api/models"
	"net/http"
)

func ExibeTodosAlunos(ctx *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	ctx.JSON(http.StatusOK, alunos)
}

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
