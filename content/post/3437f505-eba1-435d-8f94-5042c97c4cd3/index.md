---
title: firebaseをいつでも剥がせるように疎結合に保つ
description: 
slug: 3437f505-eba1-435d-8f94-5042c97c4cd3
date: 2024-02-26T14:27:37+09:00
image: 
categories:
    - firebase
    - flutter
    - 思索
    - 技術ポエム
weight: 1      
---


これは何もfirebaseに限ったことではなく、saas全般に言えるものだと思ってます。

たた追えば、サービスが死んだ、料金体系が変わって使うのをやめなくてはならなくなった。
こういったケースは往々にしてあると思います。

なのでこういう輩は出来る限りいつでも剥がせるようにしたいわけです。


## 疎結合ってなに?

仰々しい名前使っておきながら僕は正直その辺の用語を正確に理解しているわけではないので、僕の中で思っている疎な状態というのを考えてみます。

ざっくり言って、コアな修正に入った時に影響箇所を少なく保つことができれば疎な状態と言えるんじゃないかなと思っています。

## 複数のファイルに散らばる実装を集約する

今回の主題だと、firebaseをやめるとなった時に他の箇所に影響が少なければいいわけです。

例として、

./a.dart
```dart

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  var di = DependencyInjection();
  di.configure();
  await Firebase.initializeApp();

  runApp(const MyApp());
}

```

./b.dart
```dart
onPressed: () async {
    // Google認証
    GoogleSignInAccount? signInAccount = await googleLogin.signIn();
    if (signInAccount == null) return;
    GoogleSignInAuthentication auth =
        await signInAccount.authentication;
    final OAuthCredential credential =
        GoogleAuthProvider.credential(
            idToken: auth.idToken, accessToken: auth.accessToken);
    // 認証情報をFirebaseに登録
    await FirebaseAuth.instance.signInWithCredential(credential);
},

```

./c.dart
```dart
} else if (index == 2) {
    // アカウント アイコンがタップされたときのアクション

    try {
        //ログアウト成功時の処理
        await FirebaseAuth.instance.signOut();
    } catch (e) {
        // ログアウト失敗時のエラー処理
        print("ログアウトエラー: $e");
    }
}

```

というように複数のファイルでバラバラに書かれています。

これを一つにまとめちゃうわけです。

```dart
import 'package:firebase_auth/firebase_auth.dart';
import 'package:firebase_core/firebase_core.dart';
import 'package:google_sign_in/google_sign_in.dart';

class FirebaseServiceUserStream {
  late Stream<User?> _st;

  FirebaseServiceUserStream(Stream<User?> st) {
    _st = st;
  }

  Stream<User?> get() {
    return _st;
  }
}

class FirebaseService {
  final googleLogin = GoogleSignIn(scopes: [
    'email',
  ]);
  static FirebaseService? _instance;

  static Future<FirebaseService> init() async {
    await Firebase.initializeApp();
    _instance = FirebaseService();
    return _instance!;
  }

  FirebaseServiceUserStream listener() {
    return FirebaseServiceUserStream(FirebaseAuth.instance.authStateChanges());
  }

  Future<void> login() async {
    // Google認証
    GoogleSignInAccount? signInAccount = await googleLogin.signIn();
    if (signInAccount == null) return;
    GoogleSignInAuthentication auth = await signInAccount.authentication;
    final OAuthCredential credential = GoogleAuthProvider.credential(
        idToken: auth.idToken, accessToken: auth.accessToken);
    // 認証情報をFirebaseに登録
    await FirebaseAuth.instance.signInWithCredential(credential);
  }

  Future<void> logout() async {
    await FirebaseAuth.instance.signOut();
  }
}

```

くそ雑な実装ですが例えばこんな感じかなと思います。
ポイントは、libraryの呼び元を一箇所に集約しているという点です。
firebaseのライブラリを使って行う操作は常にこのクラスが使われるようになるというわけです。

ここでポイントはStreamのラップです。

```
FirebaseAuth.instance.authStateChanges()
```

の戻り値は、Stream\<User?\>です。
このUser型は、

```
import 'package:firebase_auth/firebase_auth.dart';
```

から読み込まれています。

これをこのままauthで呼び出してしまうと、auth側でも型を合わせるために、firebase_auth.dart  を呼び出してしまうような事態になりかねません。

なぜ呼び出すかというと、実際にこのstream処理が行われるのは、もっと外側だからです。
受け渡していくためには戻り値の指定が必要ですし、そのために毎回firebaseを呼び出すのは避けたいと思いラップしています。


さて、これで集約ができたぞわーい〜あとはこれを色んなとこで呼べばいいや！とはならないわけですね。
それを踏まえて次に行きます。


##  この集約したクラスをラップする

上記でまとめたファイルを色んなところで適当に呼ぶだけでは疎結合とは言えませんな。

このままでは本当にfirebaseが必要がなくなった時に結局たくさんの箇所で変更が必要になってしまいます。

今回の場合、firebaseはauthとして使ってるので、「auth」をまとめるサービスを作ってそちらから呼ぶようにして、UIはauthとして振る舞うようにし、裏側でfirebaseをを呼んでいるということは隠蔽したいと思います。


```

class AuthServiceUserStream {
  late FirebaseServiceUserStream _st;

  AuthServiceUserStream(FirebaseServiceUserStream st) {
    _st = st;
  }

  Stream<User?> get() {
    return _st.get();
  }
}

class AuthService {
  final firebaseService = DependencyInjection.getIt.get<FirebaseService>();

  Future<void> login()async  {
    await firebaseService.login();
  }

  Future<void> logout() async {
    await firebaseService.logout();

  }
 
  AuthServiceUserStream listener() {
    return AuthServiceUserStream(firebaseService.listener());
  }

  getData() {
    // 取得したfirebaseのデータがそれぞれ専用の型を持ってたとしてもここでキャストして使う
  }
}

```

viewModel以降では、このserviceを介して、呼ばれることになります。
こうすることで、firebaseを削除したとしてもその主な影響はauthに集約され、他のファイルでは、authの型で呼ばれているので、authに破壊的な変更をしない限りはさほど影響がないという状況を作ることができます。




## まとめ

firebaseなりBaaSは詰め替え可能というか捨てる可能性が出てくるということを見越した上で書いた方がいいかなと思っています。
でも、listenっていうfirebaseのmethodをラップしちゃってるので、今後これが使えなくなった時が問題。。。かなあ。

