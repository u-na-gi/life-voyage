---
title: flutter tips
description: 
slug: 9692cddc-832c-4ff5-b4ce-7912d5758596
date: 2024-02-18T14:05:38+09:00
image: 
categories:
    - flutter
weight: 1      
---


### flutterでasync constructorがしたい。

TS(JS)のasync constructorしたい時と同じような感じね。

```
class SecureStorageService {
  late FlutterSecureStorage _storage;
  static SecureStorageService? _instance;

  SecureStorageService(FlutterSecureStorage storage) {
    _storage = storage;
  }

  FlutterSecureStorage getStorage(){
    return _storage;
  }

  static Future<SecureStorageService> init() async{
    if (_instance != null) {
      return _instance!;
    }else{
      const storage = FlutterSecureStorage();
      _instance = SecureStorageService(storage);

      final res = _instance!.getStorage();
      await res.write(key: "GOOGLE_CLIENT_ID", value: "");
      return _instance!;
    }

  }

}
```

また、上記のようにすることでシングルトンにすることもできる

参考: https://www.rm48.net/post/%E3%80%90flutter-dart%E3%80%91%E3%82%B3%E3%83%B3%E3%82%B9%E3%83%88%E3%83%A9%E3%82%AF%E3%82%BF%E3%81%A7%E9%9D%9E%E5%90%8C%E6%9C%9F%E5%87%A6%E7%90%86%E3%82%92%E8%A1%8C%E3%81%84%E3%81%9F%E3%81%84%E5%A0%B4%E5%90%88%E3%81%AE%E5%AF%BE%E5%BF%9C




