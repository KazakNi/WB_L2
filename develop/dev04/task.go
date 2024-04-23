package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	res := AnagramSet([]string{"Ласков", "словак", "славок", "сковал", "Марина", "Армани", "ранами", "ранима"})
	fmt.Println(res)
}

func AnagramSet(arr []string) *map[string][]string {
	var res = make(map[string][]string)
	var set = make(map[string]bool)

	for _, val := range arr {
		valToLower := strings.ToLower(val)
		sortedval := []rune(valToLower)

		sort.Slice(sortedval, func(i, j int) bool {
			return sortedval[i] < sortedval[j]
		})

		key := string(sortedval)

		if _, ok := set[valToLower]; !ok {
			set[valToLower] = true
			res[key] = append(res[key], valToLower)
		}

	}

	var ans = make(map[string][]string)

	for _, val := range res {
		if len(val) != 1 {
			ans[val[0]] = val
		}
	}

	return &ans
}
