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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


CREATE TABLE `auctions` (
  `id` int NOT NULL AUTO_INCREMENT,
  `ad_space_id` int NOT NULL,
  `end_time` datetime NOT NULL,
  `status` enum('PENDING','STARTED','COMPLETED') NOT NULL,
  `is_active` tinyint NOT NULL DEFAULT '1',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `auctions_FK` (`ad_space_id`),
  CONSTRAINT `auctions_FK` FOREIGN KEY (`ad_space_id`) REFERENCES `ad_spaces` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


-- auctions.bidder definition

CREATE TABLE `bidder` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT NULL,
  `email` varchar(100) DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_active` tinyint NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


-- auctions.biddings definition

CREATE TABLE `biddings` (
  `id` int NOT NULL AUTO_INCREMENT,
  `bidder_id` int NOT NULL,
  `auction_id` int NOT NULL,
  `is_active` tinyint NOT NULL DEFAULT '1',
  `amount` float DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `biddings_UN` (`auction_id`,`amount`),
  KEY `biddings_FK` (`bidder_id`),
  CONSTRAINT `biddings_FK` FOREIGN KEY (`bidder_id`) REFERENCES `bidder` (`id`),
  CONSTRAINT `biddings_FK_1` FOREIGN KEY (`auction_id`) REFERENCES `auctions` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;