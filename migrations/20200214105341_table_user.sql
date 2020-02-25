-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE users (
    id serial NOT NULL PRIMARY KEY,
    username varchar NOT NULL UNIQUE,
    phone_number int not null,
    password varchar not null,
    email text NOT NULL,
    created_at timestamp with time zone,
    deleted_at timestamp with time zone,
    updated_at timestamp with time zone
);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE users;
