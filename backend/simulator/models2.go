package simulator

type Tournament struct {
	Id      int
	Name    string
	Groups  []Group2
	Teams   []Team
	Matches []Match2
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
	Id    int
	Team1 Team
	Team2 Team
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
