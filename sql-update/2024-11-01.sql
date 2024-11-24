ALTER TABLE `business`.`swust_asset`
DROP COLUMN `order_id`,
ADD COLUMN `order_id` varchar(19) NULL COMMENT '采购订单号' AFTER `type_id`;