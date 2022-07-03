package simulator

import "fmt"

func TournamentSimulator() {
	teams := GetAllCountries()
	groups := GetGroups(teams)
	playdays := GetPlaydays(groups)

	for i := range playdays {
		fmt.Printf("Day %d \n", i+1)
		for i, teampair := range playdays[i] {
			_ = i
			fmt.Printf("%s vs. %s \n", teampair[0].name, teampair[1].name)
		}
	}
}
