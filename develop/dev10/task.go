package main

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port
go-telnet mysite.ru 8080
go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/reiver/go-telnet"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	host := flag.String("host", "", "host for connection")
	port := flag.String("port", "", "port for connection")

	flag.Parse()

	var caller telnet.Caller = telnet.StandardCaller

	telnet.DialToAndCall(fmt.Sprintf("%s:%s", *host, *port), caller)

	for {
		select {
		case <-c:
			os.Exit(0)
		case <-time.After(time.Second * 10):
			fmt.Println("timeout!")
			os.Exit(0)
		}
	}

}
