-- 删除已存在的 ciot_user 表（如果存在）
DROP TABLE IF EXISTS ciot_user;

-- 创建 ciot_user 表
CREATE TABLE ciot_user (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,        -- ID字段，自增
    ext JSON,                                    -- 扩展字段，JSON类型
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- 创建时间
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- 更新时间
    deleted_at TIMESTAMP NULL,                   -- 删除时间
    username VARCHAR(100) NOT NULL UNIQUE,        -- 用户名，最大长度 100，不能为空，唯一
    password VARCHAR(255) NOT NULL                -- 密码，最大长度 255，不能为空
);

DROP TABLE IF EXISTS ciot_product;
CREATE TABLE `ciot_product` (
    `id` BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `user_id` BIGINT,
    `product_key` VARCHAR(255) NOT NULL UNIQUE,
    `product_name` VARCHAR(255) NOT NULL,
    `product_secret` VARCHAR(255) NOT NULL,
    `desc` TEXT,
    `ext` JSON,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP NULL,
    INDEX `idx_deleted_at` (`deleted_at`)
);


DROP TABLE IF EXISTS `ciot_device`;
CREATE TABLE `ciot_device` (
    `id` BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `product_id` BIGINT NOT NULL,
    `device_name` VARCHAR(255) NOT NULL,
    `device_secret` VARCHAR(255),
    `module_id` BIGINT,
    `desc` TEXT,
    `status` BIGINT DEFAULT 0,
    `last_online` TIMESTAMP NULL,
    `ext` JSON,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP NULL,
    INDEX `idx_deleted_at` (`deleted_at`)
);


-- 删除旧的 `ciot_module` 表（如果存在）
DROP TABLE IF EXISTS `ciot_module`;
CREATE TABLE `ciot_module` (
    `id` BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `module_name` VARCHAR(255) NOT NULL,
    `desc` TEXT,
    `ext` JSON,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP NULL,
    INDEX `idx_deleted_at` (`deleted_at`)
);

-- 删除旧的 `ciot_property` 表（如果存在）
DROP TABLE IF EXISTS `ciot_property`;
CREATE TABLE `ciot_property` (
    `id` BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `property_name` VARCHAR(255) NOT NULL,
    `identifier` VARCHAR(255) NOT NULL,
    `data_type` VARCHAR(255) NOT NULL,
    `desc` TEXT,
    `ext` JSON,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP NULL,
    INDEX `idx_deleted_at` (`deleted_at`)
);

-- 删除旧的 `ciot_module_property` 表（如果存在）
DROP TABLE IF EXISTS `ciot_module_property`;
CREATE TABLE `ciot_module_property` (
    `id` BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `module_id` BIGINT NOT NULL,
    `property_id` BIGINT NOT NULL,
    `ext` JSON,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP NULL,
    INDEX `idx_deleted_at` (`deleted_at`),
    INDEX `idx_module_id` (`module_id`),
    INDEX `idx_property_id` (`property_id`)
);

DROP TABLE IF EXISTS `ciot_device_video`;
CREATE TABLE `ciot_device_video` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `device_id` INT(11) NOT NULL COMMENT '关联的设备ID',
  `enable_video` TINYINT(1) DEFAULT 0 COMMENT '是否启用视频功能，0-禁用，1-启用',
  `video_url` VARCHAR(255) DEFAULT NULL COMMENT '视频流地址或 ffmpeg 服务地址',
  `stream_type` VARCHAR(50) DEFAULT NULL COMMENT '视频流类型，如 RTSP、HLS 等',
  `status` TINYINT(1) DEFAULT 0 COMMENT '状态，0-离线，1-在线',
  `ext` JSON,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` TIMESTAMP NULL,
  PRIMARY KEY (`id`)
)