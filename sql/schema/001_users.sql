-- +goose Up 
CREATE TABLE users (
id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
name VARCHAR(50) UNIQUE NOT NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE users;
