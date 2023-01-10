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

	/* math
	pl("5 + 4 = ", 5+4)
	pl("5 - 4 = ", 5-4)
	pl("5 * 4 = ", 5*4)
	pl("5 / 4 = ", 5/4)
	pl("5 % 4 = ", 5%4)
	mInt := 1
	mInt++ // == mInt += 1, mInt = mInt + 1
	pl(mInt)
	pl("Float Precision =", 0.111111111111111111+0.111111111111111111)
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs) // これを固定すれば同じ値が出る。
	pl(seedSecs)
	randNum := rand.Intn(50) + 1 // Intn は 0から引数までの間でランダム値を生成する。
	pl("Random :", randNum)
	pl("Abs(-10 =", math.Abs(-10))
	pl("Pow(4, 2) =", math.Pow(4, 2))
	pl("Sqrt(16) =", math.Sqrt(16))
	pl("Cbrt(8) =", math.Cbrt(8))
	pl("Ceil(4.4) =", math.Ceil(4.4))
	pl("Floor(4.4) =", math.Floor(4.4))
	pl("Roud(4.4) =", math.Round(4.4))
	pl("Log2(8) =", math.Log2(8))
	pl("Log10(100) =", math.Log10(100))
	// eの2乗の対数をとる
	pl("Log(7.389) =", math.Log(math.Exp(2)))
	pl("Max(5,4) =", math.Max(5, 4))
	pl("Min(5,4) =", math.Min(5, 4))
	// 90度をラジアンに変換する
	r90 := 90 * math.Pi / 180
	d90 := r90 * (180 / math.Pi)
	fmt.Printf("%.2f radians = %.2f degrees\n", r90, d90)
	pl("Sin(90) =", math.Sin(r90))
	// Cos, Tan, Acos, Asin, Atan, Asinh, Acosh, Atanh, Atan2, Cosh, Sinh, Sincos, Htpot, etc..
	*/

	/* format
	// %d : Integer
	// %c : Character
	// %f : Float
	// %t : Boolean
	// %s : String
	// %o : Base 8
	// %x : Base 16
	// %v : Guess based on data type
	// %T : Type of supplied value

	fmt.Printf("%s %d %c %f %t %o %x\n",
		"Stuff", 1, 'A', 3.14, true, 1, 1)
	var num float64 = 123456789.123456789
	fmt.Printf("9f: %9f\n", num)
	fmt.Printf(".9f: %.9f\n", num)
	fmt.Printf("9.f: %9.f\n", num)
	fmt.Printf("5f: %5f\n", num)
	fmt.Printf("5.f: %5.f\n", num)
	fmt.Printf(".5f :%.5f\n", num)
	fmt.Printf("2f: %2f\n", num)
	fmt.Printf("2.f: %2.f\n", num)
	fmt.Printf(".2f: %.2f\n", num)

	sp1 := fmt.Sprintf("%9.f\n", 3.141592)
	pl(sp1)
	*/

	/* for
	for x := 1; x <= 5; x++ {
		pl(x)
	}
	pl("---------")
	for x := 5; x >= 1; x-- {
		pl(x)
	}
	*/

	/* while
	fx := 0
	for fx < 5 {
		pl(fx)
		fx++
	}

	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randNum := rand.Intn(50)
	for true {
		fmt.Print("Guess a number between 0 and 50 :")
		// pl("Random Number is :", randNum)
		reader := bufio.NewReader(os.Stdin)
		guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		guess = strings.TrimSpace(guess)
		iGuess, err := strconv.Atoi(guess)
		if err != nil {
			log.Fatal(err)
		}
		if iGuess > randNum {
			pl("Pick a Lower Value")
		} else if iGuess < randNum {
			pl("Pick a Higher Value")
		} else {
			pl("You Guessed it")
			break
		}
	}
	*/

	/* range
	aNums := []int{1, 2, 3}
	for _, num := range aNums {
		pl(num)
	}
	*/

	/* arrays
	var arr1 [5]int
	arr1[0] = 1

	arr2 := [5]int{1, 2, 3, 4, 5}
	pl("Index 0 :", arr2[0])
	pl("Arr Length :", len(arr2))
	for i := 0; i < len(arr2); i++ {
		pl(arr2[i])
	}
	for i, v := range arr2 {
		fmt.Printf("index: %d, value:%d\n", i, v)
	}
	arr3 := [2][2]int{
		{1, 2},
		{3, 4},
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("arr3[%d][%d] = %d\n", i, j, arr3[i][j])
		}
	}

	aStr1 := "abcde"
	rArr := []rune(aStr1)
	pl(rArr) // [97 98 99 100 101]
	for _, v := range rArr {
		fmt.Printf("Rune Array : %d\n", v)
	}
	byteArr := []byte{'a', 'b', 'c'}
	pl(byteArr)
	bStr := string(byteArr[:])
	pl("I'm a string :", bStr)
	*/

	/* slices
	// var name []datatype, make([]type, size)
	sl1 := make([]string, 6)
	sl1[0] = "Society"
	sl1[1] = "of"
	sl1[2] = "the"
	sl1[3] = "Simulated"
	sl1[4] = "Universe"
	pl("Slice Size: ", len(sl1))
	for i := 0; i < len(sl1); i++ {
		pl(sl1[i])
	}
	for _, x := range sl1 {
		pl(x)
	}
	sArr := [5]int{1, 2, 3, 4, 5}
	sl2 := sArr[0:2]
	pl("1st 3 :", sArr[:3])
	pl("Last 3 :", sArr[2:])
	sArr[0] = 10
	// sl2 は sArrのコピーの一部なので、sArrを変更するとsl2も変更される。
	pl("sl2 :", sl2)
	// 同様に、sl2を変更すれば、sArrも変更される。
	sl2[0] = 1
	pl("sArr :", sArr)

	// ここでも同様に、コピー先が変更されれば、コピー元も変更される。
	sl3 := sArr[0:2]
	sl3 = append(sl3, 12)
	pl("sl3 :", sl3)
	pl("sArr :", sArr)

	sl4 := make([]string, 6)
	pl("sl4 :", sl4)       // sl4 : [     ]
	pl("sl4[0] :", sl4[0]) // sl4[0] :
	*/

}
