package simulator

import (
	"encoding/json"
	"time"

	"github.com/derdem/iamtoolazytotip/simulator/readTournamentFromDb"
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

func LoadGroupFromDb() []Group {
	groupsRaw := readTournamentFromDb.ReadTournament1()
	groups := make([]Group, 0)
	for _, groupRaw := range groupsRaw {
		countries, countryIdIndexMap := getCountriesFromDbEntries(groupRaw.Teams)
		matches := getMatchesFromDbEntries(groupRaw.Matches, countryIdIndexMap, countries, groupRaw.Name)

		group := Group{
			Name:      groupRaw.Name,
			Countries: countries,
			Matches:   matches,
		}
		groups = append(groups, group)
	}
	return groups
}

func getCountriesFromDbEntries(teams []readTournamentFromDb.Team) ([4]*Country, map[int]int) {
	return [4]*Country{
			{
				Name:     teams[0].Name,
				Strength: ConvertStrengthStringToInt(teams[0].Strength),
				Points:   0,
				Goals:    0,
			},
			{
				Name:     teams[1].Name,
				Strength: ConvertStrengthStringToInt(teams[1].Strength),
				Points:   0,
				Goals:    0,
			},
			{
				Name:     teams[2].Name,
				Strength: ConvertStrengthStringToInt(teams[2].Strength),
				Points:   0,
				Goals:    0,
			},
			{
				Name:     teams[3].Name,
				Strength: ConvertStrengthStringToInt(teams[3].Strength),
				Points:   0,
				Goals:    0,
			},
		},
		map[int]int{
			teams[0].Id: 0,
			teams[1].Id: 1,
			teams[2].Id: 2,
			teams[3].Id: 3,
		}
}

func ConvertStrengthStringToInt(strength readTournamentFromDb.Strength) int {
	switch strength {
	case readTournamentFromDb.LowStrength:
		return 1
	case readTournamentFromDb.MediumStrength:
		return 2
	case readTournamentFromDb.HighStrength:
		return 3
	default:
		panic("Country Unknown strength")
	}
}

func getMatchesFromDbEntries(matches []readTournamentFromDb.Match, countryIdIndexMap map[int]int, countries [4]*Country, groupName string) []GroupMatch {
	groupMatches := make([]GroupMatch, 0)
	for _, match := range matches {
		country1Index := countryIdIndexMap[match.Team1Id]
		country2Index := countryIdIndexMap[match.Team2Id]
		country1 := countries[country1Index]
		country2 := countries[country2Index]

		groupMatch := GroupMatch{
			Match: Match{
				Team1:             country1,
				Team2:             country2,
				Playtime:          time.Now(),
				GoalsTeam1:        0,
				PenaltyScoreTeam1: 0,
				GoalsTeam2:        0,
				PenaltyScoreTeam2: 0,
				Winner:            nil,
			},
			GroupName: groupName,
		}
		groupMatches = append(groupMatches, groupMatch)
	}

	return groupMatches
}
