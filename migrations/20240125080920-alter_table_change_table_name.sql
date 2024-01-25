
-- +migrate Up
ALTER TABLE food_category RENAME TO menu_category;
-- +migrate Down
ALTER TABLE menu_category RENAME TO food_category;
