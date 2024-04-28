package simulator_test

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
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

func TestPrepareDataDump(t *testing.T) {
	readTournamentFromDb.GetConnection = postgres_connection.GetConnectionForTest

	tournament := readTournamentFromDb.GetTournament(2)
	jsonString, err := json.Marshal(tournament)
	if err != nil {
		t.Errorf("Error marshalling tournament: %v", err)
	}
	fmt.Println(string(jsonString))

	os.WriteFile("../dumps/tournament2.json", jsonString, 0644)

}

func TestReadTournamentFromJson(t *testing.T) {
	jsonFile, err := os.Open("../dumps/tournament2.json")

	if err != nil {
		t.Errorf("Error opening file: %v", err)
	}

	defer jsonFile.Close()
	jsonBytes, err := io.ReadAll(jsonFile)

	if err != nil {
		t.Errorf("Error reading file: %v", err)
	}
	var Tournament simulator.Tournament
	json.Unmarshal(jsonBytes, &Tournament)

	fmt.Println(Tournament)

}

func TestRunFullTournament(t *testing.T) {
	tournament := simulator.ReadTournamentFromFile("../dumps/tournament2.json")

	tournament = simulator.TournamentSimulator2(tournament)

	if len(tournament.GroupRankings) != 30 {
		t.Errorf("Expected 30 group rankings, got %v", tournament.GroupRankings)
	}
}

func TestCreateMatchFromKoMatchHappyCase(t *testing.T) {
	koMatch := []simulator.KoMatch{{
		Id:       1,
		GroupId:  3,
		GroupId1: 1,
		GroupId2: 2,
		Ranking1: 1,
		Ranking2: 1,
	}}

	rankingsSortedIntoGroups := make(map[int][]simulator.GroupRanking)
	rankingsSortedIntoGroups[1] = []simulator.GroupRanking{
		{
			GroupId: 1,
			Team:    simulator.Team{Name: "Team1"},
			Ranking: 1,
			Points:  3,
			Goals:   3,
		},
	}
	rankingsSortedIntoGroups[2] = []simulator.GroupRanking{
		{
			GroupId: 2,
			Team:    simulator.Team{Name: "Team2"},
			Ranking: 1,
			Points:  3,
			Goals:   3,
		},
	}

	matches := simulator.CreateMatchFromKoMatch(koMatch, rankingsSortedIntoGroups)

	numberOfMatches := len(matches)
	if numberOfMatches != 1 {
		t.Errorf("Expected number of matches to be 1, got %v", numberOfMatches)
	}

	match := matches[0]
	if match.Team1.Name != rankingsSortedIntoGroups[1][0].Team.Name {
		t.Errorf("Expected %v, got %v", rankingsSortedIntoGroups[1][0].Team.Name, match.Team1.Name)
	}

	if match.Team2.Name != rankingsSortedIntoGroups[2][0].Team.Name {
		t.Errorf("Expected %v, got %v", rankingsSortedIntoGroups[2][0].Team.Name, match.Team2.Name)
	}
}

func TestCreateMatchFromKoMatchUnhappyCase(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	koMatch := []simulator.KoMatch{{
		Id:       1,
		GroupId:  3,
		GroupId1: 1,
		GroupId2: 2,
		Ranking1: 1,
		Ranking2: 1,
	}}

	rankingsSortedIntoGroups := make(map[int][]simulator.GroupRanking)
	_ = simulator.CreateMatchFromKoMatch(koMatch, rankingsSortedIntoGroups)
}
