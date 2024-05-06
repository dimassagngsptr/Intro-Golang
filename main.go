package main

import (
	"fmt"
	"golang-3/utils"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	num1 := []int{3, 3, 3, 3, 3}
	num2 := []int{8, 8, 8, 8, 8}

	wg := &sync.WaitGroup{}
	result := make(chan int)

	wg.Add(2)
	go utils.CalculateSum(&num1, result, wg)
	go utils.CalculateSum(&num2, result, wg)

	go func() {
		wg.Wait()
		close(result)
	}()

	var sum1, sum2, total int

	for sum := range result {
		if sum1 == 0 {
			sum1 = sum
		} else {
			sum2 = sum
		}
		total += sum
	}
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Printf("Jumlah deret bilangan pertama:%v\nJumlah deret bilangan kedua:%v\nTotal:%v\n", sum1, sum2, total)
	fmt.Printf("Waktu yang dibutuhkan: %s\n", elapsed)
}
