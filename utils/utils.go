package utils

import (
	"fmt"
	"regexp"
	"strconv"
)

var dotPattern = regexp.MustCompile(`\.`)

// CheckError prints an error, if it exists
func CheckError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

// IntFromHumanRedableNumber "3.698" -> 3978
func IntFromHumanRedableNumber(humanNumberString string) int {
	numberStringWithNoDots := dotPattern.ReplaceAllString(humanNumberString, "")
	number, err := strconv.Atoi(numberStringWithNoDots)
	CheckError(err)
	return number
}
