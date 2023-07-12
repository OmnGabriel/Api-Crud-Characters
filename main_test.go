package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/OmnGabriel/go-api-rest.git/controllers"
)

var w http.ResponseWriter
var r *http.Request

func TestToCheckFunctionStatusCodeOfHome(t *testing.T) {
	//Assert
	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.Home)

	//Act
	handler.ServeHTTP(rr, req)

	//Assert
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Home Page"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

}

func TestToCheckTheContentHomeFunction(t *testing.T) {
	//Assert

	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatal(err)
		controllers.Home(w, r)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.Home)

	//Act
	handler.ServeHTTP(rr, req)

	//Assert

	expected := "Home Page"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

}
