package timeutil

import (
	"errors"
	"time"
)

var (
	ErrDateTooOld   = errors.New("дата должна быть после 2000 года")
	ErrInvalidDate  = errors.New("неверный формат даты")
	ErrDateInFuture = errors.New("дата не может быть в будущем")
)

var now = time.Now()

func ParseDate(date string) (string, error) {

	if date == "" {
		return time.Now().Format("2006/01/02"), nil
	}

	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return "", ErrInvalidDate
	}

	if t.Year() < 2000 {
		return "", ErrDateTooOld
	}

	if t.After(now) {
		return "", ErrDateInFuture
	}

	return t.Format("2006/01/02"), nil
}
