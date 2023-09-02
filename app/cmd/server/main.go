package main

// https://qiita.com/melty_go/items/eca50ac224ec42b85af1
import (
	// フォーマットI/Oを扱うパッケージです。
	// C言語のprintfおよびscanfと似た関数を持ちます。
	"fmt"

	// ロギングを行うシンプルなパッケージです。
	// 出力をフォーマットするメソッドを持つLogger型が定義されています。
	"log"

	// HTTPを扱うパッケージです。
	// HTTPクライアントとHTTPサーバーを実装するために必要な機能が提供されています
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)

	// ポート番号8080で待ち受けるHTTPサーバを起動します。
	// http.ListenAndServe(":8080", nil)
	fmt.Println("Waiting port 8080 ...")
	log.Fatal("[Fatal] [main] ", http.ListenAndServe(":8080", nil))
}
