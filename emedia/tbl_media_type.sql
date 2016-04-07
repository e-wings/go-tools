/*
Navicat MySQL Data Transfer

Source Server         : 192.168.3.6
Source Server Version : 50623
Source Host           : 192.168.3.6:3306
Source Database       : zhongan

Target Server Type    : MYSQL
Target Server Version : 50623
File Encoding         : 65001

Date: 2016-04-07 16:19:33
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for tbl_media_type
-- ----------------------------
DROP TABLE IF EXISTS `tbl_media_type`;
CREATE TABLE `tbl_media_type` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(30) DEFAULT NULL COMMENT '类型名称',
  `width` tinyint(4) DEFAULT NULL COMMENT '宽',
  `height` tinyint(4) DEFAULT NULL COMMENT '高',
  `remark` text COMMENT '备注',
  `status` tinyint(4) DEFAULT '1' COMMENT '状态；0-删除；1-可用',
  `add_time` char(10) DEFAULT NULL,
  `update_time` char(10) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of tbl_media_type
-- ----------------------------
