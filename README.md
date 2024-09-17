# GRB-Server
基于Gin + Gorm + Mysql + Redis + Meilisearch的个人博客后端
## 依赖
| Environment                                | Version |
| ------------------------------------------ | ------- |
| Golang                                     | 1.22.4  |
| Mysql                                      | 8.0.37  |
| Redis                                      | 7.0.15  |
| [Meilisearch](https://www.meilisearch.com) | 1.9.0   |
## 远程依赖
| Environment | Name           | Description          |
| ----------- | -------------- | -------------------- |
| Oss         | 七牛云          | 用于博客内图片存储，文件保存    |
| Ai          | 腾讯混元/百度千帆  | 用于自动生成文章摘要，均为免费lite模型   |

## 工具
### swag用于生成接口文档
```shell
go install github.com/swaggo/swag/cmd/swag@latest
swag init
```
## 配置文件
自行新建开发环境配置文件/config/dev.yaml.
参考/config/pro.yaml编写配置文件,默认配置为开发环境.
## 数据库
新建mysql数据库``databasename``.
```sql
CREATE DATABASE databasename DEFAULT CHARACTER SET = 'utf8mb4';
```
在dev/pro.yaml中指定``databasename``数据库名，运行项目即可利用Gorm自动初始化数据库表结构.
## 运行
```shell
make dev # 开发环境  args: -config ./config/dev.yaml
make pro # 生产环境  args: -config ./config/pro.yaml
```
## docker部署
TODO
## 接口文档

