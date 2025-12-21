package main

import (
	// "strings"
	"testing"
)

func TestGetTasks(t *testing.T) {
	chatHistory := `
TICKET-12345_Паша Попов_Готово_2024-01-01
TICKET-12346_Иван Иванов_В работе_2024-01-02
TICKET-12347_Анна Смирнова_Не будет сделано_2024-01-03
TICKET-12348_Паша Попов_В работе_2024-01-04
`

	// Тест 1: поиск всех задач Паши Попова
	user := "Паша Попов"
	tasks := GetTasks(chatHistory, &user, nil)
	if len(tasks) != 2 {
		t.Errorf("Ожидалось 2 задачи для Паши Попова, найдено %d: text %v, user %v, status %v", len(tasks), chatHistory, user, nil)
	}
	for _, task := range tasks {
		if task.User != "Паша Попов" {
			t.Errorf("Найденная задача не принадлежит Паше Попову: %v", task)
		}
	}

	// Тест 2: поиск всех задач со статусом "В работе"
	stat := "В работе"
	workTasks := GetTasks(chatHistory, nil, &stat)
	if len(workTasks) != 2 {
		t.Errorf("Ожидалось 2 задачи со статусом 'В работе', найдено %d: text %v, user %v, status %v", len(workTasks), chatHistory, nil, stat)
	}
	for _, task := range workTasks {
		if task.Status != "В работе" {
			t.Errorf("Найденная задача не имеет статус 'В работе': %v", task)
		}
	}

	// Тест 3: поиск всех задач со статусом "В работе" для Паши Попова
	generalTasks := GetTasks(chatHistory, &user, &stat)
	if len(generalTasks) != 1 {
		t.Errorf("Ожидалась 1 задача для Паши Попова со статусом 'В работе', найдено %v: %v ", len(generalTasks), chatHistory)
	}
	for _, task := range generalTasks {

		if task.Status != "В работе" {
			t.Errorf("Найденная задача не имеет статус 'В работе': %v", task)
		}

		if task.User != "Паша Попов" {
			t.Errorf("Найденная задача не принадлежит Паше Попову: %v", task)

		}
	}
}
