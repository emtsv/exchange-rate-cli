package timeutil

import (
	"fmt"
	"time"
)

var now = time.Now()

func ParseDate(date string) (string, error) {

	if date == "" {
		return time.Now().Format("2006/01/02"), nil
	}

	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return "", err
	}

	if t.Year() < 2000 {
		return "", fmt.Errorf("дата должна быть после 2000 года")
	}

	if t.After(now) {
		return "", fmt.Errorf("дата не может быть в будущем")
	}

	return t.Format("2006/01/02"), nil
}
