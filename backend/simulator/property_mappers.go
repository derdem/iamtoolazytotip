package simulator

func GroupsMapTeams(teams []Team) map[int][]Team {
	var teamsSortedIntoGroups = make(map[int][]Team) // groupId -> teams
	for _, team := range teams {
		teamsSortedIntoGroups[team.GroupId] = append(teamsSortedIntoGroups[team.GroupId], team)
	}
	return teamsSortedIntoGroups
}

func GroupsMapKoMatches(koMatches []KoMatch) map[int][]KoMatch {
	var koMatchesMappedToGroups = make(map[int][]KoMatch) // groupId -> koMatches
	for _, koMatch := range koMatches {
		koMatchesMappedToGroups[koMatch.GroupId] = append(koMatchesMappedToGroups[koMatch.GroupId], koMatch)
	}
	return koMatchesMappedToGroups
}

func GroupsMapGroupRankings(rankings []GroupRanking) map[int][]GroupRanking {
	var rankingsSortedIntoGroups = make(map[int][]GroupRanking) // groupId -> rankings
	for _, ranking := range rankings {
		rankingsSortedIntoGroups[ranking.GroupId] = append(rankingsSortedIntoGroups[ranking.GroupId], ranking)
	}
	return rankingsSortedIntoGroups
}
