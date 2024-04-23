package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"regexp"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	f := flag.Int("f", 0, "col")
	d := flag.String("d", " ", "delim")
	s := flag.String("s", "\t", "separated")

	flag.Parse()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	reader := bufio.NewReader(os.Stdin)

	r := regexp.MustCompile(fmt.Sprintf(`.+%s.+`, *s))
	for {
		select {
		case <-c:
			fmt.Println("Exit the programm")
			return
		default:
			var s string
			fmt.Fscanln(reader, &s)

			if r.MatchString(s) {
				fmt.Println(strings.Split(s, *d)[*f])
			}

		}
	}
}
