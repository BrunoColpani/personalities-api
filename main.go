package main

import (
	"fmt"
	"personalities-api/database"
	"personalities-api/models"
	"personalities-api/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{Id: 1, Name: "Nome 1", History: "Historia 1"},
		{Id: 2, Name: "Nome 2", History: "Historia 2"},
	}
	database.ConnectDb()
	fmt.Println("Iniciando o servidor")
	routes.HandleRequests()
}
