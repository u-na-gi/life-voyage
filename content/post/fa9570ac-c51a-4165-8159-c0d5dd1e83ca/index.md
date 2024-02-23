---
title: flutterfire configure した時に「unhandled exception」が起こる
description: 
slug: fa9570ac-c51a-4165-8159-c0d5dd1e83ca
date: 2024-02-23T15:27:23+09:00
image: 
categories:
    - flutter
    - firebase
weight: 1      
---

環境

```
❰yuunag1❙~/workspace/life-voyage(git:main)❱✔≻ sw_vers
ProductName:            macOS
ProductVersion:         14.3
BuildVersion:           23D56

❰yuunag1❙~/workspace/life-voyage(git:main)❱✔≻ flutter --version
Flutter 3.16.7 • channel stable • https://github.com/flutter/flutter.git
Framework • revision ef1af02aea (6 weeks ago) • 2024-01-11 15:19:26 -0600
Engine • revision 4a585b7929
Tools • Dart 3.2.4 • DevTools 2.28.5
```

```
flutterfire configure --project=<your project>
```
を実行した際に以下のようなエラーが出た。


```
Unhandled exception:
Exception: /System/Library/Frameworks/Ruby.framework/Versions/2.6/usr/lib/ruby/2.6.0/rubygems/core_ext/kernel_require.rb:54:in `require': cannot load such file -- xcodeproj (LoadError)
        from /System/Library/Frameworks/Ruby.framework/Versions/2.6/usr/lib/ruby/2.6.0/rubygems/core_ext/kernel_require.rb:54:in `require'
        from -e:1:in `<main>'

#0      ConfigCommand.run (package:flutterfire_cli/src/commands/config.dart:540:11)
<asynchronous suspension>
#1      CommandRunner.runCommand (package:args/command_runner.dart:212:13)
<asynchronous suspension>
#2      main (file:///Users/yuunag1/.pub-cache/hosted/pub.dev/flutterfire_cli-0.2.7/bin/flutterfire.dart:57:5)
<asynchronous suspension>
```

### やったこと

```
sudo gem install xcodeproj
```

をやってみた。


### 解決策

```
sudo gem install xcodeproj
```

したら治った。