ALTER TABLE triage_ticket ADD COLUMN creator_id INT(11) UNSIGNED ;
ALTER TABLE triage_ticket ADD CONSTRAINT ticket_creator FOREIGN KEY (creator_id) REFERENCES triage_user (id) ON DELETE NO ACTION ON UPDATE CASCADE;
