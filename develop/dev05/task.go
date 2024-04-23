package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/icza/backscanner"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type Flags struct {
	A int
	B int
	C int
	c bool
	i bool
	v bool
	F bool
	n bool
}

func main() {
	A := flag.Int("A", 0, `"after" печатать +N строк после совпадения"`)
	B := flag.Int("B", 0, `"before" печатать +N строк до совпадения`)
	C := flag.Int("С", 0, `"context" (A+B) печатать ±N строк вокруг совпадения`)
	cnt := flag.Bool("c", false, `"count" (количество строк)`)
	i := flag.Bool("i", false, `"ignore-case" (игнорировать регистр)`)
	v := flag.Bool("v", false, `"invert" (вместо совпадения, исключать)`)
	F := flag.Bool("F", false, `"fixed", точное совпадение с строкой, не паттерн`)
	n := flag.Bool("n", false, `"line num", печатать номер строки`)
	file := flag.String("file", "test.txt", `path to file`)
	pattern := flag.String("pattern", ".+", `regexp`)

	flag.Parse()

	grep(*file, *pattern, Flags{A: *A, B: *B, C: *C, c: *cnt, i: *i, v: *v, F: *F, n: *n})
}

func grep(file string, pattern string, flags Flags) [][]string {
	var res [][]string
	re := regexp.MustCompile(pattern)

	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	if flags.A != 0 {
		row := 0
		for scanner.Scan() {
			line := strings.Split(scanner.Text(), " ")

			for _, word := range line {
				if row != 0 && row < flags.A {
					res = append(res, line)
					row++
				}
				if re.MatchString(word) {
					row = 1
					res = append(res, line)
					break
				}
				if row == flags.A {
					row = 0
				}
			}

		}
		return res
	}

	if flags.B != 0 {

		fileStatus, err := f.Stat()
		if err != nil {
			panic(err)
		}

		scanner := backscanner.New(f, int(fileStatus.Size()))
		row := 0
		for {
			line, _, err := scanner.LineBytes()
			if err == io.EOF {
				break
			}
			for _, word := range string(line) {
				if row != 0 && row < flags.A {
					res = append(res, strings.Split(string(line), " "))
					row++
				}
				if re.MatchString(string(word)) {
					row = 1
					res = append(res, strings.Split(string(line), " "))
					break
				}
				if row == flags.B {
					row = 0
				}
			}

		}
		return *reverseSliceOfStringSlices(res)
	}

	if flags.C != 0 {
		ans := [][]string{}
		// пройтись по всем строкам и отметить строки с совпадениями
		// преобразовать номера строк в интервалы строк
		// добавить в результат новым проходом интересующие строки
		for scanner.Scan() {
			line := strings.Split(scanner.Text(), " ")
			lineNums := []int{}

			for i, word := range line {
				res = append(res, line)
				if re.MatchString(string(word)) {
					lineNums = append(lineNums, i)
				}
			}
			var prev, next, x1, x2 int
			for i, val := range lineNums {

				if val-flags.C >= 0 {

					prev = val - flags.C

				} else {

					prev, x1 = 0, 0
				}
				if val+flags.C >= len(res) {
					next = len(res) - 1
					ans = append(ans, res[x1:next]...)
					return ans
				} else {

					if prev <= next && i != 0 {
						next = val + flags.C
						x2 = next
					} else {
						ans = append(ans, res[x1:x2]...)
						next = val + flags.C

						x1, x2 = prev, next

					}

				}

			}
		}

	}

	if flags.c {
		ans, _ := countRows(*f)
		fmt.Println(ans)
	}
	if flags.i {
		re := regexp.MustCompile(strings.ToLower(pattern))
		for scanner.Scan() {
			line := strings.Split(scanner.Text(), " ")

			for _, word := range line {
				if re.MatchString(strings.ToLower(word)) {
					res = append(res, line)
					break
				}
			}
		}
		return res
	}

	if flags.v {
		for scanner.Scan() {
			line := strings.Split(scanner.Text(), " ")

			for _, word := range line {
				if re.MatchString(strings.ToLower(word)) {
					break
				} else {
					res = append(res, line)
				}
			}
		}
	}

	if flags.F {
		for scanner.Scan() {
			line := scanner.Text()

			if line == pattern {
				res = append(res, strings.Split(line, " "))
			}
		}
	}

	if flags.n {
		row := 0
		for scanner.Scan() {
			row++
			r := strconv.Itoa(row)
			line := strings.Split(r+" "+scanner.Text(), " ")

			for _, word := range line {
				if re.MatchString(word) {
					res = append(res, line)
					break
				}
			}
		}

		return res

	}
	for scanner.Scan() {

		line := strings.Split(scanner.Text(), " ")

		for _, word := range line {
			if re.MatchString(word) {
				res = append(res, line)
				break
			}
		}
	}
	return res
}

func countRows(f os.File) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := f.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}

}

func reverseSliceOfStringSlices(slice [][]string) *[][]string {
	newSlice := make([][]string, len(slice))

	for i := len(slice) - 1; i > 0; i-- {
		newSlice[len(slice)-i-1] = slice[i]
	}
	return &newSlice
}
