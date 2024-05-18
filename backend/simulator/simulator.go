package simulator

import (
	"errors"
	"fmt"
	"sort"
)

const (
	NumberOfGroupsInGroupPhase    int     = 6
	NumberOfMatchesInFirstKoRound int     = 8
	lambda                        float64 = 1.3
)

func TournamentSimulator(tournament Tournament) Tournament {
	fmt.Println("Start")
	tournamentAfterGroupPhase := PlayGroupRounds(tournament)
	tournamentWithWinner := PlayKoRounds(tournamentAfterGroupPhase)
	fmt.Println("Winner of the tournament is", tournamentWithWinner.Winner.Name)

	return tournamentWithWinner
}

func PlayGroupRounds(tournament Tournament) Tournament {
	matchResults := PlayGroupMatches(tournament)
	tournament.MatchResults = matchResults

	// matchResults are evaluated and groupRankings are determined
	groupPhaseGroups := FilterByGroupPhase(tournament.Groups)
	teamsSortedIntoGroups := GroupsMapTeams(tournament.Teams)
	groupRankings := DetermineGroupRanking(matchResults, teamsSortedIntoGroups, groupPhaseGroups)
	tournament.GroupRankings = groupRankings

	// update ko matches with the group of thirds
	updatedKoMatches := UpdateKoMatchesWithThirds(tournament)

	tournament.KoMatches = updatedKoMatches

	return tournament
}

func PlayGroupMatches(tournament Tournament) []MatchResult {
	groupMatchResults := make([]MatchResult, 0)
	matchResultChannel := make(chan MatchResult, len(tournament.Matches))

	for _, match := range tournament.Matches {
		go PlayGroupMatch(match, matchResultChannel)
	}

	for i := 0; i < len(tournament.Matches); i++ {
		groupMatchResults = append(groupMatchResults, <-matchResultChannel)
	}

	return groupMatchResults
}

func DetermineGroupRanking(matchResults []MatchResult, teamsSortedIntoGroups map[int][]Team, groups []Group) []GroupRanking {
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

func sortRankingsByPointsAndGoals(rankings []GroupRanking) []GroupRanking {
	sortedRankings := make([]GroupRanking, len(rankings))
	copy(sortedRankings, rankings)
	sort.Slice(sortedRankings, func(i, j int) bool {
		if sortedRankings[i].Points == sortedRankings[j].Points {
			return sortedRankings[i].Goals > sortedRankings[j].Goals
		}
		return sortedRankings[i].Points > sortedRankings[j].Points
	})
	return sortedRankings
}

func getHighestMatchId(matches []Match) int {
	highestMatchId := 0
	for _, match := range matches {
		if match.Id > highestMatchId {
			highestMatchId = match.Id
		}
	}
	return highestMatchId
}

func getNextMatchId(matches []Match) int {
	return getHighestMatchId(matches) + 1
}

func UpdateKoMatchesWithThirds(tournament Tournament) []KoMatch {
	koMatches := tournament.KoMatches
	groupRankings := tournament.GroupRankings

	rankingThirds := FilterByThirdRank(groupRankings)
	sortedRankingThirds := sortRankingsByPointsAndGoals(rankingThirds)

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
	for index, rankingOfThird := range sortedRankingThirds {
		if index == 4 {
			break
		}
		totalWeight += groupIdWeight[GroupId(rankingOfThird.GroupId)]
	}

	thirdsEvaluationRules := tournament.ThirdsEvaluationRules
	var rankingOfThirdPattern []int
	for _, rule := range thirdsEvaluationRules {
		if rule.BestFourTeamsId == totalWeight {
			rankingOfThirdPattern = rule.BestFourTeamsArrangement
			break
		}
	}

	updatesKoMatches := make([]KoMatch, 0)
	updatedKoMatchIndex := 0
	numberFirstKoGroupMatches := 0
	for _, koMatch := range koMatches {
		if koMatch.GroupId1 == 0 {
			panic("Error: GroupId1 is unspecified for KoMatch with Id " + fmt.Sprint(koMatch.Id))
		}
		if koMatch.GroupId2 == 0 {
			rankingOfThirdIndex := rankingOfThirdPattern[updatedKoMatchIndex]
			rankingOfThird := rankingThirds[rankingOfThirdIndex]
			koMatch.GroupId2 = rankingOfThird.GroupId
			updatedKoMatchIndex++
		}
		if koMatch.GroupId == firstKoGroupId {
			numberFirstKoGroupMatches++
		}

		updatesKoMatches = append(updatesKoMatches, koMatch)
	}

	if numberFirstKoGroupMatches != NumberOfMatchesInFirstKoRound {
		panic("Error: Expected " + fmt.Sprint(NumberOfMatchesInFirstKoRound) + " matches in first Ko round, instead found " + fmt.Sprint(numberFirstKoGroupMatches))
	}

	return updatesKoMatches
}

func PlayKoRounds(tournament Tournament) Tournament {
	koMatchMap := GroupsMapKoMatches(tournament.KoMatches)
	koGroups := FilterByKoRound(tournament.Groups)
	for _, koGroup := range koGroups {
		fmt.Println("Playing Ko Group", koGroup.Name)
		koMatches := koMatchMap[koGroup.Id]
		if len(koMatches) == 0 {
			panic("No matches found for Ko Group " + koGroup.Name + " with Id " + fmt.Sprint(koGroup.Id))
		}
		rankingsSortedIntoGroups := GroupsMapGroupRankings(tournament.GroupRankings)
		matches := CreateMatchFromKoMatch(koMatches, rankingsSortedIntoGroups)
		tournament.Matches = append(tournament.Matches, matches...)
		matchResults := PlayKoGroupsMatches(matches)
		tournament.MatchResults = append(tournament.MatchResults, matchResults...)

		teamsInGroup := GetTeamsFromMatches(matches)
		teamsMap := make(map[int][]Team)
		teamsMap[koGroup.Id] = teamsInGroup
		groupRankings := DetermineGroupRanking(matchResults, teamsMap, []Group{koGroup})
		tournament.GroupRankings = append(tournament.GroupRankings, groupRankings...)
	}

	finalGroupId := koGroups[len(koGroups)-1].Id
	winner := DetermineTournamentWinner(finalGroupId, tournament.GroupRankings)
	tournament.Winner = winner

	return tournament
}

func CreateMatchFromKoMatch(koMatches []KoMatch, rankingsSortedIntoGroups map[int][]GroupRanking) []Match {
	matches := make([]Match, 0)
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

		match := Match{
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

func PlayKoGroupsMatches_(matches []Match) []MatchResult {
	var matchResults []MatchResult
	matchResultChannel := make(chan MatchResult, len(matches))
	numberOfMatches := len(matches)
	pointsForWinner := numberOfMatches // maybe not a good solution as it ignores the match setup defined in the DB, need to check
	for _, match := range matches {
		go PlayEliminationMatch(match, pointsForWinner, matchResultChannel)
		pointsForWinner--
	}

	for i := 0; i < numberOfMatches; i++ {
		matchResult := <-matchResultChannel
		matchResults = append(matchResults, matchResult)
	}

	return matchResults
}

var PlayKoGroupsMatches = PlayKoGroupsMatches_

func ResolveDrawInEliminationMatch_(match Match, pointsForWinner int) MatchResult {
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

func GetTeamsFromMatches(matches []Match) []Team {
	teams := make([]Team, 0)
	for _, match := range matches {
		teams = append(teams, match.Team1, match.Team2)
	}
	return teams
}

func DetermineTournamentWinner(finalGroupId int, groupRankings []GroupRanking) Team {
	rankingsSortedIntoGroups := GroupsMapGroupRankings(groupRankings)
	finalRankings := rankingsSortedIntoGroups[finalGroupId]

	if len(finalRankings) != 2 {
		panic("Error: Final group does not have 2 teams, instead has " + fmt.Sprint(len(finalRankings)) + " teams.")
	}
	winner := finalRankings[0].Team

	return winner
}
