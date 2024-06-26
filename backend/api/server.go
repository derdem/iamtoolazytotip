package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/derdem/iamtoolazytotip/readTournamentFromDb"
	"github.com/derdem/iamtoolazytotip/simulator"
)

func Start() *http.ServeMux {
	router := http.NewServeMux()
	corsHandler := enableCors(router)
	router.HandleFunc("/api/2021", run2021Tournament())
	router.HandleFunc("/api/2021-from-json", runTournamentFromJson("/app/dumps/tournament2.json"))
	router.HandleFunc("/api/2024", run2024Tournament())
	router.HandleFunc("/api/2024-from-json", runTournamentFromJson("/app/dumps/tournament3.json"))
	router.HandleFunc("/api/read-tournament-1", readTournament1())
	http.ListenAndServe(":8080", corsHandler)
	return router
}

func enableCors(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Replace * with the specific origins allowed to access the API
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			// Preflight request handling
			return
		}

		handler.ServeHTTP(w, r)
	})
}

func run2021Tournament() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		groups := readTournamentFromDb.GetTournament(2)
		matchOutcome := simulator.TournamentSimulator(groups)

		js, err := json.Marshal(matchOutcome)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}

		//fmt.Println(matchOutcome)
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}

func run2024Tournament() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tournament := readTournamentFromDb.GetTournament(3)
		finishedTournament := simulator.TournamentSimulator(tournament)

		js, err := json.Marshal(finishedTournament)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}

		//fmt.Println(matchOutcome)
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}

func readTournament1() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		groups := readTournamentFromDb.GetTournament(1)
		finishedTournament := simulator.TournamentSimulator(groups)
		js, err := json.Marshal(finishedTournament)

		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}

func runTournamentFromJson(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tournament := readTournamentFromDb.ReadTournamentFromFile(path)
		finishedTournament := simulator.TournamentSimulator(tournament)
		js, err := json.Marshal(finishedTournament)

		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}
