## gin-blog
此项目从煎鱼大佬[go-gin-example](https://github.com/eddycjy/go-gin-example)学习笔记实践

感谢煎鱼大佬的心力之作，让我慢慢一路前行跟随。


#### 数据库准备
```mysql
CREATE DATABASE gin_web CHARACTER SET utf8 COLLATE utf8_bin;
grant all on gin_web.* to gin@"%" identified by "C6jiaAb68fjeYPjR4u";
grant all on gin_web.* to gin@"localhost" identified by "C6jiaAb68fjeYPjR4u";
flush privileges;
```

```mysql
# 标签表
CREATE TABLE `blog_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT '标签名称',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章标签管理';
```

```mysql
# 文章表
CREATE TABLE `blog_article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `tag_id` int(10) unsigned DEFAULT '0' COMMENT '标签ID',
  `title` varchar(100) DEFAULT '' COMMENT '文章标题',
  `desc` varchar(255) DEFAULT '' COMMENT '简述',
  `content` text,
  `created_on` int(11) DEFAULT NULL,
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(255) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用1为启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章管理';
```

```mysql
# 认证表
CREATE TABLE `blog_auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT '' COMMENT '账号',
  `password` varchar(50) DEFAULT '' COMMENT '密码',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='账号管理';

INSERT INTO `gin_web`.`blog_auth` (`id`, `username`, `password`) VALUES (null, 'test', 'test123456');
```


#### 项目描述

- 项目启用go mod, golang版本 1.13.1
- 项目目录描述
  + conf
    `配置文件目录`
  + middleware
    `gin中间件目录` 
  + models
    `应用数据库模型`
  + pkg
    `第三方包`
  + routers
    `路由处理逻辑`
  + runtime
    `应用运行时数据`
    
- 项目编写
1. 从main.go入口开始自定义加载routers包下面的router.go
2. router.go 初始gin.Engine导入路由组apiv1到业务api接口

    - api/v1/tag.go     `tag的增删改查，字段合法性判断等`
    - api/v1/article.go `文章的增删改查，字段合法性判断等`
  
3. models导入业务模型

    - models/models.go `初始哈模型公共结构体和字段`
    - models/tag.go    `tag增删改查逻辑`
    - models/article.go `文章增删改查逻辑`


#### 项目功能列表

1. 基于gin REST API 实现tag、article的增删改查功能.
2. 使用JWT进行Gin REST API身份鉴权.
    
    - /auth 提交用户名和密码, 获取token.
    - 在请求URL上需要添加token字段

3. 增加一个简单的文件日志
    
    - 为routers/api/auth.go 使用自定义logger写日志. 

4. 使用endless优雅的重启服务

5. 使用Swagger为Gin REST Api生成文档

    - 文档地址: `http://127.0.0.1:8000/swagger/index.html`

6. 将Golang应用部署到Docker

    - 构建镜像: `docker build -t gin-blog-docker .`
    - 创建运行为容器: `docker run -d -p8000:8000 gin-blog-docker`

7. 定制 GORM Callbacks
    
    - 使用GORM Callback替换BeforeCreate、BeforeUpdate
    - 注册Callbacks将其注册进GORM的钩子中.
    - 实现项目软删除.
    
8. Cron定时任务

    - 使用Cron清理标记删除的数据.    

#### 参考链接

- [go-gin-example](https://github.com/eddycjy/go-gin-example)
- [gin-docs](https://gin-gonic.com/docs/)
- [gin-core](https://github.com/gin-gonic/gin)
- [gin-swagger](https://github.com/swaggo/gin-swagger)
- [swag](https://github.com/swaggo/swag)
- [swagger](https://swagger.io/specification/)
- [gorm](http://gorm.io/docs/write_plugins.html)
- [cron](https://github.com/robfig/cron)