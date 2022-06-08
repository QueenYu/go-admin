## go admin

## 技术选型

#### go相关

1. web框架 gin
2. 数据库orm gorm
3. 日志记录 zap
4. 配置文件读取 viper
5. API文档 swagger

#### 其他技术栈

1. 数据库 mysql
2. 缓存 redis

## 2022-06-02 一期规划

1. 完成用户、角色、菜单基本功能的开发
2. 用户登录、并且动态菜单及权限校验
3. 完成api管理，不同的角色可访问不同的api

详细计划:

- 初始化gorm、viper、zap,并且读取配置文件
- 实现用户的增、删、改、查、分页,注意事务相关
- 实现用户的登录、登录

问题:

1. gorm使用zap打印日志
2. 如果未连接数据库，或者未连接redis，如何中断启动
- gorm会进行连通性测试，如果未连接，会中断启动,但是在gin中是否会呢。会，因为gin还未启动,不会阻止pannic
3. viper区分环境变量

