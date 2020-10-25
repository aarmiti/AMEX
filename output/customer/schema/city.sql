CREATE SCHEMA IF NOT EXISTS `CUSTOMER` DEFAULT CHARACTER SET utf8;

USE `CUSTOMER`;

DROP TABLE IF EXISTS `CUSTOMER`.`CITIES`;

CREATE TABLE `CUSTOMER`.`CITIES` (
  `ID` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,

  -- Main columns --
  `NAME` VARCHAR(10) NOT NULL,

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
  `COUNTRY_ID` BIGINT UNSIGNED,

  -- Constraint --
  CONSTRAINT `FK_COUNTRY_CITY` FOREIGN KEY (`COUNTRY_ID`) REFERENCES `CUSTOMER`.`COUNTRIES` (`ID`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  PRIMARY KEY (`ID`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;