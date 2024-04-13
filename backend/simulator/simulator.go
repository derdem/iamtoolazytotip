package simulator

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"

	"gonum.org/v1/gonum/stat/distuv"
)

const lambda = 1.3

func RunSimulator() TournamentMatches {
	groups := LoadGroupFromDb()
	return TournamentSimulator(groups)
}

func TournamentSimulator(groups []Group) TournamentMatches {
	fmt.Println("Start")

	allGroupMatches := make([]GroupMatch, 0)

	numberMatchesInGroupPhase := CountAllGroupMatches(groups)
	groupMatchChannel := make(chan GroupMatch, numberMatchesInGroupPhase)

	for _, group := range groups {
		for _, match := range group.Matches {
			go playGroupMatch(match, groupMatchChannel)
		}
	}

	for i := 0; i < numberMatchesInGroupPhase; i++ {
		allGroupMatches = append(allGroupMatches, <-groupMatchChannel)
	}

	fmt.Println("Round of 16")
	matches16 := getRoudOfSixteenMatches(groups)
	playedMatches16 := make([]Match, 0)
	for _, match := range matches16 {
		playedMatches16 = append(playedMatches16, playEliminationMatch(match))
	}

	fmt.Println("Round of 8")
	matches8 := getRoundOfEightMatches(playedMatches16)
	playedMatches8 := make([]Match, 0)
	for _, match := range matches8 {
		playedMatches8 = append(playedMatches8, playEliminationMatch(match))
	}

	fmt.Println("Round of 4")
	matches4 := getRoundOfFourMatches(playedMatches8)
	playedMatches4 := make([]Match, 0)
	for _, match := range matches4 {
		playedMatches4 = append(playedMatches4, playEliminationMatch(match))
	}

	fmt.Println("Final Match")
	matchFinal := CreateMatch(playedMatches4[0].Winner, playedMatches4[1].Winner, time.Date(2021, 7, 11, 19, 0, 0, 0, time.UTC))
	playedMatchFinal := playEliminationMatch(matchFinal)

	tournamentMatches := TournamentMatches{
		Group:   allGroupMatches,
		Sixteen: playedMatches16,
		Eight:   playedMatches8,
		Four:    playedMatches4,
		Final:   playedMatchFinal,
	}

	return tournamentMatches
}

func playGroupMatch(match GroupMatch, c chan GroupMatch) {
	team1 := match.Team1
	team2 := match.Team2
	fmt.Println(team1.Name + " - " + team2.Name)
	outcomeProbabilies := assignProbabilities(team1.Strength, team2.Strength)
	winnerCode := determineWinner(outcomeProbabilies)
	var team1Score int
	var team2Score int

	switch winnerCode {
	case 0:
		team1Score, team2Score = setRemisScore()
		team1.Points = team1.Points + 1
		team2.Points = team2.Points + 1
		match.Winner = nil
	case 1:
		team1Score = randomResult()
		team2Score = randomResultLoser(team1Score, team1.Strength-team2.Strength)
		team1.Points = team1.Points + 3
		team2.Points = team2.Points + 0
		match.Winner = team1
	case 2:
		team2Score = randomResult()
		team1Score = randomResultLoser(team2Score, team2.Strength-team1.Strength)
		team1.Points = team1.Points + 0
		team2.Points = team2.Points + 3
		match.Winner = team2
	}

	match.GoalsTeam1 = team1Score
	match.GoalsTeam2 = team2Score

	team1.Goals = team1.Goals + team1Score
	team2.Goals = team2.Goals + team2Score

	rand.Seed(time.Now().UnixNano())
	multiplier := time.Duration(rand.Intn(100))
	time.Sleep(time.Millisecond * multiplier)

	c <- match
}

func getRoudOfSixteenMatches(groups []Group) []Match {
	rankedGroups := make([][4]*Country, 0)
	for _, group := range groups {
		rankedGroup := determineGroupRanking(group)
		rankedGroups = append(rankedGroups, rankedGroup)
	}

	allThirds := make([]*Country, 0)
	for _, rankedGroup := range rankedGroups {
		allThirds = append(allThirds, rankedGroup[2])
	}

	bestFourThirds := getBestFourThirds(allThirds)

	matches := make([]Match, 0)
	matches = append(matches, CreateMatch(rankedGroups[0][1], rankedGroups[1][1], time.Date(2021, 6, 26, 16, 0, 0, 0, time.UTC)))
	matches = append(matches, CreateMatch(rankedGroups[0][0], rankedGroups[2][1], time.Date(2021, 6, 26, 19, 0, 0, 0, time.UTC)))
	matches = append(matches, CreateMatch(rankedGroups[2][0], bestFourThirds[0], time.Date(2021, 6, 27, 16, 0, 0, 0, time.UTC)))
	matches = append(matches, CreateMatch(rankedGroups[1][0], bestFourThirds[1], time.Date(2021, 6, 27, 19, 0, 0, 0, time.UTC)))
	matches = append(matches, CreateMatch(rankedGroups[3][1], rankedGroups[4][1], time.Date(2021, 6, 28, 16, 0, 0, 0, time.UTC)))
	matches = append(matches, CreateMatch(rankedGroups[5][0], bestFourThirds[2], time.Date(2021, 6, 28, 19, 0, 0, 0, time.UTC)))
	matches = append(matches, CreateMatch(rankedGroups[3][0], rankedGroups[5][1], time.Date(2021, 6, 29, 16, 0, 0, 0, time.UTC)))
	matches = append(matches, CreateMatch(rankedGroups[4][0], bestFourThirds[3], time.Date(2021, 6, 29, 19, 0, 0, 0, time.UTC)))
	return matches
}

func getRoundOfEightMatches(matches []Match) []Match {
	nextMatches := make([]Match, 0)
	nextMatches = append(nextMatches, CreateMatch(matches[5].Winner, matches[4].Winner, time.Date(2021, 7, 2, 16, 0, 0, 0, time.UTC)))
	nextMatches = append(nextMatches, CreateMatch(matches[3].Winner, matches[1].Winner, time.Date(2021, 7, 2, 19, 0, 0, 0, time.UTC)))
	nextMatches = append(nextMatches, CreateMatch(matches[2].Winner, matches[0].Winner, time.Date(2021, 7, 3, 16, 0, 0, 0, time.UTC)))
	nextMatches = append(nextMatches, CreateMatch(matches[7].Winner, matches[6].Winner, time.Date(2021, 7, 3, 19, 0, 0, 0, time.UTC)))

	return nextMatches
}

func getRoundOfFourMatches(matches []Match) []Match {
	nextMatches := make([]Match, 0)
	nextMatches = append(nextMatches, CreateMatch(matches[0].Winner, matches[1].Winner, time.Date(2021, 7, 6, 19, 0, 0, 0, time.UTC)))
	nextMatches = append(nextMatches, CreateMatch(matches[2].Winner, matches[3].Winner, time.Date(2021, 7, 7, 19, 0, 0, 0, time.UTC)))

	return nextMatches
}

func playEliminationMatch(match Match) Match {
	var team1 = match.Team1
	var team2 = match.Team2
	var team1Score int
	var team2Score int
	var team1PenaltyScore int = 0
	var team2PenaltyScore int = 0
	var winnerTeam *Country
	outcomeProbabilies := assignProbabilities(team1.Strength, team2.Strength)
	winnerCode := determineWinner(outcomeProbabilies)

	fmt.Println(team1.Name + " vs. " + team2.Name)
	switch winnerCode {
	case 0:
		team1Score, team2Score = setRemisScore()
		team1PenaltyScore, team2PenaltyScore = playPenalty(0, 0)
		resultString := fmt.Sprintf("%d (%d) - %d (%d)", team1Score, team1Score+team1PenaltyScore, team2Score, team1Score+team2PenaltyScore)
		fmt.Println(resultString)
		if team1Score+team1PenaltyScore > team2Score+team2PenaltyScore {
			winnerTeam = team1
		} else {
			winnerTeam = team2
		}
	case 1:
		team1Score = randomResult()
		team2Score = randomResultLoser(team1Score, team1.Strength-team2.Strength)
		resultString := fmt.Sprintf("%d - %d", team1Score, team2Score)
		fmt.Println(resultString)
		winnerTeam = team1
	case 2:
		team2Score = randomResult()
		team1Score = randomResultLoser(team2Score, team2.Strength-team1.Strength)
		resultString := fmt.Sprintf("%d - %d", team1Score, team2Score)
		fmt.Println(resultString)
		winnerTeam = team2
	}

	match.GoalsTeam1 = team1Score
	match.PenaltyScoreTeam1 = team1PenaltyScore
	match.GoalsTeam2 = team2Score
	match.PenaltyScoreTeam2 = team2PenaltyScore
	match.Winner = winnerTeam

	team1.Goals = team1.Goals + team1Score
	team1.PenaltyGoals = team1.PenaltyGoals + team1PenaltyScore
	team2.Goals = team2.Goals + team2Score
	team2.PenaltyGoals = team2.PenaltyGoals + team2PenaltyScore

	return match
}

func assignProbabilities(strength1 int, strength2 int) OutcomeProbabilities {
	probab1, probab2 := convertStrengthToProbabilities(strength1, strength2)
	allProbabilities := OutcomeProbabilities{Team1: probab1, Team2: probab2}
	allProbabilities.Remis = 1 - allProbabilities.Team1 - allProbabilities.Team2
	return allProbabilities
}

func convertStrengthToProbabilities(strength1 int, strength2 int) (float64, float64) {
	strengthDiff := strength1 - strength2

	if strengthDiff == -2 {
		return 0.1, 0.7
	} else if strengthDiff == -1 {
		return 0.2, 0.5
	} else if strengthDiff == 0 {
		return 0.3, 0.3
	} else if strengthDiff == 1 {
		return 0.5, 0.2
	} else if strengthDiff == 2 {
		return 0.7, 0.1
	} else {
		return 0, 0
	}
}

func determineWinner(outcomeChanges OutcomeProbabilities) int {
	rand.Seed(time.Now().UnixNano())
	randomResult := rand.Float64()
	if randomResult < outcomeChanges.Remis {
		return 0 // remis
	} else if randomResult < outcomeChanges.Remis+outcomeChanges.Team1 {
		return 1 // team 1 wins
	} else {
		return 2 // team 2 wins
	}
}

func setRemisScore() (int, int) {
	result := randomResult()
	return result, result
}

func randomResult() int {
	rand.Seed(time.Now().UnixNano())
	probab := rand.Float64()
	k := findK(probab, 0, lambda)
	return k + 1
}

func findK(probab float64, k int, lambda float64) int {
	p := distuv.Poisson{Lambda: lambda}
	x1 := distuv.Poisson.CDF(p, float64(k))
	x2 := distuv.Poisson.CDF(p, float64(k+1))

	if x2 < probab {
		return findK(probab, k+1, lambda)
	} else {
		gap1 := probab - x1
		gap2 := x2 - probab
		if gap1 < gap2 {
			return k
		} else {
			return k + 1
		}
	}
}

func randomResultLoser(resultWinner int, strengthDifference int) int {
	if resultWinner == 1 {
		return 0
	}
	rand.Seed(time.Now().UnixNano())
	if strengthDifference == 2 {
		dilute := math.Pow(0.8+(0.2*rand.Float64()), 2)
		p := rand.Float64() * dilute
		return int(math.Round(p * (float64(resultWinner - 1))))
	} else if strengthDifference == 1 {
		dilute := 0.8 + (0.2 * rand.Float64())
		p := rand.Float64() * dilute
		return int(math.Round(p * (float64(resultWinner - 1))))
	} else if strengthDifference == 0 {
		p := rand.Float64()
		return int(math.Round(p * (float64(resultWinner - 1))))
	} else if strengthDifference == -1 {
		concentrate := 1.2 - (0.2 * rand.Float64())
		p := rand.Float64() * concentrate
		if p > 1 {
			return resultWinner - 1
		} else {
			return int(math.Round(p * (float64(resultWinner - 1))))
		}
	} else if strengthDifference == -2 {
		concentrate := 1.5 - (0.5 * rand.Float64())
		p := rand.Float64() * concentrate
		if p > 1 {
			return resultWinner - 1
		} else {
			return int(math.Round(p * (float64(resultWinner - 1))))
		}
	} else {
		return 0
	}

}

func playPenalty(score1, score2 int) (int, int) {
	score1Increment := randomScoreBetween0And5() + score1
	score2Increment := randomScoreBetween0And5() + score2
	if score1Increment == score2Increment {
		return playPenalty(score1Increment, score2Increment)
	} else {
		return score1Increment, score2Increment
	}
}

func randomScoreBetween0And5() int {
	rand.Seed(time.Now().UnixNano())
	goals := 0
	for i := 1; i <= 5; i++ {
		randomResult := rand.Float64()
		if randomResult > .25 {
			goals++
		}
	}
	return goals
}

func CountAllGroupMatches(groups []Group) int {
	var numberMatches int = 0

	for _, group := range groups {
		numberMatches += len(group.Matches)
	}

	return numberMatches
}

func determineGroupRanking(group Group) [4]*Country {
	countries := group.Countries
	sort.Slice(countries[:], func(i, j int) bool {
		if countries[i].Points == countries[j].Points && countries[i].Goals > countries[j].Goals {
			return true
		}
		return countries[i].Points > countries[j].Points
	})

	return countries
}

func getBestFourThirds(thirds []*Country) [4]*Country {
	var thirdsSlice []*Country = thirds[:]
	sort.Slice(thirdsSlice, func(i, j int) bool {
		if thirdsSlice[i].Points == thirdsSlice[j].Points && thirdsSlice[i].Goals > thirdsSlice[j].Goals {
			return true
		}
		return thirdsSlice[i].Points > thirdsSlice[j].Points
	})
	var bestFourThirds = [4]*Country{thirdsSlice[0], thirdsSlice[1], thirdsSlice[2], thirdsSlice[3]}
	return bestFourThirds
}
