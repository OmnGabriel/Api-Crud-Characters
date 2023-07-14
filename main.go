package main

import (
	"fmt"

	"github.com/OmnGabriel/go-api-rest.git/database"
	"github.com/OmnGabriel/go-api-rest.git/routes"
)

func main() {
	database.ConnectWithDatabase()
	fmt.Println("Iniciando o servidor server com Go")
	routes.HandleRequest()
}
