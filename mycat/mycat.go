/* 作成するcatコマンドの仕様
引数でファイルパスの一覧を貰い、そのファイルを与えられた順に標準出力にそのまま出力するコマンドを作ってください。
また、-nオプションを指定すると、行番号を各行につけて表示されるようにしてください。
なお、行番号はすべてのファイルで通し番号にしてください。
*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var n bool

func init() {
	// ポインタを指定して設定を予約
	// BoolVar(p *bool, name string, value bool, usage string)
	flag.BoolVar(&n, "n", false, "行番号の有無")
}

func readAndWrite(fn string, countPtr *int) error {
	f, err := os.Open(fn)
	if err != nil {
		return err
	}
	defer f.Close()

	// 標準入力から読み込む
	scanner := bufio.NewScanner(f)
	// 1行ずつ読み込んで繰り返す
	for scanner.Scan() {
		if n {
			_, err := fmt.Fprintln(os.Stdout, fmt.Sprintf("%d:", *countPtr), scanner.Text())
			if err != nil {
				return err
			}
			*countPtr++
		} else {
			_, err := fmt.Fprintln(os.Stdout, scanner.Text())
			if err != nil {
				return err
			}
		}
	}
	// まとめてエラー処理をする
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "読み込みに失敗しました:", err)
	}

	return nil
}

func main() {
	flag.Parse()

	count := 1

	for _, fn := range flag.Args() {
		err := readAndWrite(fn, &count)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

}
