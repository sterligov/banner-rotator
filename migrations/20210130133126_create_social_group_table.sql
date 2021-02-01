-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS social_group (
    id INT(11) PRIMARY KEY AUTO_INCREMENT,
    description VARCHAR(255)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS social_group;