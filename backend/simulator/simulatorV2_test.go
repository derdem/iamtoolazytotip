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

func recoverFromPanic(t *testing.T, message string) {
	if message == "" {
		message = "The code did not panic"
	}
	if r := recover(); r == nil {
		t.Errorf(message)
	}
}

func _TestTournamentSimulatorV2(t *testing.T) {
	readTournamentFromDb.GetConnection = postgres_connection.GetConnectionForTest
	defer func() {
		readTournamentFromDb.GetConnection = postgres_connection.GetConnection
	}()

	tournament := readTournamentFromDb.GetTournament(2)
	simulator.TournamentSimulator2(tournament)
}

func _TestPrepareDataDump(t *testing.T) {
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

	if len(tournament.GroupRankings) != 54 {
		t.Errorf("Expected 30 group rankings, got %v", len(tournament.GroupRankings))
	}

}

func TestUpdateKoMatchesWithThirds(t *testing.T) {
	tournament := simulator.ReadTournamentFromFile("../dumps/tournament2.json")

	matchResults := simulator.PlayGroupMatches(tournament)
	tournament.MatchResults = matchResults

	// matchResults are evaluated and groupRankings are determined
	groupPhaseGroups := simulator.FilterByGroupPhase(tournament.Groups)
	teamsSortedIntoGroups := simulator.GetTeamsSortedIntoGroups(tournament.Teams)
	groupRankings := simulator.DetermineGroupRanking2(matchResults, teamsSortedIntoGroups, groupPhaseGroups)
	tournament.GroupRankings = groupRankings

	// Ko matches are updated with the third placed teams
	updatedKoMatches := simulator.UpdateKoMatchesWithThirds(tournament)

	koMatchesWithThirds := make([]simulator.KoMatch, 0)
	for _, match := range updatedKoMatches {
		if match.Ranking2 == 3 {
			koMatchesWithThirds = append(koMatchesWithThirds, match)
		}
	}

	if len(koMatchesWithThirds) != 4 {
		t.Errorf("Expected 4 ko matches with thirds, got %v", len(koMatchesWithThirds))
	}
}

func TestPlayKoRounds_HappyCase(t *testing.T) {
	tournamentId := 1
	groupRankings := []simulator.GroupRanking{
		{
			GroupId: 1,
			Team: simulator.Team{
				Name: "Team1",
			},
			Ranking: 1,
			Points:  3,
			Goals:   3,
		},
		{
			GroupId: 2,
			Team: simulator.Team{
				Name: "Team2",
			},
			Ranking: 1,
			Points:  3,
			Goals:   3,
		},
	}
	groups := []simulator.Group2{
		{
			Id:           3,
			Name:         "Group1",
			TournamentId: tournamentId,
			GroupType:    simulator.KoPhaseGroupType,
		},
	}

	koMatches := []simulator.KoMatch{
		{
			Id:       1,
			GroupId:  3,
			GroupId1: 1,
			GroupId2: 2,
			Ranking1: 1,
			Ranking2: 1,
		},
	}

	tournament := simulator.Tournament{
		Id:            tournamentId,
		Name:          "Test",
		GroupRankings: groupRankings,
		Groups:        groups,
		KoMatches:     koMatches,
	}

	finishedTournament := simulator.PlayKoRounds(tournament)
	winner := finishedTournament.Winner

	if winner.Name != "Team1" && winner.Name != "Team2" {
		t.Errorf("Expected Team1 or Team2, got %v", winner.Name)
	}
}

func TestPlayKoRounds_UnhappyCase_NoKoMatches(t *testing.T) {
	defer recoverFromPanic(t, "")

	tournamentId := 1
	groups := []simulator.Group2{
		{
			Id:           1,
			Name:         "Group1",
			TournamentId: tournamentId,
			GroupType:    simulator.KoPhaseGroupType,
		},
	}

	tournament := simulator.Tournament{
		Id:     tournamentId,
		Name:   "Test",
		Groups: groups,
	}

	simulator.PlayKoRounds(tournament)
}

func TestPlayKoRounds_UnhappyCase_WrongRankingGroupId(t *testing.T) {
	defer recoverFromPanic(t, "")

	tournamentId := 1
	groupRankings := []simulator.GroupRanking{
		{
			GroupId: 1,
			Team: simulator.Team{
				Name: "Team1",
			},
			Ranking: 1,
			Points:  3,
			Goals:   3,
		},
	}
	groups := []simulator.Group2{
		{
			Id:           3,
			Name:         "Group1",
			TournamentId: tournamentId,
			GroupType:    simulator.KoPhaseGroupType,
		},
	}

	koMatches := []simulator.KoMatch{
		{
			Id:       1,
			GroupId:  3,
			GroupId1: 1,
			GroupId2: 2,
			Ranking1: 1,
			Ranking2: 1,
		},
	}

	tournament := simulator.Tournament{
		Id:            tournamentId,
		Name:          "Test",
		GroupRankings: groupRankings,
		Groups:        groups,
		KoMatches:     koMatches,
	}

	simulator.PlayKoRounds(tournament)
}

func TestPlayKoRounds_UnhappyCase_WrongRankingNumber(t *testing.T) {
	defer recoverFromPanic(t, "")

	tournamentId := 1
	groupRankings := []simulator.GroupRanking{
		{
			GroupId: 1,
			Team: simulator.Team{
				Name: "Team1",
			},
			Ranking: 2, // Wrong ranking
			Points:  3,
			Goals:   3,
		},
		{
			GroupId: 2,
			Team: simulator.Team{
				Name: "Team2",
			},
			Ranking: 1,
			Points:  3,
			Goals:   3,
		},
	}
	groups := []simulator.Group2{
		{
			Id:           3,
			Name:         "Group1",
			TournamentId: tournamentId,
			GroupType:    simulator.KoPhaseGroupType,
		},
	}

	koMatches := []simulator.KoMatch{
		{
			Id:       1,
			GroupId:  3,
			GroupId1: 1,
			GroupId2: 2,
			Ranking1: 1, // Wrong ranking
			Ranking2: 1,
		},
	}

	tournament := simulator.Tournament{
		Id:            tournamentId,
		Name:          "Test",
		GroupRankings: groupRankings,
		Groups:        groups,
		KoMatches:     koMatches,
	}

	simulator.PlayKoRounds(tournament)
}

func TestMapKoMatchesToGroups(t *testing.T) {
	koMatches := getKoMatches()
	koMatchesMap := simulator.MapKoMatchesToGroups(koMatches)

	if len(koMatchesMap[3]) != 2 {
		t.Errorf("Expected 2, got %v", len(koMatchesMap[3]))
	}
	if len(koMatchesMap[4]) != 1 {
		t.Errorf("Expected 1, got %v", len(koMatchesMap[4]))
	}
	if len(koMatchesMap[5]) != 0 {
		t.Errorf("Expected 0, got %v", len(koMatchesMap[5]))
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
	defer recoverFromPanic(t, "")

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

func TestPlayKoGroupsMatches(t *testing.T) {
	matches := getMatches()

	simulator.PlayEliminationMatch2 = func(match simulator.Match2, pointsForWinner int, matchChannel chan simulator.MatchResult) {
		matchChannel <- simulator.MatchResult{
			Winner:            match.Team1,
			Team1Goals:        1,
			Team2Goals:        0,
			Team1PenaltyGoals: 0,
			Team2PenaltyGoals: 0,
			Team1PointsGained: pointsForWinner,
			Team2PointsGained: 0,
			Match:             match,
		}
	}
	defer func() {
		simulator.PlayEliminationMatch2 = simulator.PlayEliminationMatch2_
	}()

	results := simulator.PlayKoGroupsMatches(matches)

	resultsLength := len(results)
	if resultsLength != 2 {
		t.Errorf("Expected 2 results, got %v", len(results))
	}

	for _, result := range results {
		if result.Team1Goals != 1 {
			t.Errorf("Expected 1, got %v", result.Team1Goals)
		}
		if result.Team2Goals != 0 {
			t.Errorf("Expected 0, got %v", result.Team2Goals)
		}
		if result.Team1PointsGained > resultsLength || result.Team1PointsGained <= 0 {
			t.Errorf("Expected Points between 0 and %v, got %v", resultsLength, result.Team1PointsGained)
		}
		if result.Team2PointsGained != 0 {
			t.Errorf("Expected 0, got %v", result.Team2PointsGained)
		}
	}

}

func TestPlayEliminationMatch2_Penalty(t *testing.T) {
	match := getMatch()
	matchChannel := make(chan simulator.MatchResult)
	matchResult := simulator.MatchResult{
		Winner:            match.Team1,
		Team1Goals:        1,
		Team2Goals:        1,
		Team1PenaltyGoals: 1,
		Team2PenaltyGoals: 0,
		Team1PointsGained: 1,
		Team2PointsGained: 0,
	}

	simulator.DetermineWinner = func(a simulator.OutcomeProbabilities) int {
		return 0
	}
	defer func() {
		simulator.DetermineWinner = simulator.DetermineWinner_
	}()
	simulator.ResolveDrawInEliminationMatch = func(match simulator.Match2, pointsForWinner int) simulator.MatchResult {
		return matchResult
	}
	defer func() {
		simulator.ResolveDrawInEliminationMatch = simulator.ResolveDrawInEliminationMatch_
	}()

	go simulator.PlayEliminationMatch2(match, 1, matchChannel)

	result := <-matchChannel

	if result.Winner.Name != "Team1" {
		t.Errorf("Expected Team1, got %v", result.Winner.Name)
	}
	if result.Team1Goals != 1 {
		t.Errorf("Expected 1, got %v", result.Team1Goals)
	}
	if result.Team1PenaltyGoals != 1 {
		t.Errorf("Expected 1, got %v", result.Team1PenaltyGoals)
	}
	if result.Team1PointsGained != 1 {
		t.Errorf("Expected 1, got %v", result.Team1PointsGained)
	}
	if result.Team2Goals != 1 {
		t.Errorf("Expected 1, got %v", result.Team2Goals)
	}
	if result.Team2PenaltyGoals != 0 {
		t.Errorf("Expected 0, got %v", result.Team2PenaltyGoals)
	}
	if result.Team2PointsGained != 0 {
		t.Errorf("Expected 0, got %v", result.Team2PointsGained)
	}

}

func TestPlayEliminationMatch2_Team1Wins(t *testing.T) {
	match := getMatch()
	matchChannel := make(chan simulator.MatchResult)

	simulator.DetermineWinner = func(a simulator.OutcomeProbabilities) int {
		return 1
	}
	defer func() {
		simulator.DetermineWinner = simulator.DetermineWinner_
	}()
	simulator.RandomResult = func() int {
		return 1
	}
	defer func() {
		simulator.RandomResult = simulator.RandomResult_
	}()
	simulator.RandomResultLoser = func(rw int, sd int) int {
		return 0
	}
	defer func() {
		simulator.RandomResultLoser = simulator.RandomResultLoser_
	}()

	go simulator.PlayEliminationMatch2(match, 1, matchChannel)

	result := <-matchChannel

	if result.Winner.Name != "Team1" {
		t.Errorf("Expected Team1, got %v", result.Winner.Name)
	}
	if result.Team1Goals != 1 {
		t.Errorf("Expected 1, got %v", result.Team1Goals)
	}
	if result.Team1PenaltyGoals != 0 {
		t.Errorf("Expected 0, got %v", result.Team1PenaltyGoals)
	}
	if result.Team1PointsGained != 1 {
		t.Errorf("Expected 1, got %v", result.Team2Goals)
	}

	if result.Team2Goals != 0 {
		t.Errorf("Expected 0, got %v", result.Team2Goals)
	}
	if result.Team2PenaltyGoals != 0 {
		t.Errorf("Expected 0, got %v", result.Team2PenaltyGoals)
	}
	if result.Team2PointsGained != 0 {
		t.Errorf("Expected 0, got %v", result.Team2PointsGained)
	}

}

func TestPlayEliminationMatch2_Team2Wins(t *testing.T) {
	match := getMatch()
	matchChannel := make(chan simulator.MatchResult)

	simulator.DetermineWinner = func(a simulator.OutcomeProbabilities) int {
		return 2
	}
	defer func() {
		simulator.DetermineWinner = simulator.DetermineWinner_
	}()
	simulator.RandomResult = func() int {
		return 1
	}
	defer func() {
		simulator.RandomResult = simulator.RandomResult_
	}()
	simulator.RandomResultLoser = func(rw int, sd int) int {
		return 0
	}
	defer func() {
		simulator.RandomResultLoser = simulator.RandomResultLoser_
	}()

	go simulator.PlayEliminationMatch2(match, 1, matchChannel)

	result := <-matchChannel

	if result.Winner.Name != "Team2" {
		t.Errorf("Expected Team2, got %v", result.Winner.Name)
	}
	if result.Team1Goals != 0 {
		t.Errorf("Expected 0, got %v", result.Team1Goals)
	}
	if result.Team1PenaltyGoals != 0 {
		t.Errorf("Expected 0, got %v", result.Team1PenaltyGoals)
	}
	if result.Team1Goals != 0 {
		t.Errorf("Expected 0, got %v", result.Team2Goals)
	}

	if result.Team2Goals != 1 {
		t.Errorf("Expected 1, got %v", result.Team2Goals)
	}
	if result.Team2PenaltyGoals != 0 {
		t.Errorf("Expected 0, got %v", result.Team2PenaltyGoals)
	}
	if result.Team2PointsGained != 1 {
		t.Errorf("Expected 1, got %v", result.Team2PointsGained)
	}
}

func TestResolveDrawInEliminationMatch_Team1Wins(t *testing.T) {
	match := getMatch()

	simulator.SetRemisScore = func() (int, int) {
		return 1, 1
	}
	defer func() {
		simulator.SetRemisScore = simulator.SetRemisScore_
	}()
	simulator.PlayPenalty = func(s1 int, s2 int) (int, int) {
		return 1, 0
	}
	defer func() {
		simulator.PlayPenalty = simulator.PlayPenalty_
	}()

	result := simulator.ResolveDrawInEliminationMatch(match, 1)

	if result.Winner.Name != "Team1" {
		t.Errorf("Expected Team1, got %v", result.Winner.Name)
	}

	if result.Team1Goals != 1 {
		t.Errorf("Expected 1, got %v", result.Team1Goals)
	}
	if result.Team1PenaltyGoals != 1 {
		t.Errorf("Expected 1, got %v", result.Team1PenaltyGoals)
	}
	if result.Team1PointsGained != 1 {
		t.Errorf("Expected 1, got %v", result.Team1PointsGained)
	}

	if result.Team2Goals != 1 {
		t.Errorf("Expected 1, got %v", result.Team2Goals)
	}
	if result.Team2PenaltyGoals != 0 {
		t.Errorf("Expected 0, got %v", result.Team2PenaltyGoals)
	}
	if result.Team2PointsGained != 0 {
		t.Errorf("Expected 0, got %v", result.Team2PointsGained)
	}
}

func TestResolveDrawInEliminationMatch_Team2Wins(t *testing.T) {
	match := getMatch()

	simulator.SetRemisScore = func() (int, int) {
		return 1, 1
	}
	defer func() {
		simulator.SetRemisScore = simulator.SetRemisScore_
	}()
	simulator.PlayPenalty = func(s1 int, s2 int) (int, int) {
		return 0, 1
	}
	defer func() {
		simulator.PlayPenalty = simulator.PlayPenalty_
	}()

	result := simulator.ResolveDrawInEliminationMatch(match, 1)

	if result.Winner.Name != "Team2" {
		t.Errorf("Expected Team1, got %v", result.Winner.Name)
	}

	if result.Team1Goals != 1 {
		t.Errorf("Expected 1, got %v", result.Team1Goals)
	}
	if result.Team1PenaltyGoals != 0 {
		t.Errorf("Expected 0, got %v", result.Team1PenaltyGoals)
	}
	if result.Team1PointsGained != 0 {
		t.Errorf("Expected 0, got %v", result.Team1PointsGained)
	}

	if result.Team2Goals != 1 {
		t.Errorf("Expected 1, got %v", result.Team2Goals)
	}
	if result.Team2PenaltyGoals != 1 {
		t.Errorf("Expected 1, got %v", result.Team2PenaltyGoals)
	}
	if result.Team2PointsGained != 1 {
		t.Errorf("Expected 1, got %v", result.Team2PointsGained)
	}

}

func TestGetTeamsFromMatches(t *testing.T) {
	matches := getMatches()

	teams := simulator.GetTeamsFromMatches(matches)

	if len(teams) != 4 {
		t.Errorf("Expected 4 teams, got %v", len(teams))
	}

}

func getMatch() simulator.Match2 {
	return simulator.Match2{
		Id:      1,
		GroupId: 1,
		Team1:   simulator.Team{Name: "Team1"},
		Team2:   simulator.Team{Name: "Team2"},
	}
}

func getMatches() []simulator.Match2 {
	return []simulator.Match2{
		{
			Id:      1,
			GroupId: 1,
			Team1:   simulator.Team{Name: "Team1"},
			Team2:   simulator.Team{Name: "Team2"},
		},
		{
			Id:      2,
			GroupId: 1,
			Team1:   simulator.Team{Name: "Team3"},
			Team2:   simulator.Team{Name: "Team4"},
		},
	}
}

func getKoMatches() []simulator.KoMatch {
	return []simulator.KoMatch{
		{
			Id:       1,
			GroupId:  3,
			GroupId1: 1,
			GroupId2: 2,
			Ranking1: 1,
			Ranking2: 1,
		},
		{
			Id:       2,
			GroupId:  3,
			GroupId1: 3,
			GroupId2: 4,
			Ranking1: 1,
			Ranking2: 1,
		},
		{
			Id:       3,
			GroupId:  4,
			GroupId1: 5,
			GroupId2: 6,
			Ranking1: 1,
			Ranking2: 1,
		},
	}
}

func TestDetermineWinnerHappyCase(t *testing.T) {
	finalGroupId := 3
	groupRankings := getGroupRankings_DetermineWinnerHappyCase()
	winner := simulator.DetermineWinner2(finalGroupId, groupRankings)

	if winner.Name != groupRankings[0].Team.Name {
		t.Errorf("Expected %v, got %v", groupRankings[0].Team.Name, winner.Name)
	}
}

func getGroupRankings_DetermineWinnerHappyCase() []simulator.GroupRanking {
	return []simulator.GroupRanking{
		{
			GroupId: 3,
			Team:    simulator.Team{Name: "Team1"},
			Ranking: 1,
			Points:  3,
			Goals:   3,
		},
		{
			GroupId: 3,
			Team:    simulator.Team{Name: "Team2"},
			Ranking: 2,
			Points:  3,
			Goals:   2,
		},
	}
}

func TestDetermineWinnerUnhappyCase(t *testing.T) {
	defer recoverFromPanic(t, "")

	finalGroupId := 3
	groupRankings := getGroupRankings_DetermineWinnerUnhappyCase()

	_ = simulator.DetermineWinner2(finalGroupId, groupRankings)
}

func getGroupRankings_DetermineWinnerUnhappyCase() []simulator.GroupRanking {
	return []simulator.GroupRanking{
		{
			GroupId: 3,
			Team:    simulator.Team{Name: "Team1"},
			Ranking: 1,
			Points:  3,
			Goals:   3,
		},
	}
}
