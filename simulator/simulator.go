package simulator

import (
	//"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"time"

	"gonum.org/v1/gonum/stat/distuv"
)

// const PhaseRegistry struct {
// 	qualification: "Qualification"
// }

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

// type MatchDetails struct {
// 	Opponent string `json:"opponent"`
// 	Phase
// }

// type TournamentOutcome struct {
// 	Countries
// 	QualificationPhase
// 	RoundOf16
// 	RoundOf8
// 	RoundOf4
// 	Final
// }

const lambda = 1.3

func TournamentSimulator() []MatchOutcome {
	fmt.Println("Start")
	c := make(chan MatchOutcome, 100)
	teams := GetAllCountries()
	playdays := GetPlaydays()
	var playdayOutcomes []MatchOutcome

	go func() {
		var closeC = false
		for i, playday := range playdays {
			_ = i
			for j, teampair := range playday {
				_ = j
				if i == len(playdays)-1 && j == len(playday)-1 {
					closeC = true
				}
				go playGroupMatch(teampair[0], teampair[1], c, closeC)

			}
		}
	}()

	for matchOutcome := range c {
		UpdateCountry(&teams, matchOutcome.Team1)
		UpdateCountry(&teams, matchOutcome.Team2)
		playdayOutcomes = append(playdayOutcomes, matchOutcome)
	}

	fmt.Println("Round of 16")
	groups := GetGroups(teams)
	matchesRoundOf16 := getRoundOf16Matches(groups)
	var roundOf16Winners [8]Country
	for i, matchPair := range matchesRoundOf16 {
		winningCountry := playEliminationMatch(matchPair[0], matchPair[1])
		roundOf16Winners[i] = winningCountry
	}

	fmt.Println("Round of 8")
	matchesRoundOf8 := getRoundOf8Matches(roundOf16Winners)
	var roundOf8Winners [4]Country
	for i, matchPair := range matchesRoundOf8 {
		winningCountry := playEliminationMatch(matchPair[0], matchPair[1])
		roundOf8Winners[i] = winningCountry
	}

	fmt.Println("Round of 4")
	matchesRoundOf4 := getRoundOf4Matches(roundOf8Winners)
	var roundOf4Winners [2]Country
	for i, matchPair := range matchesRoundOf4 {
		winningCountry := playEliminationMatch(matchPair[0], matchPair[1])
		roundOf4Winners[i] = winningCountry
	}

	fmt.Println("Final Match")
	playEliminationMatch(roundOf4Winners[0], roundOf4Winners[1])

	return playdayOutcomes
}

func playGroupMatch(team1 Country, team2 Country, c chan MatchOutcome, closeC bool) {
	fmt.Println(team1.Name + " - " + team2.Name)
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

	rand.Seed(time.Now().UnixNano())
	multiplier := time.Duration(rand.Intn(100))
	time.Sleep(time.Millisecond * multiplier)

	c <- MatchOutcome{Team1: team1, Team1Score: team1Score, Team2: team2, Team2Score: team2Score}

	if closeC {
		time.Sleep(time.Millisecond * 100)
		close(c)
	}

}

func playEliminationMatch(team1 Country, team2 Country) Country {
	var team1Score int
	var team2Score int
	var team1PenaltyScore int
	var team2PenaltyScore int
	outcomeProbabilies := assignProbabilities(team1.Strength, team2.Strength)
	winnerCode := determineWinner(outcomeProbabilies)

	fmt.Println(team1.Name + " vs. " + team2.Name)
	if winnerCode == 0 {
		team1Score, team2Score = setRemisScore()
		team1PenaltyScore, team2PenaltyScore = playPenalty(0, 0)
		resultString := fmt.Sprintf("%d (%d) - %d (%d)", team1Score, team1Score+team1PenaltyScore, team2Score, team1Score+team2PenaltyScore)
		fmt.Println(resultString)
		if team1Score+team1PenaltyScore > team2Score+team2PenaltyScore {
			return team1
		} else {
			return team2
		}
	} else if winnerCode == 1 {
		team1Score = randomResult()
		team2Score = randomResultLoser(team1Score, team1.Strength-team2.Strength)
		resultString := fmt.Sprintf("%d - %d", team1Score, team2Score)
		fmt.Println(resultString)
		return team1
	} else {
		team2Score = randomResult()
		team1Score = randomResultLoser(team2Score, team2.Strength-team1.Strength)
		resultString := fmt.Sprintf("%d - %d", team1Score, team2Score)
		fmt.Println(resultString)
		return team2
	}
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
		return 0 // remis
	} else if randomResult < outcomeChanges.remis+outcomeChanges.team1 {
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
	score1Increment := randomScoreBetween0And5()
	score2Increment := randomScoreBetween0And5()
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
