CREATE DATABASE  IF NOT EXISTS `thools` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci */;
USE `thools`;
-- MySQL dump 10.13  Distrib 5.7.17, for macos10.12 (x86_64)
--
-- Host: localhost    Database: article
-- ------------------------------------------------------
-- Server version	5.7.18

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `user_name` VARCHAR(50) NOT NULL,
  `password` VARCHAR(200) NULL,
  `work_email` VARCHAR(50) NULL,
  `first_name` VARCHAR(45) NULL,
  `last_name` VARCHAR(45) NULL,
  `emp_number` VARCHAR(20) NULL,
  `address` VARCHAR(500) NULL,
  `position` VARCHAR(45) NULL,
  `role_id` VARCHAR(45) NULL,
  `hire_date` VARCHAR(100) NULL,
  `end_date` VARCHAR(100) NULL,
  `about_user` VARCHAR(500) NULL,
  `work_phone_number` VARCHAR(45) NULL,
  `social_security_number` VARCHAR(45) NULL,
  `project_start_date` VARCHAR(45) NULL,
  `project_name` VARCHAR(45) NULL,
  `project_end_date` VARCHAR(45) NULL,
  `suffix` VARCHAR(45) NULL,
  `prefix` VARCHAR(45) NULL,
  `post_code` VARCHAR(45) NULL,
  `personal_email` VARCHAR(45) NULL,
  `passport_number` VARCHAR(45) NULL,
  `insurance_number` VARCHAR(45) NULL,
  `home_phone_number` VARCHAR(45) NULL,
  `country` VARCHAR(45) NULL,
  `client_name` VARCHAR(45) NULL,
  `city` VARCHAR(45) NULL,
  `tax_id` VARCHAR(255) NULL,
  `skills` VARCHAR(255) NULL,
  `active` INT NULL,
  `gender` VARCHAR(5) NULL,
  `birth_day` VARCHAR(100) NULL,
  PRIMARY KEY (`id`, `user_name`));

CREATE TABLE `role`
    (
  `id` int
    (11) NOT NULL AUTO_INCREMENT,
  `role` varchar
    (50) DEFAULT NULL,
  PRIMARY KEY
    (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;


CREATE TABLE `user_role`
    (
  `id` int
    (11) NOT NULL AUTO_INCREMENT,
  `user_id` int
    (11) DEFAULT NULL,
  `role_id` int
    (11) DEFAULT NULL,
  PRIMARY KEY
    (`id`),
  KEY `user_id`
    (`user_id`),
  KEY `role_id`
    (`role_id`),
  CONSTRAINT `user_role_ibfk_1` FOREIGN KEY
    (`user_id`) REFERENCES `user`
    (`id`),
  CONSTRAINT `user_role_ibfk_2` FOREIGN KEY
    (`role_id`) REFERENCES `role`
    (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

CREATE TABLE `project`
    (
  `id` int
    (11) NOT NULL AUTO_INCREMENT,
  `project_code` VARCHAR(100) NULL,
  `name` VARCHAR(100) NULL,
  `manager_id` int
    (11) DEFAULT NULL,
  `start_date` DATE NULL,
  `end_date` DATE NULL,
  PRIMARY KEY
    (`id`),
  KEY `manager_id`
    (`manager_id`),
  CONSTRAINT `manager_id_ibfk_1` FOREIGN KEY
    (`manager_id`) REFERENCES `user`
    (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;



CREATE TABLE `timesheet`
    (
  `id` int
    (11) NOT NULL AUTO_INCREMENT,
  `timesheet_date` DATE NULL,
  `employee_id` int
    (11) DEFAULT NULL,
	`time_from` DATETIME NULL,
	`time_to` DATETIME NULL,
	`duration` DATETIME NULL,
	`break` DATETIME NULL,
	`net` DATETIME NULL,
	`captured_by` VARCHAR(100) NULL,
	`status` VARCHAR(50) NULL,
	`comments` VARCHAR(255) NULL,
	`submission_timestamp` DATETIME NULL,
	`project_id` int
    (11) DEFAULT NULL,
  PRIMARY KEY
    (`id`),
  KEY `employee_id`
    (`employee_id`),
	 KEY `project_id`
    (`project_id`),
  CONSTRAINT `employee_id_ibfk_1` FOREIGN KEY
    (`employee_id`) REFERENCES `user`
    (`id`),
	 CONSTRAINT `project_id_ibfk_1` FOREIGN KEY
    (`project_id`) REFERENCES `project`
    (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;


CREATE TABLE `timesheet_log`
    (
  `id` int
    (11) NOT NULL AUTO_INCREMENT,
  `timesheet_id` int
    (11) DEFAULT NULL,
  `timesheet_date` DATE NULL,
  `employee_id` int
    (11) DEFAULT NULL,
	`time_from` DATETIME NULL,
	`time_to` DATETIME NULL,
	`duration` DATETIME NULL,
	`break` DATETIME NULL,
	`net` DATETIME NULL,
	`captured_by` VARCHAR(100) NULL,
	`status` VARCHAR(50) NULL,
	`comments` VARCHAR(255) NULL,
	`submission_timestamp` DATETIME NULL,
	`project_id` int
    (11) DEFAULT NULL,
  PRIMARY KEY
    (`id`),
  KEY `timesheet_id`
    (`timesheet_id`),
  KEY `employee_id`
    (`employee_id`),
  KEY `project_id`
    (`project_id`),
 CONSTRAINT `timesheet_id_ibfk_1` FOREIGN KEY
    (`timesheet_id`) REFERENCES `timesheet`
    (`id`),
  CONSTRAINT `employee_id_ibfk_2` FOREIGN KEY
    (`employee_id`) REFERENCES `user`
    (`id`),
  CONSTRAINT `project_id_ibfk_2` FOREIGN KEY
    (`project_id`) REFERENCES `project`
    (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;



CREATE TABLE `approval`
    (
  `id` int
    (11) NOT NULL AUTO_INCREMENT,
  `employee_id` int
    (11) DEFAULT NULL,
	`start_time` DATETIME NULL,
	`end_time` DATETIME NULL,
	`submission_timestamp` DATETIME NULL,
	`comments` VARCHAR(255) NULL,
	`approved_by` VARCHAR(100) NULL,
	`approved_timestamp` DATETIME NULL,
	`approver_id` int
    (11) DEFAULT NULL,
  PRIMARY KEY
    (`id`),
  KEY `employee_id`
    (`employee_id`),
	 KEY `approver_id`
    (`approver_id`),
  CONSTRAINT `employee_id_ibfk_3` FOREIGN KEY
    (`employee_id`) REFERENCES `user`
    (`id`),
	 CONSTRAINT `approver_id_ibfk_1` FOREIGN KEY
    (`approver_id`) REFERENCES `user`
    (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;




CREATE TABLE `approval_log`
    (
  `id` int
    (11) NOT NULL AUTO_INCREMENT,
  `approval_id` int
    (11) DEFAULT NULL,
  `employee_id` int
    (11) DEFAULT NULL,
	`start_time` DATETIME NULL,
	`end_time` DATETIME NULL,
	`submission_timestamp` DATETIME NULL,
	`comments` VARCHAR(255) NULL,
	`approved_by` VARCHAR(100) NULL,
	`approved_timestamp` DATETIME NULL,
	`approver_id` int
    (11) DEFAULT NULL,
  PRIMARY KEY
    (`id`),
  KEY `approval_id`
    (`approval_id`),
  KEY `employee_id`
    (`employee_id`),
	 KEY `approver_id`
    (`approver_id`),
  CONSTRAINT `employee_id_ibfk_5` FOREIGN KEY
    (`employee_id`) REFERENCES `user`
    (`id`),
 CONSTRAINT `approver_id_ibfk_2` FOREIGN KEY
    (`approver_id`) REFERENCES `user`
    (`id`),
 CONSTRAINT `approval_id_ibfk_2` FOREIGN KEY
    (`approval_id`) REFERENCES `approval`
    (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;


create user 'admin'@'%' identified by 'password#123';

GRANT ALL PRIVILEGES ON thools.* TO  'admin'@'%';



insert into thools.role(role) values ('ADMIN');
insert into thools.role(role) values ('USER');