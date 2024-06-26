Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v, ok := <-a:
			if ok{
				c <- v
			} else {
				break
			}

			case v, ok := <-b:
				if ok{
				c <- v
			} else {
				break
			}
		}
	}()
	return c
}
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
1
2
3
4
5
6
7
8
0
0
...
```

Так как в функции `merge` в `select` нет проверки на то, возвращается или нет стандартное значение из канала, после закрытия входного канала (`a` и `b`) в выходной канал будут отправляться нули (стандартное значение для `int`). 

В go нет способа проверить закрыт ли канал не вычитывая из него значения, поэтому можно применить другой подход для решения этой задачи:

```go

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		var signalA, signalB bool
		for {
			select {
			case v, ok := <-a:
				if ok {
					c <- v
				} else {
					signalA = true
				}

			case v, ok := <-b:
				if ok {
					c <- v
				} else {
					signalB = true
				}

			}
			if signalA && signalB {
				close(c)
				break
			}
		}
	}()
	return c
}



```