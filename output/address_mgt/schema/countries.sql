CREATE SCHEMA IF NOT EXISTS `ADDRESS_MGT` DEFAULT CHARACTER SET utf8;

USE `ADDRESS_MGT`;

DROP TABLE IF EXISTS `ADDRESS_MGT`.`COUNTRIES`;

CREATE TABLE `ADDRESS_MGT`.`COUNTRIES` (
  `ID` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,

  -- Persistent columns --
  `C_NAME` VARCHAR(10) NOT NULL,

  -- Relational columns --

  -- Constraint --
  PRIMARY KEY (`ID`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
