CREATE TABLE IF NOT EXISTS `projects` (
    `id` varchar(32) PRIMARY KEY NOT NULL,
    `owner_id` varchar(32) NOT NULL,
    `name` varchar(128) NOT NULL,
    `elevator_pitch` varchar(1024) NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX `idx_projects_owner_id` (`owner_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;
