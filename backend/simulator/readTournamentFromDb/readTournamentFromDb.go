package readTournamentFromDb

import (
	"github.com/derdem/iamtoolazytotip/postgres_connection"
)

var GetConnection = postgres_connection.GetConnection

func ReadTournament(tournament_id int) []Group {
	conn := GetConnection()

	rows, err := conn.Query(`
        SELECT groups.*, json_agg(DISTINCT teams.*) as teams, json_agg(DISTINCT matches.*) as matches
        FROM groups
        INNER JOIN teams ON groups.id = teams.group_id
        INNER JOIN matches ON groups.id = matches.group_id
        WHERE tournament_id = $1 GROUP BY groups.id;
    `, tournament_id)

	if err != nil {
		panic(err)
	}

	groups := []Group{}

	defer rows.Close()
	for rows.Next() {
		var group Group
		err := rows.Scan(
			&group.Id,
			&group.Name,
			&group.TournamentId,
			&group.Teams,
			&group.Matches)
		if err != nil {
			panic(err)
		}
		groups = append(groups, group)
	}

	return groups

}
