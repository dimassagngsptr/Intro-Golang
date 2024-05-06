package utils

import (
	"fmt"
	"sync"
	"time"
)

func CalculateSum(numbers *[]int, result chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	sum := 0
	if len(*numbers) < 1 {
		fmt.Println("Error: jumlah deret dalam bilangan tidak boleh kosong")
		close(result)
		return
	}
	for _, num := range *numbers {
		if num <= 0 {
			fmt.Println("Error: number dalam slice tidak boleh bernilai minus atau tidak boleh bernilai 0 (nol)")
			close(result)
			return
		}
		sum += num
	}
	result <- sum
}

//buat method untuk validasi
//validasi tipe data
