package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/leonardogoandete/go-gin-api/controllers"
	"github.com/leonardogoandete/go-gin-api/middlewares"
)

func HandleRequests() {
	r := gin.Default()
	r.Use(middlewares.ConfigureContentType())
	r.Use(middlewares.ConfigureLogger())
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.PATCH("/alunos/:id", controllers.AtualizaAluno)
	r.GET("/alunos/:id", controllers.ExibeAlunoPorID)
	r.GET("/alunos/nome/:nome", controllers.ExibeAlunoPorNome)
	err := r.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
