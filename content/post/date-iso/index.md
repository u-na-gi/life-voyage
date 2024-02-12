
---
title: osxでコマンドでiso8601形式で現在日時を取得する
description: 
slug: command-date-iso
date: 2024-02-12T23:59:50+09:00
image: 
categories:
    - Linux/Unix command
weight: 0.5      
---

osxでコマンドでiso8601形式で現在日時を取得する

```
date --iso-8601=seconds
```

このoptionはgnuのやつなので、osx標準のやつでは使えない。


```
brew install coreutils

alias date='gdate'
```

aliasはお使いのシェルの設定ファイルに書き込んどいてください（次回シェル立ち上げ時にもaliasを聞かせたい場合。）