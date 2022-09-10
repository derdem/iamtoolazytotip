package simulator

import (
	"sort"
)

type Group struct {
	name    string
	member1 Country
	member2 Country
	member3 Country
	member4 Country
}

type Groups struct {
	A Group
	B Group
	C Group
	D Group
	E Group
	F Group
}

type RoundOf16 struct {
	member1  Country
	member2  Country
	member3  Country
	member4  Country
	member5  Country
	member6  Country
	member7  Country
	member8  Country
	member9  Country
	member10 Country
	member11 Country
	member12 Country
	member13 Country
	member14 Country
	member15 Country
	member16 Country
}

func GetGroups(allCountries AllCountries) Groups {
	groupA := Group{
		name:    "Group A",
		member1: allCountries.italy,
		member2: allCountries.switzerland,
		member3: allCountries.turkey,
		member4: allCountries.wales,
	}
	groupB := Group{
		name:    "Group B",
		member1: allCountries.belgium,
		member2: allCountries.denmark,
		member3: allCountries.finland,
		member4: allCountries.russia,
	}

	groupC := Group{
		name:    "Group C",
		member1: allCountries.netherlands,
		member2: allCountries.northmazedonia,
		member3: allCountries.ukraine,
		member4: allCountries.austria,
	}

	groupD := Group{
		name:    "Group D",
		member1: allCountries.england,
		member2: allCountries.kroatia,
		member3: allCountries.scotland,
		member4: allCountries.czechrepublic,
	}

	groupE := Group{
		name:    "Group E",
		member1: allCountries.poland,
		member2: allCountries.sweden,
		member3: allCountries.slowakia,
		member4: allCountries.spain,
	}

	groupF := Group{
		name:    "Group F",
		member1: allCountries.germany,
		member2: allCountries.france,
		member3: allCountries.portugal,
		member4: allCountries.hungry,
	}

	var groups = Groups{A: groupA, B: groupB, C: groupC, D: groupD, E: groupE, F: groupF}
	return groups

}

func determineGroupWinner(group Group) []Country {
	groupRanking := make([]Country, 0, 4)
	// add first
	groupRanking = append(groupRanking, group.member1)

	// add second
	if isBetterFirstCountry(group.member2, groupRanking[0]) {
		groupRanking = append([]Country{group.member2}, groupRanking...)
	} else {
		groupRanking = append(groupRanking, group.member2)
	}

	// add third
	if isBetterFirstCountry(group.member3, groupRanking[0]) {
		groupRanking = append([]Country{group.member3}, groupRanking...)
	} else if isBetterFirstCountry(group.member3, groupRanking[1]) {
		groupRanking = append([]Country{groupRanking[0]}, group.member3, groupRanking[1])
	} else {
		groupRanking = append(groupRanking, group.member3)
	}

	// add fourth
	if isBetterFirstCountry(group.member4, groupRanking[0]) {
		groupRanking = append([]Country{group.member4}, groupRanking...)
	} else if isBetterFirstCountry(group.member4, groupRanking[1]) {
		groupRanking = append([]Country{groupRanking[0]}, group.member4, groupRanking[1], groupRanking[2])
	} else if isBetterFirstCountry(group.member4, groupRanking[2]) {
		groupRanking = append([]Country{groupRanking[0], groupRanking[1]}, group.member4, groupRanking[2])
	} else {
		groupRanking = append(groupRanking, group.member4)
	}
	return groupRanking
}

func isBetterFirstCountry(country1, country2 Country) bool {
	if country1.Points == country2.Points && country1.Goals > country2.Goals {
		return true
	}
	return country1.Points > country2.Points
}

func getBestFourThirds(thirds [6]Country) [4]Country {
	var thirdsSlice []Country = thirds[:]
	sort.Slice(thirdsSlice, func(i, j int) bool {
		if thirdsSlice[i].Points == thirdsSlice[j].Points && thirdsSlice[i].Goals > thirdsSlice[j].Goals {
			return true
		}
		return thirdsSlice[i].Points > thirdsSlice[j].Points
	})
	var bestFourThirds = [4]Country{thirdsSlice[0], thirdsSlice[1], thirdsSlice[2], thirdsSlice[3]}
	return bestFourThirds
}

func GetRoudOfSixteen(groups Groups) RoundOf16 {
	groupAranked := determineGroupWinner(groups.A)
	groupBranked := determineGroupWinner(groups.B)
	groupCranked := determineGroupWinner(groups.C)
	groupDranked := determineGroupWinner(groups.D)
	groupEranked := determineGroupWinner(groups.E)
	groupFranked := determineGroupWinner(groups.F)
	allThirds := [6]Country{groupAranked[2], groupBranked[2], groupCranked[2], groupDranked[2], groupEranked[2], groupFranked[2]}
	bestFourThirds := getBestFourThirds(allThirds)
	// fmt.Println("Group A")
	// fmt.Println(groupAranked)
	// fmt.Println("---------------------------------")
	// fmt.Println("Group B")
	// fmt.Println(groupBranked)
	// fmt.Println("---------------------------------")
	// fmt.Println("Group C")
	// fmt.Println(groupCranked)
	// fmt.Println("---------------------------------")
	// fmt.Println("Group D")
	// fmt.Println(groupDranked)
	// fmt.Println("---------------------------------")
	// fmt.Println("Group E")
	// fmt.Println(groupEranked)
	// fmt.Println("---------------------------------")
	// fmt.Println("Group F")
	// fmt.Println(groupFranked)
	// fmt.Println("---------------------------------")
	// fmt.Println("best 4 Thirds")
	// fmt.Println(bestFourThirds)
	// fmt.Println("---------------------------------")
	return RoundOf16{
		groupAranked[1], groupBranked[1],
		groupAranked[0], groupCranked[1],
		groupCranked[0], bestFourThirds[0],
		groupBranked[0], bestFourThirds[1],
		groupDranked[1], groupEranked[1],
		groupFranked[0], bestFourThirds[2],
		groupDranked[0], groupFranked[1],
		groupEranked[0], bestFourThirds[3],
	}
}
