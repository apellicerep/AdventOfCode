package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	miFunc()
}

func miFunc() {
	file, err := os.Open("/Users/arpel4/Documents/Go/AdventOfCode/2020/Go/Day4_Passport_Processing/input4.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer file.Close()

	numFields := 0
	validPassports := 0
	var fields = make(map[string]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if ok := validNumFields(&numFields, fields); ok {
				////////////--------part1--------/////////////
				//validPassports++
				////////////--------part1--------/////////////

				////////////--------part2--------/////////////
				if ok := validateValueFields(fields, listFunctions); ok {
					validPassports++
				}
				////////////--------part2--------/////////////
			}
			numFields = 0
			fields = make(map[string]string)
		} else {
			createMapFields(line, fields, &numFields)
		}
	}
	//last row
	if ok := validNumFields(&numFields, fields); ok {
		////////////--------part1--------/////////////
		//validPassports++
		////////////--------part1--------/////////////

		////////////--------part2--------/////////////
		if ok := validateValueFields(fields, listFunctions); ok {
			validPassports++
		}
		////////////--------part2--------/////////////
	}
	fmt.Println(validPassports)
}

func createMapFields(l string, f map[string]string, nF *int) {
	splitLine := strings.Split(l, " ")
	for _, v := range splitLine {
		i := strings.Index(v, ":")
		key := v[:i]
		value := v[i+1:]
		f[key] = value
		*nF++
	}
}

func validNumFields(nF *int, f map[string]string) bool {
	switch {
	case *nF == 8:
		return true
	case *nF >= 7:
		if _, exist := f["cid"]; !exist {
			return true
		}
	}
	return false
}

func validateValueFields(fields map[string]string, listFunctions map[string]func(string) bool) bool {
	for k, v := range fields {
		if _, exist := listFunctions[k]; exist {
			if ok := listFunctions[k](v); !ok {
				return false
			}
		} else {
			fmt.Printf("function: %s - not implemented", k)
			return false
		}
	}
	return true
}

var listFunctions = map[string]func(string) bool{
	"byr": validateBry,
	"ecl": validateEcl,
	"eyr": validateEyr,
	"hcl": validateHcl,
	"hgt": validateHgt,
	"iyr": validateIyr,
	"pid": validatePid,
	"cid": validateCid,
}

var validateBry = func(value string) bool {
	valFields := valBirth{4, 1920, 2002}
	if len(value) == valFields.digits {
		if num, err := stringToInt(value); err == nil {
			if num >= valFields.min && num <= valFields.max {
				return true
			}
		}
	}
	return false
}
var validateIyr = func(value string) bool {
	valFields := valIssue{4, 2010, 2020}
	if len(value) == valFields.digits {
		if num, err := stringToInt(value); err == nil {
			if num >= valFields.min && num <= valFields.max {
				return true
			}
		}
	}
	return false
}

var validateEyr = func(value string) bool {
	valFields := valExp{4, 2020, 2030}
	if len(value) == valFields.digits {
		if num, err := stringToInt(value); err == nil {
			if num >= valFields.min && num <= valFields.max {
				return true
			}
		}
	}
	return false
}

var validateHgt = func(value string) bool {
	valFields := valHeight{[]heightOpt{{"cm", 150, 193}, {"in", 59, 76}}}
	for _, v := range valFields.heightOpts {
		if value[len(value)-2:] != v.prefix {
			continue
		}
		if num, err := stringToInt(value[:len(value)-2]); err == nil {
			if num >= v.min && num <= v.max {
				return true
			}
		}
	}
	return false
}

var validateHcl = func(value string) bool {
	valFields := valHairColor{"#", "0-9", "a-f"}
	diMin := valFields.digits[:1]
	diMax := valFields.digits[2:]
	leMin := valFields.letters[:1]
	leMax := valFields.letters[2:]

	if value[:1] == valFields.prefix {
		for _, v := range value[1:] {
			if (string(v) >= diMin && string(v) <= diMax) ||
				(string(v) >= leMin && string(v) <= leMax) {
				continue
			} else {
				return false
			}
		}
		return true
	}
	return false
}

var validateEcl = func(value string) bool {
	valFields := valEyeColor{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	if ok := contains(valFields, value); ok {
		return true
	}
	return false
}
var validatePid = func(value string) bool {
	valFields := valPassportId(9)
	if len(value) == int(valFields) {
		return true
	}
	return false
}

var validateCid = func(value string) bool {
	return true
}

func stringToInt(s string) (int, error) {
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
