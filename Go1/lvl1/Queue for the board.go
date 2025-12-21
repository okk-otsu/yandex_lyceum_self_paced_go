package main

import (
	"fmt"
)

func queue(arr []string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d. %s\n", i+1, arr[i])
	}
}

func main() {
	var count int = 0
	arr := [5]string{"-", "-", "-", "-", "-"}
	for {
		var s string
		fmt.Scan(&s)
		if s == "конец" {
			queue(arr[:])
			break
		} else if s == "очередь" {
			queue(arr[:])
			continue
		} else if s == "количество" {
			fmt.Printf("Осталось свободных мест: %d\n", 5-count)
			fmt.Printf("Всего человек в очереди: %d\n", count)
			continue
		}

		var n int
		fmt.Scan(&n)
		if n < 1 || n > 5 {
			fmt.Printf("Запись на место номер %d невозможна: некорректный ввод\n", n)
			continue
		}
		if count == 5 {
			fmt.Printf("Запись на место номер %d невозможна: очередь переполнена\n", n)
			continue
		}
		idx := n - 1
		if arr[idx] != "-" {
			fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", n)
			continue
		}

		arr[idx] = s
		count++
	}
}
