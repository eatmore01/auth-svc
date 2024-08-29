-- +goose Up
CREATE TABLE IF NOT EXIST users ( 
    id UUID DEFAULT gen_random_uuid() NOT NULL PRIMARY KEY,
    user_name VARCHAR(100) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL
);


INSERT INTO users (user_name, email, password) VALUES ('eatmore_auth', 'test@gmail.com', 'ef6689061e2cf0bcacddd13befd98a8b4b666284ac184a4c0d3c246a1ebc7656');

-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE users;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
