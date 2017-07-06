package main

import (
	"fmt"

	"strconv"
	"strings"
)

func main() {
	IntSliceToString([]int{17, 23, 100500})
}
func ReturnInt() int {
	var x int = 1
	fmt.Println(x)
	return x
}

func ReturnFloat() float32 {
	var p float32 = 1.1
	fmt.Println(p)
	return p
}

func ReturnIntArray() [3]int {
	ar := [3]int{1, 3, 4}
	fmt.Println(ar)
	return ar
}

func ReturnIntSlice() []int {
	sli := []int{1, 2, 3}
	fmt.Println(sli)
	return sli
}

func IntSliceToString(sli2 []int) string {

	// Срез int мы преобразуем в строку.
	values := sli2
	valuesText := []string{}

	// Создаем строковый срез с помощью strconv.Itoa.
	// ... Добавляем к нему строки.
	for i := range values {
		number := values[i]
		text := strconv.Itoa(number)
		valuesText = append(valuesText, text)
	}

	// Присоединяем наш срез строки.
	result := strings.Join(valuesText, "")
	fmt.Println(result)

	return result
}

func MergeSlices(sli5 []float32, sli6 []int32) []int {
	zhir := make([]int, 0)
	for i := 0; i < len(sli5); i++ {
		zhir = append(zhir, int(sli5[i]))
	}
	for i := 0; i < len(sli6); i++ {
		zhir = append(zhir, int(sli6[i]))
	}
	fmt.Println(zhir)
	return zhir

}

func GetMapValuesSortedByKey(ma map[int]string) []string {

	var keys2 []int

	for i := 1; i <= 40; i++ {
		for k := range ma {
			if i == k {
				keys2 = append(keys2, k)
			}
		}

	}

	var value []string

	for _, k := range keys2 {
		value = append(value, ma[k])
	}
	fmt.Println(value)
	return value

}
