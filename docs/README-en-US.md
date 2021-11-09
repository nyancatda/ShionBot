[中文](https://github.com/nyancatda/MediaWiki-Bot) | English | [日本語](README-ja-JP.md)
# MediaWiki-Bot
A robot that queries MediaWiki via chat software
Support a variety of languages, cross-platform compatibility, support QQ\Telegram\Line.

Based on Gin snd [mirai-api-http](https://github.com/project-mirai/mirai-api-http)

*This project is still under development with many problems and low scalability. I will continue to improve it in the future, please use it with caution now. *  
*The code is terrible. Functions and variables are used casually. I'm sorry if it makes you uncomfortable. *

## How to ues

## 💮 Start  
1. Download the latest [Releases](https://github.com/nyancatda/MediaWiki-Bot/releases). 
1. Create [config.yml](#configyml-template) in the program's sibling directory. And follow the template to fill in the information. 
1. Configure [Chat Software](#Chat Software Configuration)
1. Run the program. 

## 🛠️ Chat Software Configuration
*Please configure at least one, otherwise the robot will not work*
### mirai-api-http(QQ)
1. Enable the http and webhook
1. Enable the enableVerify and enter your VerifyKey
1. Set the webhook address as http://<IP/URL>:<Port>/<Robot Key>
  Example:
  ```
  http://127.0.0.1:8000/32eeAme5lwEG0KL
  ```

setting.yml   
*The template is for reference only*
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
1. Set the Telegram WebHook report address as the robot receving address (https://<IP/URL>:<Port>/<Key>)，具体请查看[官方文档](https://core.telegram.org/bots/api#setwebhook)
  WebHook address example:
  ```
  https://127.0.0.1:8000/32eeAme5lwEG0KL
  ```
  *Warning, Telegram WebHook report address needs' HTTPS ', which may need to do reverse proxy for the address robot receives report*
1. If your server is located in mainland China, you also need to set up the reverse proxy service of Telegram Bot API. How to set up, please read [Set up the TelegramBotAPI reverse proxy server](docs/Telegram/ReverseProxyAPI.md)
### Line
1. 设置Line Bot WebHook上报地址(https://<机器人IP/URL地址>:<指定的机器人运行端口>/<指定的机器人密钥>)，可以在[Developers的控制台](https://developers.line.biz/console/)里设置，也可以[使用API设置](https://developers.line.biz/en/reference/messaging-api/#set-webhook-endpoint-url)  
WebHook地址例子:
```
https://127.0.0.1:8000/32eeAme5lwEG0KL
```
*Warning, Line WebHook report address needs' HTTPS ', which may need to do reverse proxy for the address robot receives report*

2. If your server is located in mainland China, you also need to set up the reverse proxy service of Line Bot API. How to set up, please read [Set up the LineBotAPI reverse proxy server](docs/Line/ReverseProxyAPI.md)

## config.yml template
```
Run:
  #Specify Webhook receiving port
  WebHookPort: 8000
  #指定机器人的WebHook密钥(只能使用字母与数字)
  WebHookKey: 32eeAme5lwEG0KL
  #Language
  #Chinese:zh-CN,English:en-US,Japanese:ja-JP
  Language: zh-CN
SNS:
  QQ:
    #是否启用QQ机器人部分
    Switch: true
    #HttpAPI address
    APILink: http://127.0.0.1:8888
    #The robot QQ number
    BotQQNumber: 1000000000
    #HttpAPI‘s VerifyKey
    VerifyKey: 5eadce46qw58
  Telegram:
    #是否启用Telegram机器人部分
    Switch: true
    #机器人toekn
    Token: 688975899:DDFqpsdMwunUvwAsxzDTzl8z_UkYzStrewM
    #TelegramAPI地址
    BotAPILink: https://api.telegram.org/
#Wiki urls. Multiple URLs can be added. The first one is the default Wiki
Wiki:
  - 
    #Wiki name, the prefix of the command，example: mw:home
    WikiName: mw
    #Wiki's URLs
    WikiLink: https://minewiki.net
  - 
    WikiName: mg
    WikiLink: https://zh.moegirl.org.cn
```

## 🔣 Command
0. 帮助
```
/help
```

1. Inquire the Wiki
```
Wiki name:What to search
```
Example:
```
mw:home
```

```
[[What to search]]
```
Example:
```
[[home]]
```

2. Change language  
*The modification only takes effect for a single user. Please modify the default language in the configuration file.*
```
/language language
```
Example
```
/language zh-CN
```

## 🌐 Multi-language
Adaptation progress: 
- [x] zh-CN(中文/简体)
- [ ] zh-HK(中文/香港)
- [x] en-US(English)
- [x] ja-JP(日本語)
- [ ] ru_RU(русский язык)

If you want to add more languages to this project, please fork the repository and create a new language file in the `language` directory. After the translation is completed, you can pull the request to the main repository. 

## 🎐 Thanks  
Thanks to all contributors  
*Names not listed in order*
1. [SuperYYT](https://github.com/SuperYYT)  
  English translation provider
2. [java23333](https://github.com/java23333)  
  Japanese translation provider
