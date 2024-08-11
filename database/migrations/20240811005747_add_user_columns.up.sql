ALTER TABLE users
ADD COLUMN gender VARCHAR(10),
ADD COLUMN is_membership BOOLEAN DEFAULT FALSE,
ADD COLUMN is_admin BOOLEAN DEFAULT FALSE,
ADD COLUMN created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
ADD COLUMN updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP;
