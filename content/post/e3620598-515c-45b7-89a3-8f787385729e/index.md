---
title: "Extension host terminated unexpectedly."と急に言われ出した時
description: 
slug: e3620598-515c-45b7-89a3-8f787385729e
date: 2024-02-26T02:24:41+09:00
image: 
categories:
    - vscode
    - Mac(intel)
weight: 1      
---

# 環境

```
❰yuunag1❙~/workspace/life-voyage(git:main)❱✔≻ sw_vers
ProductName:            macOS
ProductVersion:         14.3
BuildVersion:           23D56
```

なお、僕のMacはintelです。

# 結論

一時的な策の可能性があるっぽいがこちらで言われてる通りにdocker desktop for Macの設定を変更したらなおった。

https://github.com/microsoft/vscode/issues/194458#issuecomment-1912849373

どうやらMac(intel)の仮想化の何かが問題っぽい