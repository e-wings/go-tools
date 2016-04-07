/*
Navicat MySQL Data Transfer

Source Server         : 192.168.3.6
Source Server Version : 50623
Source Host           : 192.168.3.6:3306
Source Database       : zhongan

Target Server Type    : MYSQL
Target Server Version : 50623
File Encoding         : 65001

Date: 2016-04-07 16:19:18
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for tbl_key
-- ----------------------------
DROP TABLE IF EXISTS `tbl_key`;
CREATE TABLE `tbl_key` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(300) DEFAULT NULL COMMENT '关键字名称',
  `remark` text COMMENT '备注',
  `status` tinyint(4) DEFAULT '1' COMMENT '状态；0-关闭；1-可用',
  `weight` tinyint(4) DEFAULT '50' COMMENT '权重；排序',
  `add_time` char(19) DEFAULT NULL,
  `update_time` char(19) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of tbl_key
-- ----------------------------
