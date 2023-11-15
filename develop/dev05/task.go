package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
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

// grepOptions - структура для хранения параметров утилиты grep
type grepOptions struct {
	After      int
	Before     int
	Context    int
	Count      bool
	IgnoreCase bool
	Invert     bool
	Fixed      bool
	LineNum    bool
	Pattern    string
	FilePaths  []string
}

func main() {
	options := parseCommandLineArguments()

	if len(options.FilePaths) == 0 {
		fmt.Println("Пожалуйста, укажите путь(и) к входному файлу(ам).")
		return
	}

	for _, filePath := range options.FilePaths {
		lines, err := readLinesFromFile(filePath)
		if err != nil {
			fmt.Printf("Ошибка при чтении файла %s: %v\n", filePath, err)
			return
		}

		matchedLines := grep(lines, options)
		printResults(matchedLines, options)
	}
}

// parseCommandLineArguments - функция для парсинга параметров командной строки
func parseCommandLineArguments() grepOptions {
	options := grepOptions{}

	flag.IntVar(&options.After, "A", 0, "Печатать +N строк после совпадения")
	flag.IntVar(&options.Before, "B", 0, "Печатать +N строк до совпадения")
	flag.IntVar(&options.Context, "C", 0, "Печатать ±N строк вокруг совпадения")
	flag.BoolVar(&options.Count, "c", false, "Количество строк")
	flag.BoolVar(&options.IgnoreCase, "i", false, "Игнорировать регистр")
	flag.BoolVar(&options.Invert, "v", false, "Вместо совпадения, исключать")
	flag.BoolVar(&options.Fixed, "F", false, "Точное совпадение со строкой, не паттерн")
	flag.BoolVar(&options.LineNum, "n", false, "Печатать номера строк")

	flag.Parse()

	// Оставляем только пути к файлам после обработки флагов
	options.FilePaths = flag.Args()

	// Первый аргумент после обработки флагов - шаблон поиска
	if len(options.FilePaths) > 0 {
		options.Pattern = options.FilePaths[0]
		options.FilePaths = options.FilePaths[1:]
	}

	return options
}

// readLinesFromFile - функция для чтения данных из файла
func readLinesFromFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

// grep - функция для фильтрации строк по шаблону
func grep(lines []string, options grepOptions) []string {
	var result []string
	pattern := options.Pattern

	// Игнорирование регистра при поиске, если установлен соответствующий флаг
	if options.IgnoreCase {
		pattern = "(?i)" + pattern
	}

	// Фиксированное совпадение строки, если установлен соответствующий флаг
	if options.Fixed {
		pattern = regexp.QuoteMeta(pattern)
	}

	// Создание регулярного выражения для поиска
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Printf("Ошибка в регулярном выражении: %v\n", err)
		return result
	}

	for i, line := range lines {
		matched := re.MatchString(line)

		// Инвертированный поиск, исключаем совпадающие строки
		if options.Invert {
			matched = !matched
		}

		if matched {
			result = append(result, line)

			// Печать N строк после совпадения
			if options.After > 0 && i+options.After < len(lines) {
				result = append(result, lines[i+1:i+1+options.After]...)
			}

			// Печать N строк до совпадения
			if options.Before > 0 && i-options.Before >= 0 {
				result = append(result, lines[i-options.Before:i]...)
			}

			// Печать ±N строк вокруг совпадения
			if options.Context > 0 && i-options.Context >= 0 && i+options.Context < len(lines) {
				result = append(result, lines[i-options.Context:i]...)
				result = append(result, lines[i+1:i+1+options.Context]...)
			}
		}
	}

	return result
}

// printResults - функция для печати результатов поиска
func printResults(lines []string, options grepOptions) {
	if options.Count {
		fmt.Printf("Количество совпадений: %d\n", len(lines))
	} else {
		for i, line := range lines {
			// Печать номера строки
			if options.LineNum {
				fmt.Printf("%d: ", i+1)
			}

			fmt.Println(line)
		}
	}
}
