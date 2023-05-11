package simulator

import (
	"sort"
	"time"
)

type Country struct {
	Strength int    `json:"strength"`
	Name     string `json:"name"`
	Points   int    `json:"points"`
	Goals    int    `json:"goals"`
}

type Match struct {
	team1      *Country
	team2      *Country
	playtime   time.Time
	goalsTeam1 int
	goalsTeam2 int
	winner     *Country
}

type PlaydayMatches []Match

type Group struct {
	name      string
	countries []*Country
	playplan  map[int]PlaydayMatches
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
	country := Country{Name: name, Strength: strength, Points: 0, Goals: 0}
	return &country
}

func defineMatch(team1 *Country, team2 *Country) Match {
	return Match{
		team1:      team1,
		team2:      team2,
		playtime:   time.Now().UTC(),
		goalsTeam1: 0,
		goalsTeam2: 0,
		winner:     nil,
	}
}

func GetGroupA() Group {
	italy := defineCountry("Italy", 3)
	switzerland := defineCountry("Switzerland", 2)
	turkey := defineCountry("Turkey", 2)
	wales := defineCountry("Wales", 1)
	countries := []*Country{italy, switzerland, turkey, wales}

	turkey_italy := defineMatch(turkey, italy)
	wales_switzerland := defineMatch(wales, switzerland)
	turkey_wales := defineMatch(turkey, wales)
	italy_switzerland := defineMatch(italy, switzerland)
	italy_wales := defineMatch(italy, wales)
	switzerland_turkey := defineMatch(switzerland, turkey)

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

	denmark_finland := defineMatch(denmark, finland)
	belgium_russia := defineMatch(belgium, russia)
	finland_russia := defineMatch(finland, russia)
	denmark_belgium := defineMatch(denmark, belgium)
	russia_denmark := defineMatch(russia, denmark)
	finland_belgium := defineMatch(finland, belgium)

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

	austria_northmazedonia := defineMatch(austria, northmazedonia)
	netherlands_ukraine := defineMatch(netherlands, ukraine)
	ukraine_northmazedonia := defineMatch(ukraine, northmazedonia)
	netherlands_austria := defineMatch(netherlands, austria)
	ukraine_austria := defineMatch(ukraine, austria)
	northmazedonia_netherlands := defineMatch(northmazedonia, netherlands)

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

	england_kroatia := defineMatch(england, kroatia)
	scotland_czechRepublic := defineMatch(scotland, czechRepublic)
	kroatia_czechRepublic := defineMatch(kroatia, czechRepublic)
	england_scotland := defineMatch(england, scotland)
	kroatia_scotland := defineMatch(kroatia, scotland)
	czechRepublic_england := defineMatch(czechRepublic, england)

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

	poland_slowakia := defineMatch(poland, slowakia)
	spain_sweden := defineMatch(spain, sweden)
	sweden_slowakia := defineMatch(sweden, slowakia)
	spain_poland := defineMatch(spain, poland)
	slowakia_spain := defineMatch(slowakia, spain)
	sweden_poland := defineMatch(sweden, poland)

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

	hungry_portugal := defineMatch(hungry, portugal)
	france_germany := defineMatch(france, germany)
	hungry_france := defineMatch(hungry, france)
	portugal_germany := defineMatch(portugal, germany)
	portugal_france := defineMatch(portugal, france)
	germany_hungry := defineMatch(germany, hungry)

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

func determineGroupWinner(group Group) []*Country {
	groupRanking := make([]*Country, 0, 4)
	// add first
	groupRanking = append(groupRanking, group.countries[0])

	// add second
	if isBetterFirstCountry(group.countries[1], groupRanking[0]) {
		groupRanking = append([]*Country{group.countries[1]}, groupRanking...)
	} else {
		groupRanking = append(groupRanking, group.countries[1])
	}

	// add third
	if isBetterFirstCountry(group.countries[2], groupRanking[0]) {
		groupRanking = append([]*Country{group.countries[2]}, groupRanking...)
	} else if isBetterFirstCountry(group.countries[2], groupRanking[1]) {
		groupRanking = append([]*Country{groupRanking[0]}, group.countries[2], groupRanking[1])
	} else {
		groupRanking = append(groupRanking, group.countries[2])
	}

	// add fourth
	if isBetterFirstCountry(group.countries[3], groupRanking[0]) {
		groupRanking = append([]*Country{group.countries[3]}, groupRanking...)
	} else if isBetterFirstCountry(group.countries[3], groupRanking[1]) {
		groupRanking = append([]*Country{groupRanking[0]}, group.countries[3], groupRanking[1], groupRanking[2])
	} else if isBetterFirstCountry(group.countries[3], groupRanking[2]) {
		groupRanking = append([]*Country{groupRanking[0], groupRanking[1]}, group.countries[3], groupRanking[2])
	} else {
		groupRanking = append(groupRanking, group.countries[3])
	}
	return groupRanking
}

func isBetterFirstCountry(country1, country2 *Country) bool {
	if country1.Points == country2.Points && country1.Goals > country2.Goals {
		return true
	}
	return country1.Points > country2.Points
}

func getBestFourThirds(thirds [6]*Country) [4]*Country {
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

func GetRoudOfSixteen(groups Groups) RoundOf16 {
	groupAranked := determineGroupWinner(groups.A)
	groupBranked := determineGroupWinner(groups.B)
	groupCranked := determineGroupWinner(groups.C)
	groupDranked := determineGroupWinner(groups.D)
	groupEranked := determineGroupWinner(groups.E)
	groupFranked := determineGroupWinner(groups.F)
	allThirds := [6]*Country{groupAranked[2], groupBranked[2], groupCranked[2], groupDranked[2], groupEranked[2], groupFranked[2]}
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
