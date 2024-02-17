---
title: goのメモ
description: 
slug: try-go-memo
date: 2024-02-18T04:59:40+09:00
image: 
categories:
    - go
weight: 1   
---


### ファイルパスを連結する

```
// post配下にファイルを作成する
postPath := filepath.Join(repo, "content", "post")

```

参考: https://takuroooooo.hatenablog.com/entry/2020/08/15/Go_path/filepath

### フォルダを作成する

```
if err := os.Mkdir(filepath.Join(postPath, slug), 0755); err != nil {
	log.Fatal(err)
}
```

参考: https://iketechblog.com/go-file-operations/

### ISO 8601形式でローカル時刻で現在日時を出力する

↓こういう感じのフォーマットのやつ
2024-02-18T05:38:42+09:00

```
// 現在の日時をローカルタイムゾーンで取得
localTime := time.Now().Local()
// ISO 8601形式で出力
date := localTime.Format(time.RFC3339)
```
