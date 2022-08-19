package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/derdem/iamtoolazytotip/simulator"
	"github.com/gorilla/mux"
)

func Start() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", runTournament())
	http.ListenAndServe("localhost:8080", r)
	return r
}

type returnValue struct {
	id   string
	name string
	age  int
}

func runTournament() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		matchOutcome := simulator.TournamentSimulator()

		js, err := json.Marshal(matchOutcome)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}

		fmt.Println(matchOutcome)
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}
