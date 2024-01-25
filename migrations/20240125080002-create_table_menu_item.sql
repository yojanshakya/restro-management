-- +migrate Up
CREATE TABLE IF NOT EXISTS `menu_item`(
    `id` INTEGER PRIMARY KEY AUTO_INCREMENT,
    `description` VARCHAR(200),
    `quantity` INTEGER,
    `price` INTEGER,
    `name` VARCHAR(200),
    `category_id` INTEGER,
    FOREIGN KEY (category_id) REFERENCES food_category(id)
);
-- +migrate Down
DROP TABLE IF EXISTS `menu_item`;

