package simulator

import (
	"errors"
	"fmt"
	"sort"
)

const (
	NumberOfGroupsInGroupPhase    int = 6
	NumberOfMatchesInFirstKoRound int = 8
)

func TournamentSimulator2(tournament Tournament) Tournament {
	fmt.Println("Start")

	// group matches are played and results are stored in matchResults
	matchResults := playGroupMatches(tournament)
	tournament.MatchResults = matchResults

	// matchResults are evaluated and groupRankings are determined
	groupPhaseGroups := filterByGroupPhase(tournament.Groups)
	teamsSortedIntoGroups := getTeamsSortedIntoGroups(tournament.Teams)
	groupRankings := determineGroupRanking2(matchResults, teamsSortedIntoGroups, groupPhaseGroups)
	tournament.GroupRankings = groupRankings

	// group of thirds is created and added to the tournament
	// This group is a help to setup the round of 16
	groupOfThirds := createGroupOfThirds(tournament.Groups, tournament.Id)
	// update round of 16 matches with the group of thirds
	tournament.Groups = append(tournament.Groups, groupOfThirds)
	// third teams are determined and added to the groupRankings
	rankingOfThirds := determineRankingOfThirds(groupRankings, groupOfThirds.Id)
	tournament.GroupRankings = append(tournament.GroupRankings, rankingOfThirds...)

	// update round of 16 matches with the group of thirds
	UpdateFirstKoRoundWithGroupOfThirds(tournament)

	winner := PlayKoRounds(tournament)
	fmt.Println("Winner of the tournament is", winner.Name)

	return tournament

}

func playGroupMatches(tournament Tournament) []MatchResult {
	groupMatchResults := make([]MatchResult, 0)
	matchResultChannel := make(chan MatchResult, len(tournament.Matches))

	for _, match := range tournament.Matches {
		go playGroupMatch2(match, matchResultChannel)
	}

	for i := 0; i < len(tournament.Matches); i++ {
		groupMatchResults = append(groupMatchResults, <-matchResultChannel)
	}

	return groupMatchResults
}

func playGroupMatch2(match Match2, matchResultChannel chan MatchResult) {
	fmt.Println("Playing match", match.Id)
	outcomeProbabilies := assignProbabilities(match.Team1.Strength, match.Team2.Strength)
	winnerCode := DetermineWinner(outcomeProbabilies)

	var result MatchResult

	switch winnerCode {
	case 0:
		team1Score, team2Score := SetRemisScore()
		result = MatchResult{
			Match:             match,
			Team1Goals:        team1Score,
			Team2Goals:        team2Score,
			Team1PointsGained: 1,
			Team2PointsGained: 1,
		}
	case 1:
		team1Score := RandomResult()
		team2Score := RandomResultLoser(team1Score, match.Team1.Strength-match.Team2.Strength)
		result = MatchResult{
			Match:             match,
			Team1Goals:        team1Score,
			Team2Goals:        team2Score,
			Team1PointsGained: 3,
			Team2PointsGained: 0,
			Winner:            match.Team1,
		}
	case 2:
		team2Score := RandomResult()
		team1Score := RandomResultLoser(team2Score, match.Team1.Strength-match.Team2.Strength)
		result = MatchResult{
			Match:             match,
			Team1Goals:        team1Score,
			Team2Goals:        team2Score,
			Team1PointsGained: 0,
			Team2PointsGained: 3,
			Winner:            match.Team2,
		}
	}

	matchResultChannel <- result
}

func determineGroupRanking2(matchResults []MatchResult, teamsSortedIntoGroups map[int][]Team, groups []Group2) []GroupRanking {
	var teamPoints = make(map[int]int) // teamId -> points
	var teamGoals = make(map[int]int)  // teamId -> goals
	var groupRankings = make([]GroupRanking, 0)

	for _, matchResult := range matchResults {
		teamPoints[matchResult.Match.Team1.Id] += matchResult.Team1PointsGained
		teamPoints[matchResult.Match.Team2.Id] += matchResult.Team2PointsGained

		teamGoals[matchResult.Match.Team1.Id] += matchResult.Team1Goals
		teamGoals[matchResult.Match.Team2.Id] += matchResult.Team2Goals
	}

	for _, group := range groups {
		sortedTeams := determineRankingPerGroup(
			group.Id,
			teamsSortedIntoGroups[group.Id],
			teamPoints,
			teamGoals,
		)
		groupRankings = append(
			groupRankings, sortedTeams...)
	}

	return groupRankings

}

func filterGroup(groupType GroupType) func(groups []Group2) []Group2 {
	return func(groups []Group2) []Group2 {
		filteredGroups := make([]Group2, 0)
		for _, group := range groups {
			if group.GroupType == groupType {
				filteredGroups = append(filteredGroups, group)
			}
		}
		return filteredGroups
	}
}

var filterByGroupPhase = filterGroup(GroupPhaseGroupType)
var filterByKoRound = filterGroup(KoPhaseGroupType)

func getTeamsSortedIntoGroups(teams []Team) map[int][]Team {
	var teamsSortedIntoGroups = make(map[int][]Team) // groupId -> teams
	for _, team := range teams {
		teamsSortedIntoGroups[team.GroupId] = append(teamsSortedIntoGroups[team.GroupId], team)
	}
	return teamsSortedIntoGroups
}

func getRankingsSortedIntoGroups(rankings []GroupRanking) map[int][]GroupRanking {
	var rankingsSortedIntoGroups = make(map[int][]GroupRanking) // groupId -> rankings
	for _, ranking := range rankings {
		rankingsSortedIntoGroups[ranking.GroupId] = append(rankingsSortedIntoGroups[ranking.GroupId], ranking)
	}
	return rankingsSortedIntoGroups
}

func determineRankingPerGroup(groupId int, teams []Team, teamPoints map[int]int, teamGoals map[int]int) []GroupRanking {
	teams = sortTeamsByPointsAndGoals(teams, teamPoints, teamGoals)

	var groupRankings []GroupRanking
	for index, team := range teams {
		groupRankings = append(groupRankings, GroupRanking{
			GroupId: groupId,
			Team:    team,
			Points:  teamPoints[team.Id],
			Goals:   teamGoals[team.Id],
			Ranking: index + 1,
		})
	}

	return groupRankings
}

func sortTeamsByPointsAndGoals(teams []Team, teamPoints map[int]int, teamGoals map[int]int) []Team {
	sort.Slice(teams, func(i, j int) bool {
		// If the points are equal, the team with the most goals scored is ranked higher
		if teamPoints[teams[i].Id] == teamPoints[teams[j].Id] {
			return teamGoals[teams[i].Id] > teamGoals[teams[j].Id]
		}
		// Otherwise, the team with the most points is ranked higher
		return teamPoints[teams[i].Id] > teamPoints[teams[j].Id]
	})

	return teams
}

func determineRankingOfThirds(rankings []GroupRanking, groupId int) []GroupRanking {
	thirds := make([]Team, 0)
	var teamPoints = make(map[int]int) // teamId -> points
	var teamGoals = make(map[int]int)  // teamId -> goals
	for _, ranking := range rankings {
		if ranking.Ranking == 3 {
			third := ranking.Team
			teamPoints[ranking.Team.Id] += ranking.Points
			teamGoals[ranking.Team.Id] += ranking.Goals
			thirds = append(thirds, third)
		}
	}

	return determineRankingPerGroup(groupId, thirds, teamPoints, teamGoals)
}

func createGroupOfThirds(groups []Group2, tournamentId int) Group2 {
	groupOfThirds := Group2{
		Id:           getNextGroupId(groups),
		Name:         "Group of Thirds",
		TournamentId: tournamentId,
	}

	return groupOfThirds
}

func getHighstGroupId(groups []Group2) int {
	highestGroupId := 0
	for _, ranking := range groups {
		if ranking.Id > highestGroupId {
			highestGroupId = ranking.Id
		}
	}
	return highestGroupId

}

func getNextGroupId(groups []Group2) int {
	return getHighstGroupId(groups) + 1
}

func getHighestMatchId(matches []Match2) int {
	highestMatchId := 0
	for _, match := range matches {
		if match.Id > highestMatchId {
			highestMatchId = match.Id
		}
	}
	return highestMatchId
}

func getNextMatchId(matches []Match2) int {
	return getHighestMatchId(matches) + 1
}

func UpdateFirstKoRoundWithGroupOfThirds(tournament Tournament) []KoMatch {
	koMatches := tournament.KoMatches
	groupRankings := tournament.GroupRankings

	groupOfThirds := createGroupOfThirds(tournament.Groups, tournament.Id)
	rankingOfThirds := determineRankingOfThirds(groupRankings, groupOfThirds.Id)

	var firstKoGroupId int
	for _, group := range tournament.Groups {
		if group.GroupType == KoPhaseGroupType {
			firstKoGroupId = group.Id
			break
		}
	}

	firstKoGroupMatches := make([]KoMatch, 0)
	for _, koMatch := range koMatches {
		if koMatch.GroupId == firstKoGroupId {
			firstKoGroupMatches = append(firstKoGroupMatches, koMatch)
		}
	}
	if len(firstKoGroupMatches) != NumberOfMatchesInFirstKoRound {
		panic("Error: Expected " + fmt.Sprint(NumberOfMatchesInFirstKoRound) + " matches in first Ko round, instead found " + fmt.Sprint(len(firstKoGroupMatches)))
	}

	// Get a list of the group-phase Groups
	type GroupId int
	groupPhaseGroupIdsSet := make(map[GroupId]bool, 0)
	for _, groupRanking := range groupRankings {
		if groupRanking.Ranking == 3 {
			groupPhaseGroupIdsSet[GroupId(groupRanking.GroupId)] = true
		}
	}

	groupIds := make([]GroupId, len(groupPhaseGroupIdsSet))
	i := 0
	for k := range groupPhaseGroupIdsSet {
		groupIds[i] = k
		i++
	}

	if len(groupIds) != NumberOfGroupsInGroupPhase {
		panic(
			"Error: Expected " +
				fmt.Sprint(NumberOfGroupsInGroupPhase) +
				" groups in tournament, instead found " +
				fmt.Sprint(len(groupIds)),
		)
	}

	groupIdWeight := map[GroupId]int{}
	weight := 1
	for _, groupId := range groupIds {
		groupIdWeight[groupId] = weight
		weight *= 2
	}

	totalWeight := 0
	for index, rankingOfThird := range rankingOfThirds {
		if index == 5 {
			break
		}
		totalWeight += groupIdWeight[GroupId(rankingOfThird.GroupId)]
	}

	rankingOfThirdPattern := Best4WithRanking3Pattern(weight)

	updatesKoMatches := make([]KoMatch, 0)
	updatedKoMatchIndex := 0
	numberFirstKoGroupMatches := 0
	for _, koMatch := range koMatches {
		if koMatch.GroupId == 0 {
			panic("Error: GroupId1 is unspecified for KoMatch with Id " + fmt.Sprint(koMatch.Id))
		}
		if koMatch.GroupId2 == 0 {
			rankingOfThirdIndex := rankingOfThirdPattern[updatedKoMatchIndex]
			rankingOfThird := rankingOfThirds[rankingOfThirdIndex]
			koMatch.GroupId2 = rankingOfThird.GroupId
			updatedKoMatchIndex++
		}
		if koMatch.GroupId1 == firstKoGroupId {
			numberFirstKoGroupMatches++
		}

		updatesKoMatches = append(updatesKoMatches, koMatch)
	}

	if numberFirstKoGroupMatches != NumberOfMatchesInFirstKoRound {
		panic("Error: Expected " + fmt.Sprint(NumberOfMatchesInFirstKoRound) + " matches in first Ko round, instead found " + fmt.Sprint(numberFirstKoGroupMatches))
	}

	return updatesKoMatches
}

func Best4WithRanking3Pattern(weight int) [4]int {
	pattern := map[int][4]int{
		15: {0, 3, 1, 2}, // X X X X 0  0  => 15
		23: {0, 4, 1, 2}, // X X X 0 X  0  => 23
		39: {0, 5, 1, 2}, // X X X 0 0  X  => 39
		27: {3, 4, 0, 1}, // X X 0 X X  0  => 27
		43: {3, 5, 0, 1}, // X X 0 X 0  X  => 43
		51: {4, 5, 1, 0}, // X X 0 0 X  X  => 51
		29: {4, 3, 2, 0}, // X 0 X X X  0  => 29
		45: {5, 3, 2, 0}, // X 0 X X 0  X  => 45
		53: {4, 5, 2, 0}, // X 0 X 0 X  X  => 53
		57: {4, 5, 3, 0}, // X 0 0 X X  X  => 57
		30: {4, 3, 1, 2}, // 0 X X X X  0  => 30
		46: {5, 3, 2, 1}, // 0 X X X 0  X  => 46
		54: {5, 4, 2, 1}, // 0 X X 0 X  X  => 54
		58: {5, 4, 3, 1}, // 0 X 0 X X  X  => 58
		60: {5, 4, 3, 2}, // 0 0 X X X  X  => 60
	}

	return pattern[weight]
}

func PlayKoRounds(tournament Tournament) Team {
	koMatchMap := MapKoMatchesToGroups(tournament.KoMatches)
	koGroups := filterByKoRound(tournament.Groups)
	for _, koGroup := range koGroups {
		fmt.Println("Playing Ko Group", koGroup.Name)
		koMatches := koMatchMap[koGroup.Id]
		if len(koMatches) == 0 {
			panic("No matches found for Ko Group " + koGroup.Name + " with Id " + fmt.Sprint(koGroup.Id))
		}
		rankingsSortedIntoGroups := getRankingsSortedIntoGroups(tournament.GroupRankings)
		matches := CreateMatchFromKoMatch(koMatches, rankingsSortedIntoGroups)
		matchResults := PlayKoGroupsMatches(matches)
		tournament.MatchResults = append(tournament.MatchResults, matchResults...)

		teamsInGroup := GetTeamsFromMatches(matches)
		teamsMap := make(map[int][]Team)
		teamsMap[koGroup.Id] = teamsInGroup
		groupRankings := determineGroupRanking2(matchResults, teamsMap, []Group2{koGroup})
		tournament.GroupRankings = append(tournament.GroupRankings, groupRankings...)
	}

	finalGroupId := koGroups[len(koGroups)-1].Id
	winner := DetermineWinner2(finalGroupId, tournament.GroupRankings)

	return winner
}

func MapKoMatchesToGroups(koMatches []KoMatch) map[int][]KoMatch {
	var koMatchesMappedToGroups = make(map[int][]KoMatch) // groupId -> koMatches
	for _, koMatch := range koMatches {
		koMatchesMappedToGroups[koMatch.GroupId] = append(koMatchesMappedToGroups[koMatch.GroupId], koMatch)
	}
	return koMatchesMappedToGroups
}

func CreateMatchFromKoMatch(koMatches []KoMatch, rankingsSortedIntoGroups map[int][]GroupRanking) []Match2 {
	matches := make([]Match2, 0)
	for _, koMatch := range koMatches {
		group1GroupRankings := rankingsSortedIntoGroups[koMatch.GroupId1]
		group2GroupRankings := rankingsSortedIntoGroups[koMatch.GroupId2]

		team1, err1 := findTeamBasedOnRanking(group1GroupRankings, koMatch.Ranking1)
		team2, err2 := findTeamBasedOnRanking(group2GroupRankings, koMatch.Ranking2)

		if err1 != nil {
			panic("Error: " + err1.Error() + " for GroupId: " + fmt.Sprint(koMatch.GroupId1))
		}
		if err2 != nil {
			panic("Error: " + err2.Error() + " for GroupId: " + fmt.Sprint(koMatch.GroupId2))
		}

		match := Match2{
			Id:      getNextMatchId(matches),
			Team1:   team1,
			Team2:   team2,
			GroupId: koMatch.GroupId,
		}
		matches = append(matches, match)
	}
	return matches
}

func findTeamBasedOnRanking(rankings []GroupRanking, rankingNumber int) (Team, error) {
	for _, ranking := range rankings {
		if ranking.Ranking == rankingNumber {
			return ranking.Team, nil
		}
	}
	return Team{}, errors.New("Team with ranking " + fmt.Sprint(rankingNumber) + " not found in rankings")
}

func PlayKoGroupsMatches_(matches []Match2) []MatchResult {
	var matchResults []MatchResult
	matchResultChannel := make(chan MatchResult, len(matches))
	numberOfMatches := len(matches)
	pointsForWinner := numberOfMatches // maybe not a good solution as it ignores the match setup defined in the DB, need to check
	for _, match := range matches {
		go PlayEliminationMatch2(match, pointsForWinner, matchResultChannel)
		pointsForWinner--
	}

	for i := 0; i < numberOfMatches; i++ {
		matchResult := <-matchResultChannel
		matchResults = append(matchResults, matchResult)
	}

	return matchResults

}

var PlayKoGroupsMatches = PlayKoGroupsMatches_

func PlayEliminationMatch2_(match Match2, pointsForWinner int, matchResultChannel chan MatchResult) {
	var team1 = match.Team1
	var team2 = match.Team2
	var team1Score int
	var team2Score int
	var team1PenaltyScore int = 0
	var team2PenaltyScore int = 0
	var result MatchResult

	outcomeProbabilies := assignProbabilities(team1.Strength, team2.Strength)
	winnerCode := DetermineWinner(outcomeProbabilies)

	fmt.Println(team1.Name + " vs. " + team2.Name)
	switch winnerCode {
	case 0:
		result = ResolveDrawInEliminationMatch(match, pointsForWinner)
		resultString := fmt.Sprintf("%d (%d) - %d (%d)", team1Score, team1Score+team1PenaltyScore, team2Score, team1Score+team2PenaltyScore)
		fmt.Println(resultString)

	case 1:
		team1Score = RandomResult()
		team2Score = RandomResultLoser(team1Score, team1.Strength-team2.Strength)
		result = MatchResult{
			Match:             match,
			Team1Goals:        team1Score,
			Team2Goals:        team2Score,
			Team1PointsGained: pointsForWinner,
			Team2PointsGained: 0,
			Winner:            team1,
		}
		resultString := fmt.Sprintf("%d - %d", team1Score, team2Score)
		fmt.Println(resultString)
	case 2:
		team2Score = RandomResult()
		team1Score = RandomResultLoser(team2Score, team2.Strength-team1.Strength)
		result = MatchResult{
			Match:             match,
			Team1Goals:        team1Score,
			Team2Goals:        team2Score,
			Team1PointsGained: 0,
			Team2PointsGained: pointsForWinner,
			Winner:            team2,
		}
		resultString := fmt.Sprintf("%d - %d", team1Score, team2Score)
		fmt.Println(resultString)
	}

	matchResultChannel <- result
}

var PlayEliminationMatch2 = PlayEliminationMatch2_

func ResolveDrawInEliminationMatch_(match Match2, pointsForWinner int) MatchResult {
	team1 := match.Team1
	team2 := match.Team2
	var winnerTeam Team
	var team1PointsGained int
	var team2PointsGained int

	team1Score, team2Score := SetRemisScore()
	team1PenaltyScore, team2PenaltyScore := PlayPenalty(0, 0)
	if team1Score+team1PenaltyScore > team2Score+team2PenaltyScore {
		winnerTeam = team1
		team1PointsGained = pointsForWinner
		team2PointsGained = 0
	} else {
		winnerTeam = team2
		team1PointsGained = 0
		team2PointsGained = pointsForWinner
	}
	result := MatchResult{
		Match:             match,
		Team1Goals:        team1Score,
		Team2Goals:        team2Score,
		Team1PenaltyGoals: team1PenaltyScore,
		Team2PenaltyGoals: team2PenaltyScore,
		Team1PointsGained: team1PointsGained,
		Team2PointsGained: team2PointsGained,
		Winner:            winnerTeam,
	}

	return result
}

var ResolveDrawInEliminationMatch = ResolveDrawInEliminationMatch_

func GetTeamsFromMatches(matches []Match2) []Team {
	teams := make([]Team, 0)
	for _, match := range matches {
		teams = append(teams, match.Team1, match.Team2)
	}
	return teams
}

func DetermineWinner2(finalGroupId int, groupRankings []GroupRanking) Team {
	rankingsSortedIntoGroups := getRankingsSortedIntoGroups(groupRankings)
	finalRankings := rankingsSortedIntoGroups[finalGroupId]

	if len(finalRankings) != 2 {
		panic("Error: Final group does not have 2 teams, instead has " + fmt.Sprint(len(finalRankings)) + " teams.")
	}
	winner := finalRankings[0].Team

	return winner
}
