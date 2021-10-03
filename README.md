# MediaWiki-Bot
MediaWiki的QQ查询机器人

基于Gin和[mirai-api-http](https://github.com/project-mirai/mirai-api-http)制作

*项目目前处于开发阶段，存在很多问题，可扩展性也很差，只能说堪堪能用，以后慢慢完善吧*

## 如何使用

## 启动  
1. 从[Releases](https://github.com/nyancatda/MediaWiki-Bot/releases)下载最新构建
1. 在程序同级目录创建config.yml，并按照模板填写信息
1. 配置mirai-api-http
1. 运行程序

## 配置mirai-api-http
1. 启用http和webhook
1. 启用enableVerify，并设置VerifyKey
1. 将webhook地址设置为http://127.0.0.1:8000/

setting.yml模板*仅供参考*
```
adapters:
  - http
  - webhook
enableVerify: true
verifyKey: 5eadce46qw58
debug: false
singleMode: false
cacheSize: 4096
adapterSettings:
  http:
    host: 0.0.0.0
    port: 8888
    cors: [*]
  webhook:
    destinations: 
    - 'http://127.0.0.1:8000/'
```

## config.yml文件模板
```
QQBot:
  #HttpAPI地址
  APILink: http://127.0.0.1:8888
  #机器人QQ号
  BotQQNumber: 1000000000
  #HttpAPI的VerifyKey
  VerifyKey: 5eadce46qw58
Wiki:
  #需要对接的wiki的地址
  WikiLink: https://minewiki.net
```