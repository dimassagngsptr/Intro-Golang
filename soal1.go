package main

// import (
// 	"fmt"
// 	"reflect"
// 	"strconv"
// )

// func main() {
// 	//buat variable berdasarkan tipe data
// 	var name string = "123"
// 	var age int = 24
// 	var height float64 = 175.2
// 	var isMeried bool = false
// 	var interestInCoding bool = true

// 	fmt.Println("tipe variable name:", reflect.TypeOf(name), "\nvalue dari name :", name)
// 	fmt.Println("tipe variable age:", reflect.TypeOf(age), "\nvalue dari age :", age)
// 	fmt.Println("tipe variable height:", reflect.TypeOf(height), "\nvalue dari height :", height)
// 	fmt.Println("tipe variable isMeried:", reflect.TypeOf(isMeried), "\nvalue dari isMeried :", isMeried)
// 	fmt.Println("tipe variable interestInCoding:", reflect.TypeOf(interestInCoding), "\nvalue dari interestInCoding :", interestInCoding)

// 	//konversi int => str
// 	str := strconv.Itoa(age)
// 	fmt.Println("nilai string: ",str)

// 	//konversi str => int
// 	num, err := strconv.Atoi(name)
// 	if err != nil {
// 		fmt.Println("Tidak dapat mengkonversi ke integer")
// 	} else {
// 		fmt.Println("nilai integer: ", num)
// 	}

// 	//konversi float => str
// 	strFloat := strconv.FormatFloat(height, 'f', 2, 64)
// 	fmt.Println("Nilai string:", strFloat)

// 	//konversi str => float
// 	floatStr, err := strconv.ParseFloat(name, 64)
// 	if err != nil {
// 		fmt.Println("Tidak dapat mengkonversi ke string")
// 	} else {
// 		fmt.Println("Nilai float:", floatStr)
// 	}

// }
