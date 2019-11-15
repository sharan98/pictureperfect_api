package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"rest-api/homepage"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("It works!")
	logger := log.New(os.Stdout, "rest-api", log.LstdFlags|log.Lshortfile)
	h := homepage.NewHandlers(logger)
	mux := mux.NewRouter()

	h.SetupRoutes(mux)
	http.ListenAndServe(":8000", mux)
}
