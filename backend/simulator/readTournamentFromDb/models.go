package readTournamentFromDb

type Group struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	TournamentId int     `json:"tournament_id"`
	Teams        []Team  `json:"teams"`
	Matches      []Match `json:"matches"`
}

type Team struct {
	Id       int      `json:"id"`
	Name     string   `json:"name"`
	GroupId  int      `json:"group_id"`
	Strength Strength `json:"strength"`
}

type Match struct {
	Id                int `json:"id"`
	GroupId           int `json:"group_id"`
	Team1Id           int `json:"team1_id"`
	Team1Goals        int `json:"team1_goals"`
	Team1PenaltyGoals int `json:"team1_penalty_goals"`
	Team2Id           int `json:"team2_id"`
	Team2Goals        int `json:"team2_goals"`
	Team2PenaltyGoals int `json:"team2_penalty_goals"`
}

type Strength string

const (
	LowStrength    Strength = "low"
	MediumStrength Strength = "medium"
	HighStrength   Strength = "high"
)
