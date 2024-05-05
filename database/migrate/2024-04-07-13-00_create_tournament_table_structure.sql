CREATE TABLE IF NOT EXISTS tournaments (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TYPE strength AS ENUM ('low', 'medium', 'high');
CREATE TYPE group_type AS ENUM ('group_phase', 'knockout_phase');


CREATE TABLE IF NOT EXISTS groups (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    tournament_id INT NOT NULL,
    group_type group_type NOT NULL,
    FOREIGN KEY (tournament_id) REFERENCES tournaments(id)
);

CREATE TABLE IF NOT EXISTS teams (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    strength strength NOT NULL,
    group_id INT NOT NULL,
    FOREIGN KEY (group_id) REFERENCES groups(id)
);

CREATE TABLE IF NOT EXISTS matches (
    id SERIAL PRIMARY KEY,
    team1_id INT NOT NULL,
    team2_id INT NOT NULL,
    group_id INT NOT NULL,
    FOREIGN KEY (team1_id) REFERENCES teams(id),
    FOREIGN KEY (team2_id) REFERENCES teams(id),
    FOREIGN KEY (group_id) REFERENCES groups(id)
);

CREATE TABLE IF NOT EXISTS match_results (
    id SERIAL PRIMARY KEY,
    match_id INT NOT NULL,
    team1_goals INT NOT NULL,
    team1_penalty_goals INT DEFAULT NULL,
    team2_goals INT NOT NULL,
    team2_penalty_goals INT DEFAULT NULL,
    FOREIGN KEY (match_id) REFERENCES matches(id),
    UNIQUE(match_id)
);

CREATE TABLE IF NOT EXISTS group_rankings (
    id SERIAL PRIMARY KEY,
    group_id INT NOT NULL,
    team_id INT NOT NULL,
    ranking INT NOT NULL,
    FOREIGN KEY (group_id) REFERENCES groups(id),
    FOREIGN KEY (team_id) REFERENCES teams(id),
    UNIQUE(group_id, team_id)
);


CREATE TABLE IF NOT EXISTS ko_matches (
    id SERIAL PRIMARY KEY,
    group_id INT NOT NULL,
    group_id1 INT, -- can be null when the value depends on the outcome of previous matches
    group_id2 INT, -- can be null when the value depends on the outcome of previous matches
    ranking1 INT NOT NULL,
    ranking2 INT NOT NULL,
    FOREIGN KEY (group_id) REFERENCES groups(id),
    FOREIGN KEY (group_id1) REFERENCES groups(id),
    FOREIGN KEY (group_id2) REFERENCES groups(id)
);
