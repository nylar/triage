ALTER TABLE project ADD COLUMN created_at TIMESTAMP NOT NULL DEFAULT NOW();
ALTER TABLE project ADD COLUMN updated_at TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE NOW();