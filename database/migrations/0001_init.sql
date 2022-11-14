-- +goose Up
CREATE TABLE IF NOT EXISTS `members` (
	`id`   	     INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
	`name`    varchar(255),
	`birthday`   TIMESTAMP,
	`created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	`updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	PRIMARY key (id)
);
-- +goose Down
DROP TABLE members;