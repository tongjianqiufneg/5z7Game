/* 模板表，所有表都必须有以下字段 */
drop table if exists temp;
create table temp (
        id int unsigned not null primary key auto_increment comment '主键',
        created_at  timestamp    default CURRENT_TIMESTAMP not null comment '创建时间',
        updated_at  timestamp    default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
        key `idx_updated_at`(`updated_at`)
)charset=utf8mb4 and comment='temp';

/* 用户基本信息 */
drop table if exists users;
create table users (
        id int unsigned not null primary key auto_increment comment '主键',
        username varchar(50) not null default '' comment '用户名',
        password varchar(100) not null default '' comment '密码',
        status tinyint unsigned not null default 1 comment '状态(1:启用,2:停用)',
        created_at  timestamp    default CURRENT_TIMESTAMP not null comment '创建时间',
        updated_at  timestamp    default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
        key `idx_updated_at`(`updated_at`)
)charset=utf8mb4 and comment='用户表';

/* 用户扩展信息 */
drop table if exists user_info;
create table user_info (
        id int unsigned not null primary key auto_increment comment '主键',
        user_id int unsigned not null default 0 comment '用户ID',
        user_level tinyint unsigned not null default 1 comment '用户段位',
        user_score int not null default 0 comment '分数',
        created_at  timestamp    default CURRENT_TIMESTAMP not null comment '创建时间',
        updated_at  timestamp    default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
        key `idx_updated_at`(`updated_at`)
)charset=utf8mb4 and comment='用户信息表';

/* 记录用户对战棋局的信息 */
drop table if exists chess;
create table chess (
        id int unsigned not null primary key auto_increment comment '主键',
        created_at  timestamp    default CURRENT_TIMESTAMP not null comment '创建时间',
        updated_at  timestamp    default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
        key `idx_updated_at`(`updated_at`)
)charset=utf8mb4 and comment='棋局';

/* 记录用户落子信息 */
drop table if exists board;
create table board (
        id int unsigned not null primary key auto_increment comment '主键',
        created_at  timestamp    default CURRENT_TIMESTAMP not null comment '创建时间',
        updated_at  timestamp    default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
        key `idx_updated_at`(`updated_at`)
)charset=utf8mb4 and comment='棋盘';