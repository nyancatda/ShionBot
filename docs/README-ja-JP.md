[中文](https://github.com/nyancatda/MediaWiki-Bot) | [English](README-en-US.md) | 日本語
# MediaWiki-Bot
通过聊天软件对MediaWiki进行信息查询的机器人  
可以对使用聊天软件对MediaWiki搭建的站点进行信息查询，支持多种语言，跨平台兼容，支持QQ，Telegram，Line

Ginと[mirai-api-http](https://github.com/project-mirai/mirai-api-http)に基づいて作られた

*このプロジェクトは現在開発中なので、様々な問題があって、スケーラビリティも悪い。まあまあしか言えない。これからどんどん改善しよう*  
*書かれたコードが…あまり上手とは言えない。関数もほとんど思い付いたらパッケージングしておく。変数も思い付いたら書いておく。怒ったらごめんww*

## 使い方

##   スタートアップ
1. [Releases](https://github.com/nyancatda/MediaWiki-Bot/releases)から最新バージョンの構築をダウンロードする
1. プログラムの同じディレクトリで[config.yml](#configyml%E3%83%95%E3%82%A1%E3%82%A4%E3%83%AB%E3%83%86%E3%83%B3%E3%83%97%E3%83%AC%E3%83%BC%E3%83%88)を作成して、それからテンプレートにしたがってメッセージを入力する
1. 配置[聊天软件](#聊天软件配置)
1. プログラムを実行する

## 聊天软件配置
*请至少配置一个聊天软件，否则机器人将无法工作*
### mirai-api-http(QQ)
1. httpとwebhookを起動する
1. enableVerifyを起動してから、VerifyKeyを設定する
1. 将webhook地址设置为http://<机器人IP/URL地址>:<指定的机器人运行端口>/<指定的机器人密钥>
  例子:
  ```
  http://127.0.0.1:8000/32eeAme5lwEG0KL
  ```

setting.ymlテンプレート  
*参考だけ*
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
### Line
1. 设置Line Bot WebHook上报地址(https://<机器人IP/URL地址>:<指定的机器人运行端口>/<指定的机器人密钥>)，可以在[Developers的控制台](https://developers.line.biz/console/)里设置，也可以[使用API设置](https://developers.line.biz/en/reference/messaging-api/#set-webhook-endpoint-url)  
WebHook地址例子:
```
https://127.0.0.1:8000/32eeAme5lwEG0KL
```
*注意，Line的WebHook上报地址需要`https`，这可能需要需要对机器人接收上报的地址做反向代理*

2. 如果你的服务器位于中国大陆，你还需要搭建Line Bot API的反向代理服务，关于如何搭建，请查看[LineBotAPI反向代理服务器搭建](docs/Line/ReverseProxyAPI.md)

## config.ymlファイルテンプレート
```
Run:
  #指定されたボットのWebHookが受信ボット
  WebHookPort: 8000
  #指定机器人的WebHook密钥(只能使用字母与数字)
  WebHookKey: 32eeAme5lwEG0KL
  #ボットの言語を選ぶ
  #中国語:zh-CN,英語:en-US,日本語ja-JP
  Language: zh-CN
SNS:
  QQ:
    #是否启用QQ机器人部分
    Switch: true
    #HttpAPIアドレス
    APILink: http://127.0.0.1:8888
    #ボットのQQアプリ番号
    BotQQNumber: 1000000000
    #HttpAPIのVerifyKey
    VerifyKey: 5eadce46qw58
  Telegram:
    #是否启用Telegram机器人部分
    Switch: true
    #机器人toekn
    Token: 688975899:DDFqpsdMwunUvwAsxzDTzl8z_UkYzStrewM
    #TelegramAPI地址
    BotAPILink: https://api.telegram.org/
#Wikiアドレス 複数、一番目が優先のWiki
Wiki:
  - 
    #Wiki名前，コマンドのプレフィックス，例:mw:◯◯
    WikiName: mw
    #Wikiアドレス
    WikiLink: https://minewiki.net
  - 
    WikiName: me
    WikiLink: https://zh.moegirl.org.cn
```

## コマンド
0. 帮助
```
/help
```

1. 検索Wiki
```
Wiki名前:検索したいこと
```
例:
```
mw:◯◯
```

```
[[検索したいこと]]
```
例:
```
[[◯◯]]
```

2. 言語の変更
*変更は1人のユーザーに対してのみ有効。デフォルトの言語を変更するには、構成ファイルで変更してください*
```
/language 言語
```
例:
```
/language ja-JP
```

## 多语言アダプテーション
現在使える言語は:  
進度:  
- [x] zh-CN(中文/简体)
- [ ] zh-HK(中文/香港)
- [x] en-US(English)
- [x] ja-JP(日本語)
- [ ] ru_RU(русский язык)

このプロジェクトにさらに言語を追加したい場合は、repositoryをforkして、`language`ディレクトリでファイルを作成する。翻訳ができたらmaster repositoriesに提出したら助かります

##   最後
このプロジェクトの翻訳の提供者に感謝します
*以下のリストは優先順位つけていない*
1. [SuperYYT](https://github.com/SuperYYT)  
  英語翻訳提供者
2. [java23333](https://github.com/java23333)  
  日本語翻訳提供者
