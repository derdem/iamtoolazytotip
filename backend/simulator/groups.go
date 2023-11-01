package simulator

import (
	"sort"
)

type Country struct {
	Strength     int    `json:"strength"`
	Name         string `json:"name"`
	Points       int    `json:"points"`
	Goals        int    `json:"goals"`
	PenaltyGoals int    `json:"penaltyGoals"`
}

type Match struct {
	Team1             *Country `json:"team1"`
	Team2             *Country `json:"team2"`
	GroupNumber       int      `json:"groupNumber"`
	GroupName         string   `json:"groupName"`
	Playday           int      `json:"playday"`
	GoalsTeam1        int      `json:"goalsTeam1"`
	PenaltyScoreTeam1 int      `json:"penaltyScoreTeam1"`
	PenaltyScoreTeam2 int      `json:"penaltyScoreTeam2"`
	GoalsTeam2        int      `json:"goalsTeam2"`
	Winner            *Country `json:"winner"`
}

type TournamentMatches struct {
	Group   []Match `json:"group"`
	Sixteen []Match `json:"sixteen"`
	Eight   []Match `json:"eight"`
	Four    []Match `json:"four"`
	Final   Match   `json:"final"`
}

type PlaydayMatches []Match

type Group struct {
	name      string
	countries []*Country
	playplan  map[int]PlaydayMatches
	ranking   []*Country
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
	member1  *Country
	member2  *Country
	member3  *Country
	member4  *Country
	member5  *Country
	member6  *Country
	member7  *Country
	member8  *Country
	member9  *Country
	member10 *Country
	member11 *Country
	member12 *Country
	member13 *Country
	member14 *Country
	member15 *Country
	member16 *Country
}

func defineCountry(name string, strength int) *Country {
	country := Country{Name: name, Strength: strength, Points: 0, Goals: 0, PenaltyGoals: 0}
	return &country
}

func defineMatch(team1 *Country, team2 *Country, groupNumber int, groupName string, playday int) Match {
	return Match{
		Team1:       team1,
		Team2:       team2,
		GroupNumber: groupNumber,
		GroupName:   groupName,
		Playday:     playday,
		GoalsTeam1:  0,
		GoalsTeam2:  0,
		Winner:      nil,
	}
}

func GetGroupA() Group {
	italy := defineCountry("Italy", 3)
	switzerland := defineCountry("Switzerland", 2)
	turkey := defineCountry("Turkey", 2)
	wales := defineCountry("Wales", 1)
	countries := []*Country{italy, switzerland, turkey, wales}

	turkey_italy := defineMatch(turkey, italy, 1, "Group A", 1)
	wales_switzerland := defineMatch(wales, switzerland, 1, "Group A", 1)
	turkey_wales := defineMatch(turkey, wales, 1, "Group A", 3)
	italy_switzerland := defineMatch(italy, switzerland, 1, "Group A", 3)
	italy_wales := defineMatch(italy, wales, 1, "Group A", 5)
	switzerland_turkey := defineMatch(switzerland, turkey, 1, "Group A", 5)

	playplan := make(map[int]PlaydayMatches)
	playplan[1] = PlaydayMatches{turkey_italy, wales_switzerland}
	playplan[3] = PlaydayMatches{turkey_wales, italy_switzerland}
	playplan[5] = PlaydayMatches{italy_wales, switzerland_turkey}

	return Group{
		name:      "Group A",
		countries: countries,
		playplan:  playplan,
	}
}

func GetGroupB() Group {
	belgium := defineCountry("Belgium", 2)
	denmark := defineCountry("Denmark", 2)
	finland := defineCountry("Finland", 1)
	russia := defineCountry("Russia", 2)
	countries := []*Country{belgium, denmark, finland, russia}

	denmark_finland := defineMatch(denmark, finland, 2, "Group B", 1)
	belgium_russia := defineMatch(belgium, russia, 2, "Group B", 1)
	finland_russia := defineMatch(finland, russia, 2, "Group B", 3)
	denmark_belgium := defineMatch(denmark, belgium, 2, "Group B", 3)
	russia_denmark := defineMatch(russia, denmark, 2, "Group B", 5)
	finland_belgium := defineMatch(finland, belgium, 2, "Group B", 5)

	playplan := make(map[int]PlaydayMatches)
	playplan[1] = PlaydayMatches{denmark_finland, belgium_russia}
	playplan[3] = PlaydayMatches{finland_russia, denmark_belgium}
	playplan[5] = PlaydayMatches{russia_denmark, finland_belgium}

	return Group{
		name:      "Group B",
		countries: countries,
		playplan:  playplan,
	}
}

func GetGroupC() Group {
	netherlands := defineCountry("Netherlands", 3)
	northmazedonia := defineCountry("Northmazedonia", 1)
	ukraine := defineCountry("Ukraine", 1)
	austria := defineCountry("Austria", 2)
	countries := []*Country{netherlands, northmazedonia, ukraine, austria}

	austria_northmazedonia := defineMatch(austria, northmazedonia, 3, "Group C", 1)
	netherlands_ukraine := defineMatch(netherlands, ukraine, 3, "Group C", 1)
	ukraine_northmazedonia := defineMatch(ukraine, northmazedonia, 3, "Group C", 3)
	netherlands_austria := defineMatch(netherlands, austria, 3, "Group C", 3)
	ukraine_austria := defineMatch(ukraine, austria, 3, "Group C", 5)
	northmazedonia_netherlands := defineMatch(northmazedonia, netherlands, 3, "Group C", 5)

	playplan := make(map[int]PlaydayMatches)
	playplan[1] = PlaydayMatches{austria_northmazedonia, netherlands_ukraine}
	playplan[3] = PlaydayMatches{ukraine_northmazedonia, netherlands_austria}
	playplan[5] = PlaydayMatches{ukraine_austria, northmazedonia_netherlands}

	return Group{
		name:      "Group C",
		countries: countries,
		playplan:  playplan,
	}
}

func GetGroupD() Group {
	england := defineCountry("England", 3)
	kroatia := defineCountry("Kroatia", 2)
	scotland := defineCountry("Scotland", 2)
	czechRepublic := defineCountry("Czech Republic", 2)
	countries := []*Country{england, kroatia, scotland, czechRepublic}

	england_kroatia := defineMatch(england, kroatia, 4, "Group D", 1)
	scotland_czechRepublic := defineMatch(scotland, czechRepublic, 4, "Group D", 2)
	kroatia_czechRepublic := defineMatch(kroatia, czechRepublic, 4, "Group D", 4)
	england_scotland := defineMatch(england, scotland, 4, "Group D", 4)
	kroatia_scotland := defineMatch(kroatia, scotland, 4, "Group D", 6)
	czechRepublic_england := defineMatch(czechRepublic, england, 4, "Group D", 6)

	playday := make(map[int]PlaydayMatches)
	playday[1] = PlaydayMatches{england_kroatia}
	playday[2] = PlaydayMatches{scotland_czechRepublic}
	playday[4] = PlaydayMatches{kroatia_czechRepublic, england_scotland}
	playday[6] = PlaydayMatches{kroatia_scotland, czechRepublic_england}

	return Group{
		name:      "Group D",
		countries: countries,
		playplan:  playday,
	}
}

func GetGroupE() Group {
	poland := defineCountry("Poland", 2)
	sweden := defineCountry("Sweden", 2)
	slowakia := defineCountry("Slowakia", 2)
	spain := defineCountry("Spain", 1)
	countries := []*Country{poland, sweden, slowakia, spain}

	poland_slowakia := defineMatch(poland, slowakia, 5, "Group E", 2)
	spain_sweden := defineMatch(spain, sweden, 5, "Group E", 2)
	sweden_slowakia := defineMatch(sweden, slowakia, 5, "Group E", 4)
	spain_poland := defineMatch(spain, poland, 5, "Group E", 4)
	slowakia_spain := defineMatch(slowakia, spain, 5, "Group E", 6)
	sweden_poland := defineMatch(sweden, poland, 5, "Group E", 6)

	playday := make(map[int]PlaydayMatches)
	playday[2] = PlaydayMatches{poland_slowakia, spain_sweden}
	playday[4] = PlaydayMatches{sweden_slowakia, spain_poland}
	playday[6] = PlaydayMatches{slowakia_spain, sweden_poland}

	return Group{
		name:      "Group E",
		countries: countries,
		playplan:  playday,
	}
}

func GetGroupF() Group {
	germany := defineCountry("Germany", 3)
	france := defineCountry("France", 3)
	portugal := defineCountry("Portugal", 3)
	hungry := defineCountry("Hungry", 1)
	countries := []*Country{germany, france, portugal, hungry}

	hungry_portugal := defineMatch(hungry, portugal, 6, "Group F", 2)
	france_germany := defineMatch(france, germany, 6, "Group F", 2)
	hungry_france := defineMatch(hungry, france, 6, "Group F", 4)
	portugal_germany := defineMatch(portugal, germany, 6, "Group F", 4)
	portugal_france := defineMatch(portugal, france, 6, "Group F", 6)
	germany_hungry := defineMatch(germany, hungry, 6, "Group F", 6)

	playday := make(map[int]PlaydayMatches)
	playday[2] = PlaydayMatches{hungry_portugal, france_germany}
	playday[4] = PlaydayMatches{hungry_france, portugal_germany}
	playday[6] = PlaydayMatches{portugal_france, germany_hungry}

	return Group{
		name:      "Group F",
		countries: countries,
		playplan:  playday,
	}
}

func GetGroups() []Group {
	groupA := GetGroupA()
	groupB := GetGroupB()
	groupC := GetGroupC()
	groupD := GetGroupD()
	groupE := GetGroupE()
	groupF := GetGroupF()

	var groups = []Group{groupA, groupB, groupC, groupD, groupE, groupF}
	return groups

}

func determineGroupRanking(group Group) []*Country {
	countries := group.countries
	sort.Slice(countries, func(i, j int) bool {
		if countries[i].Points == countries[j].Points && countries[i].Goals > countries[j].Goals {
			return true
		}
		return countries[i].Points > countries[j].Points
	})

	return countries
}

func getBestFourThirds(thirds []*Country) [4]*Country {
	var thirdsSlice []*Country = thirds[:]
	sort.Slice(thirdsSlice, func(i, j int) bool {
		if thirdsSlice[i].Points == thirdsSlice[j].Points && thirdsSlice[i].Goals > thirdsSlice[j].Goals {
			return true
		}
		return thirdsSlice[i].Points > thirdsSlice[j].Points
	})
	var bestFourThirds = [4]*Country{thirdsSlice[0], thirdsSlice[1], thirdsSlice[2], thirdsSlice[3]}
	return bestFourThirds
}
