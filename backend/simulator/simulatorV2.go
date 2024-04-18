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

	playKoRounds(tournament)

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

func filterGroup(groupType GroupType) {
    return func (groups Group2) {
        filteredGroups := make([]Group2, 0)
        for _, group := range groups {
            if group.Type == groupType {
                filteredGroups = append(filteredGroups, group)
            }
        }
        return filteredGroups
    }
}

var filterByGroupPhase = filterBy(GroupType.GroupPhase)
var filterByKoRound = filterBy(GroupType.KoRound)

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

func playKoRounds(tournament Tournament) {
	// roundOf16 := playRoundOf16(tournament)
	// quarterFinals := playQuarterFinals(tournament)
	// semiFinals := playSemiFinals(tournament)
	// final := playFinal(tournament)

}

func playRoundOf16(tournament Tournament) []MatchResult {
	roundOf16Group := Group2{
		Id:           getNextGroupId(tournament.Groups),
		Name:         "Round of 16",
		TournamentId: tournament.Id,
	}
	tournament.Groups = append(tournament.Groups, roundOf16Group)
	roundOf16Matches := setupRoundOf16Matches(tournament, tournament.GroupRankings)
}

func setupRoundOf16Matches(tournament Tournament) {

	teams := make([]Team, 0)
	for _, ranking := range rankingScore {
		teams = append(teams, ranking.Team)
	}

	nextMatchId := getNextMatchId(tournament.Matches)
    rankingsSortedIntoGroups := getRankingsSortedIntoGroups(tournament.GroupRankings)

	roundOf16 := []Match2{
		Match2{
			Id:    nextMatchId,
			Team1: rankingsSortedIntoGroups[],
			Team2: teams[3],
		},
		Match2{
			Id:    nextMatchId + 1,
			Team1: teams[1],
			Team2: teams[2],
		},
		Match2{
			Id:    nextMatchId + 2,
			Team1: teams[4],
			Team2: teams[7],
		},
		Match2{
			Id:    nextMatchId + 3,
			Team1: teams[5],
			Team2: teams[6],
		},
		Match2{
			Id:    nextMatchId + 4,
			Team1: teams[8],
			Team2: teams[11],
		},
		Match2{
			Id:    nextMatchId + 5,
			Team1: teams[9],
			Team2: teams[10],
		},
		Match2{
			Id:    nextMatchId + 6,
			Team1: teams[12],
			Team2: teams[15],
		},
		Match2{
			Id:    nextMatchId + 7,
			Team1: teams[13],
			Team2: teams[14],
		},
	}

	tournament.Matches = append(tournament.Matches, roundOf16...)

	return tournament
}
