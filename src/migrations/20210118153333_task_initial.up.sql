CREATE TABLE IF NOT EXISTS `tasks` (
    `id` varchar(32) PRIMARY KEY NOT NULL,
    `name` varchar(256) NOT NULL,
    `description` text NOT NULL,
    `owner_id` varchar(32) DEFAULT NULL,
    `assignee_id` varchar(32) DEFAULT NULL,
    `status` varchar(16) NOT NULL,
    `deadline` datetime,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX `idx_tasks_owner_id` (`owner_id`),
    INDEX `idx_tasks_assignee_id` (`assignee_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;
