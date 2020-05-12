# 数据监控后端
基于 go gin 搭建的一个数据监控平台后端代码仓库，毕设作品

目前仅完成了用户注册登录验证部分功能，其他功能持续完善中


---

### require
- postgresql


---
### 项目架构重构
```
gin-blog/
├── config //配置文件
├── middleware //应用中间件
├── models // 数据模型
├── utils // 应用包
├── router // 路由
└── runtime // 应用运行时数据
└── docs // swagger文件
└── static // 静态文件

```
