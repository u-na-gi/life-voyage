---
title: 所有権について1
description: 
slug: 6726b01c-c446-4d07-bb85-e0cba81d582d
date: 2024-03-03T02:08:27+09:00
image: 
categories:
    - Rust
    - 所有権
    - tips
weight: 1      
---

Rustの所有権周りで詰まったのでメモ

```
let r: PostMessageRequest = request.into_inner().clone();
```

こういうコードがあったとして、into_inner()を使った時にすでに所有権が移動しているとのこと。

なので、たいていの言語で何気なく書くような以下のようなコードを書くとエラーになる。

```
let user_id = request.into_inner().user_id;
let message = request.into_inner().message;
// 一度目のrequestで、所有権が移動しているため、上記のように二回目は呼び出せない。
// なお、以下のようなエラーが発生している
// use of moved value: `request`
// value used here after moverustcClick for full compiler diagnostic
// server.rs(28, 27): `request` moved due to this method call
// server.rs(18, 5): move occurs because `request` has type `tonic::Request<PostMessageRequest>`, which does not implement the `Copy` trait
// request.rs(140, 23): `tonic::Request::<T>::into_inner` takes ownership of the receiver `self`, which moves `request`
println!("posted message user id: {}", user_id);
println!("posted message: 「{}」", message);
```

これを解決するためには、into_innerを一度だけ呼び出すようにすれば良い。

```
let r: PostMessageRequest = request.into_inner();
let user_id = r.user_id;
let message = r.message;
println!("posted message user id: {}", user_id);
println!("posted message: 「{}」", message);
```

なお、chatGPTちゃんは以下のように言ってました。

> 問題は、Rustが所有権と借用をどのように扱うか、特にtonicクレートのRequestオブジェクトに関してです。into_inner()をRequestオブジェクトに対して呼び出すと、そのオブジェクトの所有権を取得し、それを移動させます。つまり、この操作後、元のオブジェクトはもう使用できなくなります。なぜなら、それは元のコンテキストから「移動」されたからです。
>
> 提供されたコードスニペットでは、requestを二回使用しようとしていますが、これが問題につながっています：
>
> - 最初に、requestをクローンしてからinto_inner()を呼び出し、user_idを取得します。この操作自体は問題ありませんが、クローンを操作しているためです。
> - しかし、元のrequestオブジェクト自体にinto_inner()を再度呼び出してmessageを取得しようとすると、Rustの所有権ルールにより、requestは最初のinto_inner()呼び出しによって既に移動されているため、これを行うことができません。
>
> この問題を修正するには、requestを一度に一回だけ移動させるようにするか、それを移動させることなくその内容を操作する方法を確保する必要があります。
>
> tonic::Requestがその内部の内容を直接借用する方法を提供していないため、直接借用して必要な部分を扱う代わりに、全体をクローンすることは避けるべきです。
>
> tonic::Requestがどのように使用されるかによりますが、一般的には、requestの消費を避けるためにアクセス方法を計画することが必要です。tonic::Request<T>に対して利用可能なメソッドについては、使用しているtonicのバージョンに依存するため、具体的なコードソリューションを提供することはできません。tonicのAPIドキュメントやソースコードにアクセスできる場合は、それを消費せずにリクエストの内容にアクセスまたは借用する方法を提供するメソッドを探してください。


reference:

- https://docs.rs/tonic/latest/tonic/struct.Request.html#method.into_inner