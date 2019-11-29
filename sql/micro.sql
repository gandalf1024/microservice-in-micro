/*
 Navicat MySQL Data Transfer

 Source Server         : windows_centos
 Source Server Type    : MySQL
 Source Server Version : 50719
 Source Host           : 192.168.59.137:3306
 Source Schema         : micro

 Target Server Type    : MySQL
 Target Server Version : 50719
 File Encoding         : 65001

 Date: 29/11/2019 18:02:18
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for inventory
-- ----------------------------
DROP TABLE IF EXISTS `inventory`;
CREATE TABLE `inventory`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `book_id` int(10) UNSIGNED NOT NULL COMMENT '书id',
  `unit_price` int(10) UNSIGNED NOT NULL COMMENT '单价',
  `stock` int(10) UNSIGNED NOT NULL COMMENT '总数',
  `version` int(10) UNSIGNED NOT NULL COMMENT '版本号',
  `created_time` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `updated_time` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '库存表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for inventory_history
-- ----------------------------
DROP TABLE IF EXISTS `inventory_history`;
CREATE TABLE `inventory_history`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `book_id` int(10) UNSIGNED NOT NULL COMMENT '书id',
  `user_id` int(10) UNSIGNED NOT NULL COMMENT '单价',
  `state` int(10) UNSIGNED NOT NULL COMMENT '未出库1，出库2',
  `created_time` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `updated_time` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '订单表历史' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for orders
-- ----------------------------
DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` int(10) UNSIGNED NULL DEFAULT NULL COMMENT '用户id',
  `book_id` int(10) NOT NULL COMMENT '书id',
  `inv_his_id` int(10) NOT NULL COMMENT '库存历史记录id',
  `state` tinyint(1) NOT NULL COMMENT '1未支付，2支付',
  `created_time` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `updated_time` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '订单表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for payment
-- ----------------------------
DROP TABLE IF EXISTS `payment`;
CREATE TABLE `payment`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` int(10) UNSIGNED NULL DEFAULT NULL COMMENT '用户id',
  `book_id` int(10) NOT NULL COMMENT '书id',
  `order_id` int(10) NOT NULL COMMENT '订单id',
  `inv_his_id` int(10) NOT NULL COMMENT '库存历史id',
  `state` tinyint(1) NOT NULL COMMENT '1：未支付，2：支付',
  `created_time` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `updated_time` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '支付表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` int(10) UNSIGNED NULL DEFAULT NULL COMMENT '用户id',
  `user_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `pwd` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
  `created_time` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `updated_time` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `user_user_name_uindex`(`user_name`) USING BTREE,
  UNIQUE INDEX `user_user_id_uindex`(`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '用户表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 10001, 'micro', '123', '2019-11-28 11:16:10.188', '2019-11-28 11:16:10.188');

SET FOREIGN_KEY_CHECKS = 1;
