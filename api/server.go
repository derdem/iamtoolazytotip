package api

import (
	"net/http"

	"github.com/derdem/iamtoolazytotip/simulator"
	"github.com/gorilla/mux"
)

func Start() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", runTournament())
	http.ListenAndServe(":8080", r)
	return r
}

func runTournament() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		simulator.TournamentSimulator()
	}
}
