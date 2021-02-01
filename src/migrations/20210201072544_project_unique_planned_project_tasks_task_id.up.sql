--
-- Delete duplicated entries
--
DELETE main
FROM `planned_project_tasks` AS main
    LEFT JOIN (
        SELECT id,
            ROW_NUMBER() over (PARTITION by planned_task_id) AS seq
        FROM `planned_project_tasks`
    ) AS temp ON main.id = temp.id
WHERE temp.seq > 1;
--
-- setting unique column
--
ALTER TABLE `planned_project_tasks`
ADD CONSTRAINT UNIQUE (`planned_task_id`);
