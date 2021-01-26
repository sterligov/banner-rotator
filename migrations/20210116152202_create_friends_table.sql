-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS friends (
    id INT(11) AUTO_INCREMENT PRIMARY KEY,
    user_id INT(11) NOT NULL,
    friend_id INT(11) NOT NULL,
    UNIQUE (user_id, friend_id),
    FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE RESTRICT,
    FOREIGN KEY (friend_id) REFERENCES user(id) ON DELETE RESTRICT
) ENGINE=INNODB;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE friends;
