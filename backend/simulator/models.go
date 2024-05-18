package simulator

type Tournament struct {
	Id                    int                     `json:"id"`
	Name                  string                  `json:"name"`
	Groups                []Group                 `json:"groups"`
	Teams                 []Team                  `json:"teams"`
	Matches               []Match                 `json:"matches"`
	MatchResults          []MatchResult           `json:"matchResults"`
	GroupRankings         []GroupRanking          `json:"groupRankings"`
	KoMatches             []KoMatch               `json:"koMatches"`
	ThirdsEvaluationRules []ThirdsEvaluationRules `json:"thirdsEvaluationRules"`
	Winner                Team                    `json:"winner"`
}

type Group struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	TournamentId int       `json:"tournamentId"`
	GroupType    GroupType `json:"groupType"`
}

type Team struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	GroupId  int    `json:"groupId"`
	Strength int    `json:"strength"`
}

type Match struct {
	Id      int  `json:"id"`
	GroupId int  `json:"groupId"`
	Team1   Team `json:"team1"`
	Team2   Team `json:"team2"`
}

type KoMatch struct {
	Id       int `json:"id"`
	GroupId  int `json:"group_id"`
	GroupId1 int `json:"group_id1"`
	GroupId2 int `json:"group_id2"`
	Ranking1 int `json:"ranking1"`
	Ranking2 int `json:"ranking2"`
}

type MatchResult struct {
	Match             Match `json:"match"`
	Team1Goals        int   `json:"team1Goals"`
	Team2Goals        int   `json:"team2Goals"`
	Team1PenaltyGoals int   `json:"team1PenaltyGoals"`
	Team2PenaltyGoals int   `json:"team2PenaltyGoals"`
	Team1PointsGained int   `json:"team1PointsGained"`
	Team2PointsGained int   `json:"team2PointsGained"`
	Winner            Team  `json:"winner"`
}

type GroupRanking struct {
	GroupId int  `json:"groupId"`
	Team    Team `json:"team"`
	Ranking int  `json:"ranking"`
	Points  int  `json:"points"`
	Goals   int  `json:"goals"`
}

type ThirdsEvaluationRules struct {
	TournamentId             int   `json:"tournament_id"`
	BestFourTeamsId          int   `json:"best_four_teams_id"`
	BestFourTeamsArrangement []int `json:"best_four_teams_arrangement"`
}

type OutcomeProbabilities struct {
	Team1 float64
	Team2 float64
	Remis float64
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
