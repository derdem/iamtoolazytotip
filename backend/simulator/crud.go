package simulator

import (
	"encoding/json"
	"time"
)

func CreateGroupsFromRaw(groupsRaw []byte) []Group {
	var groupsDeserialised []GroupDeserialisation
	json.Unmarshal(groupsRaw, &groupsDeserialised)
	groups := make([]Group, 0)
	for _, groupRaw := range groupsDeserialised {
		groups = append(groups, createGroup(groupRaw))
	}
	return groups
}

func createGroup(groupRaw GroupDeserialisation) Group {
	matches := make([]GroupMatch, 0)
	for _, match := range groupRaw.Matches {
		matches = append(matches, CreateGroupMatch(
			groupRaw.Countries[match.Team1Index],
			groupRaw.Countries[match.Team2Index],
			match.Playtime,
			groupRaw.Name),
		)
	}

	return Group{
		Name:      groupRaw.Name,
		Countries: groupRaw.Countries,
		Matches:   matches,
	}
}

func CreateMatch(team1 *Country, team2 *Country, playtime time.Time) Match {
	return Match{
		Team1:      team1,
		Team2:      team2,
		Playtime:   playtime,
		GoalsTeam1: 0,
		GoalsTeam2: 0,
		Winner:     nil,
	}
}

func CreateGroupMatch(team1 *Country, team2 *Country, playtime time.Time, groupName string) GroupMatch {
	return GroupMatch{
		Match: Match{
			Team1:      team1,
			Team2:      team2,
			Playtime:   playtime,
			GoalsTeam1: 0,
			GoalsTeam2: 0,
			Winner:     nil,
		},
		GroupName: groupName,
	}
}

func ConvertStrengthStringToInt(strength Strength) int {
	switch strength {
	case LowStrength:
		return 1
	case MediumStrength:
		return 2
	case HighStrength:
		return 3
	default:
		panic("Country Unknown strength")
	}
}
