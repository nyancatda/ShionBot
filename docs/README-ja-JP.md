[中文](https://github.com/nyancatda/MediaWiki-Bot) | [English](README-en-US.md) | 日本語
# MediaWiki-Bot
通过聊天软件对MediaWiki进行查询信息的机器人  
可以对使用聊天软件对MediaWiki搭建的站点进行信息查询，支持多种语言，跨平台兼容，支持QQ，Telegram

Ginと[mirai-api-http](https://github.com/project-mirai/mirai-api-http)に基づいて作られた

*このプロジェクトは現在開発中なので、様々な問題があって、スケーラビリティも悪い。まあまあしか言えない。これからどんどん改善しよう*  
*書かれたコードが…あまり上手とは言えない。関数もほとんど思い付いたらパッケージングしておく。変数も思い付いたら書いておく。怒ったらごめんww*

## 使い方

##   スタートアップ
1. [Releases](https://github.com/nyancatda/MediaWiki-Bot/releases)から最新バージョンの構築をダウンロードする
1. プログラムの同じディレクトリで[config.yml](https://github.com/nyancatda/MediaWiki-Bot#configyml%E6%96%87%E4%BB%B6%E6%A8%A1%E6%9D%BF)を作成して、それからテンプレートにしたがってメッセージを入力する
1. 配置[聊天软件]()
1. プログラムを実行する

## 聊天软件配置
*请至少配置一个聊天软件，否则机器人将无法工作*
### mirai-api-http(QQ)
1. httpとwebhookを起動する
1. enableVerifyを起動してから、VerifyKeyを設定する
1. webhookアドレスをhttp://127.0.0.1:+指定されたボットのポート に設定する

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
1. 设置Telegram WebHook上报地址为机器人接收地址，具体请查看[官方文档](https://core.telegram.org/bots/api#setwebhook)
  *注意，Telegram的WebHook上报地址需要`https`，这可能需要需要对机器人接收上报的地址做反向代理*
1. 如果你的服务器位于中国大陆，你还需要搭建Telegram Bot API的反向代理服务，关于如何搭建，请查看[TelegramBotAPI反向代理服务器搭建](docs/Telegram/ReverseProxyAPI.md)

## config.ymlファイルテンプレート
```
Run:
  #指定されたボットのWebHookが受信ボット
  WebHookPort: 8000
  #ボットの言語を選ぶ
  #中国語:zh-CN,英語:en-US,日本語ja-JP
  Language: zh-CN
QQBot:
  #HttpAPIアドレス
  APILink: http://127.0.0.1:8888
  #ボットのQQアプリ番号
  BotQQNumber: 1000000000
  #HttpAPIのVerifyKey
  VerifyKey: 5eadce46qw58
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
