package simulator

import (
	"fmt"
	"math/rand"
	"time"

	"gonum.org/v1/gonum/stat/distuv"
)

type OutcomeProbabilities struct {
	team1 float64
	team2 float64
	remis float64
}

func TournamentSimulator() {
	// teams := GetAllCountries()
	// groups := GetGroups(teams)
	// playdays := GetPlaydays(groups)

	randomResult()

	// for i := range playdays {
	// 	fmt.Printf("Day %d \n", i+1)
	// 	for i, teampair := range playdays[i] {
	// 		_ = i
	// 		fmt.Printf("%s vs. %s \n", teampair[0].name, teampair[1].name)
	// 	}
	// }
}

func playGroupMatch(team1 Country, team2 Country) {
	outcomeProbabilies := assignProbabilities(team1.strength, team2.strength)
	winnerCode := determineWinner(outcomeProbabilies)
	if winnerCode == 0 {

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

func randomResult() {
	var p distuv.Poisson = distuv.Poisson{Lambda: 3.4}
	x := distuv.Poisson.CDF(p, 3.4)

	fmt.Println(x)
}

// def play_group_match(team1: Country, team2: Country):
//     team1_strength = team1.strength
//     team2_strength = team2.strength
//     outcome_chances = assign_probabilities(team1_strength, team2_strength)
//     winnercode = determine_winner(outcome_chances)
//     team1_score = None
//     team2_score = None
//     if winnercode == 0:
//         [team1_score, team2_score] = set_remis_score()
//         team1.set_goals_and_points(team1_score, 0)
//         team2.set_goals_and_points(team2_score, 0)
//     elif winnercode == 1:
//         [team1_score, team2_score] = set_score()
//         team1.set_goals_and_points(team1_score, 1)
//         team2.set_goals_and_points(team2_score, -1)
//     elif winnercode == 2:
//         [team2_score, team1_score] = set_score()
//         team1.set_goals_and_points(team1_score, -1)
//         team2.set_goals_and_points(team2_score, 1)
//     print("    " + team1.name + " vs. " + team2.name + ": " + str(team1_score) + " : " + str(team2_score))

// def assign_probabilities(team1_strength: int, team2_strength: int) -> OutcomeProbability:
//     probability1, probability2 = convert_strength_to_probabilities(team1_strength, team2_strength)
//     outcome_chances = OutcomeProbability(probability1, probability2)
//     return outcome_chances
