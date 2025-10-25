package main

import "fmt"


func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}


func main() {

	var i int = 100
	fmt.Println(i)

	var s string = "Hello, Go!"
	fmt.Println(s)		

	var t,f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	// 変数の型だけを指定して宣言　→ 初期値が入るようにする
	// intの場合は0
	var i3 int
	// stringの場合は空文字列
	var s3 string
	fmt.Println(i3, s3)

	i3 = 300
	s3 = "InitializedGo"
	fmt.Println(i3, s3)

	i = 150

	fmt.Println(i)

	// 暗黙的な定義（関数外では定義することができない）
	// 変数名 := 値（複数回 := を使用して再定義はできない）
	// 型推論される
	i4 := 400
	fmt.Println(i4)

	i4 = 450
	fmt.Println(i4)

	// エラーになる
	//i4 = "Hello Golang"

	outer()

	var s5 string = "Not User"
}