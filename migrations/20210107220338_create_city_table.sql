-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS city (
    id INT(11) AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    country_id INT(11) NOT NULL,
    FOREIGN KEY (country_id) REFERENCES country(id) ON DELETE RESTRICT
) ENGINE=INNODB;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE city;