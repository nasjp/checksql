CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '名前',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) COMMENT = 'ユーザー' ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin;

INSERT INTO
  `users` (`id`, `name`, `created_at`, `updated_at`)
VALUES
  (
    1,
    'admin',
    '2022-01-01 01:00:00',
    '2022-01-01 01:00:00'
  ),
  (
    2,
    'tmp',
    '2022-02-01 00:00:00',
    '2022-02-01 00:00:00'
  );
