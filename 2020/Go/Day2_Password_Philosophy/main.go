package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//example  8-11 c: cbccrbdwdccccsk
func main() {
	miFunc()
}

func miFunc() {
	file, err := os.Open("/Users/arpel4/Documents/Go/AdventOfCode/2020/Go/Day2_Password_Philosophy/input2.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var cntOk1 int
	var cntOk2 int

	for scanner.Scan() {
		line := scanner.Text()
		lineArray := strings.Split(line, " ")
		index := strings.Index(lineArray[0], "-")
		min, _ := strconv.Atoi(lineArray[0][0:index])
		max, _ := strconv.Atoi(lineArray[0][index+1:])
		word := string(lineArray[1][0])
		password := lineArray[2]
		part1(password, word, min, max, &cntOk1)
		part2(password, word, min, max, &cntOk2)

	}
	fmt.Println(cntOk1)
	fmt.Println(cntOk2)

}

func part1(p, w string, min, max int, c *int) {
	charCount := make(map[string]int)
	for _, v := range p {
		charCount[string(v)]++
	}
	if v, ok := charCount[w]; ok {
		if v >= min && v <= max {
			*c++
		}
	}
}

func part2(p, w string, min, max int, c *int) {
	charPos1 := string(p[min-1])
	charPos2 := string(p[max-1])
	if (charPos1 == w) != (charPos2 == w) {
		*c++
	}
}
