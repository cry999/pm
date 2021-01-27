-- カラム追加
ALTER TABLE `tasks`
ADD COLUMN `project_id` varchar(32) DEFAULT NULL
AFTER `assignee_id`;
-- INDEX 追加
ALTER TABLE `tasks`
ADD INDEX `idx_tasks_project_id` (`project_id`);
