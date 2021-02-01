-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS statistic (
   id INT(11) PRIMARY KEY AUTO_INCREMENT,
   banner_slot_id INT(11) NOT NULL,
   social_group_id INT(11) NOT NULL,
   clicks INT(11) NOT NULL,
   shows INT(11) NOT NULL,
   UNIQUE(banner_slot_id, social_group_id),
   FOREIGN KEY (banner_slot_id) REFERENCES banner_slot(id) ON DELETE CASCADE,
   FOREIGN KEY (social_group_id) REFERENCES social_group(id) ON DELETE CASCADE
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS statistic;