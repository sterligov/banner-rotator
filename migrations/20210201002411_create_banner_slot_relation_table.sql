-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS banner_slot (
  id INT(11) PRIMARY KEY AUTO_INCREMENT,
  banner_id INT(11) NOT NULL,
  slot_id INT(11) NOT NULL,
  FOREIGN KEY (banner_id) REFERENCES banner(id) ON DELETE CASCADE,
  FOREIGN KEY (slot_id) REFERENCES slot(id) ON DELETE CASCADE
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS banner_slot;