[中文](https://github.com/nyancatda/ShionBot) | [English](README-en-US.md) | 日本語
# ShionBot
チャットソフトでMediaWikiを使って検索するボット  
MediaWikiで作られたページに検索できる 多言語可能、プラットフォームを跨る、QQ、テルグラム、LINEで使う可能

Ginと[mirai-api-http](https://github.com/project-mirai/mirai-api-http)に基づいて作られた

*このプロジェクトは現在開発中なので、様々な問題があって、スケーラビリティも悪い。まあまあしか言えない。これからどんどん改善しよう*  
*書かれたコードが…あまり上手とは言えない。関数もほとんど思い付いたらパッケージングしておく。変数も思い付いたら書いておく。怒ったらごめんww*

## 使い方

## 💮 スタートアップ
1. [Releases](https://github.com/nyancatda/ShionBot/releases)から最新バージョンの構築をダウンロードする
1. プログラムの同じディレクトリで[config.yml](#configyml%E3%83%95%E3%82%A1%E3%82%A4%E3%83%AB%E3%83%86%E3%83%B3%E3%83%97%E3%83%AC%E3%83%BC%E3%83%88)を作成して、それからテンプレートにしたがってメッセージを入力する
1. [チャットソフト]のコンフィグ(#チャットソフトのコンフィグ)
1. プログラムを実行する

## 🛠️ チャットソフトのコンフィグ
*最低一つのチャットソフトを設定してください そうしないと、ボットは実行できない*
### mirai-api-http(QQ)
1. httpとwebhookを起動する
1. enableVerifyを起動してから、VerifyKeyを設定する
1. webhookアドレスは http://<ボットIP/URL>:<指定されたボットの実行ポート1>/<指定されたボットのパスワード> に設定して
  例:
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
### テルグラム
1. テルグラム WebHookのアップロードするアドレスはボットの受信アドレスに設定して(https://<ボットIP/URL>:<指定されたボットの実行ポート>/<指定されたボットのパスワード>)，詳しいのは[公式ファイル](https://core.telegram.org/bots/api#setwebhook)を見てください
  WebHookアドレスの例:
  ```
  https://127.0.0.1:8000/32eeAme5lwEG0KL
  ```
  *注意:LineのWebHookアップロードアドレスは`https`が必要なので、ボットのアップロードしたアドレスを逆方向プロキシする可能性もある*
1. もしあなたのサーバは中国にいたら、Telegram Bot APIの逆方向プロキシを構築するも必要 どうやって構築できることに関して、これを見てください[TelegramBotAPIプロキシするサーバの構築](docs/Telegram/ReverseProxyAPI.md)
### Line
1. Line Bot WebHookのアップロードするアドレスを設定して(https://<ボットIP/URL>:<設定されたボットの実行ポート>/<設定されたボットのパスワード>)、[Developersのコンソール](https://developers.line.biz/console/)で設定できる それに、[APIを使う設定](https://developers.line.biz/en/reference/messaging-api/#set-webhook-endpoint-url)  も使ってもいい
WebHookアドレス例:
```
https://127.0.0.1:8000/32eeAme5lwEG0KL
```
*注意:LineのWebHookアップロードアドレスは`https`が必要なので、ボットのアップロードしたアドレスを逆方向プロキシする可能性もある*

2. もしあなたのサーバは中国にいたら、Telegram Bot APIの逆方向プロキシを構築するも必要 どうやって構築できることに関して、これを見てください[LineBotAPIのプロキシするサーバの構築](docs/Line/ReverseProxyAPI.md)

## config.ymlファイルテンプレート
```
Run:
  #指定されたボットのWebHookが受信ボット
  WebHookPort: 8000
  #指定されたWebHookのパスワード(ローマ字と数字だけ使える)
  WebHookKey: 32eeAme5lwEG0KL
  #ボットの言語を選ぶ
  #中国語:zh-CN,英語:en-US,日本語ja-JP
  Language: zh-CN
SNS:
  QQ:
    #QQボット部分をONにするか
    Switch: true
    #HttpAPIアドレス
    APILink: http://127.0.0.1:8888
    #ボットのQQアプリ番号
    BotQQNumber: 1000000000
    #HttpAPIのVerifyKey
    VerifyKey: 5eadce46qw58
  テルグラム:
    #テルグラムボット部分をONにするか
    Switch: true
    #ボットtoekn
    Token: 688975899:DDFqpsdMwunUvwAsxzDTzl8z_UkYzStrewM
    #TelegramAPIアドレス
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

## 🔣 コマンド
0. ヘルプ
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

## 🌐 多语言アダプテーション
現在使える言語は:  
進度:  
- [x] zh-CN(中文/简体)
- [ ] zh-HK(中文/香港)
- [x] en-US(English)
- [x] ja-JP(日本語)
- [ ] ru_RU(русский язык)

このプロジェクトにさらに言語を追加したい場合は、repositoryをforkして、`language`ディレクトリでファイルを作成する。翻訳ができたらmaster repositoriesに提出したら助かります

## 🎐 最後
このプロジェクトの翻訳の提供者に感謝します
*以下のリストは優先順位つけていない*
1. [SuperYYT](https://github.com/SuperYYT)  
  英語翻訳提供者
2. [java23333](https://github.com/java23333)  
  日本語翻訳提供者
