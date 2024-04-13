package em2021

import (
	"encoding/json"

	"github.com/derdem/iamtoolazytotip/simulator"
)

func getGroupsRaw() []byte {
	groupsRaw := []byte(`[
        {
            "name": "Group A",
            "countries": [
                {"strength": 2, "name": "Turkey", "points": 0, "goals": 0, "penaltyGoals": 0},
                {"strength": 3, "name": "Italy", "points": 0, "goals": 0, "penaltyGoals": 0},
                {"strength": 1, "name": "Wales", "points": 0, "goals": 0, "penaltyGoals": 0},
                {"strength": 2, "name": "Switzerland", "points": 0, "goals": 0, "penaltyGoals": 0}
            ],
            "matches": [
                {"team1": 0, "team2": 1, "playtime": "2021-06-11T19:00:00Z"},
                {"team1": 2, "team2": 3, "playtime": "2021-06-12T13:00:00Z"},
                {"team1": 0, "team2": 2, "playtime": "2021-06-16T16:00:00Z"},
                {"team1": 1, "team2": 3, "playtime": "2021-06-16T19:00:00Z"},
                {"team1": 3, "team2": 0, "playtime": "2021-06-20T16:00:00Z"},
                {"team1": 1, "team2": 2, "playtime": "2021-06-20T16:00:00Z"}
            ]
        },
        {
            "name": "Group B",
            "countries": [
                {"strength": 2, "name": "Denmark", "points": 0, "goals": 0, "penaltyGoals": 0},
                {"strength": 1, "name": "Finland", "points": 0, "goals": 0, "penaltyGoals": 0},
                {"strength": 2, "name": "Belgien", "points": 0, "goals": 0, "penaltyGoals": 0},
                {"strength": 2, "name": "Russia", "points": 0, "goals": 0, "penaltyGoals": 0}
            ],
            "matches": [
                {"team1": 0, "team2": 1, "playtime": "2021-06-12T16:00:00Z"},
                {"team1": 2, "team2": 3, "playtime": "2021-06-12T19:00:00Z"},
                {"team1": 1, "team2": 3, "playtime": "2021-06-16T13:00:00Z"},
                {"team1": 0, "team2": 2, "playtime": "2021-06-17T16:00:00Z"},
                {"team1": 3, "team2": 0, "playtime": "2021-06-21T19:00:00Z"},
                {"team1": 1, "team2": 2, "playtime": "2021-06-21T19:00:00Z"}
            ]
        },
        {
            "name": "Group C",
            "countries": [
                {"strength": 3, "name": "Netherlands", "points": 0, "goals": 0, "penaltyGoals": 0},
                {"strength": 1, "name": "Ukraine", "points": 0, "goals": 0, "penaltyGoals": 0},
                {"strength": 2, "name": "Austria", "points": 0, "goals": 0, "penaltyGoals": 0},
                {"strength": 1, "name": "North Mazedonia", "points": 0, "goals": 0, "penaltyGoals": 0}
            ],
            "matches": [
                {"team1": 2, "team2": 3, "playtime": "2021-06-13T16:00:00Z"},
                {"team1": 0, "team2": 1, "playtime": "2021-06-13T19:00:00Z"},
                {"team1": 1, "team2": 3, "playtime": "2021-06-17T13:00:00Z"},
                {"team1": 0, "team2": 2, "playtime": "2021-06-17T19:00:00Z"},
                {"team1": 3, "team2": 0, "playtime": "2021-06-21T16:00:00Z"},
                {"team1": 1, "team2": 3, "playtime": "2021-06-21T16:00:00Z"}
            ]
        },
        {
            "name": "Group D",
            "countries": [
                {"strength": 3, "name": "England", "points": 0, "goals": 0, "penaltyGoals": 0},
                {"strength": 2, "name": "Kroatia", "points": 0, "goals": 0, "penaltyGoals": 0},
                {"strength": 2, "name": "Scotland", "points": 0, "goals": 0, "penaltyGoals": 0},
                {"strength": 2, "name": "Czech Republic", "points": 0, "goals": 0, "penaltyGoals": 0}
            ],
            "matches": [
                {"team1": 0, "team2": 1, "playtime": "2021-06-13T13:00:00Z"},
                {"team1": 2, "team2": 3, "playtime": "2021-06-14T13:00:00Z"},
                {"team1": 1, "team2": 3, "playtime": "2021-06-18T16:00:00Z"},
                {"team1": 0, "team2": 2, "playtime": "2021-06-18T19:00:00Z"},
                {"team1": 1, "team2": 2, "playtime": "2021-06-22T19:00:00Z"},
                {"team1": 3, "team2": 0, "playtime": "2021-06-22T19:00:00Z"}
            ]
        },
        {
            "name": "Group E",
            "countries": [
                {"strength": 3, "name": "Spain", "points": 0, "goals": 0, "penaltyGoals": 0},
                {"strength": 2, "name": "Sweden", "points": 0, "goals": 0, "penaltyGoals": 0},
                {"strength": 2, "name": "Poland", "points": 0, "goals": 0, "penaltyGoals": 0},
                {"strength": 2, "name": "Slowakia", "points": 0, "goals": 0, "penaltyGoals": 0}
            ],
            "matches": [
                {"team1": 2, "team2": 3, "playtime": "2021-06-14T16:00:00Z"},
                {"team1": 0, "team2": 1, "playtime": "2021-06-14T19:00:00Z"},
                {"team1": 1, "team2": 3, "playtime": "2021-06-18T13:00:00Z"},
                {"team1": 0, "team2": 2, "playtime": "2021-06-19T19:00:00Z"},
                {"team1": 3, "team2": 0, "playtime": "2021-06-23T16:00:00Z"},
                {"team1": 1, "team2": 2, "playtime": "2021-06-23T16:00:00Z"}
            ]
        },
        {
            "name": "Group F",
            "countries": [
                {"strength": 1, "name": "Hungry", "points": 0, "goals": 0, "penaltyGoals": 0},
                {"strength": 3, "name": "Portugal", "points": 0, "goals": 0, "penaltyGoals": 0},
                {"strength": 3, "name": "France", "points": 0, "goals": 0, "penaltyGoals": 0},
                {"strength": 3, "name": "Germany", "points": 0, "goals": 0, "penaltyGoals": 0}
            ],
            "matches": [
                {"team1": 0, "team2": 1, "playtime": "2021-06-15T16:00:00Z"},
                {"team1": 2, "team2": 3, "playtime": "2021-06-15T19:00:00Z"},
                {"team1": 0, "team2": 2, "playtime": "2021-06-19T13:00:00Z"},
                {"team1": 1, "team2": 3, "playtime": "2021-06-19T16:00:00Z"},
                {"team1": 1, "team2": 2, "playtime": "2021-06-23T19:00:00Z"},
                {"team1": 3, "team2": 0, "playtime": "2021-06-23T19:00:00Z"}
            ]
        }
    ]`)

	// check if groupsRaw is correct json
	if !json.Valid(groupsRaw) {
		panic("Invalid json")
	}

	return groupsRaw
}

func Run2021Tournament() simulator.TournamentMatches {
	groups := simulator.CreateGroupsFromRaw(getGroupsRaw())
	tournamentMatches := simulator.TournamentSimulator(groups)
	return tournamentMatches
}
