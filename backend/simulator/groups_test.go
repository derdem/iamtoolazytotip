package simulator_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/derdem/iamtoolazytotip/simulator"
)

func TestSorting(t *testing.T) {
	italy := simulator.Country{
		Name:     "Italy",
		Strength: 3,
		Points:   7,
		Goals:    12,
	}

	switzerland := simulator.Country{
		Name:     "Switzerland",
		Strength: 2,
		Points:   5,
		Goals:    6,
	}

	turkey := simulator.Country{
		Name:     "Turkey",
		Strength: 2,
		Points:   5,
		Goals:    5,
	}

	wales := simulator.Country{
		Name:     "Wales",
		Strength: 1,
		Points:   3,
		Goals:    2,
	}

	countries := []simulator.Country{wales, switzerland, italy, turkey}

	fmt.Println(countries)

	sort.Slice(countries, func(i, j int) bool {
		if countries[i].Points == countries[j].Points && countries[i].Goals > countries[j].Goals {
			return true
		}
		return countries[i].Points > countries[j].Points
	})

	fmt.Println(countries)

}
