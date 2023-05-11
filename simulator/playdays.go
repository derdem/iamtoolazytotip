package simulator

func DeterminePlaydaysFromGroup(groups []Group) []PlaydayMatches {
	// combine all matches per group into one map
	playdaysMap := make(map[int]PlaydayMatches)

	for _, group := range groups {
		playplan := group.playplan
		for day, matches := range playplan {
			_, ok := playdaysMap[day]
			if !ok {
				playdaysMap[day] = PlaydayMatches{}
			}
			playdaysMap[day] = append(playdaysMap[day], matches...)
		}
	}

	// convert the map to a slice
	playdays := []PlaydayMatches{}
	for _, matchesOfTheDay := range playdaysMap {
		playdays = append(playdays, matchesOfTheDay)
	}

	return playdays
}

func CountAllGroupMatches(playdays []PlaydayMatches) int {
	var numberMatches int = 0

	for _, playday := range playdays {
		numberMatches += len(playday)
	}

	return numberMatches
}

func getRoundOf16Matches(groups Groups) [8][2]*Country {
	roundOf16 := GetRoudOfSixteen(groups)
	return [8][2]*Country{
		{roundOf16.member1, roundOf16.member2},
		{roundOf16.member3, roundOf16.member4},
		{roundOf16.member5, roundOf16.member6},
		{roundOf16.member7, roundOf16.member8},
		{roundOf16.member9, roundOf16.member10},
		{roundOf16.member11, roundOf16.member12},
		{roundOf16.member13, roundOf16.member14},
		{roundOf16.member15, roundOf16.member16},
	}
}

func getRoundOf8Matches(teams [8]Country) [4][2]Country {
	return [4][2]Country{
		{teams[5], teams[4]},
		{teams[3], teams[1]},
		{teams[2], teams[0]},
		{teams[7], teams[6]},
	}
}

func getRoundOf4Matches(teams [4]Country) [2][2]Country {
	return [2][2]Country{
		{teams[0], teams[1]},
		{teams[2], teams[3]},
	}
}
