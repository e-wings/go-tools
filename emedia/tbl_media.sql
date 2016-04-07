/*
Navicat MySQL Data Transfer

Source Server         : 192.168.3.6
Source Server Version : 50623
Source Host           : 192.168.3.6:3306
Source Database       : zhongan

Target Server Type    : MYSQL
Target Server Version : 50623
File Encoding         : 65001

Date: 2016-04-07 16:19:26
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for tbl_media
-- ----------------------------
DROP TABLE IF EXISTS `tbl_media`;
CREATE TABLE `tbl_media` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `type_id` int(11) DEFAULT '0' COMMENT '图片类型id',
  `member_id` int(11) DEFAULT '0' COMMENT '上传用户id',
  `key_id` int(11) DEFAULT '0' COMMENT '关键字id',
  `url` varchar(300) DEFAULT NULL COMMENT '图片路径；名称',
  `status` tinyint(4) DEFAULT '1' COMMENT '状态；0-删除；1-只有自己可见；2-所有人都可见',
  `remark` text COMMENT '备注',
  `weight` tinyint(4) DEFAULT '50' COMMENT '权重；排序',
  `add_time` char(10) DEFAULT NULL,
  `update_time` char(10) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of tbl_media
-- ----------------------------
