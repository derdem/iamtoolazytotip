package simulator

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"time"

	"gonum.org/v1/gonum/stat/distuv"
)

type OutcomeProbabilities struct {
	team1 float64
	team2 float64
	remis float64
}

type MatchOutcome struct {
	Team1      Country `json:"team1"`
	Team1Score int     `json:"team1Score"`
	Team2      Country `json:"team2"`
	Team2Score int     `json:"team2Score"`
}

const lambda = 1.3

func TournamentSimulator() []MatchOutcome {
	teams := GetAllCountries()
	groups := GetGroups(teams)
	playdays := GetPlaydays(groups)
	var playdayOutcomes []MatchOutcome

	for i := range playdays {
		fmt.Printf("Day %d \n", i+1)
		for j, teampair := range playdays[i] {
			_ = j
			matchOutcome := playGroupMatch(teampair[0], teampair[1])
			matchOutcomePrintable, _ := json.Marshal(matchOutcome)
			fmt.Println(string(matchOutcomePrintable))
			playdayOutcomes = append(playdayOutcomes, matchOutcome)
		}
	}
	return playdayOutcomes
}

func playGroupMatch(team1 Country, team2 Country) MatchOutcome {
	outcomeProbabilies := assignProbabilities(team1.Strength, team2.Strength)
	winnerCode := determineWinner(outcomeProbabilies)
	var team1Score int
	var team2Score int

	if winnerCode == 0 {
		team1Score, team2Score = setRemisScore()
		team1.Points = team1.Points + 1
		team2.Points = team2.Points + 1
	} else if winnerCode == 1 {
		team1Score = randomResult()
		team2Score = randomResultLoser(team1Score, team1.Strength-team2.Strength)
		team1.Points = team1.Points + 3
		team2.Points = team2.Points + 0
	} else if winnerCode == 2 {
		team2Score = randomResult()
		team1Score = randomResultLoser(team2Score, team2.Strength-team1.Strength)
		team1.Points = team1.Points + 0
		team2.Points = team2.Points + 3
	}
	team1.Goals = team1.Goals + team1Score
	team2.Goals = team2.Goals + team2Score

	return MatchOutcome{Team1: team1, Team1Score: team1Score, Team2: team2, Team2Score: team2Score}

}

func assignProbabilities(strength1 int, strength2 int) OutcomeProbabilities {
	probab1, probab2 := convertStrengthToProbabilities(strength1, strength2)
	allProbabilities := OutcomeProbabilities{team1: probab1, team2: probab2}
	allProbabilities.remis = 1 - allProbabilities.team1 - allProbabilities.team2
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
	if randomResult < outcomeChanges.remis {
		// remis
		return 0
	} else if randomResult < outcomeChanges.remis+outcomeChanges.team1 {
		// team 1 wins
		return 1
	} else {
		// team 2 wins
		return 2
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
