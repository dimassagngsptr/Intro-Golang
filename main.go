package main

import (
	"fmt"
	"golang-2/utils"
)

func main() {
	//method soal 1
	functions := utils.Method()
	app := functions["slice"].(func() []int)()
	fmt.Println("append:", app)
	length := functions["length"].(func() []int)()
	fmt.Println("len:", len(length))
	exampleMake := functions["exMake"].(func() map[string]string)()
	fmt.Println("make:", exampleMake)
	copies := functions["copy"].(func() []string)()
	fmt.Println("copy:", copies)
	println := functions["print"].(func() string)()
	fmt.Println("print:", println)
	printf := functions["print"].(func() string)()
	fmt.Printf("printf : %v\n", printf)
	newPointer := functions["newPtr"].(func() *int)()
	fmt.Println("new Pointer:", &newPointer)
	capacity := functions["capacity"].(func() []int)()
	fmt.Println("capacity:", cap(capacity))
	maxMethod := functions["max"].(func() int)()
	fmt.Println("max :", maxMethod)
	minMethod := functions["min"].(func() int)()
	fmt.Println("min :", minMethod)

	// Rectangle soal 2
	// rectangle := &utils.Rectangle{W: 10, H: 10}
	// results, err := utils.CalculateArea(rectangle)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Printf("lebar: %.2f \ntinggi %.2f \nluas: %v\n", rectangle.W, rectangle.H, results)

	// Searching name soal 3
	// query := string("an")
	// limit := uint64(2)
	// utils.SearchingName(&query, &limit, func(results []string, err error) {
	// 	if err != nil {
	// 		fmt.Println("Error:", err)
	// 		return
	// 	}
	// 	for i, res := range results {
	// 		fmt.Printf("%d. %s\n", i+1, res)
	// 	}
	// })

	// Seleksi nilai soal 4
	// dataArray := []uint{2, 25, 4, 14, 17, 30, 8}
	// nilaiAwal := uint(5)
	// nilaiAkhir := uint(20)
	// result, err := utils.SeleksiNilai(&nilaiAwal, &nilaiAkhir, &dataArray)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Println(result)
}
