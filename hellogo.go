package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	/* userInput
	pl("Hello Go")
	pl("What is your name?")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err == nil {
		pl("Hello", name)
	} else {
		log.Fatal(err)
	}
	*/

	/* variables
	// var <name> <type>
	// 変数名や関数名、型名が大文字で始まる場合、エクスポートすることを表す。
	// 基本の変数名はキャメルケース
	// var vName string = "Taichi" // 型を明示
	// var v1, v2 = 1.2, 3.4       //
	// var v3 = "hello"
	// v4 := 2.4
	// v4 = 5.4
	*/

	/* data types
	// pl(reflect.TypeOf(25))
	// pl(reflect.TypeOf(3.14))
	// pl(reflect.TypeOf(true))
	// pl(reflect.TypeOf("Hello"))
	// pl(reflect.TypeOf('🦍'))
	*/

	/* casting strings
	cV1 := 1.5
	cV2 := int(cV1)
	pl(cV2)
	cV3 := "50000000"
	// Atoi means ascii to integer
	cV4, err := strconv.Atoi(cV3)
	pl(cV4, err, reflect.TypeOf(cV4))
	cV5 := 50000000
	// Itoa means integer to ascii
	cV6 := strconv.Itoa(cV5)
	pl(cV6, reflect.TypeOf(cV6))
	cV7 := "3.14"
	if cV8, err := strconv.ParseFloat(cV7, 64); err == nil {
		pl(cV8, reflect.TypeOf(cV8))
	}
	cV9 := fmt.Sprintf("%f", 3.14)
	pl(cV9, reflect.TypeOf(cV9))
	*/

	/* conditional operators
	iAge := 8
	if (iAge >= 1) && (iAge <= 18) {
		pl("Important Birthday")
	} else if (iAge == 21) || (iAge == 50) {
		pl("Important Birthday")
	} else if iAge >= 65 {
		pl("Important Birthday")
	} else {
		pl("Not an Important Birthday")
	}
	pl("!true =", !true)
	*/

	/* strings
	sV1 := "A word" // []byte
	replacer := strings.NewReplacer("A", "Another")
	sV2 := replacer.Replace(sV1)
	pl(sV2)
	pl("Length :", len(sV2))
	pl("Contains Another :", strings.Contains(sV2, "Another"))
	pl("o index :", strings.Index(sV2, "o"))
	pl("Replace :", strings.Replace(sV2, "o", "0", -1))
	// strings.Replaceは、文字列sのコピーに、重複しない最初のn個のoldをnewで置き換えたものを返す。oldが空の場合、文字列の先頭とUTF-8シーケンスの後にマッチし、 k個のルーン文字列に対して最大k+1個の置換が行われる。n < 0 の場合、置換の回数に制限はない。
	sV3 := "\nSome Words\n" // \t, \" \\
	pl(sV3)
	sV3 = strings.TrimSpace(sV3)
	// TrimSpace は，文字列 s から Unicode で定義されているように先頭と末尾の空白をすべて取り除いたスライスを返します。
	pl(sV3)
	pl("Split :", strings.Split("a-b-c-d", "-"))
	pl("Lower :", strings.ToLower(sV2))
	pl("Upper :", strings.ToUpper(sV2))
	// HasPrefix 第一引数が第二引数の文字列で始まっているかを返す。
	pl("Prefix :", strings.HasPrefix("tacocat", "taco"))
	// HasSuffix 第一引数が第二引数の文字列で終わっているかを返す。
	pl("Suffix :", strings.HasSuffix("tococat", "cat"))
	*/

	/* printf
	rStr := "abcdefg"
	pl("Rune Count :", utf8.RuneCountInString(rStr))
	for i, runeVal := range rStr {
		fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal)
	}
	*/

	/* time
	now := time.Now()
	pl(now.Year(), now.Month(), now.Day())
	pl(now.Hour(), now.Minute(), now.Second())
	*/
}
