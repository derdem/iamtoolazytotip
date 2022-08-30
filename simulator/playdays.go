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

// ko_rounds = [
//     Eighth,
//     Quarter,
//     Half,
//     Final
// ]
