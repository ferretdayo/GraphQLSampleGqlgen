
CREATE TABLE `users` (
  `id`              int(10) unsigned  NOT NULL AUTO_INCREMENT,
  `is_unsubscribed` tinyint(1)        NOT NULL DEFAULT 0 COMMENT '退会フラグ',
  `created_at`      datetime          NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at`      datetime          NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ユーザ';

CREATE TABLE `user_unsubscribes` (
  `user_id`     int(10) unsigned      NOT NULL,
  `type`        ENUM('test', 'aaaa')  NOT NULL,
  `created_at`  datetime              NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at`  datetime              NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  CONSTRAINT `user_unsubscribes_users_id_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ユーザ退会';

CREATE TABLE `hobbies` (
  `id`          int(10) unsigned  NOT NULL AUTO_INCREMENT,
  `name`        varchar(255)      NOT NULL            COMMENT '趣味',
  `is_delete`   tinyint(1)        NOT NULL DEFAULT 0  COMMENT '削除済みフラグ',
  `created_at`  datetime          NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at`  datetime          NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='趣味マスタ';

CREATE TABLE `user_details` (
  `user_id`   int(10) unsigned  NOT NULL,
  `nickname`  varchar(255)      NOT NULL COMMENT 'ニックネーム',
  `birthday`  DATE              NOT NULL COMMENT '誕生日',
  `hobby_id`  int(10) unsigned  NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  CONSTRAINT `user_details_users_id_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `user_details_hobbies_id_fk` FOREIGN KEY (`hobby_id`) REFERENCES `hobbies` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ユーザ詳細';
