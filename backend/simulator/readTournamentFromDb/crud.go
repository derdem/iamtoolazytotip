package readTournamentFromDb

import (
	"time"

	"github.com/derdem/iamtoolazytotip/postgres_connection"
	"github.com/derdem/iamtoolazytotip/simulator"
	"github.com/jackc/pgx"
)

var GetConnection = postgres_connection.GetConnection

func PanicAtError(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadTournamentGrouped(tournament_id int) []GroupDb {
	conn := GetConnection()

	rows, err := conn.Query(`
        SELECT groups.*, json_agg(DISTINCT teams.*) as teams, json_agg(DISTINCT matches.*) as matches
        FROM groups
        INNER JOIN teams ON groups.id = teams.group_id
        INNER JOIN matches ON groups.id = matches.group_id
        WHERE tournament_id = $1 GROUP BY groups.id;
    `, tournament_id)
	PanicAtError(err)

	groups := []GroupDb{}

	defer rows.Close()
	for rows.Next() {
		var group GroupDb
		err := rows.Scan(
			&group.Id,
			&group.Name,
			&group.TournamentId,
			&group.Teams,
			&group.Matches,
		)
		PanicAtError(err)
		groups = append(groups, group)
	}

	return groups

}

func GetTournament(tournament_id int) simulator.Tournament {
	tournamentDb := ReadTournament(tournament_id)
	return ConvertTournamentDbToModel(tournamentDb)
}

func ConvertTournamentDbToModel(tournamentDb TournamentDb) simulator.Tournament {
	var groups []simulator.Group2
	for _, groupDb := range tournamentDb.Groups {
		groups = append(groups, simulator.Group2{
			Id:           groupDb.Id,
			Name:         groupDb.Name,
			TournamentId: groupDb.TournamentId,
		})
	}

	teamMap := make(map[int]simulator.Team)
	var teams []simulator.Team
	for _, teamDb := range tournamentDb.Teams {
		team := simulator.Team{
			Id:       teamDb.Id,
			Name:     teamDb.Name,
			GroupId:  teamDb.GroupId,
			Strength: simulator.ConvertStrengthStringToInt(teamDb.Strength),
		}
		teamMap[team.Id] = team
		teams = append(teams, team)
	}

	var matches []simulator.Match2
	for _, match := range tournamentDb.Matches {
		team1 := teamMap[match.Team1Id]
		team2 := teamMap[match.Team2Id]
		matches = append(matches, simulator.Match2{
			Id:      match.Id,
			Team1:   team1,
			Team2:   team2,
			GroupId: match.GroupId,
		})
	}

	return simulator.Tournament{
		Id:      tournamentDb.Id,
		Name:    tournamentDb.Name,
		Groups:  groups,
		Teams:   teams,
		Matches: matches,
	}

}

func ReadTournament(tournament_id int) TournamentDb {
	conn := GetConnection()
	tx, err := conn.Begin()
	PanicAtError(err)
	id, name := readTournament(tx, tournament_id)
	groups := readGroups(tx, tournament_id)
	group_ids := getGroupIds(groups)
	teams := readTeams(tx, group_ids)
	matches := readMatches(tx, group_ids)
	tx.Rollback()

	return TournamentDb{
		Id:      id,
		Name:    name,
		Groups:  groups,
		Teams:   teams,
		Matches: matches,
	}
}

func readTournament(tx *pgx.Tx, tournament_id int) (int, string) {
	var id int
	var name string
	err := tx.QueryRow(`
        SELECT id, name
        FROM tournaments WHERE id = $1;
    `, tournament_id).Scan(&id, &name)
	PanicAtError(err)
	return id, name
}

func readGroups(tx *pgx.Tx, tournament_id int) []GroupLightDb {
	groupRows, err := tx.Query(`
        SELECT groups.*
        FROM groups WHERE tournament_id = $1;
    `, tournament_id)
	PanicAtError(err)
	defer groupRows.Close()

	var groups []GroupLightDb
	for groupRows.Next() {
		var group GroupLightDb
		err := groupRows.Scan(
			&group.Id,
			&group.Name,
			&group.TournamentId,
		)
		PanicAtError(err)
		groups = append(groups, group)
	}

	return groups
}

func getGroupIds(groups []GroupLightDb) []int {
	groupIds := make([]int, len(groups))
	for i, group := range groups {
		groupIds[i] = group.Id
	}
	return groupIds
}

func readTeams(tx *pgx.Tx, group_ids []int) []TeamDb {
	teamRows, err := tx.Query(`
        SELECT id, name, group_id, strength
        FROM teams WHERE group_id = ANY($1);
    `, group_ids)
	PanicAtError(err)
	defer teamRows.Close()

	var teams []TeamDb
	for teamRows.Next() {
		var team TeamDb
		err := teamRows.Scan(
			&team.Id,
			&team.Name,
			&team.GroupId,
			&team.Strength,
		)
		PanicAtError(err)
		teams = append(teams, team)
	}

	return teams
}

func readMatches(tx *pgx.Tx, group_ids []int) []MatchDb {
	matchRows, err := tx.Query(`
        SELECT id, group_id, team1_id, team2_id
        FROM matches WHERE group_id = ANY($1);
    `, group_ids)
	PanicAtError(err)
	defer matchRows.Close()

	var matches []MatchDb
	for matchRows.Next() {
		var match MatchDb
		err := matchRows.Scan(
			&match.Id,
			&match.GroupId,
			&match.Team1Id,
			&match.Team2Id,
		)
		PanicAtError(err)
		matches = append(matches, match)
	}

	return matches
}

func LoadGroupFromDb(tournament_id int) []simulator.Group {
	groupsRaw := ReadTournamentGrouped(tournament_id)
	groups := make([]simulator.Group, 0)
	for _, groupRaw := range groupsRaw {
		countries, countryIdIndexMap := getCountriesFromDbEntries(groupRaw.Teams)
		matches := getMatchesFromDbEntries(groupRaw.Matches, countryIdIndexMap, countries, groupRaw.Name)

		group := simulator.Group{
			Name:      groupRaw.Name,
			Countries: countries,
			Matches:   matches,
		}
		groups = append(groups, group)
	}
	return groups
}

func LoadTournamentFromDb(tournament_id int) TournamentDb {
	tournament := ReadTournament(tournament_id)
	return tournament
}

func getCountriesFromDbEntries(teams []TeamDb) ([4]*simulator.Country, map[int]int) {
	return [4]*simulator.Country{
			{
				Name:     teams[0].Name,
				Strength: simulator.ConvertStrengthStringToInt(teams[0].Strength),
				Points:   0,
				Goals:    0,
			},
			{
				Name:     teams[1].Name,
				Strength: simulator.ConvertStrengthStringToInt(teams[1].Strength),
				Points:   0,
				Goals:    0,
			},
			{
				Name:     teams[2].Name,
				Strength: simulator.ConvertStrengthStringToInt(teams[2].Strength),
				Points:   0,
				Goals:    0,
			},
			{
				Name:     teams[3].Name,
				Strength: simulator.ConvertStrengthStringToInt(teams[3].Strength),
				Points:   0,
				Goals:    0,
			},
		},
		map[int]int{
			teams[0].Id: 0,
			teams[1].Id: 1,
			teams[2].Id: 2,
			teams[3].Id: 3,
		}
}

func getMatchesFromDbEntries(matches []MatchDb, countryIdIndexMap map[int]int, countries [4]*simulator.Country, groupName string) []simulator.GroupMatch {
	groupMatches := make([]simulator.GroupMatch, 0)
	for _, match := range matches {
		country1Index := countryIdIndexMap[match.Team1Id]
		country2Index := countryIdIndexMap[match.Team2Id]
		country1 := countries[country1Index]
		country2 := countries[country2Index]

		groupMatch := simulator.GroupMatch{
			Match: simulator.Match{
				Team1:             country1,
				Team2:             country2,
				Playtime:          time.Now(),
				GoalsTeam1:        0,
				PenaltyScoreTeam1: 0,
				GoalsTeam2:        0,
				PenaltyScoreTeam2: 0,
				Winner:            nil,
			},
			GroupName: groupName,
		}
		groupMatches = append(groupMatches, groupMatch)
	}

	return groupMatches
}
