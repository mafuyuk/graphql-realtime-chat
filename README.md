# GraphQL RealTimeChat 
## 📌 概要
https://outcrawl.com/go-graphql-realtime-chat/ を参考にリアルタイムチャットを作成する
構成はクライアント(React + Redux) - API(Golang + GraphQL) - Redis

## 🌐 動作環境 
* Golang v1.10.0
* Docker Latest Version

## ▶️ 実行方法
### 開発の仕方
```bash
$ make dstart
$ make run

```
### GraphQLサーバーの呼び出し
```bash
# 作成
curl http://localhost:8080/graphql -XPOST -H 'Content-Type:application/json' \
-d \
'
{
  "query": "mutation postMessage { postMessage(user:\"taro\", text:\"test text\") { user id text }}"
}
'


 

# 取得
curl http://localhost:8080/graphql -XPOST -H 'Content-Type:application/json' \
-d \
'
{
  "query": "query messages { messages { user id text } }"
}
'
 

```

