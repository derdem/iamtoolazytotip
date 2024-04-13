package simulator

import (
	"fmt"
	"sort"
)

func TournamentSimulator2(tournament Tournament) {
	fmt.Println("Start")

	matchResults := playGroupMatches(tournament)

	rankingScore := determineGroupRanking2(matchResults, tournament)

	fmt.Println("Ranking")
	for _, ranking := range rankingScore {
		fmt.Printf("%+v\n", ranking)
	}

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
	var teamsSortedIntoGroups = make(map[int][]Team)
	var groupRankings = make([]GroupRanking, 0)

	for _, matchResult := range matchResults {
		teamPoints[matchResult.Match.Team1.Id] += matchResult.Team1PointsGained
		teamPoints[matchResult.Match.Team2.Id] += matchResult.Team2PointsGained

		teamGoals[matchResult.Match.Team1.Id] += matchResult.Team1Goals
		teamGoals[matchResult.Match.Team2.Id] += matchResult.Team2Goals
	}

	for _, team := range tournament.Teams {
		teamsSortedIntoGroups[team.GroupId] = append(teamsSortedIntoGroups[team.GroupId], team)
	}

	for _, group := range tournament.Groups {
		groupRankings = append(
			groupRankings,
			determineRankingPerGroup(
				group.Id,
				teamsSortedIntoGroups[group.Id],
				teamPoints,
				teamGoals)...,
		)
	}

	return groupRankings

}

func determineRankingPerGroup(groupId int, teams []Team, teamPoints map[int]int, teamGoals map[int]int) []GroupRanking {
	sort.Slice(teams, func(i, j int) bool {
		// If the points are equal, the team with the most goals scored is ranked higher
		if teamPoints[teams[i].Id] == teamPoints[teams[j].Id] {
			return teamGoals[teams[i].Id] > teamGoals[teams[j].Id]
		}
		// Otherwise, the team with the most points is ranked higher
		return teamPoints[teams[i].Id] > teamPoints[teams[j].Id]
	})

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
