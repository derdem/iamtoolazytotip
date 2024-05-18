package simulator

import (
	"fmt"
	"math"
	"math/rand"

	"gonum.org/v1/gonum/stat/distuv"
)

func PlayGroupMatch_(match Match2, matchResultChannel chan MatchResult) {
	fmt.Println("Playing match", match.Id)
	outcomeProbabilies := AssignProbabilities(match.Team1.Strength, match.Team2.Strength)
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
		team2Score := ScoreLooser(team1Score, match.Team1.Strength-match.Team2.Strength)
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
		team1Score := ScoreLooser(team2Score, match.Team1.Strength-match.Team2.Strength)
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

var PlayGroupMatch = PlayGroupMatch_

func PlayEliminationMatch_(match Match2, pointsForWinner int, matchResultChannel chan MatchResult) {
	var team1 = match.Team1
	var team2 = match.Team2
	var team1Score int
	var team2Score int
	var team1PenaltyScore int = 0
	var team2PenaltyScore int = 0
	var result MatchResult

	outcomeProbabilies := AssignProbabilities(team1.Strength, team2.Strength)
	winnerCode := DetermineWinner(outcomeProbabilies)

	fmt.Println(team1.Name + " vs. " + team2.Name)
	switch winnerCode {
	case 0:
		result = ResolveDrawInEliminationMatch(match, pointsForWinner)
		resultString := fmt.Sprintf("%d (%d) - %d (%d)", team1Score, team1Score+team1PenaltyScore, team2Score, team1Score+team2PenaltyScore)
		fmt.Println(resultString)

	case 1:
		team1Score = RandomResult()
		team2Score = ScoreLooser(team1Score, team1.Strength-team2.Strength)
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
		team1Score = ScoreLooser(team2Score, team2.Strength-team1.Strength)
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

var PlayEliminationMatch = PlayEliminationMatch_

func AssignProbabilities(strength1 int, strength2 int) OutcomeProbabilities {
	probab1, probab2 := ConvertStrengthToProbabilities(strength1, strength2)
	allProbabilities := OutcomeProbabilities{Team1: probab1, Team2: probab2}
	allProbabilities.Remis = 1 - allProbabilities.Team1 - allProbabilities.Team2
	return allProbabilities
}

func ConvertStrengthToProbabilities(strength1 int, strength2 int) (float64, float64) {
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

func DetermineWinner_(outcomeChanges OutcomeProbabilities) int {
	randomResult := rand.Float64()
	if randomResult < outcomeChanges.Remis {
		return 0 // remis
	} else if randomResult < outcomeChanges.Remis+outcomeChanges.Team1 {
		return 1 // team 1 wins
	} else {
		return 2 // team 2 wins
	}
}

var DetermineWinner = DetermineWinner_

func SetRemisScore_() (int, int) {
	result := RandomResult()
	return result, result
}

var SetRemisScore = SetRemisScore_

func RandomResult_() int {
	probab := rand.Float64()
	k := FindK(probab, 0, lambda)
	return k + 1
}

var RandomResult = RandomResult_

func FindK(probab float64, k int, lambda float64) int {
	p := distuv.Poisson{Lambda: lambda}
	x1 := distuv.Poisson.CDF(p, float64(k))
	x2 := distuv.Poisson.CDF(p, float64(k+1))

	if x2 < probab {
		return FindK(probab, k+1, lambda)
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

func ScoreLooser_(resultWinner int, strengthDifference int) int {
	if resultWinner == 1 {
		return 0
	}
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

var ScoreLooser = ScoreLooser_

func PlayPenalty_(score1, score2 int) (int, int) {
	score1Increment := RandomScoreBetween0And5() + score1
	score2Increment := RandomScoreBetween0And5() + score2
	if score1Increment == score2Increment {
		return PlayPenalty_(score1Increment, score2Increment)
	} else {
		return score1Increment, score2Increment
	}
}

var PlayPenalty = PlayPenalty_

func RandomScoreBetween0And5() int {
	goals := 0
	for i := 1; i <= 5; i++ {
		randomResult := rand.Float64()
		if randomResult > .25 {
			goals++
		}
	}
	return goals
}
