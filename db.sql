CREATE TABLE `posts` (
	`id` INT NOT NULL AUTO_INCREMENT,
	`title` VARCHAR(255) NOT NULL DEFAULT 'Untitled',
	`created_at` INT NOT NULL,
	`created_by` INT NOT NULL,
	PRIMARY KEY (`id`)
);
