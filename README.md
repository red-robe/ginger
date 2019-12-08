# ginger

## 简介
Ginger is a scaffold for building gin framework application.

Ginger 是一个构建gin应用的脚手架。

#### 特性

- 可适应开发、测试、生成环境的配置；
- 可追踪请求调用链；
- 统一的输出格式
- 开箱即用的jwt鉴权；
- 整合sql builder方式的dao层，给不喜欢orm和原生sql的程序猿一条生路；
- 整合mgo三方库的连接池及简易调用方法；
- 整合redigo三方库的连接池及简易调用方法、管道调用方法；
- 提供通用的utils工具包
- 使用ginger-cli客户端生成通用代码，轻松搭建gin应用


#### 目录结构


    app_root
    |_asset 可直接访问的静态资源目录js、css、image
    |_boot 应用启动时的初始化逻辑
    |_cache 缓存层目录：Redis实现，保存转态信息，热点数据，数据库信息缓存等
    |_common 公用目录：编写公共处理函数，退出错误处理等
    |_config 配置目录:存放yaml配置文件，yaml解析代码及常量级的配置
        |_dev 存放开发环境配置文件
        |_prod 存放生产环境配置文件
        |_test 存放测试环境配置文件
    |_cron 定时任务代码目录
    |_dao 数据访问层目录
            |_mysql 该目录实现mysql连接池初始化，以及自动生成的基于sql builder基本数据库表curd
            |_redis 该目录实现redis连接池初始化，以及通用的redis访问方法R()；
            |_mongodb 该目录实现mongodb连接池初始化，以及通用的mongodb访问方法M()；
            |_...
    |_handler 业务处理函数目录
    |_logger 实现基于zap的日志记录器
    |_middleware 中间件目录
    |_model 业务数据模型目录，编写mysql相关的业务存储逻辑
    |_mq 实现消息中间件系统的client，这里整合了redis pubsub和nats，可根据业务随意使用，在配置文件mq.yaml配置必要项即可
    |_repository 数据仓库目录，编写mongo存储相关的业务逻辑
    |_router 路由设置目录
    |_subscriber 整合了redis pubsub和nats的消息订阅器，可随系统启动一并运行，也可抽离出来独立运行
    |_util 工具包目录：编写工具包方法，提供基于日志记录的错误处理包，jwt编解码包，业务日志记录包
    |_views 页面模板目录
    |_validator 自定义验证器目录


#### 工具
##### ginger-cli
脚手架代码生成工具
https://github.com/gofuncchan/ginger-cli

- init 初始化项目目录
- handler 生成基本的handler代码
- dao 整合gingger-forge工具，生成go struct和curd代码
- config 根据yaml文件生成go的解析代码
- repo 生成基于mongodb的repository相关代码
- model 生成基于mysql的业务数据模型相关代码
- cache 生成基于redis的缓存操作代码

##### gingger-forge
基于didi/gendry的dao代码生成工具
https://github.com/gofuncchan/ginger-forge

- 可根据数据库schema映射生成go struct和curd代码

#### 用例
##### 方式一：download
- 下载项目:

    `git clone https://github.com/gofuncchan/ginger.git {your project name}`

- 由于使用go module,请自定义go.mod文件的replace本地代码目录，或运行以下代码：

    ` cd {your project name} && go mod edit -replace github.com/gofuncchan/ginger={your project directory}`

- 在/config目录下重置或新增配置项，并解析到全局变量

- 运行程序：

    `go run main.go`

##### 方式二：使用ginger-cli工具

- 安装工具:

    `go get -u github.com/gofuncchan/ginger-cli`
    
- 初始化项目:

    请确保你的$GOPATH/bin已设置到全局变量$PATH,切换到你的代码目录，执行init命令，会在你的代码目录创建项目脚手架
    
    `ginger-cli init {your project name} [-g]`
    
- 在/config目录下重置或新增配置项，并解析到全局变量

- 运行程序：

    `go run main.go`
    
#### Tips:
   
   脚手架项目代码内默认使用github.com/gofuncchan/ginger作为内部根包名，如需使用可fork本项目并自行修改代码，或下载本项目自行全局替换成你的自定义包名，并修改go.mod文件



