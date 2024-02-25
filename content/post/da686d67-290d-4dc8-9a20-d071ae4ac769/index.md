---
title: python tips
description: 
slug: da686d67-290d-4dc8-9a20-d071ae4ac769
date: 2024-02-26T02:36:35+09:00
image: 
categories:
    - python
weight: 1      
---

数億年ぶりにpython触ってて何もかも忘れたので開発しながら参考にした記事とかをまとめとこうと思います。


### VSCodeにPoetryの仮想環境を認識させる

https://zenn.dev/bee2/articles/74b975c70ae6ed

- pathを通したい。

```
poetry env info
```

この中で、executable pathだけ取ってきたい

```
poetry env info -e
```

詳細はhelpを見てね。

### 仮想環境に入る

https://qiita.com/ksato9700/items/b893cf1db83605898d8a#5-%E4%BB%AE%E6%83%B3%E7%92%B0%E5%A2%83%E3%81%A7%E5%AE%9F%E8%A1%8C
