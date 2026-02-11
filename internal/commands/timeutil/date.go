package timeutil

import "time"

func ParseDate(date string) (string, error) {
	if date == "" {
		return time.Now().Format("2006/01/02"), nil
	}

	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return "", err
	}

	return t.Format("2006/01/02"), nil
}
