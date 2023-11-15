package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/
type telnetOptions struct {
	Timeout time.Duration
}

func main() {
	options := parseCommandLineArguments()

	if flag.NArg() != 2 {
		fmt.Println("Используйте: go-telnet [--timeout=10s] host port")
		os.Exit(1)
	}

	host := flag.Arg(0)
	port := flag.Arg(1)

	address := fmt.Sprintf("%s:%s", host, port)

	conn, err := net.DialTimeout("tcp", address, options.Timeout)
	if err != nil {
		fmt.Printf("Ошибка соединения: %s: %v\n", address, err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Подключено к", address)

	go readFromServer(conn)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Fprintf(conn, "%s\n", line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при чтении данных: %v\n", err)
		os.Exit(1)
	}
}

func readFromServer(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	fmt.Println("Соединение было прервано")
	os.Exit(0)
}

// parseCommandLineArguments - функция для парсинга параметров командной строки
func parseCommandLineArguments() telnetOptions {
	options := telnetOptions{}

	timeout := flag.Duration("timeout", 10*time.Second, "Connection timeout")
	flag.Parse()

	options.Timeout = *timeout

	return options
}
