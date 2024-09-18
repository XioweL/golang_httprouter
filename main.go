package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Hello GET")
	})

	server := http.Server{
		Handler: router,
		Addr:    "localhost:8080",
	}
	server.ListenAndServe()

}
