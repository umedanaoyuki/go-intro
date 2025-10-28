package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()
 	number, _ := strconv.Atoi(sc.Text())
 	
 	if number < 100 {
 	    fmt.Println(strconv.Itoa(number) + "は100より小さい")
 	} else if number < 200 {
 	    fmt.Println(strconv.Itoa(number) + "は100以上200より小さい")
 	} else {
 	    fmt.Println(strconv.Itoa(number) + "は200以上")
 	}
}
