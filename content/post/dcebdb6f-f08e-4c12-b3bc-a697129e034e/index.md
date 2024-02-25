---
title: statefulなクラスでコンストラクタで値を渡したい
description: 
slug: dcebdb6f-f08e-4c12-b3bc-a697129e034e
date: 2024-02-26T04:41:45+09:00
image: 
categories:
    - flutter
weight: 1      
---

親クラスのコンストラクタに値を入れる

```
class Chat extends StatefulWidget {
  final String chatTypeKey; // 追加するプロパティ

  const Chat({super.key, required this.chatTypeKey}); // コンストラクタで初期化

  @override
  State<Chat> createState() => _ChatState();

}

```

この場合、stateクラス内でwidget.chatTypeKeyを通じてこの値にアクセスできる。

```

class _ChatState extends State<Chat> {
  List<types.Message> _messages = [];

  // _ChatState({required String chatTypeKey});

  @override
  void initState() {
    super.initState();
    // 画面立ち上げ時にデータをとってくる
    print("Chat Type Key: ${widget.chatTypeKey}");
    _loadMessages();
  }
  

```