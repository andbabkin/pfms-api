CREATE TABLE `users` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(150) NOT NULL,
  `pswd` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `public_id` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `admin` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `name_idx` (`name`),
  KEY `public_id_idx` (`public_id`)
) ENGINE=InnoDB;
