CREATE TABLE IF NOT EXISTS `worker_profiles` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `worker_id` int(11) unsigned NOT NULL,
  `birthdate` date DEFAULT NULL,
  `city` varchar(255) DEFAULT NULL,
  `locality` varchar(255) DEFAULT NULL,
  `country` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `worker_id` (`worker_id`)
)