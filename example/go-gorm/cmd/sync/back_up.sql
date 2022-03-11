PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE `users` (`id` integer,`name` text,`age` integer,`birthday` datetime,PRIMARY KEY (`id`));
INSERT INTO users VALUES(1,'xx',0,'0001-01-01 00:00:00+00:00');
INSERT INTO users VALUES(2,'2x',0,'0001-01-01 00:00:00+00:00');
COMMIT;
