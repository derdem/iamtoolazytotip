package readTournamentFromDb

import (
	"github.com/derdem/iamtoolazytotip/postgres_connection"
	"github.com/derdem/iamtoolazytotip/simulator"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/pgtype"
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
	var groups []simulator.Group
	for _, groupDb := range tournamentDb.Groups {
		groups = append(groups, simulator.Group{
			Id:           groupDb.Id,
			Name:         groupDb.Name,
			TournamentId: groupDb.TournamentId,
			GroupType:    simulator.GroupType(groupDb.GroupType),
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

	var matches []simulator.Match
	for _, match := range tournamentDb.Matches {
		team1 := teamMap[match.Team1Id]
		team2 := teamMap[match.Team2Id]
		matches = append(matches, simulator.Match{
			Id:      match.Id,
			Team1:   team1,
			Team2:   team2,
			GroupId: match.GroupId,
		})
	}

	var koMatches []simulator.KoMatch
	for _, koMatch := range tournamentDb.KoMatches {
		koMatches = append(koMatches, simulator.KoMatch{
			Id:       koMatch.Id,
			GroupId:  koMatch.GroupId,
			GroupId1: koMatch.GroupId1,
			GroupId2: koMatch.GroupId2,
			Ranking1: koMatch.Ranking1,
			Ranking2: koMatch.Ranking2,
		})
	}

	var thirdsEvaluationRules []simulator.ThirdsEvaluationRules
	for _, thirdEvaluationRule := range tournamentDb.thirdsEvaluationRules {
		thirdsEvaluationRules = append(thirdsEvaluationRules, simulator.ThirdsEvaluationRules{
			TournamentId:             thirdEvaluationRule.TournamentId,
			BestFourTeamsId:          thirdEvaluationRule.BestFourTeamsId,
			BestFourTeamsArrangement: thirdEvaluationRule.BestFourTeamsArrangement,
		})
	}

	return simulator.Tournament{
		Id:                    tournamentDb.Id,
		Name:                  tournamentDb.Name,
		Groups:                groups,
		Teams:                 teams,
		Matches:               matches,
		KoMatches:             koMatches,
		ThirdsEvaluationRules: thirdsEvaluationRules,
	}

}

func ReadTournament(tournament_id int) TournamentDb {
	conn := GetConnection()
	tx, err := conn.Begin()
	PanicAtError(err)
	id, name := readTournament(tx, tournament_id)
	groups := readGroups(tx, tournament_id)
	group_ids := getGroupIdsMap(groups)
	teams := readTeams(tx, group_ids[simulator.GroupPhaseGroupType])
	matches := readMatches(tx, group_ids[simulator.GroupPhaseGroupType])
	koMatches := readKoMatches(tx, group_ids[simulator.KoPhaseGroupType])
	thirdsEvaluationRules := readThirdEvaluationRules(tx, tournament_id)
	tx.Rollback()

	return TournamentDb{
		Id:                    id,
		Name:                  name,
		Groups:                groups,
		Teams:                 teams,
		Matches:               matches,
		KoMatches:             koMatches,
		thirdsEvaluationRules: thirdsEvaluationRules,
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
			&group.GroupType,
		)
		PanicAtError(err)
		groups = append(groups, group)
	}

	return groups
}

func getGroupIdsMap(groups []GroupLightDb) map[simulator.GroupType][]int {
	groupIds := make(map[simulator.GroupType][]int, len(groups))
	for _, group := range groups {
		groupIds[group.GroupType] = append(groupIds[group.GroupType], group.Id)
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

func readKoMatches(tx *pgx.Tx, group_ids []int) []KoMatchDb {
	matchRows, err := tx.Query(`
        SELECT id, group_id, COALESCE(group_id1, 0), COALESCE(group_id2, 0), ranking1, ranking2
        FROM ko_matches WHERE group_id = ANY($1);
    `, group_ids)
	PanicAtError(err)
	defer matchRows.Close()

	var matches []KoMatchDb
	for matchRows.Next() {
		var match KoMatchDb
		err := matchRows.Scan(
			&match.Id,
			&match.GroupId,
			&match.GroupId1,
			&match.GroupId2,
			&match.Ranking1,
			&match.Ranking2,
		)
		PanicAtError(err)
		matches = append(matches, match)
	}

	return matches
}

func readThirdEvaluationRules(tx *pgx.Tx, tournament_id int) []ThirdEvaluationRulesDb {
	thirdsEvaluationRulesRows, err := tx.Query(`
    SELECT tournament_id, best_four_teams_id, best_four_teams_arrangement
    FROM thirds_evaluation_rules WHERE tournament_id = $1;
    `, tournament_id)
	PanicAtError(err)
	defer thirdsEvaluationRulesRows.Close()

	var thirdEvaluationRules []ThirdEvaluationRulesDb
	var bestFourTeamsArrangementPgtype pgtype.Int4Array
	for thirdsEvaluationRulesRows.Next() {
		var thirdEvaluationRule ThirdEvaluationRulesDb
		err := thirdsEvaluationRulesRows.Scan(
			&thirdEvaluationRule.TournamentId,
			&thirdEvaluationRule.BestFourTeamsId,
			&bestFourTeamsArrangementPgtype,
		)
		PanicAtError(err)
		bestFourTeamsArrangement32 := make([]int32, len(bestFourTeamsArrangementPgtype.Elements))
		bestFourTeamsArrangement := make([]int, len(bestFourTeamsArrangementPgtype.Elements))
		err = bestFourTeamsArrangementPgtype.AssignTo(&bestFourTeamsArrangement32)
		if err != nil {
			panic(err)
		}
		for i, v := range bestFourTeamsArrangement32 {
			bestFourTeamsArrangement[i] = int(v)
		}

		thirdEvaluationRule.BestFourTeamsArrangement = bestFourTeamsArrangement
		thirdEvaluationRules = append(thirdEvaluationRules, thirdEvaluationRule)
	}

	PanicAtError(err)

	return thirdEvaluationRules
}

func LoadTournamentFromDb(tournament_id int) TournamentDb {
	tournament := ReadTournament(tournament_id)
	return tournament
}
