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
	//part1()
	part2()

}

func part1() {
	file, err := os.Open("/Users/arpel4/Documents/Go/AdventOfCode/2020/Go/Day2_Password_Philosophy/input2.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var cntOk int

	for scanner.Scan() {
		line := scanner.Text()
		lineArray := strings.Split(line, " ")
		index := strings.Index(lineArray[0], "-")
		min, _ := strconv.Atoi(lineArray[0][0:index])
		max, _ := strconv.Atoi(lineArray[0][index+1:])
		word := string(lineArray[1][0])
		password := lineArray[2]
		charCount := make(map[string]int)
		for _, v := range password {
			charCount[string(v)]++
		}
		if v, ok := charCount[word]; ok {
			if v >= min && v <= max {
				cntOk++
			}
		}
	}
	fmt.Println(cntOk)

}

func part2() {
	file, err := os.Open("/Users/arpel4/Documents/Go/AdventOfCode/2020/Go/Day2_Password_Philosophy/input2.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var cntOk int

	for scanner.Scan() {
		line := scanner.Text()
		lineArray := strings.Split(line, " ")
		index := strings.Index(lineArray[0], "-")
		pos1, _ := strconv.Atoi(lineArray[0][0:index])
		pos2, _ := strconv.Atoi(lineArray[0][index+1:])
		word := string(lineArray[1][0])
		password := lineArray[2]
		charPos1 := string(password[pos1-1])
		charPos2 := string(password[pos2-1])
		if (charPos1 == word) != (charPos2 == word) {
			cntOk++
		}
	}
	fmt.Println(cntOk)
}
