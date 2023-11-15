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

// findAnagrams ищет множества анаграмм в заданном словаре
func findAnagrams(words *[]string) map[string][]string {
	anagramSets := make(map[string][]string)

	for _, word := range *words {
		// Приводим слово к нижнему регистру и сортируем в нем буквы
		sortedWord := sortString(strings.ToLower(word))

		// Проверяем, есть ли уже множество анаграмм для отсортированного слова
		if set, found := anagramSets[sortedWord]; found {
			// Добавляем слово к существующему множеству
			anagramSets[sortedWord] = append(set, word)
		} else {
			// Если нет, тогда создаем новое множество анаграмм для отсортированного слова
			anagramSets[sortedWord] = []string{word}
		}
	}

	// Удаляем множества из одного элемента
	for key, value := range anagramSets {
		if len(value) <= 1 {
			delete(anagramSets, key)
		} else {
			// Сортируем слова в множестве по возрастанию
			sort.Strings(anagramSets[key])
		}
	}

	return anagramSets
}

// sortString сортирует буквы в строке
func sortString(s string) string {
	sortedRunes := []rune(s)
	sort.Slice(sortedRunes, func(i, j int) bool {
		return sortedRunes[i] < sortedRunes[j]
	})
	return string(sortedRunes)
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	result := findAnagrams(&words)

	for key, value := range result {
		if len(value) > 1 {
			fmt.Printf("Множество анаграмм для %s: %v\n", key, value)
		}
	}
}
