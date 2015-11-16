package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

var port string = "3000"

func main() {
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}

	router := mux.NewRouter()

	router.HandleFunc("/", hello).Methods("GET")

	n := negroni.New()
	n.UseHandler(router)
	n.Run(fmt.Sprintf(":%s", port))
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test" + os.Getenv("MESSAGE")))
}
