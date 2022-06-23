DROP DATABASE IF EXISTS `jwt_db`;
CREATE DATABASE IF NOT EXISTS `jwt_db`;
use `jwt_db`;
CREATE TABLE IF NOT EXISTS `members`
(
    `id`   INT(20) AUTO_INCREMENT,
    `first_name` VARCHAR(255) NOT NULL,
    `last_name` VARCHAR(255) NOT NULL,
    `password` VARCHAR(128) NOT NULL,
    `email` VARCHAR(50) NOT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO members VALUES (1, 'minoru', 'tanaka', 'password', 'tanaka@gmail.com');
INSERT INTO members VALUES (2, 'tadashi', 'sato', 'password', 'sato@gmail.com');
INSERT INTO members VALUES (3, 'sachiko', 'suzuki', 'password','suzuki@yahoo.co.jp');


