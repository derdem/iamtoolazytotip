package simulator

type Country struct {
	Strength int    `json:"strength"`
	Name     string `json:"name"`
	Points   int    `json:"points"`
	Goals    int    `json:"goals"`
}

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
		italy:          defineCountry("Italy", 3),
		switzerland:    defineCountry("Switzerland", 2),
		turkey:         defineCountry("Turkey", 2),
		wales:          defineCountry("Wales", 1),
		belgium:        defineCountry("Belgium", 2),
		denmark:        defineCountry("Denmark", 2),
		finland:        defineCountry("Finland", 1),
		russia:         defineCountry("Russia", 2),
		netherlands:    defineCountry("Netherlands", 3),
		northmazedonia: defineCountry("Northmazedonia", 1),
		ukraine:        defineCountry("Ukraine", 1),
		austria:        defineCountry("Austria", 2),
		england:        defineCountry("England", 3),
		kroatia:        defineCountry("Kroatia", 2),
		scotland:       defineCountry("Scotland", 2),
		czechrepublic:  defineCountry("Czech Republic", 2),
		poland:         defineCountry("Poland", 2),
		sweden:         defineCountry("Sweden", 2),
		slowakia:       defineCountry("Slowakia", 1),
		spain:          defineCountry("Spain", 3),
		germany:        defineCountry("Germany", 3),
		france:         defineCountry("France", 3),
		portugal:       defineCountry("Portugal", 3),
		hungry:         defineCountry("Hungry", 1),
	}

	return allCountries
}
