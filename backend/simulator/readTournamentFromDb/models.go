package readTournamentFromDb

type Group struct {
	Id           int
	Name         string
	TournamentId int
	Teams        []Team
	Matches      []Match
}

type Team struct {
	Id       int
	Name     string
	GroupId  int
	Strength string
}

type Match struct {
	Id                int
	GroupId           int
	Team1Id           int
	Team1Goals        int
	Team1PenaltyGoals int
	Team2Id           int
	Team2Goals        int
	Team2PenaltyGoals int
}
