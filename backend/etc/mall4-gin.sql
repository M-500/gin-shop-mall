/*
 Navicat Premium Data Transfer

 Source Server         : 小新
 Source Server Type    : MySQL
 Source Server Version : 80033
 Source Host           : 192.168.1.52:3306
 Source Schema         : mall4-gin

 Target Server Type    : MySQL
 Target Server Version : 80033
 File Encoding         : 65001

 Date: 09/01/2024 18:14:55
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for m4_sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `m4_sys_menu`;
CREATE TABLE `m4_sys_menu`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键自增长ID',
  `create_at` datetime(3) NULL DEFAULT NULL COMMENT '记录创建时间',
  `update_at` datetime(3) NULL DEFAULT NULL COMMENT '记录更新时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `parent_id` bigint NULL DEFAULT NULL COMMENT '父ID',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '菜单名称',
  `url` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '菜单URL',
  `perms` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '授权(多个用逗号分隔，如：user:list,user:create)',
  `type` bigint NULL DEFAULT NULL COMMENT '类型   0：目录   1：菜单   2：按钮',
  `icon` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '菜单图标',
  `order_num` bigint NULL DEFAULT NULL COMMENT '排序',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 317 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of m4_sys_menu
-- ----------------------------
INSERT INTO `m4_sys_menu` VALUES (1, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 0, '系统管理', '', '', 0, 'system', 3);
INSERT INTO `m4_sys_menu` VALUES (2, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 1, '管理员列表', 'sys/user', '', 1, 'admin', 1);
INSERT INTO `m4_sys_menu` VALUES (3, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 1, '角色管理', 'sys/role', '', 1, 'role', 2);
INSERT INTO `m4_sys_menu` VALUES (4, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 1, '菜单管理', 'sys/menu', '', 1, 'menu', 3);
INSERT INTO `m4_sys_menu` VALUES (15, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 2, '查看', NULL, 'sys:user:page,sys:user:info', 2, NULL, 0);
INSERT INTO `m4_sys_menu` VALUES (16, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 2, '新增', '', 'sys:user:save,sys:role:list', 2, '', 1);
INSERT INTO `m4_sys_menu` VALUES (17, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 2, '修改', '', 'sys:user:update,sys:role:list', 2, '', 2);
INSERT INTO `m4_sys_menu` VALUES (18, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 2, '删除', '', 'sys:user:delete', 2, '', 3);
INSERT INTO `m4_sys_menu` VALUES (19, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 3, '查看', NULL, 'sys:role:page,sys:role:info', 2, NULL, 0);
INSERT INTO `m4_sys_menu` VALUES (20, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 3, '新增', NULL, 'sys:role:save,sys:menu:list', 2, NULL, 0);
INSERT INTO `m4_sys_menu` VALUES (21, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 3, '修改', NULL, 'sys:role:update,sys:menu:list', 2, NULL, 0);
INSERT INTO `m4_sys_menu` VALUES (22, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 3, '删除', NULL, 'sys:role:delete', 2, NULL, 0);
INSERT INTO `m4_sys_menu` VALUES (23, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 4, '查看', NULL, 'sys:menu:list,sys:menu:info', 2, NULL, 0);
INSERT INTO `m4_sys_menu` VALUES (24, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 4, '新增', NULL, 'sys:menu:save,sys:menu:select', 2, NULL, 0);
INSERT INTO `m4_sys_menu` VALUES (25, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 4, '修改', NULL, 'sys:menu:update,sys:menu:select', 2, NULL, 0);
INSERT INTO `m4_sys_menu` VALUES (26, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 4, '删除', NULL, 'sys:menu:delete', 2, NULL, 0);
INSERT INTO `m4_sys_menu` VALUES (27, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 1, '参数管理', 'sys/config', 'sys:config:page,sys:config:info,sys:config:save,sys:config:update,sys:config:delete', 1, 'config', 6);
INSERT INTO `m4_sys_menu` VALUES (29, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 1, '系统日志', 'sys/log', 'sys:log:page', 1, 'log', 7);
INSERT INTO `m4_sys_menu` VALUES (34, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 0, '产品管理', '', '', 0, 'admin', 0);
INSERT INTO `m4_sys_menu` VALUES (51, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 34, '规格管理', 'prod/spec', '', 1, '', 2);
INSERT INTO `m4_sys_menu` VALUES (57, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 51, '查看', '', 'prod:spec:page', 2, '', 0);
INSERT INTO `m4_sys_menu` VALUES (58, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 51, '添加', '', 'prod:spec:save', 2, '', 1);
INSERT INTO `m4_sys_menu` VALUES (59, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 51, '修改', '', 'prod:spec:update,prod:spec:info', 2, '', 2);
INSERT INTO `m4_sys_menu` VALUES (60, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 51, '删除', '', 'prod:spec:delete', 2, '', 3);
INSERT INTO `m4_sys_menu` VALUES (63, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 0, '门店管理', '', '', 0, 'store', 0);
INSERT INTO `m4_sys_menu` VALUES (70, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 34, '产品管理', 'prod/prodList', '', 1, NULL, 0);
INSERT INTO `m4_sys_menu` VALUES (71, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 70, '产品管理', '', 'prod:prod:page', 2, NULL, 0);
INSERT INTO `m4_sys_menu` VALUES (72, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 70, '查看产品', '', 'prod:prod:info', 2, NULL, 0);
INSERT INTO `m4_sys_menu` VALUES (73, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 70, '新增产品', '', 'prod:prod:save', 2, NULL, 0);
INSERT INTO `m4_sys_menu` VALUES (74, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 70, '修改产品', '', 'prod:prod:update', 2, NULL, 0);
INSERT INTO `m4_sys_menu` VALUES (75, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 70, '删除产品', '', 'prod:prod:delete', 2, NULL, 0);
INSERT INTO `m4_sys_menu` VALUES (91, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 0, '订单管理', '', '', 0, 'order', 2);
INSERT INTO `m4_sys_menu` VALUES (92, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 91, '订单管理', 'order/order', '', 1, NULL, 1);
INSERT INTO `m4_sys_menu` VALUES (93, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 92, '查看', '', 'order:order:page', 2, NULL, 0);
INSERT INTO `m4_sys_menu` VALUES (99, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 92, '保存', '', 'order:order:save', 2, NULL, 0);
INSERT INTO `m4_sys_menu` VALUES (100, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 92, '修改', '', 'order:order:update', 2, NULL, 0);
INSERT INTO `m4_sys_menu` VALUES (101, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 92, '删除', '', 'order:order:delete', 2, NULL, 0);
INSERT INTO `m4_sys_menu` VALUES (107, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 92, '详情', '', 'order:order:info', 2, NULL, 0);
INSERT INTO `m4_sys_menu` VALUES (108, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 92, '支付', '', 'order:order:pay', 2, NULL, 0);
INSERT INTO `m4_sys_menu` VALUES (125, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 0, '会员管理', '', '', 0, 'vip', 0);
INSERT INTO `m4_sys_menu` VALUES (126, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 125, '会员管理', 'user/user', '', 1, NULL, 0);
INSERT INTO `m4_sys_menu` VALUES (127, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 126, '查看', '', 'admin:user:page', 2, NULL, 0);
INSERT INTO `m4_sys_menu` VALUES (128, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 126, '新增', '', 'admin:user:save', 2, NULL, 0);
INSERT INTO `m4_sys_menu` VALUES (129, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 126, '修改', '', 'admin:user:update,admin:user:info', 2, NULL, 0);
INSERT INTO `m4_sys_menu` VALUES (130, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 126, '删除', '', 'admin:user:delete', 2, NULL, 0);
INSERT INTO `m4_sys_menu` VALUES (131, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 63, '自提点管理', 'shop/pickAddr', '', 1, '', 0);
INSERT INTO `m4_sys_menu` VALUES (132, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 131, '查看', '', 'shop:pickAddr:page', 2, NULL, 0);
INSERT INTO `m4_sys_menu` VALUES (133, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 131, '保存', '', 'shop:pickAddr:save', 2, NULL, 0);
INSERT INTO `m4_sys_menu` VALUES (134, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 131, '修改', '', 'shop:pickAddr:update,shop:pickAddr:info', 2, NULL, 0);
INSERT INTO `m4_sys_menu` VALUES (135, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 131, '删除', '', 'shop:pickAddr:delete', 2, NULL, 0);
INSERT INTO `m4_sys_menu` VALUES (136, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 34, '分类管理', 'prod/category', '', 1, NULL, 0);
INSERT INTO `m4_sys_menu` VALUES (137, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 136, '查看', '', 'prod:category:page', 2, NULL, 0);
INSERT INTO `m4_sys_menu` VALUES (138, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 136, '新增', '', 'prod:category:save', 2, NULL, 0);
INSERT INTO `m4_sys_menu` VALUES (139, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 136, '修改', '', 'prod:category:info,prod:category:update', 2, NULL, 0);
INSERT INTO `m4_sys_menu` VALUES (140, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 136, '删除', '', 'prod:category:delete', 2, NULL, 0);
INSERT INTO `m4_sys_menu` VALUES (146, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 92, '发货', '', 'order:order:delivery', 2, '', 0);
INSERT INTO `m4_sys_menu` VALUES (163, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 63, '运费模板', 'shop/transport', '', 1, NULL, 0);
INSERT INTO `m4_sys_menu` VALUES (164, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 163, '查看', '', 'shop:transport:page,shop:shopDetail:info,shop:transfee:info,admin:area:page,shop:transcity:info', 2, '', 0);
INSERT INTO `m4_sys_menu` VALUES (165, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 163, '修改', '', 'shop:transport:update,shop:transport:info,shop:transfee:update,admin:area:page', 2, '', 0);
INSERT INTO `m4_sys_menu` VALUES (166, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 163, '新增', '', 'shop:transport:save,shop:transfee:save', 2, '', 0);
INSERT INTO `m4_sys_menu` VALUES (167, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 163, '删除', '', 'shop:transport:delete,shop:transfee:delete', 2, '', 0);
INSERT INTO `m4_sys_menu` VALUES (174, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 63, '轮播图管理', 'admin/indexImg', '', 1, '', 0);
INSERT INTO `m4_sys_menu` VALUES (175, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 174, '查看', '', 'admin:indexImg:page', 2, '', 0);
INSERT INTO `m4_sys_menu` VALUES (176, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 174, '新增', '', 'admin:indexImg:save', 2, '', 0);
INSERT INTO `m4_sys_menu` VALUES (177, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 174, '修改', '', 'admin:indexImg:info,admin:indexImg:update', 2, '', 0);
INSERT INTO `m4_sys_menu` VALUES (178, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 174, '删除', '', 'admin:indexImg:delete', 2, '', 0);
INSERT INTO `m4_sys_menu` VALUES (184, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 92, '导出待发货订单', '', 'order:order:waitingConsignmentExcel', 2, '', 0);
INSERT INTO `m4_sys_menu` VALUES (185, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 92, '导出销售记录', '', 'order:order:soldExcel', 2, '', 0);
INSERT INTO `m4_sys_menu` VALUES (201, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 63, '热搜管理', 'shop/hotSearch', '', 1, '', 0);
INSERT INTO `m4_sys_menu` VALUES (202, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 201, '查询热搜', '', 'admin:hotSearch:page', 2, '', 0);
INSERT INTO `m4_sys_menu` VALUES (203, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 201, '查询热搜详情', '', 'admin:hotSearch:page', 2, '', 0);
INSERT INTO `m4_sys_menu` VALUES (204, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 201, '添加热搜', '', 'admin:hotSearch:save', 2, '', 0);
INSERT INTO `m4_sys_menu` VALUES (205, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 201, '修改热搜', '', 'admin:hotSearch:update', 2, '', 0);
INSERT INTO `m4_sys_menu` VALUES (206, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 201, '删除热搜', '', 'admin:hotSearch:delete', 2, '', 0);
INSERT INTO `m4_sys_menu` VALUES (230, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 126, '添加', '', 'user:addr:save', 2, '', 0);
INSERT INTO `m4_sys_menu` VALUES (239, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 34, '分组管理', 'prod/prodTag', 'prod:prodTag', 1, '', 0);
INSERT INTO `m4_sys_menu` VALUES (240, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 239, '添加商品分组', '', 'prod:prodTag:save', 2, '', 0);
INSERT INTO `m4_sys_menu` VALUES (241, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 239, '修改商品分组', '', 'prod:prodTag:update', 2, '', 0);
INSERT INTO `m4_sys_menu` VALUES (242, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 239, '删除商品分组', '', 'prod:prodTag:delete', 2, '', 0);
INSERT INTO `m4_sys_menu` VALUES (243, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 239, '查看商品分组', '', 'prod:prodTag:info,prod:prodTag:page', 2, '', 0);
INSERT INTO `m4_sys_menu` VALUES (300, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 63, '公告管理', 'shop/notice', '', 1, '', 0);
INSERT INTO `m4_sys_menu` VALUES (301, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 300, '添加公告', '', 'shop:notice:save', 2, '', 0);
INSERT INTO `m4_sys_menu` VALUES (302, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 300, '修改公告', '', 'shop:notice:update', 2, '', 0);
INSERT INTO `m4_sys_menu` VALUES (303, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 300, '查看公告', '', 'shop:notice:info,shop:notice:page', 2, '', 0);
INSERT INTO `m4_sys_menu` VALUES (305, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 300, '删除公告', '', 'shop:notice:delete', 2, '', 0);
INSERT INTO `m4_sys_menu` VALUES (306, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 34, '评论管理', 'prod/prodComm', '', 1, '', 1);
INSERT INTO `m4_sys_menu` VALUES (307, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 306, '查看', '', 'prod:prodComm:page,prod:prodComm:info', 2, '', 0);
INSERT INTO `m4_sys_menu` VALUES (308, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 306, '添加', '', 'prod:prodComm:save', 2, '', 0);
INSERT INTO `m4_sys_menu` VALUES (309, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 306, '修改', '', 'prod:prodComm:update', 2, '', 0);
INSERT INTO `m4_sys_menu` VALUES (310, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 306, '删除', '', 'prod:prodComm:delete', 2, '', 0);
INSERT INTO `m4_sys_menu` VALUES (312, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 1, '地址管理', 'sys/area', '', 1, 'dangdifill', 0);
INSERT INTO `m4_sys_menu` VALUES (313, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 312, '新增地址', '', 'admin:area:save', 2, '', 0);
INSERT INTO `m4_sys_menu` VALUES (314, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 312, '修改地址', '', 'admin:area:update', 2, '', 0);
INSERT INTO `m4_sys_menu` VALUES (315, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 312, '删除地址', '', 'admin:area:delete', 2, '', 0);
INSERT INTO `m4_sys_menu` VALUES (316, '2024-01-09 11:07:24.000', '2024-01-09 11:07:24.000', NULL, 312, '查看地址', '', 'admin:area:info,admin:area:page,admin:area:list', 2, '', 0);

-- ----------------------------
-- Table structure for m4_sys_role
-- ----------------------------
DROP TABLE IF EXISTS `m4_sys_role`;
CREATE TABLE `m4_sys_role`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键自增长ID',
  `create_at` datetime(3) NULL DEFAULT NULL COMMENT '记录创建时间',
  `update_at` datetime(3) NULL DEFAULT NULL COMMENT '记录更新时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `role_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '角色名称',
  `remark` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注',
  `create_user_id` bigint NULL DEFAULT NULL COMMENT '创建人ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of m4_sys_role
-- ----------------------------
INSERT INTO `m4_sys_role` VALUES (1, '2019-07-03 08:39:49.000', NULL, NULL, '管理员', '测试', NULL);

-- ----------------------------
-- Table structure for m4_sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `m4_sys_role_menu`;
CREATE TABLE `m4_sys_role_menu`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键自增长ID',
  `create_at` datetime(3) NULL DEFAULT NULL COMMENT '记录创建时间',
  `update_at` datetime(3) NULL DEFAULT NULL COMMENT '记录更新时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `role_id` bigint NULL DEFAULT NULL,
  `menu_id` bigint NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 106 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of m4_sys_role_menu
-- ----------------------------
INSERT INTO `m4_sys_role_menu` VALUES (1, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 34);
INSERT INTO `m4_sys_role_menu` VALUES (2, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 70);
INSERT INTO `m4_sys_role_menu` VALUES (3, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 71);
INSERT INTO `m4_sys_role_menu` VALUES (4, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 72);
INSERT INTO `m4_sys_role_menu` VALUES (5, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 73);
INSERT INTO `m4_sys_role_menu` VALUES (6, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 74);
INSERT INTO `m4_sys_role_menu` VALUES (7, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 75);
INSERT INTO `m4_sys_role_menu` VALUES (8, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 136);
INSERT INTO `m4_sys_role_menu` VALUES (9, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 137);
INSERT INTO `m4_sys_role_menu` VALUES (10, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 138);
INSERT INTO `m4_sys_role_menu` VALUES (11, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 139);
INSERT INTO `m4_sys_role_menu` VALUES (12, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 140);
INSERT INTO `m4_sys_role_menu` VALUES (13, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 239);
INSERT INTO `m4_sys_role_menu` VALUES (14, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 240);
INSERT INTO `m4_sys_role_menu` VALUES (15, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 241);
INSERT INTO `m4_sys_role_menu` VALUES (16, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 242);
INSERT INTO `m4_sys_role_menu` VALUES (17, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 243);
INSERT INTO `m4_sys_role_menu` VALUES (18, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 306);
INSERT INTO `m4_sys_role_menu` VALUES (19, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 307);
INSERT INTO `m4_sys_role_menu` VALUES (20, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 308);
INSERT INTO `m4_sys_role_menu` VALUES (21, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 309);
INSERT INTO `m4_sys_role_menu` VALUES (22, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 310);
INSERT INTO `m4_sys_role_menu` VALUES (23, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 51);
INSERT INTO `m4_sys_role_menu` VALUES (24, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 57);
INSERT INTO `m4_sys_role_menu` VALUES (25, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 58);
INSERT INTO `m4_sys_role_menu` VALUES (26, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 59);
INSERT INTO `m4_sys_role_menu` VALUES (27, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 60);
INSERT INTO `m4_sys_role_menu` VALUES (28, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 63);
INSERT INTO `m4_sys_role_menu` VALUES (29, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 300);
INSERT INTO `m4_sys_role_menu` VALUES (30, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 301);
INSERT INTO `m4_sys_role_menu` VALUES (31, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 302);
INSERT INTO `m4_sys_role_menu` VALUES (32, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 303);
INSERT INTO `m4_sys_role_menu` VALUES (33, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 305);
INSERT INTO `m4_sys_role_menu` VALUES (34, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 131);
INSERT INTO `m4_sys_role_menu` VALUES (35, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 132);
INSERT INTO `m4_sys_role_menu` VALUES (36, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 133);
INSERT INTO `m4_sys_role_menu` VALUES (37, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 134);
INSERT INTO `m4_sys_role_menu` VALUES (38, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 135);
INSERT INTO `m4_sys_role_menu` VALUES (39, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 163);
INSERT INTO `m4_sys_role_menu` VALUES (40, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 164);
INSERT INTO `m4_sys_role_menu` VALUES (41, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 165);
INSERT INTO `m4_sys_role_menu` VALUES (42, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 166);
INSERT INTO `m4_sys_role_menu` VALUES (43, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 167);
INSERT INTO `m4_sys_role_menu` VALUES (44, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 174);
INSERT INTO `m4_sys_role_menu` VALUES (45, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 175);
INSERT INTO `m4_sys_role_menu` VALUES (46, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 176);
INSERT INTO `m4_sys_role_menu` VALUES (47, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 177);
INSERT INTO `m4_sys_role_menu` VALUES (48, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 178);
INSERT INTO `m4_sys_role_menu` VALUES (49, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 201);
INSERT INTO `m4_sys_role_menu` VALUES (50, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 202);
INSERT INTO `m4_sys_role_menu` VALUES (51, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 203);
INSERT INTO `m4_sys_role_menu` VALUES (52, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 204);
INSERT INTO `m4_sys_role_menu` VALUES (53, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 205);
INSERT INTO `m4_sys_role_menu` VALUES (54, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 206);
INSERT INTO `m4_sys_role_menu` VALUES (55, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 125);
INSERT INTO `m4_sys_role_menu` VALUES (56, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 126);
INSERT INTO `m4_sys_role_menu` VALUES (57, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 127);
INSERT INTO `m4_sys_role_menu` VALUES (58, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 128);
INSERT INTO `m4_sys_role_menu` VALUES (59, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 129);
INSERT INTO `m4_sys_role_menu` VALUES (60, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 130);
INSERT INTO `m4_sys_role_menu` VALUES (61, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 230);
INSERT INTO `m4_sys_role_menu` VALUES (62, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 91);
INSERT INTO `m4_sys_role_menu` VALUES (63, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 92);
INSERT INTO `m4_sys_role_menu` VALUES (64, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 93);
INSERT INTO `m4_sys_role_menu` VALUES (65, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 99);
INSERT INTO `m4_sys_role_menu` VALUES (66, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 100);
INSERT INTO `m4_sys_role_menu` VALUES (67, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 101);
INSERT INTO `m4_sys_role_menu` VALUES (68, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 107);
INSERT INTO `m4_sys_role_menu` VALUES (69, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 108);
INSERT INTO `m4_sys_role_menu` VALUES (70, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 146);
INSERT INTO `m4_sys_role_menu` VALUES (71, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 184);
INSERT INTO `m4_sys_role_menu` VALUES (72, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 185);
INSERT INTO `m4_sys_role_menu` VALUES (73, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 1);
INSERT INTO `m4_sys_role_menu` VALUES (74, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 312);
INSERT INTO `m4_sys_role_menu` VALUES (75, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 313);
INSERT INTO `m4_sys_role_menu` VALUES (76, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 314);
INSERT INTO `m4_sys_role_menu` VALUES (77, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 315);
INSERT INTO `m4_sys_role_menu` VALUES (78, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 316);
INSERT INTO `m4_sys_role_menu` VALUES (79, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 2);
INSERT INTO `m4_sys_role_menu` VALUES (80, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 15);
INSERT INTO `m4_sys_role_menu` VALUES (81, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 16);
INSERT INTO `m4_sys_role_menu` VALUES (82, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 17);
INSERT INTO `m4_sys_role_menu` VALUES (83, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 18);
INSERT INTO `m4_sys_role_menu` VALUES (84, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 3);
INSERT INTO `m4_sys_role_menu` VALUES (85, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 19);
INSERT INTO `m4_sys_role_menu` VALUES (86, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 20);
INSERT INTO `m4_sys_role_menu` VALUES (87, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 21);
INSERT INTO `m4_sys_role_menu` VALUES (88, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 22);
INSERT INTO `m4_sys_role_menu` VALUES (89, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 4);
INSERT INTO `m4_sys_role_menu` VALUES (90, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 23);
INSERT INTO `m4_sys_role_menu` VALUES (91, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 24);
INSERT INTO `m4_sys_role_menu` VALUES (92, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 25);
INSERT INTO `m4_sys_role_menu` VALUES (93, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 26);
INSERT INTO `m4_sys_role_menu` VALUES (94, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 6);
INSERT INTO `m4_sys_role_menu` VALUES (95, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 7);
INSERT INTO `m4_sys_role_menu` VALUES (96, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 8);
INSERT INTO `m4_sys_role_menu` VALUES (97, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 9);
INSERT INTO `m4_sys_role_menu` VALUES (98, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 10);
INSERT INTO `m4_sys_role_menu` VALUES (99, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 11);
INSERT INTO `m4_sys_role_menu` VALUES (100, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 12);
INSERT INTO `m4_sys_role_menu` VALUES (101, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 13);
INSERT INTO `m4_sys_role_menu` VALUES (102, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 14);
INSERT INTO `m4_sys_role_menu` VALUES (103, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 27);
INSERT INTO `m4_sys_role_menu` VALUES (104, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, 29);
INSERT INTO `m4_sys_role_menu` VALUES (105, '2024-01-09 11:09:58.000', '2024-01-09 11:09:58.000', NULL, 1, -666666);

-- ----------------------------
-- Table structure for m4_sys_user
-- ----------------------------
DROP TABLE IF EXISTS `m4_sys_user`;
CREATE TABLE `m4_sys_user`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键自增长ID',
  `create_at` datetime(3) NULL DEFAULT NULL COMMENT '记录创建时间',
  `update_at` datetime(3) NULL DEFAULT NULL COMMENT '记录更新时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `username` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '用户账号',
  `password` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '密码',
  `email` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '邮箱',
  `mobile` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '手机号',
  `status` tinyint NULL DEFAULT NULL COMMENT '状态 0：禁用 1:正常',
  `creatUserId` bigint NULL DEFAULT NULL COMMENT '创建人ID',
  `shopId` bigint NULL DEFAULT NULL COMMENT '用户所在的商场ID',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_m4_sys_user_username`(`username` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of m4_sys_user
-- ----------------------------
INSERT INTO `m4_sys_user` VALUES (1, '2024-01-09 18:14:10.155', '2024-01-09 18:14:10.155', NULL, 'admin', '$pbkdf2-sha512$ZH1Qqtnnok5Euzav$d5221631bdbb73b051c5ee0e2024804f58fe75b65c1f5a81c263c6fdcb43b474', '', '13333333333', 0, 0, 0);

-- ----------------------------
-- Table structure for m4_sys_user_role
-- ----------------------------
DROP TABLE IF EXISTS `m4_sys_user_role`;
CREATE TABLE `m4_sys_user_role`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键自增长ID',
  `create_at` datetime(3) NULL DEFAULT NULL COMMENT '记录创建时间',
  `update_at` datetime(3) NULL DEFAULT NULL COMMENT '记录更新时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `user_id` bigint NULL DEFAULT NULL,
  `role_id` bigint NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of m4_sys_user_role
-- ----------------------------

-- ----------------------------
-- Table structure for m4_user
-- ----------------------------
DROP TABLE IF EXISTS `m4_user`;
CREATE TABLE `m4_user`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键自增长ID',
  `create_at` datetime(3) NULL DEFAULT NULL COMMENT '记录创建时间',
  `update_at` datetime(3) NULL DEFAULT NULL COMMENT '记录更新时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `user_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '用户ID',
  `nick_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '昵称',
  `real_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '真实姓名',
  `user_email` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '邮箱',
  `login_pwd` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '登录密码',
  `pay_pwd` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '支付密码',
  `user_mobile` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '手机号',
  `modify_time` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `user_reg_ip` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT 'NULL' COMMENT '注册IP',
  `user_last_time` datetime NULL DEFAULT NULL COMMENT '最后登录时间',
  `user_memo` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '用户备注',
  `sex` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '性别M(男) or F(女)',
  `birth_day` char(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '例如：2009-11-27',
  `cover_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '头像图片路径',
  `score` bigint NULL DEFAULT NULL COMMENT '用户积分',
  `status` bigint NULL DEFAULT NULL COMMENT '状态 1 正常 0 无效',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_m4_user_user_id`(`user_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of m4_user
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
