-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS user (
    id INT(11) AUTO_INCREMENT PRIMARY KEY,
    password VARCHAR(255) NOT NULL,
    first_name VARCHAR(128) NOT NULL,
    last_name VARCHAR(128) NOT NULL,
    email VARCHAR(128) NOT NULL,
    interests TEXT NOT NULL,
    sex ENUM('M', 'F') NOT NULL DEFAULT 'M',
    birthday DATETIME NOT NULL,
    city_id INT(11) NOT NULL,
    FOREIGN KEY (city_id) REFERENCES city(id) ON DELETE RESTRICT,
    UNIQUE(email)
) ENGINE=INNODB;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE user;
