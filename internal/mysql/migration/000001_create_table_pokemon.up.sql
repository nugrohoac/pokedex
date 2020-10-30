CREATE TABLE `pokemon` (
    `id` varchar(100) NOT NULL,
    `number` varchar(100) NOT NULL,
    `name` varchar(100) NOT NULL,
    `type` varchar(100) NOT NULL,
    `created_at` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    PRIMARY KEY (`id`),
    KEY `number` (`number`),
    KEY `name` (`name`),
    KEY `type` (`type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;