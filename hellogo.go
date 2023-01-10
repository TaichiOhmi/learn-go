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
}
