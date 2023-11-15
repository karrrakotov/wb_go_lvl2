package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
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

type SortOptions struct {
	ColumnIndex   int
	Numeric       bool
	Reverse       bool
	Unique        bool
	Month         bool
	IgnoreBlanks  bool
	CheckSorted   bool
	NumericSuffix bool
}

func main() {
	filePath := flag.String("file", "", "Path to the input file")
	options := parseCommandLineArguments()
	flag.Parse()

	if *filePath == "" {
		fmt.Println("Пожалуйста, укажите путь к входному файлу, используя флаг -file")
		return
	}

	lines, err := readLinesFromFile(*filePath)
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}

	copyLines := make([]string, len(lines))
	copy(copyLines, lines)

	if options.Unique {
		lines = removeDuplicates(lines)
	}

	sort.Sort(sortLines{lines, options})

	if options.CheckSorted && isSorted(copyLines, options) {
		fmt.Println("Данные отсортированны")
		return
	} else if options.CheckSorted && !isSorted(copyLines, options) {
		fmt.Println("Данные не отсортированны")
		return
	}

	outputFilePath := "sorted_" + *filePath
	err = writeLinesToFile(outputFilePath, lines)
	if err != nil {
		fmt.Println("Ошибка при записи файла: ", err)
		return
	}

	fmt.Println("Сортировка завершена. Новые данные записаны в файл:", outputFilePath)
}

// parseCommandLineArguments - функция для парсинга данных с консоли
func parseCommandLineArguments() SortOptions {
	columnIndex := flag.Int("k", 0, "Индекс столбца для сортировки (по умолчанию 0)")
	numeric := flag.Bool("n", false, "Сортировать по числовому значению")
	reverse := flag.Bool("r", false, "Сортировать в обратном порядке")
	unique := flag.Bool("u", false, "Убрать повторяющиеся строки")
	month := flag.Bool("M", false, "Сортировать по названию месяца")
	ignoreBlanks := flag.Bool("b", false, "Игнорировать начальные пробелы")
	checkSorted := flag.Bool("c", false, "Проверить, отсортированы ли данные")
	numericSuffix := flag.Bool("h", false, "Сортировка по числовому значению с суффиксами")

	flag.Parse()

	return SortOptions{
		ColumnIndex:   *columnIndex,
		Numeric:       *numeric,
		Reverse:       *reverse,
		Unique:        *unique,
		Month:         *month,
		IgnoreBlanks:  *ignoreBlanks,
		CheckSorted:   *checkSorted,
		NumericSuffix: *numericSuffix,
	}
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

// readLinesFromFile - функция для записи данных в файла
func writeLinesToFile(filePath string, lines []string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}

// sortLines - структура для сортировки данных из файла
type sortLines struct {
	lines   []string
	options SortOptions
}

func (s sortLines) Len() int      { return len(s.lines) }
func (s sortLines) Swap(i, j int) { s.lines[i], s.lines[j] = s.lines[j], s.lines[i] }

func (s sortLines) Less(i, j int) bool {
	line1 := s.lines[i]
	line2 := s.lines[j]

	if s.options.Numeric {
		num1, err1 := extractNumericValue(line1, s.options.ColumnIndex, s.options.NumericSuffix)
		num2, err2 := extractNumericValue(line2, s.options.ColumnIndex, s.options.NumericSuffix)

		if err1 == nil && err2 == nil {
			if num1 < num2 {
				return !s.options.Reverse
			} else if num1 > num2 {
				return s.options.Reverse
			}
		}
	}

	if s.options.Month {
		month1, err1 := parseMonth(line1, s.options.ColumnIndex)
		month2, err2 := parseMonth(line2, s.options.ColumnIndex)

		if err1 == nil && err2 == nil {
			if month1 < month2 {
				return !s.options.Reverse
			} else if month1 > month2 {
				return s.options.Reverse
			}
		}
	}

	if s.options.IgnoreBlanks {
		line1 = strings.TrimSpace(line1)
		line2 = strings.TrimSpace(line2)
	}

	if !s.options.Reverse {
		return line1 < line2
	} else {
		return line1 > line2
	}
}

// extractNumericValue - функция для
func extractNumericValue(line string, columnIndex int, numericSuffix bool) (float64, error) {
	fields := strings.Fields(line)
	if columnIndex >= len(fields) {
		return 0, fmt.Errorf("индекс столбца вне диапазона")
	}

	value := fields[columnIndex]
	if numericSuffix {
		return parseNumericValueWithSuffix(value)
	}

	return strconv.ParseFloat(value, 64)
}

func parseNumericValueWithSuffix(value string) (float64, error) {
	suffixes := map[string]float64{
		"K": 1e3,
		"M": 1e6,
		"G": 1e9,
		"T": 1e12,
	}

	for suffix, multiplier := range suffixes {
		if strings.HasSuffix(value, suffix) {
			numStr := strings.TrimSuffix(value, suffix)
			num, err := strconv.ParseFloat(numStr, 64)
			if err != nil {
				return 0, err
			}
			return num * multiplier, nil
		}
	}

	return strconv.ParseFloat(value, 64)
}

// parseMonth - функция для сортировки по названию месяца
func parseMonth(line string, columnIndex int) (time.Month, error) {
	fields := strings.Fields(line)
	if columnIndex >= len(fields) {
		return 0, fmt.Errorf("индекс столбца вне диапазона")
	}

	monthStr := strings.ToLower(fields[columnIndex])
	switch monthStr {
	case "январь":
		return time.January, nil
	case "февраль":
		return time.February, nil
	case "март":
		return time.March, nil
	case "апрель":
		return time.April, nil
	case "май":
		return time.May, nil
	case "июнь":
		return time.June, nil
	case "июль":
		return time.July, nil
	case "август":
		return time.August, nil
	case "сентябрь":
		return time.September, nil
	case "октябрь":
		return time.October, nil
	case "ноябрь":
		return time.November, nil
	case "декабрь":
		return time.December, nil
	default:
		return 0, fmt.Errorf("неверный месяц: %s", monthStr)
	}
}

// isSorted - сортирует данные
func isSorted(lines []string, options SortOptions) bool {
	sorter := sortLines{lines, options}
	for i := 1; i < len(lines); i++ {
		if sorter.Less(i, i-1) {
			return false
		}
	}
	return true
}

// removeDuplicates - удаляет дубликаты (при флаге -u)
func removeDuplicates(lines []string) []string {
	seen := make(map[string]struct{})
	result := make([]string, 0, len(lines))

	for _, line := range lines {
		if _, exists := seen[line]; !exists {
			result = append(result, line)
			seen[line] = struct{}{}
		}
	}

	return result
}
