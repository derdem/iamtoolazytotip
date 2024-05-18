package simulator

func FilterGroup(groupType GroupType) func(groups []Group2) []Group2 {
	return func(groups []Group2) []Group2 {
		filteredGroups := make([]Group2, 0)
		for _, group := range groups {
			if group.GroupType == groupType {
				filteredGroups = append(filteredGroups, group)
			}
		}
		return filteredGroups
	}
}

var FilterByGroupPhase = FilterGroup(GroupPhaseGroupType)
var FilterByKoRound = FilterGroup(KoPhaseGroupType)

func FilterByThirdRank(rankings []GroupRanking) []GroupRanking {
	rankingThirds := make([]GroupRanking, 0)
	for _, ranking := range rankings {
		if ranking.Ranking == 3 {
			rankingThirds = append(rankingThirds, ranking)
		}
	}
	return rankingThirds
}
