-- Description: Insert data for Test tournament
INSERT INTO tournaments (id, name) VALUES (1, 'Test tournament');

INSERT INTO groups (id, name, tournament_id) VALUES (1, 'A', 1);
INSERT INTO groups (id, name, tournament_id) VALUES (2, 'B', 1);
INSERT INTO groups (id, name, tournament_id) VALUES (3, 'C', 1);
INSERT INTO groups (id, name, tournament_id) VALUES (4, 'D', 1);
INSERT INTO groups (id, name, tournament_id) VALUES (5, 'E', 1);
INSERT INTO groups (id, name, tournament_id) VALUES (6, 'F', 1);

INSERT INTO teams (id, name, strength, group_id) VALUES (1, 'team1', 'low', 1);
INSERT INTO teams (id, name, strength, group_id) VALUES (2, 'team2', 'low', 1);
INSERT INTO teams (id, name, strength, group_id) VALUES (3, 'team3', 'low', 1);
INSERT INTO teams (id, name, strength, group_id) VALUES (4, 'team4', 'low', 1);
INSERT INTO teams (id, name, strength, group_id) VALUES (5, 'team5', 'low', 2);
INSERT INTO teams (id, name, strength, group_id) VALUES (6, 'team6', 'low', 2);
INSERT INTO teams (id, name, strength, group_id) VALUES (7, 'team7', 'low', 2);
INSERT INTO teams (id, name, strength, group_id) VALUES (8, 'team8', 'low', 2);
INSERT INTO teams (id, name, strength, group_id) VALUES (9, 'team9', 'low', 3);
INSERT INTO teams (id, name, strength, group_id) VALUES (10, 'team10', 'low', 3);
INSERT INTO teams (id, name, strength, group_id) VALUES (11, 'team11', 'low', 3);
INSERT INTO teams (id, name, strength, group_id) VALUES (12, 'team12', 'low', 3);
INSERT INTO teams (id, name, strength, group_id) VALUES (13, 'team13', 'low', 4);
INSERT INTO teams (id, name, strength, group_id) VALUES (14, 'team14', 'low', 4);
INSERT INTO teams (id, name, strength, group_id) VALUES (15, 'team15', 'low', 4);
INSERT INTO teams (id, name, strength, group_id) VALUES (16, 'team16', 'low', 4);
INSERT INTO teams (id, name, strength, group_id) VALUES (17, 'team17', 'low', 5);
INSERT INTO teams (id, name, strength, group_id) VALUES (18, 'team18', 'low', 5);
INSERT INTO teams (id, name, strength, group_id) VALUES (19, 'team19', 'low', 5);
INSERT INTO teams (id, name, strength, group_id) VALUES (20, 'team20', 'low', 5);
INSERT INTO teams (id, name, strength, group_id) VALUES (21, 'team21', 'low', 6);
INSERT INTO teams (id, name, strength, group_id) VALUES (22, 'team22', 'low', 6);
INSERT INTO teams (id, name, strength, group_id) VALUES (23, 'team23', 'low', 6);
INSERT INTO teams (id, name, strength, group_id) VALUES (24, 'team24', 'low', 6);

-- GROUP A matches
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (1, 1, 2, 1);
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (2, 3, 4, 1);
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (3, 1, 3, 1);
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (4, 2, 4, 1);
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (5, 1, 4, 1);
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (6, 2, 3, 1);

-- GROUP B matches
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (7, 5, 6, 2);
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (8, 7, 8, 2);
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (9, 5, 7, 2);
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (10, 6, 8, 2);
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (11, 5, 8, 2);
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (12, 6, 7, 2);

-- GROUP C matches
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (13, 9, 10, 3);
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (14, 11, 12, 3);
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (15, 9, 11, 3);
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (16, 10, 12, 3);
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (17, 9, 12, 3);
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (18, 10, 11, 3);

-- GROUP D matches
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (19, 13, 14, 4);
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (20, 15, 16, 4);
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (21, 13, 15, 4);
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (22, 14, 16, 4);
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (23, 13, 16, 4);
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (24, 14, 15, 4);

-- GROUP E matches
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (25, 17, 18, 5);
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (26, 19, 20, 5);
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (27, 17, 19, 5);
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (28, 18, 20, 5);
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (29, 17, 20, 5);
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (30, 18, 19, 5);

-- GROUP F matches
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (31, 21, 22, 6);
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (32, 23, 24, 6);
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (33, 21, 23, 6);
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (34, 22, 24, 6);
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (35, 21, 24, 6);
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (36, 22, 23, 6);


-- Description: Insert data for the Euro 2021 tournament
INSERT INTO tournaments (id, name) VALUES (2, 'EM 2021');

INSERT INTO groups (id, name, tournament_id, group_type) VALUES (7, 'Group A', 2, 'group_phase');
INSERT INTO groups (id, name, tournament_id, group_type) VALUES (8, 'Group B', 2, 'group_phase');
INSERT INTO groups (id, name, tournament_id, group_type) VALUES (9, 'Group C', 2, 'group_phase');
INSERT INTO groups (id, name, tournament_id, group_type) VALUES (10, 'Group D', 2, 'group_phase');
INSERT INTO groups (id, name, tournament_id, group_type) VALUES (11, 'Group E', 2, 'group_phase');
INSERT INTO groups (id, name, tournament_id, group_type) VALUES (12, 'Group F', 2, 'group_phase');

INSERT INTO teams (id, name, strength, group_id) VALUES (25, 'Turkey', 'medium', 7);
INSERT INTO teams (id, name, strength, group_id) VALUES (26, 'Italy', 'high', 7);
INSERT INTO teams (id, name, strength, group_id) VALUES (27, 'Wales', 'low', 7);
INSERT INTO teams (id, name, strength, group_id) VALUES (28, 'Switzerland', 'medium', 7);
INSERT INTO teams (id, name, strength, group_id) VALUES (29, 'Denmark', 'medium', 8);
INSERT INTO teams (id, name, strength, group_id) VALUES (30, 'Finland', 'low', 8);
INSERT INTO teams (id, name, strength, group_id) VALUES (31, 'Belgien', 'medium', 8);
INSERT INTO teams (id, name, strength, group_id) VALUES (32, 'Russia', 'medium', 8);
INSERT INTO teams (id, name, strength, group_id) VALUES (33, 'Netherlands', 'high', 9);
INSERT INTO teams (id, name, strength, group_id) VALUES (34, 'Ukraine', 'low', 9);
INSERT INTO teams (id, name, strength, group_id) VALUES (35, 'Austria', 'medium', 9);
INSERT INTO teams (id, name, strength, group_id) VALUES (36, 'North Mazedonia', 'low', 9);
INSERT INTO teams (id, name, strength, group_id) VALUES (37, 'England', 'high', 10);
INSERT INTO teams (id, name, strength, group_id) VALUES (38, 'Croatia', 'medium', 10);
INSERT INTO teams (id, name, strength, group_id) VALUES (39, 'Scotland', 'medium', 10);
INSERT INTO teams (id, name, strength, group_id) VALUES (40, 'Czech Republic', 'medium', 10);
INSERT INTO teams (id, name, strength, group_id) VALUES (41, 'Spain', 'high', 11);
INSERT INTO teams (id, name, strength, group_id) VALUES (42, 'Sweden', 'medium', 11);
INSERT INTO teams (id, name, strength, group_id) VALUES (43, 'Poland', 'medium', 11);
INSERT INTO teams (id, name, strength, group_id) VALUES (44, 'Slovakia', 'medium', 11);
INSERT INTO teams (id, name, strength, group_id) VALUES (45, 'Hungry', 'low', 12);
INSERT INTO teams (id, name, strength, group_id) VALUES (46, 'Portugal', 'high', 12);
INSERT INTO teams (id, name, strength, group_id) VALUES (47, 'France', 'high', 12);
INSERT INTO teams (id, name, strength, group_id) VALUES (48, 'Germany', 'high', 12);

-- GROUP A matches
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (37, 25, 26, 7); -- Turkey vs Italy
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (38, 27, 28, 7); -- Wales vs Switzerland
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (39, 25, 27, 7); -- Turkey vs Wales
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (40, 26, 28, 7); -- Italy vs Switzerland
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (41, 28, 25, 7); -- Switzerland vs Turkey
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (42, 26, 27, 7); -- Italy vs Wales

-- GROUP B matches
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (43, 29, 30, 8); -- Denmark vs Finland
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (44, 31, 32, 8); -- Belgium vs Russia
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (45, 30, 32, 8); -- Finland vs Russia
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (46, 29, 31, 8); -- Denmark vs Belgium
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (47, 32, 29, 8); -- Russia vs Denmark
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (48, 30, 31, 8); -- Finland vs Belgium

-- GROUP C matches
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (49, 35, 36, 9); -- Austria vs North Macedonia
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (50, 33, 34, 9); -- Netherlands vs Ukraine
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (51, 34, 36, 9); -- Ukraine vs North Macedonia
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (52, 33, 35, 9); -- Netherlands vs Austria
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (53, 36, 33, 9); -- North Macedonia vs Netherlands
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (54, 34, 35, 9); -- Ukraine vs Austria

-- GROUP D matches
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (55, 37, 38, 10); -- England vs Croatia
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (56, 39, 40, 10); -- Scotland vs Czech Republic
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (57, 38, 40, 10); -- Croatia vs Czech Republic
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (58, 37, 39, 10); -- England vs Scotland
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (59, 38, 39, 10); -- Croatia vs Scotland
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (60, 40, 37, 10); -- Czech Republic vs England

-- GROUP E matches
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (61, 43, 44, 11); -- Poland vs Slovakia
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (62, 41, 42, 11); -- Spain vs Sweden
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (63, 42, 44, 11); -- Sweden vs Slovakia
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (64, 41, 43, 11); -- Spain vs Poland
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (65, 44, 41, 11); -- Slovakia vs Spain
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (66, 42, 43, 11); -- Sweden vs Poland

-- GROUP F matches
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (67, 45, 46, 12); -- Hungary vs Portugal
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (68, 47, 48, 12); -- France vs Germany
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (69, 45, 47, 12); -- Hungary vs France
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (70, 46, 48, 12); -- Portugal vs Germany
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (71, 46, 47, 12); -- Portugal vs France
INSERT INTO matches (id, team1_id, team2_id, group_id) VALUES (72, 48, 45, 12); -- Germany vs Hungary


-- GROUP Ko groups
INSERT INTO groups (id, name, tournament_id, group_type) VALUES (13, 'Round 16', 2, 'knockout_phase');
INSERT INTO groups (id, name, tournament_id, group_type) VALUES (14, 'Quarterfinals', 2, 'knockout_phase');
INSERT INTO groups (id, name, tournament_id, group_type) VALUES (15, 'Semifinals', 2, 'knockout_phase');
INSERT INTO groups (id, name, tournament_id, group_type) VALUES (16, 'Final', 2, 'knockout_phase');

-- GROUP Round 16 matches
INSERT INTO ko_matches (id, group_id, group_id1, group_id2, ranking1, ranking2) VALUES (1, 13,  8, 12, 1, 3); -- B1 vs F3
INSERT INTO ko_matches (id, group_id, group_id1, group_id2, ranking1, ranking2) VALUES (2, 13,  7, 9,  1, 2); -- A1 vs C2
INSERT INTO ko_matches (id, group_id, group_id1, group_id2, ranking1, ranking2) VALUES (3, 13, 12, 7,  1, 3); -- F1 vs A3
INSERT INTO ko_matches (id, group_id, group_id1, group_id2, ranking1, ranking2) VALUES (4, 13, 10, 11, 2, 2); -- D2 vs E2
INSERT INTO ko_matches (id, group_id, group_id1, group_id2, ranking1, ranking2) VALUES (5, 13, 11, 9,  1, 3); -- E1 vs C3
INSERT INTO ko_matches (id, group_id, group_id1, group_id2, ranking1, ranking2) VALUES (6, 13, 10, 12, 1, 2); -- D1 vs F2
INSERT INTO ko_matches (id, group_id, group_id1, group_id2, ranking1, ranking2) VALUES (7, 13,  9, 10, 1, 3); -- C1 vs D3
INSERT INTO ko_matches (id, group_id, group_id1, group_id2, ranking1, ranking2) VALUES (8, 13,  7, 8,  2, 2); -- A2 vs B2

-- GROUP Quarterfinals matches
INSERT INTO ko_matches (id, group_id, group_id1, group_id2, ranking1, ranking2) VALUES (9,  14,  13, 13, 1, 2); -- 1 vs 2
INSERT INTO ko_matches (id, group_id, group_id1, group_id2, ranking1, ranking2) VALUES (10, 14,  13, 13, 3, 4); -- 3 vs 4
INSERT INTO ko_matches (id, group_id, group_id1, group_id2, ranking1, ranking2) VALUES (11, 14,  13, 13, 5, 6); -- 5 vs 6
INSERT INTO ko_matches (id, group_id, group_id1, group_id2, ranking1, ranking2) VALUES (12, 14,  13, 13, 7, 8); -- 7 vs 8

-- GROUP Semifinals matches
INSERT INTO ko_matches (id, group_id, group_id1, group_id2, ranking1, ranking2) VALUES (13, 15,  14, 14, 1, 2); -- 1 vs 2
INSERT INTO ko_matches (id, group_id, group_id1, group_id2, ranking1, ranking2) VALUES (14, 15,  14, 14, 3, 4); -- 3 vs 4

-- GROUP Final matches
INSERT INTO ko_matches (id, group_id, group_id1, group_id2, ranking1, ranking2) VALUES (15, 16,  15, 15, 1, 2); -- 1 vs 2
