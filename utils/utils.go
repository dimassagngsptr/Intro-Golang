package utils

import (
	"fmt"
	"sort"
	"strings"
)

type Rectangle struct {
	W, H float64
}
type Callback func([]string, error)

func Method() map[string]interface{} {
	functions := make(map[string]interface{})
	/* append => digunakan untuk menambahkan elemen pada slice. Elemen baru tersebut diposisikan setelah indeks
	paling akhir.*/
	functions["slice"] = func() []int {
		s := []int{1, 2, 3}
		s = append(s, 4, 5)
		return s
	}
	// len => digunakan untuk menghitung jumlah elemen slice yang ada
	functions["length"] = func() []int {
		s := []int{1, 2, 3, 4, 5}
		return s
	}

	functions["exMake"] = func() map[string]string {
		m := make(map[string]string)
		m["name"] = "Dimas"
		return m
	}

	functions["copy"] = func() []string {
		slices1 := []string{"Dimas", "Bandung"}
		slices2 := make([]string, len(slices1))
		copy(slices2, slices1)
		return slices2
	}

	functions["print"] = func() string {
		name := "Dimas"
		return name
	}

	functions["newPtr"] = func() *int {
		ptr := new(int)
		*ptr = 10
		return ptr
	}
	functions["capacity"] = func() []int {
		arr := make([]int, 5)
		return arr
	}
	functions["max"] = func() int {
		res := max(10, 5)
		return res
	}
	functions["min"] = func() int {
		res := min(10, 5)
		return res
	}

	return functions
}

func CalculateArea(val *Rectangle) (*float64, error) {

	if val.H < 0 || val.W < 0 {
		err := fmt.Errorf("nilai tidak boleh minus")
		return nil, err

	}
	result := val.W * val.H
	return &result, nil
}

func SearchingName(query *string, limit *uint64, callback Callback) {
	var results []string
	var listNames = []string{
		"Abigail", "Alexandra", "Alison",
		"Amanda", "Angela", "Bella",
		"Carol", "Caroline", "Carolyn",
		"Deirdre", "Diana", "Elizabeth",
		"Ella", "Faith", "Olivia", "Penelope",
	}
	for _, names := range listNames {
		if strings.Contains(names, *query) {
			results = append(results, names)
			if len(results) == int(*limit) {
				break
			}
		}
	}

	if len(results) > 0 {
		callback(results, nil)
	} else {
		err := fmt.Errorf("nama tidak ditemukan")
		callback(nil, err)
	}
}

func SeleksiNilai(nilaiAwal, nilaiAkhir *uint, dataArray *[]uint) (interface{}, error) {
	var results []int
	if *nilaiAwal > *nilaiAkhir {
		err := fmt.Errorf("nilai awal tidak boleh lebih besar dari nilai akhir")
		return nil, err
	}
	if len(*dataArray) < 5 {
		err := fmt.Errorf("data array harus lebih atau sama dengan 5")
		return nil, err
	}
	for _, i := range *dataArray {
		if i >= *nilaiAwal && i <= *nilaiAkhir {
			results = append(results, int(i))
		}
	}
	if len(results) < 1 {
		err := fmt.Errorf("nama tidak ditemukan")
		return nil, err
	}
	sort.Ints(results)
	return results, nil
}
