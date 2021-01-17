CREATE TABLE IF NOT EXISTS `users` (
    id varchar(32) PRIMARY KEY NOT NULL,
    email varchar(256) NOT NULL UNIQUE,
    hashed_password varchar(256) NOT NULL,
    created_at datetime DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
    INDEX `idx_users_email` (`email`)
);
