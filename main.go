package main

import (
	"fmt"

	"github.com/marcosmrs8/go-rest-api/database"
	"github.com/marcosmrs8/go-rest-api/routes"
)

func main() {
	database.ConectComBancoDeDados()
	fmt.Println("Iniciando servidor...")
	routes.HandleRequest()
}
