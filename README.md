go-mux-mongoDB

## Api Server

- mux web 框架
- go_mux_mongoDB 資料庫

## 檔案配置

- main.go - 主進入點,開啟 localhost:12345 的 server api
- db/db.go - 資料庫連線
- controller/library.go - library 資料庫的操作

將`db.go`裡的 `LbCollection` 導入使用

## API 功能

| 動作   | api                   | func        | 說明         |
| ------ | --------------------- | ----------- | ------------ |
| GET    | api/                  | Index       | Welcome      |
| GET    | api/books             | GetBooks    | 查詢全部書   |
| DELETE | api/books             | DeleteBooks | 刪除全部書   |
| POST   | api/books             | NewBook     | 新增 1 本書  |
| GET    | api/books/isbn/{isbn} | GetBook     | 查詢書(isbn) |
| POST   | api/books/isbn/{isbn} | UpdateBook  | 更新書(isbn) |
| DELETE | api/books/isbn/{isbn} | DeleteBook  | 刪除書(isbn) |

## 使用外部變數的方法

```
import ("RESTful/go_mux_mongoDB/db")
```
