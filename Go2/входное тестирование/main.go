package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

type Ticket struct {
	Ticket string
	User   string
	Status string
	Date   time.Time
}

func isTicket(s string) bool {
	re := regexp.MustCompile(`^TICKET-\d+_[А-Яа-яЁёA-Za-z\d ]+_[А-Яа-яЁёA-Za-z\d ]+_\d{4}-\d{2}-\d{2}$`)

	if !re.MatchString(s) {
		return false
	}

	parts := strings.Split(s, "_")
	if len(parts) != 4 {
		return false
	}

	if parts[2] != "Готово" && parts[2] != "В работе" && parts[2] != "Не будет сделано" {
		return false
	}

	return true
}

func GetTasks(text string, user *string, status *string) []Ticket {
	myStrings := strings.Split(text, "\n")
	var res []Ticket = []Ticket{}

	for _, curString := range myStrings {
		curString = strings.TrimSpace(curString)
		if curString == "" || !isTicket(curString) {
			continue
		}
		if !isTicket(curString) {
			continue
		}

		parts := strings.Split(curString, "_")
		timeStamp, err := time.Parse("2006-01-02", parts[3])
		if err != nil {
			continue
		}
		curTic := Ticket{parts[0], parts[1], parts[2], timeStamp}

		if user != nil && status != nil && curTic.User == *user && curTic.Status == *status {
			res = append(res, curTic)
		} else if user == nil && status != nil && curTic.Status == *status {
			res = append(res, curTic)
		} else if status == nil && user != nil && curTic.User == *user {
			res = append(res, curTic)
		} else if user == nil && status == nil {
			res = append(res, curTic)
		}

	}

	return res
}

func main() {

	chatHistory := `
TICKET-12345_Паша Попов_Готово_2024-01-01
TICKET-12346_Иван Иванов_В работе_2024-01-02
TICKET-12347_Анна Смирнова_Не будет сделано_2024-01-03
TICKET-12348_Паша Попов_В работе_2024-01-04
`
	// Тест 3: поиск всех задач со статусом "В работе" для Паши Попова
	user := "Паша Попов"
	stat := "В работе"
	generalTasks := GetTasks(chatHistory, &user, &stat)
	if len(generalTasks) != 1 {
		fmt.Printf("Ожидалась 1 задача для Паши Попова со статусом 'В работе', найдено %v: %v ", len(generalTasks), chatHistory)
	}
	for _, task := range generalTasks {

		if task.Status != "В работе" {
			fmt.Printf("Найденная задача не имеет статус 'В работе': %v", task)
		}

		if task.User != "Паша Попов" {
			fmt.Printf("Найденная задача не принадлежит Паше Попову: %v", task)
		}
	}
}
