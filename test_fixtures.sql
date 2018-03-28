INSERT INTO project (id, name) VALUES (1, "Test Project");
INSERT INTO project (id, name) VALUES (2, "Other Project");

INSERT INTO status (id, name, colour) VALUES (1, "Open", "#5CA32E");
INSERT INTO status (id, name, colour) VALUES (2, "Closed", "#BB1432");
INSERT INTO status (id, name, colour) VALUES (3, "In Progress", "#F49B1D");

INSERT INTO ticket (id, subject, description, project_id, status_id) VALUES (1, "first", "", 1, 1);
INSERT INTO ticket (id, subject, description, project_id, status_id) VALUES (2, "second", "", 1, 2);
