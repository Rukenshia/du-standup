package main

import (
	"github.com/julienschmidt/httprouter"

	"encoding/json"
	"log"
	"net/http"
)

var standup Standup

func main() {
	standup = NewStandup()

	router := httprouter.New()
	router.GET("/api/standup", apiGetStandup)
	router.GET("/api/categories", apiGetCategories)
	router.GET("/api/categories/:category", apiGetCategory)
	router.GET("/api/categories/:category/entries", apiGetEntries)
	router.POST("/api/categories/:category/entries", apiCreateEntry)
	router.DELETE("/api/categories/:category/entries/:id", apiDeleteEntry)
	router.GET("/api/categories/:category/entries/:id", apiGetEntry)
	router.PUT("/api/categories/:category/entries/:id", apiUpdateEntry)

	router.ServeFiles("/web/*filepath", http.Dir("web"))

	log.Fatal(http.ListenAndServe(":8080", StandupMiddleware(router)))
}

func getJSON(r *http.Request, target interface{}) error {
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
