package main

import (
	"log"
	"net/http"
	"playlist-generator/web"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", web.GetRoutes()))
}