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
	router.GET("/api/entries", apiGetEntries)
	router.POST("/api/entries", apiCreateEntry)
	router.DELETE("/api/entries/:id", apiDeleteEntry)
	router.PUT("/api/entries/:id", apiUpdateEntry)
	router.GET("/api/standup", apiGetStandup)

	router.ServeFiles("/web/*filepath", http.Dir("web"))

	log.Fatal(http.ListenAndServe(":8080", StandupMiddleware(router)))
}

func getJSON(r *http.Request, target interface{}) error {
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
