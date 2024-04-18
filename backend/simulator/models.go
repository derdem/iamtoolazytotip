package simulator

import "time"

type OutcomeProbabilities struct {
	Team1 float64
	Team2 float64
	Remis float64
}

type MatchOutcome struct {
	Team1      Country `json:"team1"`
	Team1Score int     `json:"team1Score"`
	Team2      Country `json:"team2"`
	Team2Score int     `json:"team2Score"`
}

type Country struct {
	Strength     int    `json:"strength"`
	Name         string `json:"name"`
	Points       int    `json:"points"`
	Goals        int    `json:"goals"`
	PenaltyGoals int    `json:"penaltyGoals"`
}

type Match struct {
	Team1             *Country  `json:"team1"`
	Team2             *Country  `json:"team2"`
	Playtime          time.Time `json:"playtime"`
	GoalsTeam1        int       `json:"goalsTeam1"`
	PenaltyScoreTeam1 int       `json:"penaltyScoreTeam1"`
	PenaltyScoreTeam2 int       `json:"penaltyScoreTeam2"`
	GoalsTeam2        int       `json:"goalsTeam2"`
	Winner            *Country  `json:"winner"`
}

type GroupMatch struct {
	Match
	GroupName string `json:"groupName"`
}

type MatchDeserialized struct {
	Team1Index int       `json:"team1"`
	Team2Index int       `json:"team2"`
	Playtime   time.Time `json:"playtime"`
}

type TournamentMatches struct {
	Group   []GroupMatch `json:"group"`
	Sixteen []Match      `json:"sixteen"`
	Eight   []Match      `json:"eight"`
	Four    []Match      `json:"four"`
	Final   Match        `json:"final"`
}

type PlaydayMatches []Match

type Group struct {
	Name      string
	Countries [4]*Country
	Matches   []GroupMatch
	Ranking   []*Country
}

type GroupDeserialisation struct {
	Name      string              `json:"name"`
	Countries [4]*Country         `json:"countries"`
	Matches   []MatchDeserialized `json:"matches"`
}

type Groups struct {
	A Group
	B Group
	C Group
	D Group
	E Group
	F Group
}

type RoundOf16 struct {
	Member1  *Country
	Member2  *Country
	Member3  *Country
	Member4  *Country
	Member5  *Country
	Member6  *Country
	Member7  *Country
	Member8  *Country
	Member9  *Country
	Member10 *Country
	Member11 *Country
	Member12 *Country
	Member13 *Country
	Member14 *Country
	Member15 *Country
	Member16 *Country
}

type CountryFrontend struct {
	Name     string `json:"name"`
	Strength int    `json:"strength"`
}

type GroupFrontend struct {
	Name      string             `json:"groupName"`
	Countries [4]CountryFrontend `json:"countries"`
}

type MatchFrontend struct {
	GroupIndex int    `json:"groupIndex"`
	MatchIndex int    `json:"matchIndex"`
	Country1   string `json:"country1"`
	Country2   string `json:"country2"`
}

type GroupsAndMatches struct {
	Groups  []GroupFrontend `json:"groups"`
	Matches []MatchFrontend `json:"matches"`
}

type Strength string

const (
	LowStrength    Strength = "low"
	MediumStrength Strength = "medium"
	HighStrength   Strength = "high"
)

type GroupType string

const (
	GroupPhaseGroupType GroupType = "group_phase"
	KoPhaseGroupType    GroupType = "knockout_phase"
	SupportGroupType    GroupType = "support"
)
