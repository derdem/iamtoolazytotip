package readTournamentFromDb_test

import (
	"fmt"
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
