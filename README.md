learn https://outcrawl.com/go-graphql-realtime-chat/

# 開発時
```bash
$ make dstart
$ make run

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

