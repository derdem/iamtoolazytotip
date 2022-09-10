package simulator

type CountryPair [2]Country
type Playday []CountryPair

func GetPlaydays() []Playday {

	allCountries := GetAllCountries()

	playday1 := Playday{
		{allCountries.turkey, allCountries.italy},
		{allCountries.wales, allCountries.switzerland},
		{allCountries.denmark, allCountries.finland},
		{allCountries.belgium, allCountries.russia},
		{allCountries.england, allCountries.kroatia},
		{allCountries.austria, allCountries.northmazedonia},
		{allCountries.netherlands, allCountries.ukraine},
	}

	playday2 := Playday{
		{allCountries.scotland, allCountries.czechrepublic},
		{allCountries.poland, allCountries.slowakia},
		{allCountries.spain, allCountries.sweden},
		{allCountries.hungry, allCountries.portugal},
		{allCountries.france, allCountries.germany},
	}

	playday3 := Playday{
		{allCountries.finland, allCountries.russia},
		{allCountries.turkey, allCountries.wales},
		{allCountries.italy, allCountries.switzerland},
		{allCountries.ukraine, allCountries.northmazedonia},
		{allCountries.denmark, allCountries.belgium},
		{allCountries.netherlands, allCountries.austria},
	}

	playday4 := Playday{
		{allCountries.sweden, allCountries.slowakia},
		{allCountries.kroatia, allCountries.czechrepublic},
		{allCountries.england, allCountries.scotland},
		{allCountries.hungry, allCountries.france},
		{allCountries.portugal, allCountries.germany},
		{allCountries.spain, allCountries.poland},
	}

	playday5 := Playday{
		{allCountries.italy, allCountries.wales},
		{allCountries.switzerland, allCountries.turkey},
		{allCountries.ukraine, allCountries.austria},
		{allCountries.northmazedonia, allCountries.netherlands},
		{allCountries.russia, allCountries.denmark},
		{allCountries.finland, allCountries.belgium},
	}

	playday6 := Playday{
		{allCountries.kroatia, allCountries.scotland},
		{allCountries.czechrepublic, allCountries.england},
		{allCountries.slowakia, allCountries.spain},
		{allCountries.sweden, allCountries.poland},
		{allCountries.portugal, allCountries.france},
		{allCountries.germany, allCountries.hungry},
	}

	playdays := []Playday{playday1, playday2, playday3, playday4, playday5, playday6}
	return playdays
}

func getRoundOf16Matches(groups Groups) [8][2]Country {
	roundOf16 := GetRoudOfSixteen(groups)
	return [8][2]Country{
		{roundOf16.member1, roundOf16.member2},
		{roundOf16.member3, roundOf16.member4},
		{roundOf16.member5, roundOf16.member6},
		{roundOf16.member7, roundOf16.member8},
		{roundOf16.member9, roundOf16.member10},
		{roundOf16.member11, roundOf16.member12},
		{roundOf16.member13, roundOf16.member14},
		{roundOf16.member15, roundOf16.member16},
	}
}

// [self.teams[5], self.teams[4]],
// [self.teams[3], self.teams[1]],
// [self.teams[2], self.teams[0]],
// [self.teams[7], self.teams[6]],

func getRoundOf8Matches(teams [8]Country) [4][2]Country {
	return [4][2]Country{
		{teams[5], teams[4]},
		{teams[3], teams[1]},
		{teams[2], teams[0]},
		{teams[7], teams[6]},
	}
}

func getRoundOf4Matches(teams [4]Country) [2][2]Country {
	return [2][2]Country{
		{teams[0], teams[1]},
		{teams[2], teams[3]},
	}
}

// ko_rounds = [
//     Eighth,
//     Quarter,
//     Half,
//     Final
// ]
