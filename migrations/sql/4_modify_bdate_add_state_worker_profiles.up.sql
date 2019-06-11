ALTER TABLE `worker_profiles` 
  ADD `state` varchar(255) NOT NULL default 'incomplete',
  MODIFY `birthdate` varchar(255) default NULL;