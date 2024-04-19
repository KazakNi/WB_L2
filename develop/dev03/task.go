package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

# Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

# Дополнительное

# Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
type Params struct {
	k int
	n bool
	r bool
	u bool
}

func main() {
	file := flag.String("file", "example.txt", "file for parsing")
	k := flag.Int("k", -1, "column sorting")
	n := flag.Bool("n", false, "numeric sorting")
	r := flag.Bool("r", false, "reverse sorting")
	u := flag.Bool("u", false, "unique strings output")

	flag.Parse()

	params := Params{
		k: *k,
		n: *n,
		r: *r,
		u: *u,
	}

	result := SortMan(*file, params)

	fmt.Println(result)

}

func SortMan(file string, params Params) [][]string {

	strings, err := readFile(file, params.u)

	if err != nil {
		log.Fatal(err)
	}
	// n flag with variations
	if params.n {

		if params.k != -1 && params.r {

			if err := validateCol(strings, params.k); err != nil {
				log.Fatal(err)
			}
			sort.Slice(strings, func(i, j int) bool {
				val1, err := strconv.Atoi(strings[i][params.k])
				if err != nil {
					log.Fatal(err)
				}
				val2, err := strconv.Atoi(strings[j][params.k])
				if err != nil {
					log.Fatal(err)
				}
				return val1 > val2
			})
		} else if params.k != -1 {
			sort.Slice(strings, func(i, j int) bool {
				val1, err := strconv.Atoi(strings[i][params.k])
				if err != nil {
					log.Fatal(err)
				}
				val2, err := strconv.Atoi(strings[j][params.k])
				if err != nil {
					log.Fatal(err)
				}
				return val1 < val2
			})
			return strings
		}
		colNum := -1
		for _, strSlice := range strings {
			// search for numeric column
			for col, str := range strSlice {
				if _, err := strconv.Atoi(str); err == nil {
					colNum = col
					break
				}
			}
		}

		if colNum == -1 {
			return strings
		}

		if params.r {
			sort.Slice(strings, func(i, j int) bool {
				val1, _ := strconv.Atoi(strings[i][colNum])
				val2, _ := strconv.Atoi(strings[j][colNum])
				return val1 > val2
			})
			return strings
		}

		sort.Slice(strings, func(i, j int) bool {
			val1, _ := strconv.Atoi(strings[i][colNum])
			val2, _ := strconv.Atoi(strings[j][colNum])
			return val1 < val2
		})
		return strings
	}

	if params.k != -1 {
		if err := validateCol(strings, params.k); err != nil {
			log.Fatal(err)
		}
		if params.r {
			sort.Slice(strings, func(i, j int) bool {

				if len(strings[i]) == 0 || len(strings[j]) == 0 {
					return len(strings[i]) != 0
				}
				return strings[i][params.k] > strings[j][params.k]
			})
			return strings
		} else {

			sort.Slice(strings, func(i, j int) bool {

				if len(strings[i]) == 0 || len(strings[j]) == 0 {
					return len(strings[i]) != 0
				}
				return strings[i][params.k] < strings[j][params.k]
			})
			return strings
		}
	}

	if params.r {
		sort.Slice(strings, func(i, j int) bool {

			if len(strings[i]) == 0 || len(strings[j]) == 0 {
				return len(strings[i]) != 0
			}
			return strings[i][0] > strings[j][0]
		})

		return strings

	}
	sort.Slice(strings, func(i, j int) bool {

		if len(strings[i]) == 0 || len(strings[j]) == 0 {
			return len(strings[i]) != 0
		}
		return strings[i][0] < strings[j][0]
	})

	return strings
}

func readFile(file string, u bool) ([][]string, error) {
	var res [][]string

	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	if u {
		var d = make(map[string]bool)
		for scanner.Scan() {

			line := scanner.Text()
			if _, ok := d[line]; !ok {
				res = append(res, strings.Split(line, " "))
				d[line] = true
			}
		}
	} else {

		for scanner.Scan() {
			res = append(res, strings.Split(scanner.Text(), " "))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return [][]string{}, err
	}

	return res, nil
}

func validateCol(strings [][]string, col int) error {
	for _, slices := range strings {
		for _, slice := range slices {
			if len(slice)-1 <= col {
				return ErrColNum
			}
		}
	}
	return nil
}

var ErrColNum = errors.New("column exceeds len of slice")
