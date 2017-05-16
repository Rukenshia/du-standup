package main

import (
	"github.com/julienschmidt/httprouter"

	"encoding/json"
	"log"
	"net/http"
	"time"
)

var standup Standup
var timezone *time.Location

func main() {
	loc, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		log.Fatal(err)
	}
	timezone = loc

	standup = NewStandup()

	log.Printf("next standup expires at %v", standup.Expires)

	log.Println("adding routes")
	router := httprouter.New()
	router.GET("/api/standup", apiGetStandup)
	router.GET("/api/categories", apiGetCategories)
	router.GET("/api/categories/:category", apiGetCategory)
	router.GET("/api/categories/:category/entries", apiGetEntries)
	router.POST("/api/categories/:category/entries", apiCreateEntry)
	router.DELETE("/api/categories/:category/entries/:entry", apiDeleteEntry)
	router.GET("/api/categories/:category/entries/:entry", apiGetEntry)
	router.PUT("/api/categories/:category/entries/:entry", apiUpdateEntry)
	router.POST("/api/categories/:category/entries/:entry/vote", apiVoteEntry)

	router.ServeFiles("/web/*filepath", http.Dir("web"))

	log.Println("starting webserver")
	log.Fatal(http.ListenAndServe(":8080", StandupMiddleware(router)))
}

func getJSON(r *http.Request, target interface{}) error {
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
