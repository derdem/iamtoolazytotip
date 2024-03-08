package customtournament

import (
	"encoding/json"
	"log"

	"github.com/derdem/iamtoolazytotip/simulator"
)

func RunCustomTournament(payload []byte) {
	var groupsAndMatches simulator.GroupsAndMatches
	json.Unmarshal(payload, &groupsAndMatches)
	log.Println(groupsAndMatches)
}
