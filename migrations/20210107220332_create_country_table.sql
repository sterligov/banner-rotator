-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS country (
    id INT(11) AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL
) ENGINE=INNODB;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE country;