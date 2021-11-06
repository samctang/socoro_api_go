
-- +migrate Up
CREATE TABLE IF NOT EXISTS user_acc (
    id serial PRIMARY KEY,
    uuid text NOT NULL,
    name text NOT NULL,
    email text NOT NULL UNIQUE,
    password text,
    role int,
    isAdmin bool
);

-- +migrate Down
DROP TABLE user_acc;