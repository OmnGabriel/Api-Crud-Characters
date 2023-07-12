package main

import (
	"fmt"

	"github.com/OmnGabriel/go-api-rest.git/models"
	"github.com/OmnGabriel/go-api-rest.git/routes"
)

func main() {
	models.Characters = []models.Character{
		{Id: 1, Name: "xpto 1", Backstory: "xpto1"},
		{Id: 2, Name: "xpto 2", Backstory: "xpto2"},
	}

	fmt.Println("Iniciando o servidor server com Go")
	routes.HandleRequest()
}
