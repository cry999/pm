-- INDEX 削除
ALTER TABLE `tasks` DROP INDEX `idx_tasks_project_id`;
-- カラム削除
ALTER TABLE `tasks` DROP COLUMN `project_id`;
