[中文](https://github.com/nyancatda/MediaWiki-Bot) | [English](README-en-US.md) | 日本語
# MediaWiki-Bot
MediaWikiの多言語QQ検索ボット  
MediaWikiを使うWebページの検索が可能

Ginと[mirai-api-http](https://github.com/project-mirai/mirai-api-http)に基づいて作られた

*このプロジェクトは現在開発中なので、様々な問題があって、スケーラビリティも悪い。まあまあしか言えない。これからどんどん改善しよう*  
*書かれたコードが…あまり上手とは言えない。関数もほとんど思い付いたらパッケージングしておく。変数も思い付いたら書いておく。怒ったらごめんww*

## 使い方

##   スタートアップ
1. [Releases](https://github.com/nyancatda/MediaWiki-Bot/releases)から最新バージョンの構築をダウンロードする
1. プログラムの同じディレクトリで[config.yml](https://github.com/nyancatda/MediaWiki-Bot#configyml%E6%96%87%E4%BB%B6%E6%A8%A1%E6%9D%BF)を作成して、それからテンプレートにしたがってメッセージを入力する
1. [mirai-api-http](https://github.com/nyancatda/MediaWiki-Bot#%E9%85%8D%E7%BD%AEmirai-api-http)を設定する
1. プログラムを実行する

## mirai-api-httpの設定
1. httpとwebhookを起動する
1. enableVerifyを起動してから、VerifyKeyを設定する
1. webhookアドレスをhttp://127.0.0.1:+指定されたボットのポート に設定する

setting.ymlテンプレート*参考だけ*
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
