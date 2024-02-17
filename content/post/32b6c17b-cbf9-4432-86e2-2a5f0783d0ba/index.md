
---
title: Goで自作コマンドをインストールする
description: 
slug: 32b6c17b-cbf9-4432-86e2-2a5f0783d0ba
date: 2024-02-18T05:50:30+09:00
image: 
categories:
    - go
weight: 1      
---


Goで自作したコマンドをインストールして使いたいとします。

これ簡単で、main.goがある階層で、

```
go install
```

と実行するだけで、$GOPATHに追加されます。
GOAPTHがパス通ってれば、

```
<自作コマンド> -h
```

みたいな感じで呼び出すことが可能になります