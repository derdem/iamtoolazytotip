package simulator

type CountryPair [2]Country
type Playday []CountryPair

func GetPlaydays(groups Groups) []Playday {

	playday1 := Playday{
		{groups.A.member3, groups.A.member1},
		{groups.A.member4, groups.A.member2},
		{groups.B.member2, groups.B.member3},
		{groups.B.member1, groups.B.member4},
		{groups.D.member1, groups.D.member2},
		{groups.C.member4, groups.C.member2},
		{groups.C.member1, groups.C.member3},
	}

	playday2 := Playday{
		{groups.D.member3, groups.D.member4},
		{groups.E.member1, groups.E.member3},
		{groups.E.member4, groups.E.member2},
		{groups.F.member4, groups.F.member3},
		{groups.F.member2, groups.F.member1},
	}

	playday3 := Playday{
		{groups.B.member3, groups.B.member4},
		{groups.A.member3, groups.A.member4},
		{groups.A.member1, groups.A.member2},
		{groups.C.member3, groups.C.member2},
		{groups.B.member2, groups.B.member1},
		{groups.C.member1, groups.C.member4},
	}

	playday4 := Playday{
		{groups.E.member2, groups.E.member3},
		{groups.D.member2, groups.D.member4},
		{groups.D.member1, groups.D.member3},
		{groups.F.member4, groups.F.member2},
		{groups.F.member3, groups.F.member1},
		{groups.E.member4, groups.E.member1},
	}

	playday5 := Playday{
		{groups.A.member1, groups.A.member4},
		{groups.A.member2, groups.A.member3},
		{groups.C.member3, groups.C.member4},
		{groups.C.member2, groups.C.member1},
		{groups.B.member4, groups.B.member2},
		{groups.B.member3, groups.B.member1},
	}

	playday6 := Playday{
		{groups.D.member2, groups.D.member3},
		{groups.D.member4, groups.D.member1},
		{groups.E.member3, groups.E.member4},
		{groups.E.member2, groups.E.member1},
		{groups.F.member3, groups.F.member2},
		{groups.F.member1, groups.F.member4},
	}

	playdays := []Playday{playday1, playday2, playday3, playday4, playday5, playday6}
	return playdays
}

// ko_rounds = [
//     Eighth,
//     Quarter,
//     Half,
//     Final
// ]
