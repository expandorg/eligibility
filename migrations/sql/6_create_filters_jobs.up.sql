CREATE TABLE `filters_jobs` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `filter_id` int(11) unsigned NOT NULL,
  `job_id` int(11) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `filter_job` (`filter_id`,`job_id`),
  KEY `filter_id` (`filter_id`),
  CONSTRAINT `filters_jobs_ibfk_1` FOREIGN KEY (`filter_id`) REFERENCES `filters` (`id`)
)