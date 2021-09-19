package strings

import (
	"regexp"
    "strings"
	"strconv"
	"fmt"
)

// FindNum contribute数を文字列から取り出す
func FindNum(strVal string) int64 {
    rex := regexp.MustCompile("[0-9]+")
    strVal = rex.FindString(strings.ReplaceAll(strVal, ",", ""))
    intVal, err := strconv.ParseInt(strVal, 10, 64)
    if err != nil {
        fmt.Println(err)
    }
    return intVal
}