create table blog_tag (
    id int unsigned not null auto_increment,
    name varchar(100) default '' comment '标签名称',
    created_on int unsigned default 0 comment '创建时间',
    created_by varchar(100) default '' comment '创建人',
    modified_on int unsigned default 0 comment '修改时间',
    modified_by varchar(100) default '' comment '修改人',
    deleted_on int unsigned default 0 comment '删除时间',
    is_del tinyint unsigned default 0 comment '是否删除 0-未删除 1-已删除',
    state tinyint unsigned default 1 comment '状态 0-禁用 1-启用',
    primary key(id)
)engine=innodb default charset=utf8mb4 comment='标签管理';

create table blog_article (
    id int unsigned not null auto_increment,
    title varchar(100) default '' comment '文章标题',
    description varchar(255) default '' comment '文章简述',
    cover_image_url varchar(255) default '' comment '封面图片地址',
    content longtext comment '文章内容',
    created_on int unsigned default 0 comment '创建时间',
    created_by varchar(100) default '' comment '创建人',
    modified_on int unsigned default 0 comment '修改时间',
    modified_by varchar(100) default '' comment '修改人',
    deleted_on int unsigned default 0 comment '删除时间',
    is_del tinyint unsigned default 0 comment '是否删除 0-未删除 1-已删除',
    state tinyint unsigned default 1 comment '状态 0-禁用 1-启用',
    primary key(id)
)engine=innodb default charset=utf8mb4 comment='文章管理';

create table blog_article_tag (
    id int unsigned not null auto_increment,
    article_id int unsigned not null comment '文章id',
    tag_id int unsigned not null comment '标签id',
    created_on int unsigned default 0 comment '创建时间',
    created_by varchar(100) default '' comment '创建人',
    modified_on int unsigned default 0 comment '修改时间',
    modified_by varchar(100) default '' comment '修改人',
    deleted_on int unsigned default 0 comment '删除时间',
    is_del tinyint unsigned default 0 comment '是否删除 0-未删除 1-已删除',
    primary key(id)
)engine=innodb default charset=utf8mb4 comment='文章标签关联';