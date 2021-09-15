package main

import (
	"fmt"
	"strings"
)

func main() {
	// str := strings.Fields("My name is sunil")
	// fmt.Println(str)
	// var newStr strings.Builder
	// for i := len(str)-1; i > -1; i-- {
	// 	newStr.WriteString(str[i])
	// 	if i != 0 {
	// 		newStr.WriteString(" ")
	// 	}
	// }

	fmt.Println(getReversedStringFields("My Name is Sunil"))
	fmt.Println(getReversedStringFieldsByByte("My Name is Sunil"))
}

func getReversedStringFields(str string) string {
	r := []rune(str)
	
	var newStrs []string
	var newStr strings.Builder
	for i, c := range r {
		if c != ' ' {
			newStr.WriteRune(c)
		} else {
			newStrs = append(newStrs, newStr.String())
			newStr = strings.Builder{}
		}
		if i == len(r) -1 {
			newStrs = append(newStrs, newStr.String())
		}
	}
	fmt.Println(newStrs)
	var revStr strings.Builder
	for i := len(newStrs) -1; i > -1; i-- {
		revStr.WriteString(newStrs[i])
		if i != 0 {
			revStr.WriteString(" ")
		}
	}
	return revStr.String()
}

func getReversedStringFieldsByByte(str string) string {
	strBytes := []byte(str)
	var strSlice []string
	var newStr strings.Builder
	for i, sb := range strBytes {
		if sb != ' ' {
			newStr.WriteByte(sb)
		} else {
			strSlice = append(strSlice, newStr.String())
			newStr = strings.Builder{}
		}
		if i == len(strBytes) -1 {
			strSlice = append(strSlice, newStr.String())
		}
	}

	var revStr strings.Builder
	for i := len(strSlice) - 1; i > -1; i-- {
		revStr.WriteString(strSlice[i])
		if i != 0 {
			revStr.WriteString(" ")
		}
	}
	return revStr.String()
}