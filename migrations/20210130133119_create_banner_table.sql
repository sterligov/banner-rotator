-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS banner (
    id INT(11) PRIMARY KEY AUTO_INCREMENT,
    description VARCHAR(255)
);

INSERT INTO banner(id, description) VALUES(100, 'Banner 1'), (200, 'Banner 2'), (300, 'Banner 3');

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS banner;