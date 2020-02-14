-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE orders(
    id serial NOT NULL PRIMARY KEY,
    order_id varchar NOT NULL PRIMARY KEY,
    username varchar FOREIGN KEY,
    pick_up_loc point NOT NULL,
    drop_off_loc point NOT NULL,
    date DATE NOT NULL,
    no_off_passanger int NOT NULL,
    price int,
    vehicle_type int,
);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE orders;
