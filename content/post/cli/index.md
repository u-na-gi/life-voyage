
---
title: goでcliを作成する
description: 
slug: go-create-cli
date: 2024-02-18T04:33:44+09:00
image: 
categories:
    - Go
    - cli
    - cobra
weight: 1     
---


https://github.com/spf13/cobra を使用します。

### initialize

対象dirで、

```
cobra-cli init
```

参考: https://github.com/spf13/cobra-cli/blob/main/README.md#initalizing-a-cobra-cli-application

### 新しくコマンドを追加するとき

```
cobra-cli add <追加したいコマンド名>
```

参考: https://github.com/spf13/cobra-cli/blob/main/README.md#add-commands-to-a-project
