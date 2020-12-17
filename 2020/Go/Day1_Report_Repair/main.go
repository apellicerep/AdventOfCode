package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	withoutMapPart1() //brute force
	withMapPart1()
	withoutMapPart2() //brute force
	withMapPart2()
}

func withoutMapPart1() {
	expArr := expenses.stringToIntArray("\n")
	numbers := find2sum2020(expArr)
	_ = multiplyItemsArray(numbers)
	//fmt.Println(result)
}

func withoutMapPart2() {
	expArr := expenses.stringToIntArray("\n")
	numbers := find3sum2020(expArr)
	_ = multiplyItemsArray(numbers)
	//fmt.Println(result)
}

func withMapPart1() {
	mapInt := expenses.stringToIntArrayAndMap("\n")
	numbers := find2sum2020WithMap(mapInt)
	_ = multiplyItemsArray(numbers)
	//fmt.Println(result)
}

func withMapPart2() {
	mapInt := expenses.stringToIntArrayAndMap("\n")
	numbers := find3sum2020WithMap(mapInt)
	_ = multiplyItemsArray(numbers)
	//fmt.Println(result)
}

type ExpenseStr string

func (e ExpenseStr) toInt() int {
	if n, err := strconv.Atoi(string(e)); err == nil {
		return n
	} else {
		fmt.Println(err)
		os.Exit(-1)
	}
	return 0
}

func (e ExpenseStr) stringToIntArray(sep string) []int {
	arrayExpenses := strings.Split(string(e), sep)
	arrayInt := make([]int, 0, len(arrayExpenses))
	for _, v := range arrayExpenses {
		num := ExpenseStr(v).toInt()
		arrayInt = append(arrayInt, num)
	}
	return arrayInt
}

func (e ExpenseStr) stringToIntArrayAndMap(sep string) map[int]struct{} {
	arrayExpenses := strings.Split(string(e), sep)
	mapInt := make(map[int]struct{})
	for _, v := range arrayExpenses {
		num := ExpenseStr(v).toInt()
		mapInt[num] = struct{}{}
	}
	return mapInt
}

func find2sum2020(expArr []int) []int {
	var sum2020 []int
out:
	for _, i := range expArr {
		for _, j := range expArr {
			if j != i {
				if i+j == 2020 {
					sum2020 = append(sum2020, i, j)
					break out
				}
			}
		}
	}
	return sum2020
}

func find2sum2020WithMap(expMap map[int]struct{}) []int {
	var sum2020 []int
	for k := range expMap {
		rest := 2020 - k
		if _, ok := expMap[rest]; ok {
			sum2020 = append(sum2020, k, rest)
			break
		}
	}
	return sum2020
}

func find3sum2020(expArr []int) []int {
	var sum2020 []int
out:
	for _, i := range expArr {
		for _, j := range expArr {
			if j != i {
				if i+j < 2020 {
					for _, k := range expArr {
						if k != j && k != i {
							if i+j+k == 2020 {
								sum2020 = append(sum2020, i, j, k)
								break out
							}
						}
					}
				}
			}
		}
	}
	return sum2020
}

func find3sum2020WithMap(expMap map[int]struct{}) []int {
	var sum2020 []int
out:
	for x := range expMap {
		for y := range expMap {
			if y != x {
				rest := 2020 - x - y
				if _, ok := expMap[rest]; ok {
					sum2020 = append(sum2020, x, y, rest)
					break out
				}
			}
		}
	}
	return sum2020
}

func multiplyItemsArray(array []int) int {
	result := 1
	for _, v := range array {
		result *= v
	}
	return result
}
