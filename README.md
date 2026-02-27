# 簡単なTodoアプリをGolangで作成する (学習用)

がんばる。メモしたいこと書き連ねる。<br>
標準パッケージからGin/Echoの使い方に慣れるまで。<br>

## 日記

### 2026.2.26 <br>

for rangeで一個目:インデックス番号、二個目:実データが入ってくるのを忘れてたー。

---

### 2026.2.27 <br>

DELETEリクエストの成功時の返し方はどっちがいいんだろう。(細かいけどね) <br>
200 No Contentでシンプルにするか、200 OKでメッセージも返してあげるか。 <br>
基本前者で扱うらしい。

CRUD（追加・取得・更新・削除）のAPIをフレームワークを使わず標準ライブラリのみで実装してみたので、以下にテストするcurlコマンドをメモしておく。<br>

- Create:作成 <br>
  curl -X POST -H "Content-Type: application/json" -d "{\"id\": 3, \"title\": \"◯◯◯◯◯◯\", \"checked\": false}" http://localhost:8080/todos/

- Read:取得 <br>
  curl -X GET http://localhost:8080/todos/

- Update:更新 <br>
  curl -X PUT -H "Content-Type: application/json" -d "{\"id\": 1, \"title\": \"◯◯◯◯◯◯\", \"checked\": true}" http://localhost:8080/todos/1

- Delete:削除 <br>
  curl -X DELETE http://localhost:8080/todos/1

---

GinでのCRUD実装について <br>
go mod init gin <br>
go get -u github.com/gin-gonic/gin <br>
を必ず最初に行いパッケージをインストールする。 <br>
