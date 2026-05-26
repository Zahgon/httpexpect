package examples

import (
	"net/http"
)

type (
	fruitMap map[string]interface{}
)

// FruitsHandler creates http.Handler for the fruits server.
//
// Routes:
//
//	GET /fruits           get fruit list
//	GET /fruits/{name}    get fruit
//	PUT /fruits/{name}    add or update fruit
func FruitsHandler() http.Handler { _ = "STUB: not implemented"; return *new(http.Handler) }

func handleFruitList(fruits fruitMap, w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func handleFruit(fruits fruitMap, w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
