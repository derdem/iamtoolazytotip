package simulator_test

import (
	"testing"

	"github.com/derdem/iamtoolazytotip/postgres_connection"
	"github.com/derdem/iamtoolazytotip/simulator"
	"github.com/derdem/iamtoolazytotip/simulator/readTournamentFromDb"
)

func TestTournamentSimulatorV2(t *testing.T) {
	readTournamentFromDb.GetConnection = postgres_connection.GetConnectionForTest
	defer func() {
		readTournamentFromDb.GetConnection = postgres_connection.GetConnection
	}()

	tournament := readTournamentFromDb.GetTournament(2)
	simulator.TournamentSimulator2(tournament)
}
