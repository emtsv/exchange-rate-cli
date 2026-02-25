package timeutil

import (
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	testTable := []struct {
		tdate    string
		expected string
	}{
		{
			tdate:    "2000-01-01",
			expected: "2000/01/01",
		},

		{
			tdate:    "1900-01-01",
			expected: "Дата должна быть после 2000 года",
		},
		{
			tdate:    "2022-12-",
			expected: "неправильная дата: %s",
		},
		{
			tdate:    "",
			expected: "",
		},
		{
			tdate:    "2022/01/01",
			expected: "неправильная дата: %s",
		},
	}
	for _, testCase := range testTable {
		result, err := ParseDate(testCase.tdate)
		if err != nil {
			t.Errorf("Ошибка %v", err)
		}

		expected := testCase.expected
		if testCase.tdate == "" {
			expected = time.Now().Format("2006/01/02")
		}

		if result != expected {
			t.Errorf("Некорректный результат. ожидание: %s, получение: %s", expected, result)
		}
	}
}
