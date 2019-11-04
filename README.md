# ginger

## 简介
Ginger is a scaffold for building gin framework application.

Ginger 是一个构建gin应用的脚手架。

#### 特性

- 可适应开发、测试、生成环境的配置；
- 可追踪调用链的日志记录；
- 开箱即用的jwk鉴权机制；
- 整合sql builder方式的dao层，给不喜欢orm和原生sql的程序猿一条生路；
- 整合mgo三方库的连接池及简易调用方法；
- 整合redigo三方库的连接池及简易调用方法、管道调用方法；
- 整合基于redis的简单消息队列
- 提供简单的utils工具包
- 使用ginger-cli客户端生成通用代码，轻松搭建gin应用


## 开始


#### 依赖


#### 安装


#### 目录结构
`

    root
    |_common 公用目录：编写公共处理函数，通用错误处理，常量设置等
    |_router 路由设置目录
    |_handlers 业务处理函数目录
    |_config 配置目录:配置项应可根据系统环境动态获取
        |_dev 存放开发环境配置文件
        |_prod 存放生产环境配置文件
        |_test 存放测试环境配置文件
    |_util 工具包目录：编写工具类的函数或方法
    |_asset 可直接访问的静态资源目录js、css、image
    |_views 页面模板目录
    |_dao 数据访问层目录
        |_mysql 该目录实现mysql连接池初始化，以及自动生成的基于sql builder基本数据库表curd
        |_redis 该目录实现redis连接池初始化，以及通用的redis访问方法R()；
        |_mongodb 该目录实现mongodb连接池初始化，以及通用的mongodb访问方法M()；
        |_...
    |_cache 缓存层目录：Redis实现，保存转态信息，热点数据，数据库信息缓存等
    |_repository 数据仓库目录，编写mongo存储相关的业务逻辑
    |_model 业务数据模型目录，编写mysql相关的业务存储逻辑
    |_middleware 中间件目录
    |_validator 自定义验证器目录
`

#### 工具


#### 简单示例

