INSERT INTO triage_project (id, name) VALUES (1, "Test Project");
INSERT INTO triage_project (id, name) VALUES (2, "Other Project");

INSERT INTO triage_status (id, name, colour) VALUES (1, "Open", "#5CA32E");
INSERT INTO triage_status (id, name, colour) VALUES (2, "Closed", "#BB1432");
INSERT INTO triage_status (id, name, colour) VALUES (3, "In Progress", "#F49B1D");

INSERT INTO triage_user (id, username, hashed_password) VALUES (1, "test", "$2a$10$IKaI0ghzshw4kzfjojozAuz9/8cC3Pqvt1Q5liFq91yWFBc1pEnKa");

INSERT INTO triage_ticket (id, subject, description, project_id, status_id) VALUES (1, "first", "", 1, 1);
INSERT INTO triage_ticket (id, subject, description, project_id, status_id) VALUES (2, "second", "", 1, 2);
