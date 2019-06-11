CREATE TABLE IF NOT EXISTS `filters_workers` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `filter_id` int(11) unsigned NOT NULL,
  `worker_id` int(11) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `filter_worker` (`filter_id`,`worker_id`),
  KEY `filter_id` (`filter_id`),
  CONSTRAINT `filters_workers_ibfk_1` FOREIGN KEY (`filter_id`) REFERENCES `filters` (`id`)
)