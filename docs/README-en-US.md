[中文](https://github.com/nyancatda/MediaWiki-Bot)|English|[日本語](docs/README-ja-JP.md)
# MediaWiki-Bot
QQ query bot on MedisWiki

Based on Gin snd [mirai-api-http](https://github.com/project-mirai/mirai-api-http)

*This project is still under development with many problems and low scalability. I will continue to improve it in the future, please use it with caution now. *  
*The code is terrible. Functions and variables are used casually. I'm sorry if it makes you uncomfortable. *

## How to ues

## Start  
1. Download the latest [Releases](https://github.com/nyancatda/MediaWiki-Bot/releases). 
1. Create [config.yml](https://github.com/nyancatda/MediaWiki-Bot#configyml%E6%96%87%E4%BB%B6%E6%A8%A1%E6%9D%BF) in the program's sibling directory. And follow the template to fill in the information. 
1. [Configure the mirai-api-http](https://github.com/nyancatda/MediaWiki-Bot/blob/main/docs/README-en-US.md#%E9%85%8D%E7%BD%AEmirai-api-http)
1. Run the program. 

## Configure the mirai-api-http
1. Enable the http and webhook
1. Enable the enableVerify and enter your VerifyKey
1. Fill the webhook address as: http://127.0.0.1:+the port which your robot is running on

setting.yml *The template is for reference only*
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

## config.yml template
```
Run:
  #Specify WebHook port
  WebHookPort: 8000
  #Language
  #Chinese:zh-CN,English:en-US,Japanese:ja-JP
  Language: zh-CN
QQBot:
  #HttpAPI address
  APILink: http://127.0.0.1:8888
  #The robot QQ number
  BotQQNumber: 1000000000
  #HttpAPI‘s VerifyKey
  VerifyKey: 5eadce46qw58
#Wiki urls. Multiple URLs can be added. The first one is the default Wiki
Wiki:
  - 
    #Wiki name, the prefix of the command，example: mw:home
    WikiName: mw
    #Wiki's URLs
    WikiLink: https://minewiki.net
  - 
    WikiName: me
    WikiLink: https://zh.moegirl.org.cn
```

## Command
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
