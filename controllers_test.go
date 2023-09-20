package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/OmnGabriel/go-api-rest.git/controllers"
	"github.com/OmnGabriel/go-api-rest.git/database"
	"github.com/OmnGabriel/go-api-rest.git/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int
var Name int

func RoutesTestingSetup() *gin.Engine {
	r := gin.Default()
	return r

}

func MakeANewCharacterMock() {
	character := models.Character{
		Name:      "Make A New Character Mock",
		Backstory: "This account was created and then deleted",
	}
	database.DB.Create(&character)
	ID = int(character.ID)
}

func DeletingCharacterMock() {
	var character models.Character
	database.DB.Unscoped().Delete(&character, ID)
}

func TestToCheckFunctionTypeOfHome(t *testing.T) {
	//Assert
	r := RoutesTestingSetup()
	r.GET("/", controllers.Home)
	req, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	responseType, _ := io.ReadAll(response.Body)
	expected := string("text/plain; charset=utf-8")
	//Act
	r.ServeHTTP(response, req)

	//Assert
	assert.Equal(t, expected, http.DetectContentType(responseType), "Return unexpected Type")
}

func TestToCheckTheContentHomeFunction(t *testing.T) {

	//Assert
	r := RoutesTestingSetup()
	r.GET("/", controllers.Home)
	req, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	expected := "Home Page"

	//Act
	r.ServeHTTP(response, req)

	//Assert
	assert.Equal(t, response.Body.String(), expected, "Return unexpected Type")
}

func TestToCheckFunctionTypeOfAllCharacters(t *testing.T) {
	//Assert
	database.ConnectWithDatabase()
	r := RoutesTestingSetup()
	r.GET("/api/characters", controllers.AllCharacters)
	req, _ := http.NewRequest("GET", "/api/characters", nil)
	response := httptest.NewRecorder()
	responseType, _ := io.ReadAll(response.Body)
	expected := string("text/plain; charset=utf-8")
	//Act
	r.ServeHTTP(response, req)

	//Assert
	assert.Equal(t, expected, http.DetectContentType(responseType), "Return unexpected Type")
}

// func TestToCheckFunctionTypeOfMakeANewCharacter(t *testing.T) {
// 	//Assert
// 	database.ConnectWithDatabase()
// 	r := RoutesTestingSetup()
// 	r.POST("/api/characters", controllers.MakeANewCharacter)
// 	req, _ := http.NewRequest("POST", "/api/characters", nil)
// 	response := httptest.NewRecorder()
// 	responseType, _ := io.ReadAll(response.Body)
// 	expected := string("text/plain; charset=utf-8")

// 	//Act
// 	r.ServeHTTP(response, req)

// 	//Assert
// 	assert.Equal(t, expected, http.DetectContentType(responseType), "Return unexpected Type")
// }

// func TestToCheckIfTheFunctionCreateANewCharacterIsReallyCreatingANewCharacter(t *testing.T) {
// 	//Assert
// 	var character = models.Character{
// 		Name:      "TesteeeeMake A New Character With The Function Make A New Character",
// 		Backstory: "This account has been created",
// 	}
// 	database.ConnectWithDatabase()
// 	defer DeletingCharacterMock()
// 	r := RoutesTestingSetup()
// 	r.POST("/api/characters", controllers.MakeANewCharacter)
// 	jsonValue, _ := json.Marshal(character)
// 	req, _ := http.NewRequest("POST", "/api/characters", bytes.NewBuffer(jsonValue))
// 	response := httptest.NewRecorder()

// 	//Act
// 	r.ServeHTTP(response, req)

// 	//Assert
// 	assert.Equal(t, http.StatusCreated, response.Code)

// }

// func TestToCheckIfTheFunctionReturnACharacterIsReallyReturningACharacter(t *testing.T) {
// 	//Assert
// 	database.ConnectWithDatabase()
// 	MakeANewCharacterMock()
// 	defer DeletingCharacterMock()
// 	r := RoutesTestingSetup()
// 	r.GET("/api/characters/:id", controllers.ReturnACharacter)
// 	searchPath := "/api/characters/" + strconv.Itoa(ID)
// 	req, _ := http.NewRequest("GET", searchPath, nil)
// 	response := httptest.NewRecorder()

// 	//Act
// 	r.ServeHTTP(response, req)

// 	//Arrange
// 	var characterMock models.Character
// 	json.Unmarshal(response.Body.Bytes(), &characterMock)
// 	assert.Equal(t, "Make A New Character Mock", characterMock.Name, "The names must be the same")
// 	assert.Equal(t, http.StatusOK, response.Code)

// }

// func TestToCheckValidRequestOfFunctionEditACharacter(t *testing.T) {
// 	//Arrange
// 	database.ConnectWithDatabase()
// 	MakeANewCharacterMock()
// 	defer DeletingCharacterMock()
// 	r := RoutesTestingSetup()
// 	r.PATCH("/api/characters/:id", controllers.EditACharacter)
// 	var character = models.Character{
// 		Name:      "user for edit character",
// 		Backstory: "not so important",
// 	}
// 	jsonValue, _ := json.Marshal(character)
// 	pathForEditing := "/api/characters/" + strconv.Itoa(ID)
// 	req, _ := http.NewRequest("PATCH", pathForEditing, bytes.NewBuffer(jsonValue))
// 	w := httptest.NewRecorder()

// 	//Act
// 	r.ServeHTTP(w, req)

// 	//Arrange
// 	assert.Equal(t, http.StatusOK, w.Code)

// }

// func TestToCheckNonValidRequestOfFunctionEditACharacter(t *testing.T) {
// 	//Arrange
// 	database.ConnectWithDatabase()
// 	MakeANewCharacterMock()
// 	defer DeletingCharacterMock()
// 	r := RoutesTestingSetup()
// 	r.PATCH("/api/characters/:id", controllers.EditACharacter)
// 	var character = models.Character{
// 		Name:      "user for edit character",
// 		Backstory: "not so important",
// 	}
// 	jsonValue, _ := json.Marshal(character)

// 	reqNotFound, _ := http.NewRequest("PATCH", "/api/characters/", bytes.NewBuffer(jsonValue))
// 	w := httptest.NewRecorder()

// 	//Act
// 	r.ServeHTTP(w, reqNotFound)

// 	// Assert
// 	assert.Equal(t, http.StatusNotFound, w.Code)
// }

// func TestToCheckIfTheFunctionDeletACharacterIsReallyDeletingACharacter(t *testing.T) {
// 	database.ConnectWithDatabase()
// 	MakeANewCharacterMock()
// 	r := RoutesTestingSetup()
// 	r.DELETE("/api/characters/:id", controllers.DeletACharacter)
// 	searchPath := "/api/characters/" + strconv.Itoa(ID)
// 	req, _ := http.NewRequest("DELETE", searchPath, nil)
// 	response := httptest.NewRecorder()
// 	r.ServeHTTP(response, req)
// 	assert.Equal(t, http.StatusOK, response.Code)
// }
