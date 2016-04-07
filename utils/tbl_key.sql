/*
Navicat MySQL Data Transfer

Source Server         : 192.168.3.6
Source Server Version : 50623
Source Host           : 192.168.3.6:3306
Source Database       : zhongan

Target Server Type    : MYSQL
Target Server Version : 50623
File Encoding         : 65001

Date: 2016-04-07 14:06:36
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for tbl_key
-- ----------------------------
DROP TABLE IF EXISTS `tbl_key`;
CREATE TABLE `tbl_key` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(300) DEFAULT NULL,
  `remark` text,
  `status` tinyint(4) DEFAULT NULL,
  `weight` tinyint(4) DEFAULT NULL,
  `add_time` char(19) DEFAULT NULL,
  `update_time` char(19) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of tbl_key
-- ----------------------------
