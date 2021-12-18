package main

import (
	"log"
	"net/http"
)

// http.メソッドの種類
// http.ListenAndServe(TCPのアドレス , http.Handler)
// -> サーバーを起動する
// http.Handler
// -> ServeHTTPメソッドを定義しているinterface
// ServeHTTP
// -> HTTPリクエストを受けてHTTPレスポンスを返す処理
// http.ServeMux
// -> HTTPリクエストのURLとそれに対応するハンドラを登録
// -> リクエストが来たらURLにマッチしたハンドラを呼び出すマルチプレクサ
//　http.HandlerFunc
// -> func(ResponseWriter, *Request) を別の型として定義したもの
// -> HandlerFunc方にはServeHTTPメソットが実装されている
// -> HandlerFuncにてキャストするだけで構造体を宣言することなくhttp.Handlerを実装できる
// -> ServeHTTPを実装するstructを作らなくて良い

func main() {
	server := NewPlayerServer(NewInMemoryPlayerStore())

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
