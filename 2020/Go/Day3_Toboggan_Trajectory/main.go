package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Slope struct {
	Right int
	Down  int
	Index int
	X     int
}

func main() {
	part1()
	part2()
}

func part1() {

	var X int
	var index = 3

	file, err := os.Open("/Users/arpel4/Documents/Go/AdventOfCode/2020/Go/Day3_Toboggan_Trajectory/input3.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	for scanner.Scan() {
		line := scanner.Text()
		miFunc1(line, &X, &index)
	}
	fmt.Println(X)
}

func miFunc1(line string, X *int, index *int) {
	if string(line[*index]) == "#" {
		*X++
	}
	*index += 3
	switch {
	case *index > len(line):
		*index = *index - len(line)
	case *index == len(line):
		*index = 0
	}
}

func part2() {

	var slopes = []*Slope{
		{1, 1, 1, 0},
		{3, 1, 3, 0},
		{5, 1, 5, 0},
		{7, 1, 7, 0},
		{1, 2, 1, 0}}

	file, err := os.Open("/Users/arpel4/Documents/Go/AdventOfCode/2020/Go/Day3_Toboggan_Trajectory/input3.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	var par int
	for scanner.Scan() {
		par += 1
		even := isEven(par)
		line := scanner.Text()
		miFunc2(line, even, slopes)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	result := multiplyResult(slopes)
	fmt.Println(result)
}

func miFunc2(line1 string, even bool, slopes []*Slope) {
	for _, v := range slopes {
		if (v.Down == 1) || (v.Down == 2 && even) {
			if string(line1[v.Index]) == "#" {
				v.X++
			}
			v.Index += v.Right
			switch {
			case v.Index > len(line1):
				v.Index = v.Index - len(line1)
			case v.Index == len(line1):
				v.Index = 0
			}
		}
	}
}

func multiplyResult(slopes []*Slope) int {
	var result = 1
	for _, v := range slopes {
		fmt.Println(v.X)
		result *= v.X
	}
	return result
}

func isEven(n int) bool {
	even := false
	if n%2 == 0 {
		even = true
	}
	return even
}
