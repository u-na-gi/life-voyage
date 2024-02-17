---
title: fishのコマンドまとめ
description: 
slug: fish-config-path
date: 2024-02-12T23:59:50+09:00
image: 
categories:
    - fish
weight: 0.5      
---


### fishで設定ファイルを書き換えたい時のパスの場所

```
~/.config/fish/config.fish
```

### fishでpathを通す

```
set -U fish_user_paths <追加したいパス> $fish_user_paths
```

参考: https://qiita.com/ledsun/items/8ca1a450b21c8ebc9670

[invalid variable nameと言われた時]

```

❰yuunag1❙~/workspace/life-voyage/cli(git✱≠main)❱✔≻ set -U fish_user_paths　/Users/yuunag1/go/bin　$fish_user_paths

set: fish_user_paths　/Users/yuunag1/go/bin　/Users/yuunag1/development/flutter/bin: invalid variable name. See `help identifiers`

```

解決策:

次のように入力し直す

```
set -U fish_user_paths $fish_user_paths /Users/yuunag1/go/bin
```

$fish_user_pathsを前に持ってくることで後続に追加されるっぽい


### fishで環境変数を設定する

```
set -Ux NAME_OF_ENVIRONMENT_VARIABLE value_of_environment_variable
```

参考: https://zenn.dev/t4aru/articles/aed6dc4312b28f