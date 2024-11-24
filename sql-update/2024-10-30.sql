ALTER TABLE `business`.`swust_asset`
    MODIFY COLUMN `id` varchar(19) NOT NULL COMMENT 'id' FIRST,
    MODIFY COLUMN `producer_id` varchar(19) NULL DEFAULT NULL COMMENT '供应商ID' AFTER `create_time`,
    MODIFY COLUMN `type_id` varchar(19) NULL DEFAULT NULL COMMENT '类型ID' AFTER `producer_id`;
ALTER TABLE `business`.`swust_asset`
    MODIFY COLUMN `user_id` varchar(19) NULL DEFAULT NULL COMMENT '责任人' AFTER `code`;
ALTER TABLE `business`.`swust_asset_type`
    MODIFY COLUMN `id` varchar(19) NOT NULL COMMENT 'id' FIRST;

ALTER TABLE `business`.`swust_order`
    MODIFY COLUMN `id` varchar(19) NOT NULL COMMENT 'id' FIRST;

ALTER TABLE `business`.`swust_producer`
    MODIFY COLUMN `id` varchar(19) NOT NULL COMMENT 'id' FIRST;
ALTER TABLE `business`.`swust_room`
    MODIFY COLUMN `id` varchar(19) NOT NULL COMMENT 'id' FIRST;
ALTER TABLE `business`.`swust_system_menu`
    MODIFY COLUMN `id` varchar(19) NOT NULL FIRST;

ALTER TABLE `business`.`swust_system_user`
    MODIFY COLUMN `id` varchar(19) NOT NULL COMMENT 'id' FIRST;

ALTER TABLE `business`.`swust_warehouse`
    CHANGE COLUMN `number` `quality` int NULL DEFAULT NULL COMMENT '存放数量' AFTER `asset_id`,
    MODIFY COLUMN `id` varchar(19) NOT NULL COMMENT 'id' FIRST,
    MODIFY COLUMN `room_id` varchar(19) NULL DEFAULT NULL COMMENT '房间号' AFTER `id`,
    MODIFY COLUMN `asset_id` varchar(19) NULL DEFAULT NULL COMMENT '存放物品数量' AFTER `room_id`;
ALTER TABLE `business`.`swust_order`
    MODIFY COLUMN `user_id` varchar NULL DEFAULT NULL COMMENT '采购申请人' AFTER `name`,
    MODIFY COLUMN `allow_user_id` varchar(19) NULL DEFAULT NULL COMMENT '审批人' AFTER `status`;
ALTER TABLE `business`.`swust_system_menu`
    MODIFY COLUMN `parent_id` varchar(19) NULL DEFAULT NULL COMMENT '父节点id' AFTER `level`;

ALTER TABLE `business`.`swust_room`
    MODIFY COLUMN `header` varchar(19) NULL DEFAULT NULL COMMENT '房间负责人Id' AFTER `location`;

-- 晚上
ALTER TABLE `business`.`swust_asset`
    ADD UNIQUE INDEX `asset_code_unique`(`code`) COMMENT '设备唯一键';
ALTER TABLE `business`.`swust_asset`
    MODIFY COLUMN `status` int NULL DEFAULT NULL COMMENT '0 在用 1 故障 2 维修 3 报废 4 闲置' AFTER `price`;