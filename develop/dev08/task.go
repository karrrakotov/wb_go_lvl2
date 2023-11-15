package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("$ ") // Приглашение шелла

		if !scanner.Scan() {
			break // Прекращаем выполнение при ошибке ввода
		}

		line := scanner.Text()
		args := strings.Fields(line)

		if len(args) == 0 {
			continue // Пропускаем пустые строки
		}

		switch args[0] {
		case "cd":
			if len(args) < 2 {
				fmt.Println("Используйте: cd <directory>")
				continue
			}
			err := os.Chdir(args[1])
			if err != nil {
				fmt.Println("Ошибка при смене каталога:", err)
			}
		case "pwd":
			cwd, err := os.Getwd()
			if err != nil {
				fmt.Println("Ошибка при получении текущего каталога:", err)
			}
			fmt.Println(cwd)
		case "echo":
			fmt.Println(strings.Join(args[1:], " "))
		case "kill":
			if len(args) < 2 {
				fmt.Println("Используйте: kill <pid>")
				continue
			}
			pid := args[1]
			err := exec.Command("kill", pid).Run()
			if err != nil {
				fmt.Println("Ошибка при убийстве процесса:", err)
			}
		case "ps":
			out, err := exec.Command("ps").Output()
			if err != nil {
				fmt.Println("Ошибка при выполнении команды ps:", err)
			}
			fmt.Println(string(out))
		default:
			cmd := exec.Command(args[0], args[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			err := cmd.Run()
			if err != nil {
				fmt.Println("Ошибка при выполнении команды:", err)
			}
		}
	}
}
