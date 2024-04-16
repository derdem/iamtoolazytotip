package simulator

type Tournament struct {
	Id            int
	Name          string
	Groups        []Group2 // need to distinguish between groups and ko rounds
	Teams         []Team
	Matches       []Match2
	MatchResults  []MatchResult
	GroupRankings []GroupRanking
	KoMatches     []KoMatch
}

type Group2 struct {
	Id           int
	Name         string
	TournamentId int
}

type Team struct {
	Id       int
	Name     string
	GroupId  int
	Strength int
}

type Match2 struct {
	Id      int
	GroupId int
	Team1   Team
	Team2   Team
}

type KoMatch struct {
	Id       int
	GroupId  int
	Group1   Group2
	Group2   Group2
	ranking1 int
	ranking2 int
}

type MatchResult struct {
	Match             Match2
	Team1Goals        int
	Team2Goals        int
	Team1PenaltyGoals int
	Team2PenaltyGoals int
	Team1PointsGained int
	Team2PointsGained int
	Winner            Team
}

type GroupRanking struct {
	GroupId int
	Team    Team
	Ranking int
	Points  int
	Goals   int
}
