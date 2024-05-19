package readTournamentFromDb_test

import (
	// "encoding/json"
	"fmt"
	// "os"
	"testing"

	"github.com/derdem/iamtoolazytotip/postgres_connection"
	"github.com/derdem/iamtoolazytotip/readTournamentFromDb"
)

func TestReadTournament(t *testing.T) {
	readTournamentFromDb.GetConnection = postgres_connection.GetConnectionForTest
	defer func() {
		readTournamentFromDb.GetConnection = postgres_connection.GetConnection
	}()

	tournament := readTournamentFromDb.ReadTournament(2)
	fmt.Println(tournament)
}

// func _TestPrepareDataDump(t *testing.T) {
// 	readTournamentFromDb.GetConnection = postgres_connection.GetConnectionForTest

// 	tournament := readTournamentFromDb.GetTournament(2)
// 	jsonString, err := json.Marshal(tournament)
// 	if err != nil {
// 		t.Errorf("Error marshalling tournament: %v", err)
// 	}
// 	fmt.Println(string(jsonString))

// 	os.WriteFile("../dumps/tournament2.json", jsonString, 0644)

// }
