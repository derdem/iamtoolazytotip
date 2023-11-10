package em2024

import (
	"github.com/derdem/iamtoolazytotip/simulator"
)

func CreateEm2024Groups() []simulator.Group {
	groupsRaw := []byte(`[
        {
            "name": "Group A",
            "countries": [
                {"strength": 3, "name": "Germany", "points": 0, "goals": 0, "penaltyGoals": 0},
                {"strength": 3, "name": "Spain", "points": 0, "goals": 0, "penaltyGoals": 0},
                {"strength": 3, "name": "France", "points": 0, "goals": 0, "penaltyGoals": 0},
                {"strength": 3, "name": "England", "points": 0, "goals": 0, "penaltyGoals": 0}
            ],
            "matches": [
                {"team1": 0, "team2": 1, "playtime": "2024-06-14T19:00:00Z"},
                {"team1": 2, "team2": 3, "playtime": "2024-06-15T19:00:00Z"},
                {"team1": 1, "team2": 3, "playtime": "2024-06-19T19:00:00Z"},
                {"team1": 0, "team2": 2, "playtime": "2024-06-19T19:00:00Z"},
                {"team1": 3, "team2": 0, "playtime": "2024-06-23T19:00:00Z"},
                {"team1": 1, "team2": 2, "playtime": "2024-06-23T19:00:00Z"}
            ]
        },
        {
            "name": "Group B",
            "countries": [
                {"strength": 2, "name": "Belgium", "points": 0, "goals": 0, "penaltyGoals": 0},
                {"strength": 3, "name": "Italy", "points": 0, "goals": 0, "penaltyGoals": 0},
                {"strength": 3, "name": "Netherlands", "points": 0, "goals": 0, "penaltyGoals": 0},
                {"strength": 1, "name": "Wales", "points": 0, "goals": 0, "penaltyGoals": 0}
            ],
            "matches": [
                {"team1": 0, "team2": 1, "playtime": "2024-06-14T19:00:00Z"},
                {"team1": 2, "team2": 3, "playtime": "2024-06-15T19:00:00Z"},
                {"team1": 1, "team2": 3, "playtime": "2024-06-19T19:00:00Z"},
                {"team1": 0, "team2": 2, "playtime": "2024-06-19T19:00:00Z"},
                {"team1": 3, "team2": 0, "playtime": "2024-06-23T19:00:00Z"},
                {"team1": 1, "team2": 2, "playtime": "2024-06-23T19:00:00Z"}
            ]
        }
    ]`)

	groups := simulator.CreateGroupsFromRaw(groupsRaw)
	return groups
}
