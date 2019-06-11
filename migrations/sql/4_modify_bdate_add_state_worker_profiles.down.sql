ALTER TABLE `worker_profiles` 
  DROP COLUMN `state`,
  MODIFY `birthdate` date DEFAULT NULL