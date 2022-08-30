package simulator

type Country struct {
	Strength int    `json:"strength"`
	Name     string `json:"name"`
	Points   int    `json:"points"`
	Goals    int    `json:"goals"`
}

const italyName = "Italy"
const switzerlandName = "Switzerland"
const turkeyName = "Turkey"
const walesName = "Wales"
const belgiumName = "Belgium"
const denmarkName = "Denmark"
const finlandName = "Finland"
const russiaName = "Russia"
const netherlandsName = "Netherlands"
const northmazedoniaName = "Northmazedonia"
const ukraineName = "Ukraine"
const austriaName = "Austria"
const englandName = "England"
const kroatiaName = "Kroatia"
const scotlandName = "Scotland"
const czechRepublicName = "Czech Republic"
const polandName = "Poland"
const swedenName = "Sweden"
const slowakiaName = "Slowakia"
const spainName = "Spain"
const germanyName = "Germany"
const franceName = "France"
const portugalName = "Portugal"
const hungryName = "Hungry"

type AllCountries struct {
	italy          Country
	switzerland    Country
	turkey         Country
	wales          Country
	belgium        Country
	denmark        Country
	finland        Country
	russia         Country
	netherlands    Country
	northmazedonia Country
	ukraine        Country
	austria        Country
	england        Country
	kroatia        Country
	scotland       Country
	czechrepublic  Country
	poland         Country
	sweden         Country
	slowakia       Country
	spain          Country
	germany        Country
	france         Country
	portugal       Country
	hungry         Country
}

func defineCountry(name string, strength int) Country {
	country := Country{Name: name, Strength: strength, Points: 0, Goals: 0}
	return country
}

func GetAllCountries() AllCountries {
	allCountries := AllCountries{
		italy:          defineCountry(italyName, 3),
		switzerland:    defineCountry(switzerlandName, 2),
		turkey:         defineCountry(turkeyName, 2),
		wales:          defineCountry(walesName, 1),
		belgium:        defineCountry(belgiumName, 2),
		denmark:        defineCountry(denmarkName, 2),
		finland:        defineCountry(finlandName, 1),
		russia:         defineCountry(russiaName, 2),
		netherlands:    defineCountry(netherlandsName, 3),
		northmazedonia: defineCountry(northmazedoniaName, 1),
		ukraine:        defineCountry(ukraineName, 1),
		austria:        defineCountry(austriaName, 2),
		england:        defineCountry(englandName, 3),
		kroatia:        defineCountry(kroatiaName, 2),
		scotland:       defineCountry(scotlandName, 2),
		czechrepublic:  defineCountry(czechRepublicName, 2),
		poland:         defineCountry(polandName, 2),
		sweden:         defineCountry(swedenName, 2),
		slowakia:       defineCountry(slowakiaName, 1),
		spain:          defineCountry(spainName, 3),
		germany:        defineCountry(germanyName, 3),
		france:         defineCountry(franceName, 3),
		portugal:       defineCountry(portugalName, 3),
		hungry:         defineCountry(hungryName, 1),
	}

	return allCountries
}

func UpdateCountry(allCountries *AllCountries, country Country) {
	if country.Name == italyName {
		allCountries.italy = Country{country.Strength, country.Name, allCountries.italy.Points + country.Points, allCountries.italy.Goals + country.Goals}
	} else if country.Name == switzerlandName {
		allCountries.switzerland = Country{country.Strength, country.Name, allCountries.switzerland.Points + country.Points, allCountries.switzerland.Goals + country.Goals}
	} else if country.Name == turkeyName {
		allCountries.turkey = Country{country.Strength, country.Name, allCountries.turkey.Points + country.Points, allCountries.turkey.Goals + country.Goals}
	} else if country.Name == walesName {
		allCountries.wales = Country{country.Strength, country.Name, allCountries.wales.Points + country.Points, allCountries.wales.Goals + country.Goals}
	} else if country.Name == belgiumName {
		allCountries.belgium = Country{country.Strength, country.Name, allCountries.belgium.Points + country.Points, allCountries.belgium.Goals + country.Goals}
	} else if country.Name == denmarkName {
		allCountries.denmark = Country{country.Strength, country.Name, allCountries.denmark.Points + country.Points, allCountries.denmark.Goals + country.Goals}
	} else if country.Name == finlandName {
		allCountries.finland = Country{country.Strength, country.Name, allCountries.finland.Points + country.Points, allCountries.finland.Goals + country.Goals}
	} else if country.Name == russiaName {
		allCountries.russia = Country{country.Strength, country.Name, allCountries.russia.Points + country.Points, allCountries.russia.Goals + country.Goals}
	} else if country.Name == netherlandsName {
		allCountries.netherlands = Country{country.Strength, country.Name, allCountries.netherlands.Points + country.Points, allCountries.netherlands.Goals + country.Goals}
	} else if country.Name == northmazedoniaName {
		allCountries.northmazedonia = Country{country.Strength, country.Name, allCountries.northmazedonia.Points + country.Points, allCountries.northmazedonia.Goals + country.Goals}
	} else if country.Name == ukraineName {
		allCountries.ukraine = Country{country.Strength, country.Name, allCountries.ukraine.Points + country.Points, allCountries.ukraine.Goals + country.Goals}
	} else if country.Name == austriaName {
		allCountries.austria = Country{country.Strength, country.Name, allCountries.austria.Points + country.Points, allCountries.austria.Goals + country.Goals}
	} else if country.Name == englandName {
		allCountries.england = Country{country.Strength, country.Name, allCountries.england.Points + country.Points, allCountries.england.Goals + country.Goals}
	} else if country.Name == kroatiaName {
		allCountries.kroatia = Country{country.Strength, country.Name, allCountries.kroatia.Points + country.Points, allCountries.kroatia.Goals + country.Goals}
	} else if country.Name == scotlandName {
		allCountries.scotland = Country{country.Strength, country.Name, allCountries.scotland.Points + country.Points, allCountries.scotland.Goals + country.Goals}
	} else if country.Name == czechRepublicName {
		allCountries.czechrepublic = Country{country.Strength, country.Name, allCountries.czechrepublic.Points + country.Points, allCountries.czechrepublic.Goals + country.Goals}
	} else if country.Name == polandName {
		allCountries.poland = Country{country.Strength, country.Name, allCountries.poland.Points + country.Points, allCountries.poland.Goals + country.Goals}
	} else if country.Name == swedenName {
		allCountries.sweden = Country{country.Strength, country.Name, allCountries.sweden.Points + country.Points, allCountries.sweden.Goals + country.Goals}
	} else if country.Name == slowakiaName {
		allCountries.slowakia = Country{country.Strength, country.Name, allCountries.slowakia.Points + country.Points, allCountries.slowakia.Goals + country.Goals}
	} else if country.Name == spainName {
		allCountries.spain = Country{country.Strength, country.Name, allCountries.spain.Points + country.Points, allCountries.spain.Goals + country.Goals}
	} else if country.Name == germanyName {
		allCountries.germany = Country{country.Strength, country.Name, allCountries.germany.Points + country.Points, allCountries.germany.Goals + country.Goals}
	} else if country.Name == franceName {
		allCountries.france = Country{country.Strength, country.Name, allCountries.france.Points + country.Points, allCountries.france.Goals + country.Goals}
	} else if country.Name == portugalName {
		allCountries.portugal = Country{country.Strength, country.Name, allCountries.portugal.Points + country.Points, allCountries.portugal.Goals + country.Goals}
	} else if country.Name == hungryName {
		allCountries.hungry = Country{country.Strength, country.Name, allCountries.hungry.Points + country.Points, allCountries.hungry.Goals + country.Goals}
	}
}
