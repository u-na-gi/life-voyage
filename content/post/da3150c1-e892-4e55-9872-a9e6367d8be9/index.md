---
title: flutterでprivate packageを作る
description: 
slug: da3150c1-e892-4e55-9872-a9e6367d8be9
date: 2024-03-03T02:49:19+09:00
image: 
categories:
    - flutter
weight: 1      
---


flutterでprivate packageを作りたいなと思ったので調べてみました。

https://blog.pentagon.tokyo/1721/#index_id0

検索では普通に出てきたのでできるっぽいですね。

ちなみにやりたいと思った理由は、protobufファイルの共通化のためです。
バックエンドと同じものを使いたいのですが、そのためのrepositoryなんて作りたいくないし、
ましてやflutter側にファイルをコピーしてつかうのもしんどいわけです。

そこで、バックエンドのprotobufを置いてるdirごとprivate packageとして切り出してflutter側からライブラリとしてダウンロードできるようにしちゃおうと考えたわけです。


https://medium.com/flutter-community/make-your-private-flutter-package-23a75ba899

と思ったけど、local開発専用感あるっぽくてなんかイメージと違うので多分やらない。
ごめんね。