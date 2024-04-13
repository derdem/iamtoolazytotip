CREATE TABLE IF NOT EXISTS tournament (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- create a strength enum
CREATE TYPE strength AS ENUM ('low', 'medium', 'high');


CREATE TABLE IF NOT EXISTS groups (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    tournament_id INT NOT NULL,
    FOREIGN KEY (tournament_id) REFERENCES tournament(id)
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
    team1_goals INT DEFAULT 0,
    team1_penalty_goals INT DEFAULT 0,
    team2_id INT NOT NULL,
    team2_goals INT DEFAULT 0,
    team2_penalty_goals INT DEFAULT 0,
    group_id INT NOT NULL,
    FOREIGN KEY (team1_id) REFERENCES teams(id),
    FOREIGN KEY (team2_id) REFERENCES teams(id),
    FOREIGN KEY (group_id) REFERENCES groups(id)
);

CREATE TABLE IF NOT EXISTS results (
    id SERIAL PRIMARY KEY,
    match_id INT NOT NULL,
    team1_score INT NOT NULL,
    team2_score INT NOT NULL,
    FOREIGN KEY (match_id) REFERENCES matches(id)
);
