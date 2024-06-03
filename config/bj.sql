CREATE TABLE `styles` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `price` decimal(10,2) NOT NULL,
  `image` blob NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

ALTER TABLE styles
ADD COLUMN created_at timestamp NULL DEFAULT NULL,
ADD COLUMN updated_at timestamp NULL DEFAULT NULL,
ADD COLUMN deleted_at TIMESTAMP NULL DEFAULT NULL,
ADD COLUMN create_by INT(11) UNSIGNED DEFAULT NULL,
ADD COLUMN update_by INT(11) UNSIGNED DEFAULT NULL;

CREATE TABLE `paints` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `price` decimal(10,2) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

ALTER TABLE paints
ADD COLUMN created_at timestamp NULL DEFAULT NULL,
ADD COLUMN updated_at timestamp NULL DEFAULT NULL,
ADD COLUMN deleted_at TIMESTAMP NULL DEFAULT NULL,
ADD COLUMN create_by INT(11) UNSIGNED DEFAULT NULL,
ADD COLUMN update_by INT(11) UNSIGNED DEFAULT NULL;

CREATE TABLE `regions` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `base1` decimal(10,2) NOT NULL,
  `base2` decimal(10,2) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

ALTER TABLE regions
ADD COLUMN created_at timestamp NULL DEFAULT NULL,
ADD COLUMN updated_at timestamp NULL DEFAULT NULL,
ADD COLUMN deleted_at TIMESTAMP NULL DEFAULT NULL,
ADD COLUMN create_by INT(11) UNSIGNED DEFAULT NULL,
ADD COLUMN update_by INT(11) UNSIGNED DEFAULT NULL;

CREATE TABLE `quotations` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `style_id` int(11) NOT NULL,
  `paint_id` int(11) NOT NULL,
  `region_id` int(11) NOT NULL,
  `length` decimal(10,2) NOT NULL,
  `width` decimal(10,2) NOT NULL,
  `height` decimal(10,2) NOT NULL,
  `quantity` int(11) NOT NULL,
  `packaging_cost` decimal(10,2) NOT NULL,
  `created_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`user_id`) REFERENCES `sys_user` (`user_id`),
  FOREIGN KEY (`style_id`) REFERENCES `styles` (`id`),
  FOREIGN KEY (`paint_id`) REFERENCES `paints` (`id`),
  FOREIGN KEY (`region_id`) REFERENCES `regions` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


