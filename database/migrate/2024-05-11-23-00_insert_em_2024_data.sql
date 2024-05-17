-- Description: Insert data for the Euro 2024 tournament
INSERT INTO tournaments (id, name) VALUES (3, 'EM 2024');

INSERT INTO thirds_evaluation_rules (tournament_id, best_four_teams_id, best_four_teams_arrangement) VALUES
    (3, 15, '{0, 3, 1, 4}'), -- X X X X 0 0  => 1 + 2 + 4  + 8  = 15
    (3, 23, '{0, 4, 1, 2}'), -- X X X 0 X 0  => 1 + 2 + 4  + 16 = 23
    (3, 39, '{0, 5, 1, 2}'), -- X X X 0 0 X  => 1 + 2 + 4  + 32 = 39
    (3, 27, '{3, 4, 0, 1}'), -- X X 0 X X 0  => 1 + 2 + 8  + 16 = 27
    (3, 43, '{3, 5, 0, 1}'), -- X X 0 X 0 X  => 1 + 2 + 8  + 32 = 43
    (3, 51, '{4, 5, 1, 0}'), -- X X 0 0 X X  => 1 + 2 + 16 + 32 = 51
    (3, 29, '{4, 3, 2, 0}'), -- X 0 X X X 0  => 1 + 4 + 8  + 16 = 29
    (3, 45, '{5, 3, 2, 0}'), -- X 0 X X 0 X  => 1 + 4 + 8  + 32 = 45
    (3, 53, '{4, 5, 2, 0}'), -- X 0 X 0 X X  => 1 + 4 + 16 + 32 = 53
    (3, 57, '{4, 5, 3, 0}'), -- X 0 0 X X X  => 1 + 8 + 16 + 32 = 57
    (3, 30, '{4, 3, 1, 2}'), -- 0 X X X X 0  => 2 + 4 + 8  + 16 = 30
    (3, 46, '{5, 3, 2, 1}'), -- 0 X X X 0 X  => 2 + 4 + 8  + 32 = 46
    (3, 54, '{5, 4, 2, 1}'), -- 0 X X 0 X X  => 2 + 4 + 16 + 32 = 54
    (3, 58, '{5, 4, 3, 1}'), -- 0 X 0 X X X  => 2 + 8 + 16 + 32 = 58
    (3, 60, '{5, 4, 3, 2}'); -- 0 0 X X X X  => 4 + 8 + 16 + 32 = 60

INSERT INTO groups (id, name, tournament_id, group_type) VALUES (17, 'Group A', 3, 'group_phase');
INSERT INTO groups (id, name, tournament_id, group_type) VALUES (18, 'Group B', 3, 'group_phase');
INSERT INTO groups (id, name, tournament_id, group_type) VALUES (19, 'Group C', 3, 'group_phase');
INSERT INTO groups (id, name, tournament_id, group_type) VALUES (20, 'Group D', 3, 'group_phase');
INSERT INTO groups (id, name, tournament_id, group_type) VALUES (21, 'Group E', 3, 'group_phase');
INSERT INTO groups (id, name, tournament_id, group_type) VALUES (22, 'Group F', 3, 'group_phase');

INSERT INTO teams (id, name, strength, group_id) VALUES (49, 'Germany', 'high', 17);
INSERT INTO teams (id, name, strength, group_id) VALUES (50, 'Scotland', 'low', 17);
INSERT INTO teams (id, name, strength, group_id) VALUES (51, 'Hungary', 'low', 17);
INSERT INTO teams (id, name, strength, group_id) VALUES (52, 'Switzerland', 'medium', 17);
INSERT INTO teams (id, name, strength, group_id) VALUES (53, 'Spain', 'high', 18);
INSERT INTO teams (id, name, strength, group_id) VALUES (54, 'Croatia', 'medium', 18);
INSERT INTO teams (id, name, strength, group_id) VALUES (55, 'Italy', 'high', 18);
INSERT INTO teams (id, name, strength, group_id) VALUES (56, 'Albania', 'low', 18);
INSERT INTO teams (id, name, strength, group_id) VALUES (57, 'Slovenia', 'low', 19);
INSERT INTO teams (id, name, strength, group_id) VALUES (58, 'Denmark', 'medium', 19);
INSERT INTO teams (id, name, strength, group_id) VALUES (59, 'Serbia', 'low', 19);
INSERT INTO teams (id, name, strength, group_id) VALUES (60, 'England', 'high', 19);
INSERT INTO teams (id, name, strength, group_id) VALUES (61, 'Netherlands', 'high', 20);
INSERT INTO teams (id, name, strength, group_id) VALUES (62, 'France', 'high', 20);
INSERT INTO teams (id, name, strength, group_id) VALUES (63, 'Poland', 'medium', 20);
INSERT INTO teams (id, name, strength, group_id) VALUES (64, 'Austria', 'medium', 20);
INSERT INTO teams (id, name, strength, group_id) VALUES (65, 'Ukraine', 'low', 21);
INSERT INTO teams (id, name, strength, group_id) VALUES (66, 'Slovakia', 'low', 21);
INSERT INTO teams (id, name, strength, group_id) VALUES (67, 'Belgium', 'medium', 21);
INSERT INTO teams (id, name, strength, group_id) VALUES (68, 'Romania', 'low', 21);
INSERT INTO teams (id, name, strength, group_id) VALUES (69, 'Portugal', 'high', 22);
INSERT INTO teams (id, name, strength, group_id) VALUES (70, 'Czechia', 'low', 22);
INSERT INTO teams (id, name, strength, group_id) VALUES (71, 'Georgia', 'low', 22);
INSERT INTO teams (id, name, strength, group_id) VALUES (72, 'Turkey', 'medium', 22);

-- Matches
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (73, 49, 50, 17); -- Germany vs Scotland 2024-06-14T19:00:00.0Z
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (74, 51, 52, 17); -- Hungary vs Switzerland 2024-06-15T13:00:00.0Z
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (75, 53, 54, 18); -- Spain vs Croatia 2024-06-15T16:00:00.0Z
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (76, 55, 56, 18); -- Italy vs Albania 2024-06-15T19:00:00.0Z
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (77, 63, 61, 20); -- Poland vs Netherlands 2024-06-16T13:00:00.0Z
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (78, 57, 58, 19); -- Slovenia vs Denmark 2024-06-16T16:00:00.0Z
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (79, 59, 60, 19); -- Serbia vs England 2024-06-16T19:00:00.0Z
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (80, 68, 65, 21); -- Romania vs Ukraine 2024-06-17T13:00:00.0Z
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (81, 67, 66, 21); -- Belgium vs Slovakia 2024-06-17T16:00:00.0Z
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (82, 64, 62, 20); -- Austria vs France 2024-06-17T19:00:00.0Z
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (83, 72, 71, 22); -- Turkey vs Georgia 2024-06-18T16:00:00.0Z
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (84, 69, 70, 22); -- Portugal vs Czechia 2024-06-18T19:00:00.0Z
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (85, 54, 56, 18); -- Croatia vs Albania 2024-06-19T13:00:00.0Z
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (86, 49, 51, 17); -- Germany vs Hungary 2024-06-19T16:00:00.0Z
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (87, 50, 52, 17); -- Scotland vs Switzerland 2024-06-19T19:00:00.0Z
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (88, 57, 59, 19); -- Slovenia vs Serbia 2024-06-20T13:00:00.0Z
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (89, 58, 60, 19); -- Denmark vs England 2024-06-20T16:00:00.0Z
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (90, 53, 55, 18); -- Spain vs Italy 2024-06-20T19:00:00.0Z
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (91, 66, 65, 21); -- Slovakia vs Ukraine 2024-06-21T13:00:00.0Z
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (92, 63, 64, 20); -- Poland vs Austria 2024-06-21T16:00:00.0Z
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (93, 61, 62, 20); -- Netherlands vs France 2024-06-21T19:00:00.0Z
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (94, 71, 70, 22); -- Georgia vs Czechia 2024-06-22T13:00:00.0Z
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (95, 72, 69, 22); -- Turkey vs Portugal 2024-06-22T16:00:00.0Z
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (96, 67, 68, 21); -- Belgium vs Romania 2024-06-22T19:00:00.0Z
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (97, 52, 49, 17); -- Switzerland vs Germany 2024-06-23T19:00:00.0Z
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (98, 50, 51, 17); -- Scotland vs Hungary 2024-06-23T19:00:00.0Z
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (99, 56, 53, 18); -- Albania vs Spain 2024-06-24T19:00:00.0Z
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (100, 54, 55, 18); -- Croatia vs Italy 2024-06-24T19:00:00.0Z
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (101, 61, 64, 20); -- Netherlands vs Austria 2024-06-25T16:00:00.0Z
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (102, 62, 63, 20); -- France vs Poland 2024-06-25T16:00:00.0Z
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (103, 60, 57, 19); -- England vs Slovenia 2024-06-25T19:00:00.0Z
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (104, 58, 59, 19); -- Denmark vs Serbia 2024-06-25T19:00:00.0Z
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (105, 66, 68, 21); -- Slovakia vs Romania 2024-06-26T16:00:00.0Z
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (106, 65, 67, 21); -- Ukraine vs Belgium 2024-06-26T16:00:00.0Z
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (107, 71, 69, 22); -- Georgia vs Portugal 2024-06-26T19:00:00.0Z
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (108, 70, 72, 22); -- Czechia vs Turkey 2024-06-26T19:00:00.0Z

-- GROUP KO groups
INSERT INTO groups (id, name, tournament_id, group_type) VALUES (23, 'Round 16', 3, 'knockout_phase');
INSERT INTO groups (id, name, tournament_id, group_type) VALUES (24, 'Quarterfinals', 3, 'knockout_phase');
INSERT INTO groups (id, name, tournament_id, group_type) VALUES (25, 'Semifinals', 3, 'knockout_phase');
INSERT INTO groups (id, name, tournament_id, group_type) VALUES (26, 'Final', 3, 'knockout_phase');

-- GROUP Round 16 matches
INSERT INTO ko_matches (id, group_id, group_id1, group_id2, ranking1, ranking2) VALUES (16, 23, 17, 18,   2, 2); -- A2 vs B2 2024-06-29T16:00:00.0Z Match 37
INSERT INTO ko_matches (id, group_id, group_id1, group_id2, ranking1, ranking2) VALUES (17, 23, 17, 19,   1, 2); -- A1 vs C2 2024-06-29T19:00:00.0Z Match 38
INSERT INTO ko_matches (id, group_id, group_id1, group_id2, ranking1, ranking2) VALUES (18, 23, 19, NULL, 1, 3); -- C1 vs *3 2024-06-30T16:00:00.0Z Match 39
INSERT INTO ko_matches (id, group_id, group_id1, group_id2, ranking1, ranking2) VALUES (19, 23, 18, NULL, 1, 3); -- B1 vs *3 2024-06-30T19:00:00.0Z Match 40
INSERT INTO ko_matches (id, group_id, group_id1, group_id2, ranking1, ranking2) VALUES (20, 23, 20, 21,   2, 2); -- D2 vs E2 2024-07-01T16:00:00.0Z Match 41
INSERT INTO ko_matches (id, group_id, group_id1, group_id2, ranking1, ranking2) VALUES (21, 23, 22, NULL, 1, 3); -- F1 vs *3 2024-07-01T19:00:00.0Z Match 42
INSERT INTO ko_matches (id, group_id, group_id1, group_id2, ranking1, ranking2) VALUES (22, 23, 21, NULL, 1, 3); -- E1 vs *3 2024-07-02T16:00:00.0Z Match 43
INSERT INTO ko_matches (id, group_id, group_id1, group_id2, ranking1, ranking2) VALUES (23, 23, 20, 22,   1, 2); -- D1 vs F2 2024-07-02T19:00:00.0Z Match 44

-- GROUP Quarterfinals matches
INSERT INTO ko_matches (id, group_id, group_id1, group_id2, ranking1, ranking2) VALUES (24, 24, 23, 23, 3, 1); -- Match 39 vs Match 37 2024-07-05T16:00:00.0Z Match 45
INSERT INTO ko_matches (id, group_id, group_id1, group_id2, ranking1, ranking2) VALUES (25, 24, 23, 23, 5, 6); -- Match 41 vs Match 42 2024-07-05T19:00:00.0Z Match 46
INSERT INTO ko_matches (id, group_id, group_id1, group_id2, ranking1, ranking2) VALUES (26, 24, 23, 23, 4, 2); -- Match 40 vs Match 38 2024-07-06T16:00:00.0Z Match 47
INSERT INTO ko_matches (id, group_id, group_id1, group_id2, ranking1, ranking2) VALUES (27, 24, 23, 23, 7, 8); -- Match 43 vs Match 44 2024-07-06T19:00:00.0Z Match 48

-- GROUP Semifinals matches
INSERT INTO ko_matches (id, group_id, group_id1, group_id2, ranking1, ranking2) VALUES (28, 25, 24, 24, 1, 2); -- Match 45 vs Match 46 2024-07-09T19:00:00.0Z Match 49
INSERT INTO ko_matches (id, group_id, group_id1, group_id2, ranking1, ranking2) VALUES (29, 25, 24, 24, 3, 4); -- Match 47 vs Match 48 2024-07-10T19:00:00.0Z Match 50

-- GROUP Final matches
INSERT INTO ko_matches (id, group_id, group_id1, group_id2, ranking1, ranking2) VALUES (30, 26, 25, 25, 1, 2); -- Match 49 vs Match 50 2024-07-14T19:00:00.0Z Match 51
