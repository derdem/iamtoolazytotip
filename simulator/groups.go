package simulator

type Group struct {
	name    string
	member1 Country
	member2 Country
	member3 Country
	member4 Country
}

type Groups struct {
	A Group
	B Group
	C Group
	D Group
	E Group
	F Group
}

func GetGroups(allCountries AllCountries) Groups {
	groupA := Group{
		name:    "Group A",
		member1: allCountries.italy,
		member2: allCountries.switzerland,
		member3: allCountries.turkey,
		member4: allCountries.wales,
	}
	groupB := Group{
		name:    "Group B",
		member1: allCountries.belgium,
		member2: allCountries.denmark,
		member3: allCountries.finland,
		member4: allCountries.russia,
	}

	groupC := Group{
		name:    "Group C",
		member1: allCountries.netherlands,
		member2: allCountries.northmazedonia,
		member3: allCountries.ukraine,
		member4: allCountries.austria,
	}

	groupD := Group{
		name:    "Group D",
		member1: allCountries.england,
		member2: allCountries.kroatia,
		member3: allCountries.scotland,
		member4: allCountries.czechrepublic,
	}

	groupE := Group{
		name:    "Group E",
		member1: allCountries.poland,
		member2: allCountries.sweden,
		member3: allCountries.slowakia,
		member4: allCountries.spain,
	}

	groupF := Group{
		name:    "Group F",
		member1: allCountries.germany,
		member2: allCountries.france,
		member3: allCountries.portugal,
		member4: allCountries.hungry,
	}

	var groups = Groups{A: groupA, B: groupB, C: groupC, D: groupD, E: groupE, F: groupF}
	return groups

}
