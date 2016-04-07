/*
Navicat MySQL Data Transfer

Source Server         : 192.168.3.6
Source Server Version : 50623
Source Host           : 192.168.3.6:3306
Source Database       : zhongan

Target Server Type    : MYSQL
Target Server Version : 50623
File Encoding         : 65001

Date: 2016-04-07 14:06:44
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for tbl_media
-- ----------------------------
DROP TABLE IF EXISTS `tbl_media`;
CREATE TABLE `tbl_media` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `type_id` int(11) DEFAULT NULL,
  `member_id` int(11) DEFAULT NULL,
  `key_id` int(11) DEFAULT NULL,
  `url` varchar(300) DEFAULT NULL,
  `status` tinyint(4) DEFAULT NULL,
  `remark` text,
  `weight` tinyint(4) DEFAULT NULL,
  `add_time` char(10) DEFAULT NULL,
  `update_time` char(10) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of tbl_media
-- ----------------------------
