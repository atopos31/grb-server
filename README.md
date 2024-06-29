# 基于Gin + Gorm + Mysql + Redis + Meilisearch的个人博客后端
## 依赖
| Environment | Version     |
| ----------- | ----------- |
| Golang      | 1.22.2      |
| Mysql       | 8.0.33      |
| Redis       | 7.0.12      |
## 配置文件
新建mydql数据库``grb``.

参考/config/pro.yaml编写配置文件,默认配置为开发环境.

使用gorm自动初始化数据库表结构.
## 运行
```shell
make dev 
make pro
```
## 接口文档

