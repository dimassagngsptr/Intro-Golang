package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	num := 5
	if reflect.TypeOf(num) != reflect.TypeOf(0) {
		fmt.Printf("Invalid data : %d \nData harus number", num)
		return
	}
	for i := num; i >= 1; i-- {
		res := ""
		for j := 1; j <= i; j++ {
			res += strconv.Itoa(j)
		}
		fmt.Println(res)
	}

}
