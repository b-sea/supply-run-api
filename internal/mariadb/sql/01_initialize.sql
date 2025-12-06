CREATE DATABASE IF NOT EXISTS `supplyrun`;

USE `supplyrun`;

CREATE TABLE IF NOT EXISTS `users` (
    `id` CHAR(36) PRIMARY KEY,
    `username` NVARCHAR(24) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS `recipes` (
    `id` CHAR(36) PRIMARY KEY,
    `name` NVARCHAR(36) NOT NULL,
    `url` NVARCHAR(36),
    `num_servings` INT NOT NULL,
    `created_at` DATETIME NOT NULL,
    `created_by` CHAR(36) NOT NULL,
    `updated_at` DATETIME NOT NULL,
    `updated_by` CHAR(36) NOT NULL,

    CONSTRAINT fk_recipe_created_by FOREIGN KEY(`created_by`) REFERENCES `users`(id),
    CONSTRAINT fk_recipe_updated_by FOREIGN KEY(`updated_by`) REFERENCES `users`(id)
);

CREATE TABLE IF NOT EXISTS `users_recipes` (
    `user_id` CHAR(36),
    `recipe_id` CHAR(36),
    `is_favorite` TINYINT(1) NOT NULL DEFAULT 0,

    PRIMARY KEY(`user_id`, `recipe_id`),
    
    CONSTRAINT fk_user_recipe_user_id FOREIGN KEY(`user_id`) REFERENCES `users`(id),
    CONSTRAINT fk_user_recipe_recipe_id FOREIGN KEY(`recipe_id`) REFERENCES `recipes`(id)
);

CREATE TABLE IF NOT EXISTS `steps` (
    `id` CHAR(36) PRIMARY KEY,
    `data` TEXT NOT NULL,
    `order_by` INT NOT NULL,
    `recipe_id` CHAR(36) NOT NULL,

    CONSTRAINT fk_step_recipe_id FOREIGN KEY(`recipe_id`) REFERENCES `recipes`(id)
);

CREATE TABLE IF NOT EXISTS `tags` (
    `id` CHAR(36) PRIMARY KEY,
    `name` NVARCHAR(36) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS `recipes_tags` (
    `recipe_id` CHAR(36),
    `tag_id` CHAR(36),

    PRIMARY KEY(`recipe_id`, `tag_id`),
    
    CONSTRAINT fk_recipe_tag_recipe_id FOREIGN KEY(`recipe_id`) REFERENCES `recipes`(id),
    CONSTRAINT fk_recipe_tag_tag_id FOREIGN KEY(`tag_id`) REFERENCES `tags`(id)
);

CREATE TABLE IF NOT EXISTS `unit_types` (
    `id` CHAR(36) PRIMARY KEY,
    `name` NVARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS `unit_systems` (
    `id` CHAR(36) PRIMARY KEY,
    `name` NVARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS `units` (
    `id` CHAR(36) PRIMARY KEY,
    `name` NVARCHAR(50) NOT NULL UNIQUE,
    `plural` NVARCHAR(50) NOT NULL UNIQUE,
    `symbol` NVARCHAR(10) NOT NULL UNIQUE,
    `type_id` CHAR(36) NOT NULL,
    `system_id` CHAR(36) NOT NULL,

    CONSTRAINT fk_unit_base_type FOREIGN KEY(`type_id`) REFERENCES `unit_types`(id),
    CONSTRAINT fk_unit_system FOREIGN KEY(`system_id`) REFERENCES `unit_systems`(id)
);

CREATE TABLE IF NOT EXISTS `ingredients` (
    `id` CHAR(36) PRIMARY KEY,
    `name` NVARCHAR(36) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS `measurements` (
    `id` CHAR(36) PRIMARY KEY,
    `quantity` DOUBLE NOT NULL,
    `recipe_id` CHAR(36) NOT NULL,
    `ingredient_id` CHAR(36) NOT NULL,
    `unit_id` CHAR(36) NOT NULL,

    CONSTRAINT fk_measurement_recipe FOREIGN KEY(`recipe_id`) REFERENCES `recipes`(id),
    CONSTRAINT fk_measurement_ingredient FOREIGN KEY(`ingredient_id`) REFERENCES `ingredients`(id),
    CONSTRAINT fk_measurement_unit FOREIGN KEY(`unit_id`) REFERENCES `units`(id)
);
