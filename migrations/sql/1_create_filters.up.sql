DROP TABLE IF EXISTS `filters`;

CREATE TABLE `filters` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `type` varchar(255) NOT NULL,
  `value` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `type_value` (`type`,`value`)
)