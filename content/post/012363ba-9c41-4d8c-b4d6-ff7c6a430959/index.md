---
title: RustとPythonでPub/Subことはじめ1
description: 
slug: 012363ba-9c41-4d8c-b4d6-ff7c6a430959
date: 2024-03-03T12:55:22+09:00
image: 
categories:
    - Rust
    - Python
    - GCP
    - Pub/Sub
    - 非同期
    - My tutorial
weight: 1      
---

RustとPythonでPub/Subを行います。
Rust側サーバーがパブリッシャーで、Python側がサブスクライバーになります

[公式のtutorial](https://cloud.google.com/pubsub/docs/building-pubsub-messaging-system?hl=ja)を参考にします。

<br>

# 事前準備

<br>

## gcloudコマンドをインストールする

https://cloud.google.com/sdk/docs/install?hl=ja
を元に行います。

mac(intel)なので次のコマンドで取ってきて解凍します。

```shell
# 任意のdir直下で
wget "https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-cli-466.0.0-darwin-x86_64.tar.gz?hl=ja"

tar -zxvf "google-cloud-cli-466.0.0-darwin-x86_64.tar.gz?hl=ja"
```

[tar.gzの解凍](https://qiita.com/supersaiakujin/items/c6b54e9add21d375161f#targz)

<br>

## gcloud の初期化とか

公式に従います。

```shell
./google-cloud-sdk/install.sh
```

筆者はfishを使っていますが、Pathの追加も問題なく行なってくれそうですね。

```shell

Enter a path to an rc file to update, or leave blank to use
[/Users/yuunag1/.config/fish/config.fish]:
Backing up [/Users/yuunag1/.config/fish/config.fish] to [/Users/yuunag1/.config/fish/config.fish.backup].
[/Users/yuunag1/.config/fish/config.fish] has been updated.

==> Start a new shell for the changes to take effect.

```

backupも撮ってくれるなんてなんて親切なんでしょう笑

```shell
❰yuunag1❙~/development❱✔≻ gcloud
fish: Unknown command: gcloud
❰yuunag1❙~/development❱✘≻ exec $SHELL
Darwin uehararyoujis-MacBook-Pro.local 23.3.0 x86_64
 9:07  up 20 days,  6:48, 9 users, load averages: 6.57 8.39 6.36
❰yuunag1❙~/development❱✔≻ gcloud -h
Usage: gcloud [optional flags] <group | command>
  group may be           access-approval | access-context-manager |
                         active-directory | ai | ai-platform | alloydb |
```

一瞬Path通ってなさそうでしたが、再読み込みしたら問題なく入ってきてくれました。

<br>

## 初期化

```shell
gcloud init
```

するとなんかいくつか聞かれるので、それにしたがって進めます。
特に問題なく終わりました。

<br>

## Pub/Sub API を有効にします。

```shell
gcloud services enable pubsub.googleapis.com
```

<br>

## Google アカウントのローカル認証情報を作成します。

```shell
gcloud auth application-default login
```

```shell
Credentials saved to file: [/Users/<user>/.config/gcloud/application_default_credentials.json]

These credentials will be used by any library that requests Application Default Credentials (ADC).

Quota project "gcp-may-sandbox" was added to ADC which can be used by Google client libraries for billing and quota. Note that some services may still bill the project owning the resource.
```

home直下の.configの中にローカル認証情報が作成されるようです

<br>

## Google アカウントにロールを付与します。

```shell
gcloud projects add-iam-policy-binding $PROJECT_ID --member="user:$EMAIL_ADDRESS" --role="roles/pubsub.publisher"

gcloud projects add-iam-policy-binding $PROJECT_ID --member="user:$EMAIL_ADDRESS" --role="roles/pubsub.subscriber"
```

$PROJECT_ID と、$EMAIL_ADDRESSはお使いのものに置き換えてください


<br>

# Pub/Sub プロジェクトを設定する

<br>

## Pub/Sub トピックの作成

```shell
gcloud pubsub topics create <トピック名>
```

<br>

## Pub/Sub サブスクリプションの作成

今回は一つだけsubscriptionが欲しいので一つだけ作成します。
種類は、ストリーミング Pull サブスクリプションです。

```shell
gcloud pubsub subscriptions create sub_python --topic=$TOPIC
```

<br>

# Rustとpythonで非同期なシステムを構築する。

ここでもう少し背景に突っ込んどきましょう。
さて、チュートリアルでは1対多のシステムを構築しています。
しかし、今回のシステムではそうではないです。

ではなぜPub/Subが必要なのでしょうか。

それは、Python側で行われる重い操作が完了する前にRust側からクライエントにResponseを返したいからです。

こういった目的のためには結構刺さるんじゃないかなと思っています。


参考のためにこれも置いときます

[ちょっとこわいGoogle Cloud Pub/Sub によるFaaS アーキテクチャ](https://recruit.gmo.jp/engineer/jisedai/blog/scary_pubsub/)

<br>

## Rust側サーバーをパブリッシャーにする

ちなみに公式のライブラリはなさそうなんですよねえ。

しかし、めっちゃ古いながらも方法はあるっぽい感じでした。

https://github.com/x1-/rust-pubsub-example


依存してるライブラリはそれぞれメンテはされてるっぽい
1. [google-apis-rs](https://github.com/Byron/google-apis-rs)
2. [yup-oauth2](https://github.com/dermesser/yup-oauth2/tree/master)


うーん公式じゃないしメンテする覚悟しても良さそうな気がしたので、forkしてきました。

1. [フォークしたgoogle-apis-rs](https://github.com/u-na-gi/google-apis-rs)
2. [フォークしたyup-oauth2](https://github.com/u-na-gi/yup-oauth2)

では書いていきましょう

### ライブラリをインストールする。

https://zenn.dev/booink/scraps/6d2a72a89448c9

cargo.tomlでGitHubリポジトリを参照するにはこうしたらいいらしい。

gitにcloneするときのurl指定してversionに。。。。ってtag切られてないじゃん。。

commit 指定でできるのかなあ。

いやアスタリスクで良さそう。

<br>

### とりあえずサンプルを動かしたい

https://github.com/x1-/rust-pubsub-example

をlocalに持ってきた頃まあそうよねえって感じですが依存関係が荒れすぎて色々動かなくなりました。

とりあえずpackageを最新にしましょう。

予想通りですが、色々ぶっ壊れてますね。

https://crates.io/
でdependancyを検索してとりあえず最新のやつをとってきます。

```

[dependencies]
base64 = "*"
google-pubsub1 = "5.0.3+20230119"
hyper = "1.2.0"
hyper-rustls = "0.26.0"
yup-oauth2 = "8.3.2"
```

公式のrepoのREADMEにversion指定あるならそれを使おうと思ったのですがないっぽかったです。


色々ゴニョゴニョしましたが動かなかったので方針を変更。


https://github.com/u-na-gi/async-google-apis/blob/master/example_crates/gcs_example/src/main.rs

色々してたら良さそうなの発見。
原理的にはこれと似たような感じでclientを生成してpub/subもできそうな気がする。



-> 次回へ続く