package commands

import (
	"fmt"
	"time"
)

func def() (day, month, year int) {
	default_time := time.Now()
	return default_time.Day(), default_time.Month(), default_time.Year()
}

func parse() {
	day, month, year := def()
	fmt.Printf("Сегодняшняя дата: \n", &day, &month, &year)
}

func convertion(date) {
	
	date:=fmt.Scan(&day,&month,&year)
	date := (year, time.Month(month),day,0,0,0,0, time.UTC)

	formatted:=date.Format("02/01/2006")


}