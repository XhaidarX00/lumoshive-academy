package utils

import (
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"runtime"
	"strings"
	"unicode"
)

func Capitalize(s string) string {
	s = strings.ToLower(s)

	capitalizeNext := true
	result := []rune(s)

	for i, r := range result {
		if capitalizeNext && unicode.IsLetter(r) {
			result[i] = unicode.ToUpper(r)
			capitalizeNext = false
		} else if unicode.IsSpace(r) {
			capitalizeNext = true
		}
	}

	return string(result)
}

func ClearScreen() {
	if runtime.GOOS == "windows" {
		module := exec.Command("cmd", "/c", "cls")
		module.Stdout = os.Stdout
		module.Run()
	} else {
		module := exec.Command("clear")
		module.Stdout = os.Stdout
		module.Run()
	}
}

func IsLenVar(input any, len_ int) bool {
	lenVal := reflect.ValueOf(input)

	switch lenVal.Kind() {
	case reflect.String, reflect.Array, reflect.Slice, reflect.Map:
		if lenVal.Len() >= len_ {
			return true
		}
	default:
		ErrorMessage(fmt.Sprintf("tipe data %s tidak mendukung operasi Len()\n", lenVal.Kind()))
	}

	return false
}
