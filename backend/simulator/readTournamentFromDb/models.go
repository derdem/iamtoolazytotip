package readTournamentFromDb

import "github.com/derdem/iamtoolazytotip/simulator"

type TournamentDb struct {
	Id        int            `json:"id"`
	Name      string         `json:"name"`
	Groups    []GroupLightDb `json:"groups"`
	Teams     []TeamDb       `json:"teams"`
	Matches   []MatchDb      `json:"matches"`
	KoMatches []KoMatchDb    `json:"ko_matches"`
}

type GroupDb struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	TournamentId int       `json:"tournament_id"`
	Teams        []TeamDb  `json:"teams"`
	Matches      []MatchDb `json:"matches"`
}

type GroupLightDb struct {
	Id           int                 `json:"id"`
	Name         string              `json:"name"`
	TournamentId int                 `json:"tournament_id"`
	GroupType    simulator.GroupType `json:"group_type"`
}

type TeamDb struct {
	Id       int                `json:"id"`
	Name     string             `json:"name"`
	GroupId  int                `json:"group_id"`
	Strength simulator.Strength `json:"strength"`
}

type MatchDb struct {
	Id      int `json:"id"`
	GroupId int `json:"group_id"`
	Team1Id int `json:"team1_id"`
	Team2Id int `json:"team2_id"`
}

type KoMatchDb struct {
	Id       int `json:"id"`
	GroupId  int `json:"group_id"`
	GroupId1 int `json:"group_id1"`
	GroupId2 int `json:"group_id2"`
	Ranking1 int `json:"ranking1"`
	Ranking2 int `json:"ranking2"`
}
