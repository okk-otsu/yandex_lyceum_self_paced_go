package main

import (
	"errors"
	"strings"
	"time"
	"unicode/utf8"
)

func TimeNow() time.Time {
	return time.Now()
}
func currentDayOfTheWeek() string {
	now := TimeNow()
	switch now.Weekday() {
	case time.Monday:
		return "Понедельник"
	case time.Tuesday:
		return "Вторник"
	case time.Wednesday:
		return "Среда"
	case time.Thursday:
		return "Четверг"
	case time.Friday:
		return "Пятница"
	case time.Saturday:
		return "Суббота"
	default:
		return "Воскресенье"
	}
}

func dayOrNight() string {
	now := TimeNow()
	hour := now.Hour()

	if hour >= 10 && hour <= 22 {
		return "День"
	}
	return "Ночь"
}

func nextFriday() int {
	now := TimeNow()
	today := now.Weekday()

	days := (int(time.Friday) - int(today) + 7) % 7
	return days
}

func CheckCurrentDayOfTheWeek(answer string) bool {
	current := currentDayOfTheWeek()
	return strings.EqualFold(current, answer)
}

func CheckNowDayOrNight(answer string) (bool, error) {
	if utf8.RuneCountInString(answer) != 4 {
		return false, errors.New("исправь свой ответ, а лучше ложись поспать")
	}

	current := dayOrNight()
	if strings.EqualFold(current, answer) {
		return true, nil
	}

	return false, nil
}
