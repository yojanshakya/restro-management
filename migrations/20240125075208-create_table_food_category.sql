
-- +migrate Up
CREATE TABLE IF NOT EXISTS `food_category` (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    category_name VARCHAR(255) NOT NULL
);
-- +migrate Down
DROP TABLE IF EXISTS `food_category`;
