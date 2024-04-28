package simulator

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func ReadTournamentFromFile(path string) Tournament {
	jsonFile, err := os.Open(path)

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}

	defer jsonFile.Close()
	jsonBytes, err := io.ReadAll(jsonFile)

	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}
	var Tournament Tournament
	json.Unmarshal(jsonBytes, &Tournament)

	return Tournament
}
