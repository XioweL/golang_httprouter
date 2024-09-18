package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		id := params.ByName("id")
		text := "Products " + id
		fmt.Fprint(w, text)
	})
	r := httptest.NewRequest("GET", "http://localhost:8080/products/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, r)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Products 1", string(body))

}

//func ParamsFunc() {
//	router := httprouter.New()
//	router.GET("/products/:id", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
//		id := params.ByName("id")
//		text := "Products " + id
//		fmt.Fprint(w, text)
//	})
//	server := http.Server{
//		Handler: router,
//		Addr:    "https://localhost:8080/products/",
//	}
//}
