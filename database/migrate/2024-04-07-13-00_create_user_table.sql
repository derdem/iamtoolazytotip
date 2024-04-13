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

INSERT INTO tournament (id, name) VALUES (1, 'Test tournament');

INSERT INTO groups (name, tournament_id) VALUES ('A', 1);
INSERT INTO groups (name, tournament_id) VALUES ('B', 1);
INSERT INTO groups (name, tournament_id) VALUES ('C', 1);
INSERT INTO groups (name, tournament_id) VALUES ('D', 1);
INSERT INTO groups (name, tournament_id) VALUES ('E', 1);
INSERT INTO groups (name, tournament_id) VALUES ('F', 1);

INSERT INTO teams (name, strength, group_id) VALUES ('team1', 'low', 1);
INSERT INTO teams (name, strength, group_id) VALUES ('team2', 'low', 1);
INSERT INTO teams (name, strength, group_id) VALUES ('team3', 'low', 1);
INSERT INTO teams (name, strength, group_id) VALUES ('team4', 'low', 1);
INSERT INTO teams (name, strength, group_id) VALUES ('team5', 'low', 2);
INSERT INTO teams (name, strength, group_id) VALUES ('team6', 'low', 2);
INSERT INTO teams (name, strength, group_id) VALUES ('team7', 'low', 2);
INSERT INTO teams (name, strength, group_id) VALUES ('team8', 'low', 2);
INSERT INTO teams (name, strength, group_id) VALUES ('team9', 'low', 3);
INSERT INTO teams (name, strength, group_id) VALUES ('team10', 'low', 3);
INSERT INTO teams (name, strength, group_id) VALUES ('team11', 'low', 3);
INSERT INTO teams (name, strength, group_id) VALUES ('team12', 'low', 3);
INSERT INTO teams (name, strength, group_id) VALUES ('team13', 'low', 4);
INSERT INTO teams (name, strength, group_id) VALUES ('team14', 'low', 4);
INSERT INTO teams (name, strength, group_id) VALUES ('team15', 'low', 4);
INSERT INTO teams (name, strength, group_id) VALUES ('team16', 'low', 4);
INSERT INTO teams (name, strength, group_id) VALUES ('team17', 'low', 5);
INSERT INTO teams (name, strength, group_id) VALUES ('team18', 'low', 5);
INSERT INTO teams (name, strength, group_id) VALUES ('team19', 'low', 5);
INSERT INTO teams (name, strength, group_id) VALUES ('team20', 'low', 5);
INSERT INTO teams (name, strength, group_id) VALUES ('team21', 'low', 6);
INSERT INTO teams (name, strength, group_id) VALUES ('team22', 'low', 6);
INSERT INTO teams (name, strength, group_id) VALUES ('team23', 'low', 6);
INSERT INTO teams (name, strength, group_id) VALUES ('team24', 'low', 6);

-- GROUP A matches
INSERT INTO matches (team1_id, team2_id, group_id) VALUES (1, 2, 1);
INSERT INTO matches (team1_id, team2_id, group_id) VALUES (3, 4, 1);
INSERT INTO matches (team1_id, team2_id, group_id) VALUES (1, 3, 1);
INSERT INTO matches (team1_id, team2_id, group_id) VALUES (2, 4, 1);
INSERT INTO matches (team1_id, team2_id, group_id) VALUES (1, 4, 1);
INSERT INTO matches (team1_id, team2_id, group_id) VALUES (2, 3, 1);

-- GROUP B matches
INSERT INTO matches (team1_id, team2_id, group_id) VALUES (5, 6, 2);
INSERT INTO matches (team1_id, team2_id, group_id) VALUES (7, 8, 2);
INSERT INTO matches (team1_id, team2_id, group_id) VALUES (5, 7, 2);
INSERT INTO matches (team1_id, team2_id, group_id) VALUES (6, 8, 2);
INSERT INTO matches (team1_id, team2_id, group_id) VALUES (5, 8, 2);
INSERT INTO matches (team1_id, team2_id, group_id) VALUES (6, 7, 2);

-- GROUP C matches
INSERT INTO matches (team1_id, team2_id, group_id) VALUES (9, 10, 3);
INSERT INTO matches (team1_id, team2_id, group_id) VALUES (11, 12, 3);
INSERT INTO matches (team1_id, team2_id, group_id) VALUES (9, 11, 3);
INSERT INTO matches (team1_id, team2_id, group_id) VALUES (10, 12, 3);
INSERT INTO matches (team1_id, team2_id, group_id) VALUES (9, 12, 3);
INSERT INTO matches (team1_id, team2_id, group_id) VALUES (10, 11, 3);

-- GROUP D matches
INSERT INTO matches (team1_id, team2_id, group_id) VALUES (13, 14, 4);
INSERT INTO matches (team1_id, team2_id, group_id) VALUES (15, 16, 4);
INSERT INTO matches (team1_id, team2_id, group_id) VALUES (13, 15, 4);
INSERT INTO matches (team1_id, team2_id, group_id) VALUES (14, 16, 4);
INSERT INTO matches (team1_id, team2_id, group_id) VALUES (13, 16, 4);
INSERT INTO matches (team1_id, team2_id, group_id) VALUES (14, 15, 4);

-- GROUP E matches
INSERT INTO matches (team1_id, team2_id, group_id) VALUES (17, 18, 5);
INSERT INTO matches (team1_id, team2_id, group_id) VALUES (19, 20, 5);
INSERT INTO matches (team1_id, team2_id, group_id) VALUES (17, 19, 5);
INSERT INTO matches (team1_id, team2_id, group_id) VALUES (18, 20, 5);
INSERT INTO matches (team1_id, team2_id, group_id) VALUES (17, 20, 5);
INSERT INTO matches (team1_id, team2_id, group_id) VALUES (18, 19, 5);

-- GROUP F matches
INSERT INTO matches (team1_id, team2_id, group_id) VALUES (21, 22, 6);
INSERT INTO matches (team1_id, team2_id, group_id) VALUES (23, 24, 6);
INSERT INTO matches (team1_id, team2_id, group_id) VALUES (21, 23, 6);
INSERT INTO matches (team1_id, team2_id, group_id) VALUES (22, 24, 6);
INSERT INTO matches (team1_id, team2_id, group_id) VALUES (21, 24, 6);
INSERT INTO matches (team1_id, team2_id, group_id) VALUES (22, 23, 6);

