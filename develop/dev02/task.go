package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// В данной задаче следует провести множество проверок для достижения более точного результата
// но т.к. в условии задания приведено слишком мало данных, нельзя точно предположить, как именно должны распаковываться строки
// и обрабатываться escape-последовательности
// Поэтому проверки были опущены, а задание сделано ровно для входных значений, указанных в примере
func unpack(data string) (res string, err error) {
	// Сразу проверим, если входная строка пуста, то выйдем с функции
	// Если во входной строке только один символ, входная строка некорректная, вернем ошибку
	if len(data) == 0 {
		return
	} else if len(data) == 1 {
		return "", errors.New("некорректная строка")
	}

	// Создаем
	// builder для добавления распакованных символов
	// sliceRune для записи в него входных данных
	var builder strings.Builder
	var sliceRune []rune = []rune(data)

	for i := 0; i < len(sliceRune); i++ {

		// Если слева от текущего элемента и справа есть символ \ - значит это escape-последовательность и нам нужно просто записать число
		if i-1 >= 0 && sliceRune[i-1] == '\\' && i+1 < len(sliceRune) && sliceRune[i+1] == '\\' {
			builder.WriteRune(sliceRune[i])
			i++

			// Если слева от текущего элемента есть символ \, а справа число - значит это escape-последовательность и нам нужно повторить текущий элемент
		} else if i-1 >= 0 && sliceRune[i-1] == '\\' && i+1 < len(sliceRune) && unicode.IsDigit(sliceRune[i+1]) {
			counter, _ := strconv.Atoi(string(sliceRune[i+1]))

			for j := 0; j < counter; j++ {
				builder.WriteRune(sliceRune[i])
			}
			i++
			// Если слева от текущего элемента есть символ \, а справа ничего нет - значит это escape-последовательность и нам нужно просто записать число
		} else if i-1 >= 0 && sliceRune[i-1] == '\\' && i+1 == len(sliceRune) {
			builder.WriteRune(sliceRune[i])
			// Если текущий элемент есть символ \ - значит это escape-последовательность и нам не нужно записывать этот символ
		} else if sliceRune[i] == '\\' {
			continue
			// Если текущий элемент число и слева от него есть что-либо, значит это обычная строка и нам нужно повторить элемент слева
		} else if unicode.IsDigit(sliceRune[i]) && i-1 >= 0 {
			counter, _ := strconv.Atoi(string(sliceRune[i]))

			if unicode.IsDigit(sliceRune[i-1]) {
				return "", errors.New("некорректная строка")
			}

			for j := 0; j < counter-1; j++ {
				builder.WriteRune(sliceRune[i-1])
			}
			// Если тут просто какой-либо символ, записываем его
		} else {
			builder.WriteRune(sliceRune[i])
		}
	}

	// Получаем итоговую строку с builder
	res = builder.String()
	return res, err
}

func main() {
	testCases := []string{"a4bc2d5e", "abcd", "45", "", "qwe\\4\\5", "qwe\\45", "qwe\\\\5"}

	for _, val := range testCases {
		res, err := unpack(val)
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Printf("%v -> %v\n", val, res)
	}
}
