中文 | [English](docs/README-en-US.md) | [日本語](docs/README-ja-JP.md)
# MediaWiki-Bot
通过聊天软件对MediaWiki进行信息查询的机器人  
可以对使用聊天软件对MediaWiki搭建的站点进行信息查询，支持多种语言，跨平台兼容，支持QQ，Telegram

基于Gin和[mirai-api-http](https://github.com/project-mirai/mirai-api-http)制作

*项目目前处于开发阶段，存在很多问题，可扩展性也很差，只能说堪堪能用，以后慢慢完善吧*  
*代码写的很烂，函数基本上想到啥封装啥，变量也是想到什么写什么，高血压请勿阅读*

## 如何使用

## 启动  
1. 从[Releases](https://github.com/nyancatda/MediaWiki-Bot/releases)下载最新构建
1. 在程序同级目录创建[config.yml](#configyml%E6%96%87%E4%BB%B6%E6%A8%A1%E6%9D%BF)，并按照模板填写信息
1. 配置[聊天软件](#聊天软件配置)
1. 运行程序

## 聊天软件配置
*请至少配置一个聊天软件，否则机器人将无法工作*
### mirai-api-http(QQ)
1. 启用http和webhook
1. 启用enableVerify，并设置VerifyKey
1. 将webhook地址设置为http://<机器人IP/URL地址>:<指定的机器人运行端口>/<指定的机器人密钥>
  例子:
  ```
  http://127.0.0.1:8000/32eeAme5lwEG0KL
  ```

setting.yml模板  
*仅供参考*
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
### Telegram
1. 设置Telegram WebHook上报地址为机器人接收地址(https://<机器人IP/URL地址>:<指定的机器人运行端口>/<指定的机器人密钥>)，具体请查看[官方文档](https://core.telegram.org/bots/api#setwebhook)  
  WebHook地址例子:
  ```
  https://127.0.0.1:8000/32eeAme5lwEG0KL
  ```
  *注意，Telegram的WebHook上报地址需要`https`，这可能需要需要对机器人接收上报的地址做反向代理*
1. 如果你的服务器位于中国大陆，你还需要搭建Telegram Bot API的反向代理服务，关于如何搭建，请查看[TelegramBotAPI反向代理服务器搭建](docs/Telegram/ReverseProxyAPI.md)

## config.yml文件模板
```
Run:
  #指定机器人的WebHook接收的端口
  WebHookPort: 8000
  #指定机器人的WebHook密钥(只能使用字母与数字)
  WebHookKey: 32eeAme5lwEG0KL
  #指定机器人的语言
  #中文:zh-CN,英语:en-US,日语ja-JP
  Language: zh-CN
SNS:
  QQ:
    #HttpAPI地址
    APILink: http://127.0.0.1:8888
    #机器人QQ号
    BotQQNumber: 1000000000
    #HttpAPI的VerifyKey
    VerifyKey: 5eadce46qw58
  Telegram:
    #机器人toekn
    Token: 688975899:DDFqpsdMwunUvwAsxzDTzl8z_UkYzStrewM
    #TelegramAPI地址
    BotAPILink: https://api.telegram.org/
#Wiki链接，支持多个，第一个为主Wiki
Wiki:
  - 
    #Wiki名字，即使命令前缀，例如mw:首页
    WikiName: mw
    #Wiki的链接
    WikiLink: https://minewiki.net
  - 
    WikiName: mg
    WikiLink: https://zh.moegirl.org.cn
```

## 命令
0. 帮助
```
/help
```

1. 查询Wiki
```
Wiki名字:需要查询的内容
```
例子:
```
mw:首页
```

```
[[需要查询的内容]]
```
例子:
```
[[首页]]
```

2. 修改语言  
*修改只对单个用户生效，默认语言修改请通过配置文件*
```
/language 语言
```
例子
```
/language zh-CN
```

## 多语言适配
多语言适配进度:  
- [x] zh-CN(中文/简体)
- [ ] zh-HK(中文/香港)
- [x] en-US(English)
- [x] ja-JP(日本語)
- [ ] ru_RU(русский язык)

如果您希望为本项目增加更多语言，请fork仓库后在`language`目录下建立目标语言文件，完成翻译后可以请求提交至主仓库

## 鸣谢  
感谢大佬们对这个项目的支持  
*排名不分先后*
1. [SuperYYT](https://github.com/SuperYYT)  
  英语翻提供译者
2. [java23333](https://github.com/java23333)  
  日语翻译提供者