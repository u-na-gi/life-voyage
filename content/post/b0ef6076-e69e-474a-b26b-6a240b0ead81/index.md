---
title: RustとPythonでPub/Subことはじめ4
description: 
slug: b0ef6076-e69e-474a-b26b-6a240b0ead81
date: 2024-03-07T09:29:57+09:00
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


# subscriberはpythonで用意してみる

repositoryは前回と同じく

https://github.com/u-na-gi/google-cloud-rust-pubsub-sample
を使っていきます。


## ryeを使って環境構築

https://dev.classmethod.jp/articles/get-start-rye-python/

ryeの公式はこっち
https://github.com/astral-sh/rye


## pythonでsubscribeする 

python側は公式ライブラリがあるので、それを使います。
https://cloud.google.com/pubsub/docs/publish-receive-messages-client-library?hl=ja


interpreterの設定うまくできないかなあと思ったのですが、vscodeでworkspaceのパスとってくれるっぽい

```
{
  "python.defaultInterpreterPath": "${workspaceFolder}/python-subscriber/.venv/bin/python"
}
```

reference: https://zenn.dev/nowa0402/articles/85833db7ff2e13

## subscription type

さて、python側はpull型で処理を行うわけですが、それにも種類があります。
ざっくり分けて

- 一回一回処理するpull型
- メッセージが利用可能になるとすぐに使用される streamingPull型

今回は以下を参考にstreamingPull型で行おうと思います。

### 気になったこと

Q. gcpのpub/subでmessageをsubscriberがstreaming pullの時、1回のトピック送信で何度も受信してる。
subscriber側が確認する応答が遅くて何回も送ってるっぽい。


これはどうやら、　非公式のpubsub-emulatorでは無理かもしれない

## pub sub emulator

公式のを使う
https://cloud.google.com/pubsub/docs/emulator?hl=ja#manually_setting_the_variables


https://cloud.google.com/sdk/gcloud/reference/beta/emulators/pubsub/start

いやごめんRust側で何回も送ってただけだったわ。。。。

今回使用したrepositoryはこちら
https://github.com/u-na-gi/google-cloud-rust-pubsub-sample/tree/main/python-subscriber

-> 次回: 本番のtopicを使って行う