---
title: チャット画面を実装する
description: 
slug: ac59d1c2-88df-498f-89a6-a14e4cf96e42
date: 2024-02-26T03:37:41+09:00
image: 
categories:
    - flutter
weight: 1      
---


https://qiita.com/atm_33/items/0b066c34280b39628910
https://pub.dev/packages/flutter_chat_ui

使用したversion: flutter_chat_ui: ^1.6.12

これバックエンドはfirebase固定か??
いや違う。
sampleかなんかはfirebase使ってるってだけで非依存みたいやな

### typesの定義

```
import 'package:flutter_chat_types/flutter_chat_types.dart' as types;
```

自動importしてくれないので自分で書くことになるのかなあ

### messageを格納する

```

class _ChatState extends State<Chat> {
  List<types.Message> _message = [];
  
```

### サンプル
https://youtube.com/shorts/BmdQzobxX1M?si=HROTCOFTYCgUH_QM

この場合、widgetの実装は大体このようになっています。
細かいところは[公式repoのsample](https://github.com/flyerhq/flutter_chat_ui/tree/main/example)を見てね
```
@override
  Widget build(BuildContext context) => Scaffold(
        body: Chat(
          messages: _messages, // チャット画面に表示されてるメッセージ。自分の分も相手の分も対象
          onAttachmentPressed: _handleAttachmentPressed, // メッセージ送信フォームの横にある➕ボタンを押した時の挙動
          onMessageTap: _handleMessageTap,  // メッセージを押した時の挙動。自分の分も相手の分も対象。
          // このサンプルでは、画像をクリックしたら拡大表示するとか、リンクならページ開くとか。

          onPreviewDataFetched: _handlePreviewDataFetched, // データをとってくる??なんか更新してるっぽいけど
          onSendPressed: _handleSendPressed, // メッセージを送信する
          showUserAvatars: true, // ユーザーのアバターを表示するかどうか
          showUserNames: true, // ユーザー名を表示するかどうか
          user: _user, // 自身を表すユーザーID
          theme: const DefaultChatTheme( // テーマ。UI設定
            seenIcon: Text(
              'read',
              style: TextStyle(
                fontSize: 10.0,
              ),
            ),
          ),
        ),
      );
}
```