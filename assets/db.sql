CREATE TABLE IF NOT EXISTS `posts` (
	`id` INT NOT NULL AUTO_INCREMENT,
	`title` VARCHAR(255) NOT NULL DEFAULT 'Untitled',
	`created_at` DATETIME NOT NULL,
	`created_by` INT NOT NULL,
	PRIMARY KEY (`id`)
);

INSERT INTO `posts` VALUES (1, 'Untitled', '2020-07-16 22:51:59', 1);

CREATE TABLE IF NOT EXISTS `users` (
	`id` int NOT NULL AUTO_INCREMENT,
	`username` varchar(50) NOT NULL,
	`password` varchar(60) NOT NULL,
	PRIMARY KEY (`id`)
);
