package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()
	greeting := sc.Text()
	
	if greeting == "Hello" {
	    fmt.Println("こんにちは")
	} else {
	    fmt.Println(greeting + "はHelloではない")
	}
}
