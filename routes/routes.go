package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/leonardogoandete/go-gin-api/controllers"
	docs "github.com/leonardogoandete/go-gin-api/docs"
	"github.com/leonardogoandete/go-gin-api/middlewares"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func HandleRequests() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.Use(middlewares.ConfigureContentType())
	r.Use(middlewares.ConfigureLogger())
	r.GET("/api/ping", controllers.Ping)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/api/alunos", controllers.ExibeTodosAlunos)
	r.POST("/api/alunos", controllers.CriaNovoAluno)
	r.DELETE("/api/alunos/:id", controllers.DeletaAluno)
	r.PATCH("/api/alunos/:id", controllers.AtualizaAluno)
	r.GET("/api/alunos/:id", controllers.ExibeAlunoPorID)
	r.GET("/api/alunos/nome/:nome", controllers.ExibeAlunoPorNome)
	r.GET("/index", controllers.ExibePaginaIndex)
	r.NoRoute(controllers.ExibePaginaNotFound)
	err := r.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
