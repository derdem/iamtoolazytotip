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

func getRoudOfSixteenMatches(groups []Group) []Match {
	rankedGroups := make([][]*Country, 0)
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
	matches = append(matches, defineMatch(rankedGroups[0][1], rankedGroups[1][1]))
	matches = append(matches, defineMatch(rankedGroups[0][0], rankedGroups[2][1]))
	matches = append(matches, defineMatch(rankedGroups[2][0], bestFourThirds[0]))
	matches = append(matches, defineMatch(rankedGroups[1][0], bestFourThirds[1]))
	matches = append(matches, defineMatch(rankedGroups[3][1], rankedGroups[4][1]))
	matches = append(matches, defineMatch(rankedGroups[5][0], bestFourThirds[2]))
	matches = append(matches, defineMatch(rankedGroups[3][0], rankedGroups[5][1]))
	matches = append(matches, defineMatch(rankedGroups[4][0], bestFourThirds[3]))
	return matches
}

func getRoundOfEightMatches(matches []Match) []Match {
	nextMatches := make([]Match, 0)
	nextMatches = append(nextMatches, defineMatch(matches[5].winner, matches[4].winner))
	nextMatches = append(nextMatches, defineMatch(matches[3].winner, matches[1].winner))
	nextMatches = append(nextMatches, defineMatch(matches[2].winner, matches[0].winner))
	nextMatches = append(nextMatches, defineMatch(matches[7].winner, matches[6].winner))

	return nextMatches
}

func getRoundOfFourMatches(matches []Match) []Match {
	nextMatches := make([]Match, 0)
	nextMatches = append(nextMatches, defineMatch(matches[0].winner, matches[1].winner))
	nextMatches = append(nextMatches, defineMatch(matches[2].winner, matches[3].winner))

	return nextMatches
}
