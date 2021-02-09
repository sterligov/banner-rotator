-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS slot (
    id INT(11) PRIMARY KEY AUTO_INCREMENT,
    description VARCHAR(255)
);

INSERT INTO slot(id, description) VALUES(100, '300x400'), (200, '970x250'), (300, '300x400');

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS slot;