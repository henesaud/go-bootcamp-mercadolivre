CREATE DATABASE library;

USE library;

CREATE TABLE `customer` (
	`id` INT PRIMARY KEY AUTO_INCREMENT,
	`name` VARCHAR(255) NOT NULL
)ENGINE=InnoDB;


CREATE TABLE `book` (
	`id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
	`name` VARCHAR(255) NOT NULL,
	`isbn` VARCHAR(255) NOT NULL
)ENGINE=InnoDB;


CREATE TABLE `rent` (
	`id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
	`rent_date` DATE NOT NULL,
	`returning_date` DATE NOT NULL,
	`fk_customer_id` INT NOT NULL,
	`fk_book_id` INT NOT NULL,
	FOREIGN KEY(`fk_customer_id`) REFERENCES `customer`(`id`),
	FOREIGN KEY(`fk_book_id`) REFERENCES `book`(`id`)
)ENGINE=InnoDB;