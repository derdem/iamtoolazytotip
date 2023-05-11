package main

import (
	"fmt"

	"github.com/derdem/iamtoolazytotip/simulator"
)

func main() {
	// api.Start()
	result := simulator.TournamentSimulator()
	fmt.Println(result)

}
