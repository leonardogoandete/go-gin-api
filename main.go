package main

import (
	"github.com/leonardogoandete/go-gin-api/database"
	"github.com/leonardogoandete/go-gin-api/routes"
)

func main() {

	database.ConectaComBancoDeDados()
	routes.HandleRequests()

}
