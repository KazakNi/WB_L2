package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var ErrInvalidString = errors.New("invalid sequence")

func main() {
	res, err := UnpackString("a4bc2d5e")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

func UnpackString(s string) (string, error) {
	var seenDigit bool
	str := []rune(s)
	var stack []string
	var res []string
	for i, val := range str {
		if unicode.IsDigit(val) {

			if seenDigit || len(stack) == 0 {
				return "", ErrInvalidString
			}
			seenDigit = true

			symbol := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			num, _ := strconv.Atoi(string(val))
			for range num {
				res = append(res, symbol)
			}

		} else {
			seenDigit = false
			if i < len(s)-1 && unicode.IsDigit(str[i+1]) {
				stack = append(stack, string(val))
			} else {
				res = append(res, string(val))
			}

		}

	}
	return strings.Join(res, ""), nil
}
