package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/derdem/iamtoolazytotip/simulator/em2021"
)

func Start() *http.ServeMux {
	router := http.NewServeMux()
	corsHandler := enableCors(router)
	router.HandleFunc("/api/2021", run2021Tournament())
	router.HandleFunc("/api/2024", run2024Tournament())
	router.HandleFunc("/api/run-custom", runCustomTournament())
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
		matchOutcome := em2021.Run2021Tournament()

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
		matchOutcome := em2021.Run2021Tournament()

		js, err := json.Marshal(matchOutcome)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}

		//fmt.Println(matchOutcome)
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}

func runCustomTournament() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resBody, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatalf("Error happened in reading request body. Err: %s", err)
		}

		log.Println(string(resBody))
		w.Header().Set("Content-Type", "application/json")
		w.Write(resBody)

	}
}
