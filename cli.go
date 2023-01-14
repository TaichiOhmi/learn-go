package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// 標準入力から読み込む
	scanner := bufio.NewScanner(os.Stdin)
	// 1行ずつ読み込んで繰り返す
	for scanner.Scan() {
		//1行分を出力する
		fmt.Println(scanner.Text())
	}
	// まとめてエラー処理をする
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "読み込みに失敗しました:", err)
	}

	// パスを結合する
	path := filepath.Join("GoTut", "cl.go")
	// 拡張子を取る
	fmt.Println(filepath.Ext(path))
	// ファイル名を取得
	fmt.Println(filepath.Base(path))
	// ディレクトリ名を取得
	fmt.Println(filepath.Dir(path))

	// Goファイルを探し出す
	err := filepath.Walk("./",
		func(path string, info os.FileInfo, err error) error {
			if filepath.Ext(path) == ".go" {
				fmt.Println(path)
			}
			return nil
		})
	if err != nil {
		log.Fatal(err)
	}

}
