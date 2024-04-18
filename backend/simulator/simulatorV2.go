package simulator

import (
	"fmt"
	"sort"
)

func TournamentSimulator2(tournament Tournament) {
	fmt.Println("Start")

	// group matches are played and results are stored in matchResults
	matchResults := playGroupMatches(tournament)
	tournament.MatchResults = matchResults

	// matchResults are evaluated and groupRankings are determined
	groupRankings := determineGroupRanking2(matchResults, tournament)
	tournament.GroupRankings = groupRankings

	// group of thirds is created and added to the tournament
	// This group is a help to setup the round of 16
	groupOfThirds := createGroupOfThirds(tournament.Groups, tournament.Id)
	tournament.Groups = append(tournament.Groups, groupOfThirds)

	// third teams are determined and added to the groupRankings
	rankingOfThirds := determineRankingOfThirds(groupRankings, groupOfThirds.Id)
	tournament.GroupRankings = append(tournament.GroupRankings, rankingOfThirds...)

	winner := playKoRounds(tournament)
	fmt.Println("Winner of the tournament is", winner.Name)

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
	winnerCode := determineWinner(outcomeProbabilies)

	var result MatchResult

	switch winnerCode {
	case 0:
		team1Score, team2Score := setRemisScore()
		result = MatchResult{
			Match:             match,
			Team1Goals:        team1Score,
			Team2Goals:        team2Score,
			Team1PointsGained: 1,
			Team2PointsGained: 1,
		}
	case 1:
		team1Score := randomResult()
		team2Score := randomResultLoser(team1Score, match.Team1.Strength-match.Team2.Strength)
		result = MatchResult{
			Match:             match,
			Team1Goals:        team1Score,
			Team2Goals:        team2Score,
			Team1PointsGained: 3,
			Team2PointsGained: 0,
			Winner:            match.Team1,
		}
	case 2:
		team2Score := randomResult()
		team1Score := randomResultLoser(team2Score, match.Team1.Strength-match.Team2.Strength)
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

func determineGroupRanking2(matchResults []MatchResult, tournament Tournament) []GroupRanking {
	var teamPoints = make(map[int]int) // teamId -> points
	var teamGoals = make(map[int]int)  // teamId -> goals
	var groupRankings = make([]GroupRanking, 0)

	for _, matchResult := range matchResults {
		teamPoints[matchResult.Match.Team1.Id] += matchResult.Team1PointsGained
		teamPoints[matchResult.Match.Team2.Id] += matchResult.Team2PointsGained

		teamGoals[matchResult.Match.Team1.Id] += matchResult.Team1Goals
		teamGoals[matchResult.Match.Team2.Id] += matchResult.Team2Goals
	}

	teamsSortedIntoGroups := getTeamsSortedIntoGroups(tournament.Teams)

	groupPhaseGroups := filterByGroupPhase(tournament.Groups)
	for _, group := range groupPhaseGroups {
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

func playKoRounds(tournament Tournament) Team {
	koMatchMap := mapKoMatchesToGroups(tournament.KoMatches)
	koGroups := filterByKoRound(tournament.Groups)
	for _, koGroup := range koGroups {
		koMatches := koMatchMap[koGroup.Id]
		matchResults := playKoGroupsMatches(koGroup, koMatches, tournament.GroupRankings)
		tournament.MatchResults = append(tournament.MatchResults, matchResults...)

		groupRankings := determineGroupRanking2(matchResults, tournament)
		tournament.GroupRankings = append(tournament.GroupRankings, groupRankings...)
	}

	finalGroupId := koGroups[len(koGroups)-1].Id
	winner := determineWinner2(finalGroupId, tournament.GroupRankings)

	return winner
}

func playKoGroupsMatches(koGroup Group2, koMatches []KoMatch, groupRankings []GroupRanking) []MatchResult {
	fmt.Println("Playing Ko Group", koGroup.Name)
	var matchResults []MatchResult
	matches := createMatchFromKoMatch(koMatches, groupRankings)
	matchResultChannel := make(chan MatchResult, len(matches))
	numberOfMatches := len(matches)
	pointsForWinner := numberOfMatches // maybe not a good solution as it ignores the match setup defined in the DB, need to check
	for _, match := range matches {
		go playEliminationMatch2(match, pointsForWinner, matchResultChannel)
		pointsForWinner--
	}

	for i := 0; i < numberOfMatches; i++ {
		matchResult := <-matchResultChannel
		matchResults = append(matchResults, matchResult)
	}

	return matchResults

}

func mapKoMatchesToGroups(koMatches []KoMatch) map[int][]KoMatch {
	var koMatchesMappedToGroups = make(map[int][]KoMatch) // groupId -> koMatches
	for _, koMatch := range koMatches {
		koMatchesMappedToGroups[koMatch.GroupId] = append(koMatchesMappedToGroups[koMatch.GroupId], koMatch)
	}
	return koMatchesMappedToGroups
}

func createMatchFromKoMatch(koMatches []KoMatch, groupRankings []GroupRanking) []Match2 {
	matches := make([]Match2, 0)
	rankingsSortedIntoGroups := getRankingsSortedIntoGroups(groupRankings)
	for _, koMatch := range koMatches {
		match := Match2{
			Id:    getNextMatchId(matches),
			Team1: rankingsSortedIntoGroups[koMatch.Group1.Id][koMatch.ranking1-1].Team,
			Team2: rankingsSortedIntoGroups[koMatch.Group2.Id][koMatch.ranking2-1].Team,
		}
		matches = append(matches, match)
	}
	return matches

}

func playEliminationMatch2(match Match2, pointsForWinner int, matchResultChannel chan MatchResult) {
	var team1 = match.Team1
	var team2 = match.Team2
	var team1Score int
	var team2Score int
	var team1PenaltyScore int = 0
	var team2PenaltyScore int = 0
	var team1PointsGained int
	var team2PointsGained int
	var winnerTeam Team
	var result MatchResult

	outcomeProbabilies := assignProbabilities(team1.Strength, team2.Strength)
	winnerCode := determineWinner(outcomeProbabilies)

	fmt.Println(team1.Name + " vs. " + team2.Name)
	switch winnerCode {
	case 0:
		team1Score, team2Score = setRemisScore()
		team1PenaltyScore, team2PenaltyScore = playPenalty(0, 0)
		if team1Score+team1PenaltyScore > team2Score+team2PenaltyScore {
			winnerTeam = team1
			team1PointsGained = pointsForWinner
			team2PointsGained = 0
		} else {
			winnerTeam = team2
			team1PointsGained = 0
			team2PointsGained = pointsForWinner
		}
		result = MatchResult{
			Match:             match,
			Team1Goals:        team1Score,
			Team2Goals:        team2Score,
			Team1PenaltyGoals: team1PenaltyScore,
			Team2PenaltyGoals: team2PenaltyScore,
			Team1PointsGained: team1PointsGained,
			Team2PointsGained: team2PointsGained,
			Winner:            winnerTeam,
		}
		resultString := fmt.Sprintf("%d (%d) - %d (%d)", team1Score, team1Score+team1PenaltyScore, team2Score, team1Score+team2PenaltyScore)
		fmt.Println(resultString)

	case 1:
		team1Score = randomResult()
		team2Score = randomResultLoser(team1Score, team1.Strength-team2.Strength)
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
		team2Score = randomResult()
		team1Score = randomResultLoser(team2Score, team2.Strength-team1.Strength)
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

func determineWinner2(finalGroupId int, groupRankings []GroupRanking) Team {
	rankingsSortedIntoGroups := getRankingsSortedIntoGroups(groupRankings)
	finalRankings := rankingsSortedIntoGroups[finalGroupId]
	winner := finalRankings[0].Team

	if len(finalRankings) != 2 {
		fmt.Println("Error: Final group does not have 2 teams")
	}

	return winner
}
