# サーバー起動方法
```bash
$ cd graph && gqlgen -schema ../schema.graphql
$ vgo build
$ ./server
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

 
 
 