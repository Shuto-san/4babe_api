CREATE TABLE `chat_room_configs` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `male_user_id` bigint(20) unsigned NOT NULL,
  `female_user_id` bigint(20) unsigned NOT NULL,
  `dating_count` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `user_id_multi_index` (`male_user_id`,`female_user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
