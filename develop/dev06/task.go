package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
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

// cutOptions - структура для хранения параметров утилиты cut
type cutOptions struct {
	Fields    string
	Delimiter string
	Separated bool
}

func main() {
	options := parseCommandLineArguments()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		if options.Separated && !strings.Contains(line, options.Delimiter) {
			continue
		}
		fields := strings.Split(line, options.Delimiter)
		outputFields := getOutputFields(fields, options.Fields)
		fmt.Println(strings.Join(outputFields, options.Delimiter))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при чтении ввода: %v\n", err)
		os.Exit(1)
	}
}

// parseCommandLineArguments - функция для парсинга параметров командной строки
func parseCommandLineArguments() cutOptions {
	options := cutOptions{}

	flag.StringVar(&options.Fields, "f", "", "Выбрать поля (колонки)")
	flag.StringVar(&options.Delimiter, "d", "\t", "Использовать другой разделитель")
	flag.BoolVar(&options.Separated, "s", false, "Только строки с разделителем")

	flag.Parse()

	return options
}

// getOutputFields - функция для выбора полей (колонок) из строки
func getOutputFields(fields []string, fieldNumbers string) []string {
	if fieldNumbers == "" {
		return fields
	}

	var outputFields []string
	selectedFields := strings.Split(fieldNumbers, ",")

	for _, field := range selectedFields {
		index := parseFieldIndex(field, len(fields))
		if index != -1 {
			outputFields = append(outputFields, fields[index])
		}
	}

	return outputFields
}

// parseFieldIndex - функция для парсинга индекса поля
func parseFieldIndex(field string, maxIndex int) int {
	index := parsePositiveInt(field)
	if index == 0 || index > maxIndex {
		return -1
	}
	return index - 1
}

func parsePositiveInt(s string) int {
	num := 0
	fmt.Sscanf(s, "%d", &num)
	if num <= 0 {
		return 0
	}
	return num
}
