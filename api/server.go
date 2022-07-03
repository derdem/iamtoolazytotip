package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", rootLogger())
	http.ListenAndServe(":8080", r)
	return r
}

func rootLogger() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("root address called")
	}
}
