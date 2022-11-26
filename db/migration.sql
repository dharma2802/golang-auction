-- auctions.ad_spaces definition

CREATE TABLE `ad_spaces` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(250) NOT NULL,
  `position` varchar(100) NOT NULL,
  `width` float NOT NULL,
  `height` float NOT NULL,
  `price` float NOT NULL,
  `is_active` tinyint DEFAULT '1',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `ad_spaces_name_UN` (`name`),
  UNIQUE KEY `ad_spaces_wh_position_UN` (`position`,`width`,`height`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;