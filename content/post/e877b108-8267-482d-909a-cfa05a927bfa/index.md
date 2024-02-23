---
title: flutterでios simulatorを立ち上げようとすると「Unable to boot the Simulator」と表示される
description: 
slug: e877b108-8267-482d-909a-cfa05a927bfa
date: 2024-02-23T14:52:35+09:00
image: 
categories:
    - flutter
    - ios
    - ios Simulator
weight: 1      
---

https://qiita.com/kokogento/items/465984cad8624f483782
この記事に従ってやってみた。

定期的に出てくるのなんでなんかな。


```
cd Library/Developer/CoreSimulator/Caches/dyld/23D56
```

23D56は人によって違うかも？？

```
rm -r ./*
```

で配下にあるキャッシュを全て削除する

