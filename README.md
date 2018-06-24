learn https://outcrawl.com/go-graphql-realtime-chat/

# 事前準備
```bash
$ go get -u golang.org/x/vgo
$ go get -u github.com/vektah/gqlgen

```

# サーバー起動方法
```bash
$ cd server
$ go generate ./...
$ vgo build
$ ./graphql-realtime-chat
```

# GraphQLサーバーの呼び出し
```bash
# 作成
curl http://localhost:8080/query -XPOST -H 'Content-Type:application/json' \
-d \
'
{
  "query": "mutation createTodo { createTodo(text:\"test\") { user {   id } text done }}"
}
'
 

# 取得
curl http://localhost:8080/query -XPOST -H 'Content-Type:application/json' \
-d \
'
{
  "query": "query findTodos { todos { text done user {   name } } }"
}
'
 

```

 