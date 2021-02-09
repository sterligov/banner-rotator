-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS banner_slot (
  id INT(11) PRIMARY KEY AUTO_INCREMENT,
  banner_id INT(11) NOT NULL,
  slot_id INT(11) NOT NULL,
  UNIQUE(banner_id, slot_id),
  FOREIGN KEY (banner_id) REFERENCES banner(id) ON DELETE CASCADE,
  FOREIGN KEY (slot_id) REFERENCES slot(id) ON DELETE CASCADE
);

INSERT INTO banner_slot(id, banner_id, slot_id) VALUES(100, 100, 100), (200, 200, 200), (300, 100, 300), (400, 300, 100);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS banner_slot;