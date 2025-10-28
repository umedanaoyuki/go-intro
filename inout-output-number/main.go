package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 整数を読み込み、そのまま出力するプログラム
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
 	number, _ := strconv.Atoi(sc.Text())
	fmt.Println(number)
}