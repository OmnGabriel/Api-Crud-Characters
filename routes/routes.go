package routes

import (
	"log"
	"net/http"

	"github.com/OmnGabriel/go-api-rest.git/controllers"
	"github.com/OmnGabriel/go-api-rest.git/middleware"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.Use(middleware.ContentTypeMiddleware())

	r.GET("/", controllers.Home)
	r.GET("/api/characters", controllers.AllCharacters)
	r.GET("/api/characters/:id", controllers.ReturnACharacter)
	r.POST("/api/characters", controllers.MakeANewCharacter)
	r.DELETE("/api/characters/:id", controllers.DeletACharacter)
	r.PATCH("/api/characters/:id", controllers.EditACharacter)
	log.Fatal(http.ListenAndServe(":8080", r))
}
