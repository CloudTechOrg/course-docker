package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func init() {
	// .envファイルから環境変数を読み込む
	godotenv.Load()
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// CORSヘッダーを設定
	w.Header().Set("Access-Control-Allow-Origin", "*")             // すべてのオリジンからのアクセスを許可
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS") // GETとOPTIONSメソッドを許可
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type") // 特定のヘッダーの使用を許可

	// リクエストメソッドがOPTIONSの場合は、プリフライトリクエストとして扱う
	if r.Method == "OPTIONS" {
		return // プリフライトリクエストにはステータス200で応答して、処理を終了する
	}

	// hello worldという文字列をレスポンスとして返す
	fmt.Fprintf(w, "API接続テストが成功しました")
}

func DbTestHandler(w http.ResponseWriter, r *http.Request) {
	// CORSヘッダーを設定
	w.Header().Set("Access-Control-Allow-Origin", "*")             // すべてのオリジンからのアクセスを許可
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS") // GETとOPTIONSメソッドを許可
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type") // 特定のヘッダーの使用を許可

	// リクエストメソッドがOPTIONSの場合は、プリフライトリクエストとして扱う
	if r.Method == "OPTIONS" {
		return // プリフライトリクエストにはステータス200で応答して、処理を終了する
	}

	reservation_count, err := database_test()

	if err != nil {
		http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// 正しいフォーマットでレスポンスを返す
	fmt.Fprintf(w, "データベース接続テストが成功しました（Reservationsの件数：%d）", reservation_count)
}

func database_test() (int, error) {
	// 環境変数からデータベース接続の各要素を取得
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	servername := os.Getenv("DB_SERVERNAME")

	// 固定となるデータベース接続変数を定義
	port := "3306"
	dbname := "testdb"

	// 接続文字列を組み立て
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, servername, port, dbname)
	if connectionString == "" {
		log.Fatal("DB connection string is not set")
	}

	// データベースに接続
	connection, err := sql.Open(
		"mysql",
		connectionString)
	if err != nil {
		return 0, err
	}
	defer connection.Close()

	// SQLの実行
	rows, err := connection.Query("SELECT COUNT(*) AS reservation_count FROM test_table")
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	// 結果の読み取り
	var reservation_count int
	for rows.Next() {
		err := rows.Scan(&reservation_count)
		if err != nil {
			return 0, err
		}
	}

	return reservation_count, nil
}

func main() {
	// APIポートを8080に固定
	apiport := "8080"

	// /パスにアクセスがあった場合に、helloHandler関数を実行するように設定
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/dbtest", DbTestHandler)

	// 8080ポートでサーバーを起動
	fmt.Println("HTTPサーバを起動しました。ポート: " + apiport)
	err := http.ListenAndServe(":"+apiport, nil)
	if err != nil {
		fmt.Println("HTTPサーバの起動に失敗しました: ", err)
	}
}
