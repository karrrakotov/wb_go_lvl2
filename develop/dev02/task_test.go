package main

import (
	"testing"
)

func TestUnpack(t *testing.T) {
	// Arrange
	testTable := []struct {
		data     string
		expected string
		isError  bool
	}{
		{
			data:     "a4bc2d5e",
			expected: "aaaabccddddde",
			isError:  false,
		},
		{
			data:     "abcd",
			expected: "abcd",
			isError:  false,
		},
		{
			data:     "45",
			expected: "",
			isError:  true,
		},
		{
			data:     "",
			expected: "",
			isError:  false,
		},
		{
			data:     "qwe\\4\\5",
			expected: "qwe45",
			isError:  false,
		},
		{
			data:     "qwe\\45",
			expected: "qwe44444",
			isError:  false,
		},
		{
			data:     "qwe\\\\5",
			expected: "qwe\\\\\\\\\\",
			isError:  false,
		},
	}

	// Act
	for _, testCase := range testTable {
		result, err := unpack(testCase.data)

		t.Logf("Входные данные теста Unpack: %v, результат: %v\n", testCase.data, result)

		// Assert
		if testCase.isError {
			if err == nil {
				t.Errorf("Для входных данных %s ожидалась ошибка, но получено значение %s", testCase.data, result)
			}
		} else {
			if err != nil {
				t.Errorf("Для входных данных %s ожидалось значение %s, но получена ошибка: %v", testCase.data, testCase.expected, err)
			} else if result != testCase.expected {
				t.Errorf("Для входных данных %s ожидалось значение %s, получено значение %s", testCase.data, testCase.expected, result)
			}
		}
	}
}
