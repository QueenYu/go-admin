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

#### 接口返回设计

API接口的响应为JSON格式，并且会响应不同的HTTP status code

1. HttpStatus 返回状态码

- 200 成功
- 400 客户端参数错误
- 401 未认证
- 403 未授权
- 500 系统未知异常

2. API返回统一格式

返回样例:
```
{
    "code": 7,
    "msg": {"后端返回消息内容"},
    "data": {}
}
```
code代表的意思:
- 0 成功
- 7 业务异常

3. 各种情况的列举

- 数据格式校验未通过 http status code为400 msg会携带未校验通过的信息 code不太重要，就无所谓了
- 业务验证未通过 http status为200 code=7 msg会携带业务验证错误信息，比如手机号已存在
- 运行时异常 http status为500 msg会携带一些错误信息 比如新建失败,code不太重要，就无所谓了

#### 总计划待做

1. 优雅地终止
2. 分布式id
3. 统一返回结果是否合理，没有进行将实体转换, json统一处理(时间、类型)


## 2022-06-02 一期规划

1. 完成用户、角色、菜单基本功能的开发
2. 用户登录、并且动态菜单及权限校验
3. 完成api管理，不同的角色可访问不同的api

详细计划:

- 实现用户的增、删、改、查、分页,注意事务相关,暂时不做角色相关的东西(完成)
- swagger文档(完成)
- 菜单的增(增加根菜单和子菜单)、删、改、查(单个查询、oneToMany)
- 角色
- 角色用户
- 实现用户的登录、登录
- 字典管理
- 操作历史
- api管理

问题:

1. gin 全局异常处理



