CREATE SCHEMA IF NOT EXISTS `CUSTOMER` DEFAULT CHARACTER SET utf8;

USE `CUSTOMER`;

DROP TABLE IF EXISTS `CUSTOMER`.`ADDRESSES`;

CREATE TABLE `CUSTOMER`.`ADDRESSES` (
  `ID` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,

  -- Main columns --
  `POST_CODE` INT NOT NULL,
  `STATE` VARCHAR(3) NOT NULL,
  `STREET` VARCHAR(3) NOT NULL,
  `NUMBER` INT NOT NULL,
  `IS_DEFAULT` BOOLEAN,

  -- Metadata columns --
  `CREATED_DATE` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `CREATED_BY` VARCHAR(100),
  `CREATED_SOURCE` VARCHAR(100),
  `LAST_MODIFIED_DATE` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LAST_MODIFIED_BY` VARCHAR(100),
  `LAST_MODIFIED_SOURCE` VARCHAR(100),
  `STATUS` VARCHAR(50),
  `VERSION` INT,

  -- Relationship columns --
  `CITY_ID` BIGINT UNSIGNED,

  -- Constraint --
  CONSTRAINT `FK_CITY_ADDRESS` FOREIGN KEY (`CITY_ID`) REFERENCES `CUSTOMER`.`CITIES` (`ID`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  PRIMARY KEY (`ID`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
