CREATE TABLE IF NOT EXISTS `planned_project_tasks` (
    `id` integer NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `project_id` varchar(32) NOT NULL,
    `planned_task_id` varchar(32) NOT NULL,
    FOREIGN KEY `fk_planned_project_tasks_project_id_and_projects_id` (`project_id`) REFERENCES `projects` (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;
