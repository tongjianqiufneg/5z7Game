# 5z7Game
实现go的五子棋游戏

## 背景
- 如何使用go实现游戏项目开发
- 要求扩展性强
- 要求有高并发处理能力


## 概念说明
五子棋游戏的元素包含用户，棋盘，棋局
- 棋盘：空间容器
- 棋局：棋子的状态
- 用户：在棋盘上操控棋子对棋局进行变更

## 系统设计
游戏就是为了输赢，所以我们发现白方胜或者黑方胜，游戏就算结束
### 流程设计

### 功能模块设计

### 数据库设计

### 缓存设计

## 项目分层
```
├─.idea
├─config
├─http
│  ├─handler
│  ├─response
│  └─route
├─pkg
│  ├─app
│  ├─constant
│  ├─dto
│  │  ├─request
│  │  └─response
│  ├─enum
│  ├─model
│  │  ├─dao
│  │  ├─data
│  │  └─entity
│  ├─service
│  └─utils
└─task
└─main.go
```

- main.go 项目入口
- http 接口模块
    - handler 接口处理模块
    - response 响应
    - router 路由
- pkg 业务模块
    - app 核心服务(logger,database,redis等)
    - constant 常量
    - dto 数据传输对象
        - request 请求对象
        - response 响应对象
    - enum 枚举
    - model 数据模型
        - dao 数据操作层(操作数据库，增删改查业务实现)
        - data 数据结构体(超过3个参数即可用data定义结构体)
        - entity 模型实体(表)
    - service 服务模块(对外提供服务,仅handler或者task能调用,服务模块调用model的方法)
    - utils 工具
- task 定时任务

