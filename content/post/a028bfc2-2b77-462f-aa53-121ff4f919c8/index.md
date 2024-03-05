---
title: vscodeの設定ファイルの位置
description: 
slug: a028bfc2-2b77-462f-aa53-121ff4f919c8
date: 2024-03-05T14:00:54+09:00
image: 
categories:
    - vscode
    - tips
weight: 1      
---

REST Clientをよく検証のために使います。

```
Name: REST Client
Id: humao.rest-client
Description: REST Client for Visual Studio Code
Version: 0.25.1
Publisher: Huachao Mao
VS Marketplace Link: https://marketplace.visualstudio.com/items?itemName=humao.rest-client
```

これで読み込む環境変数とかをvscodeのsettings.jsonに書き込むことが多いわけですがよく場所を忘れるのでメモすることにします。


Macユーザーの場合大抵ここだと思います。

```
code  /Users/$(whoami)/Library/Application\ Support/Code/User/settings.json
```