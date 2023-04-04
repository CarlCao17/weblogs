create table `t_blog`
(
    `id` bigint unsigned auto_increment,
    `created_at` datetime(3) NULL,
    `updated_at` datetime(3) NULL,
    `deleted_at` datetime(3) NULL,
    `blog_id` bigint(20) NOT NULL,
    `title` varchar(255) NOT NULL DEFAULT '',
    `author_id` bigint(20) NOT NULL,
    `content` ,
    `description` text NULL,
    `category` ,
    `status` tinyint(4) NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`)
    KEY `idx_blog_id` (`blog_id`) COMMENT 'blog_id index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='blog table';

create table `t_user`
()

create table `t_comment`
()

create table `t_label`
()

create table `t_tag`
()

create table `t_blog_tag`
()
