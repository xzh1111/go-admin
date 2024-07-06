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

CREATE TABLE `regionss` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `base1` decimal(10,2) NOT NULL,
  `base2` decimal(10,2) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

ALTER TABLE regionss
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
  `regions_id` int(11) NOT NULL,
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
  FOREIGN KEY (`regions_id`) REFERENCES `regionss` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


在Go语言的GORM库中，与MySQL数据库类型对应的GORM数据结构类型如下：

1. `TINYINT`：`int8`
2. `SMALLINT`：`int16`
3. `MEDIUMINT`：`int32`
4. `INT`：`int` 或 `int32`
5. `BIGINT`：`int64`
6. `FLOAT`：`float32`
7. `DOUBLE`：`float64`
8. `DECIMAL`：`float64` 或使用第三方库如 `github.com/shopspring/decimal` 的 `decimal.Decimal` 类型
9. `CHAR`、`VARCHAR`、`TINYTEXT`、`TEXT`、`MEDIUMTEXT`、`LONGTEXT`：`string`
10. `BINARY`、`VARBINARY`、`TINYBLOB`、`BLOB`、`MEDIUMBLOB`、`LONGBLOB`：`[]byte`
11. `DATE`：`time.Time`
12. `DATETIME`：`time.Time`
13. `TIMESTAMP`：`time.Time`
14. `TIME`：`time.Duration`
15. `YEAR`：`int16` 或 `uint8`
16. `ENUM`：自定义类型，通常使用字符串表示
17. `SET`：自定义类型，通常使用字符串表示
18. `BIT`：`[]byte` 或自定义类型，通常使用布尔数组表示
19. `JSON`：`json.RawMessage` 或自定义类型，通常使用结构体表示

请注意，这些类型映射是基于GORM的默认行为。在实际使用中，你可能需要根据具体的数据库设计和业务需求调整这些映射。此外，GORM允许你自定义类型和字段级别的映射，以便更好地适应特定的数据库设计和数据类型。



