# 簡単なTodoアプリをGolangで作成する (学習用)

がんばる。

## 日記

2026.2.26
for rangeで一個目:インデックス番号、二個目:実データが入ってくるのを忘れてたー。

2026.2.27
DELETEリクエストの成功時の返し方はどっちがいいんだろう。(細かいけどね)
200 No Contentでシンプルにするか、200 OKでメッセージも返してあげるか。
基本前者で扱うらしい。

取得・追加・削除・更新のAPIをフレームワークを使わず標準ライブラリのみで実装してみたので、以下にテストするcurlコマンドをメモしておく。
curl -X GET http://localhost:8080/todos/
curl -X POST -H "Content-Type: application/json" -d "{\"id\": 3, \"title\": \"◯◯◯◯◯◯\", \"checked\": false}" http://localhost:8080/todos/
curl -X DELETE http://localhost:8080/todos/1
curl -X PUT -H "Content-Type: application/json" -d "{\"id\": 1, \"title\": \"◯◯◯◯◯◯\", \"checked\": true}" http://localhost:8080/todos/1
