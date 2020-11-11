package services

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"quasarproject/entities"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

// TestObtainParameters tests ObtainParameters method
func TestObtainParameters(t *testing.T) {

	ctxTest, _ := gin.CreateTestContext(httptest.NewRecorder())

	body := `{
		"satelites": [
			{
				"name": "kenobi",
				"distance": 627,
				"message": [
					" ",
					"es",
					"",
					" ",
					""
				]
			},
			{
				"name": "skywalker",
				"distance": 208,
				"message": [
					"este",
					"",
					"un",
					"",
					"secreto"
				]
			},
			{
				"name": "sato",
				"distance": 450,
				"message": [
					"",
					"",
					"",
					"mensaje",
					""
				]
			}
		]
	}`

	ctxTest.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(body)))

	paramsService := NewGinParamsService()

	result, _ := paramsService.ObtainParameters(ctxTest)

	distances := make(map[string]float32)
	distances["kenobi"] = 627
	distances["skywalker"] = 208
	distances["sato"] = 450

	messages := [][]string{
		{" ", "es", "", " ", ""},
		{"este", "", "un", "", "secreto"},
		{"", "", "", "mensaje", ""},
	}

	expected := entities.ParseMessage{
		Distances: distances,
		Messages:  messages,
	}

	if reflect.DeepEqual(result, expected) {
		t.Fatal(result, " es distinto a: ", expected)
	}
}

// TestObtainParametersWrongQuantity tests ObtainParameters with error
func TestObtainParametersWrongQuantity(t *testing.T) {

	ctxTest, _ := gin.CreateTestContext(httptest.NewRecorder())

	body := `{
		"satelites": [
			{
				"name": "kenobi",
				"distance": 627,
				"message": [
					" ",
					" ",
					"es",
					"",
					" "
				]
			},
			{
				"name": "skywalker",
				"distance": 208,
				"message": [
					"este",
					"",
					"un",
					"",
					"secreto"
				]
			}
		]
	}`

	ctxTest.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(body)))

	paramsService := NewGinParamsService()

	_, err := paramsService.ObtainParameters(ctxTest)

	expected := "La cantidad de satélites ingresados es incorrecta"

	if expected != err.Error() {
		t.Fatal(err.Error(), " es distinto a: ", expected)
	}
}
