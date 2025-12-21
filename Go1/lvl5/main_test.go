package main

import (
	"testing"
)

func TestGoalsSort(t *testing.T) {
	players := []Player{
		NewPlayer("Alice", 10, 2, 5),
		NewPlayer("Bob", 15, 3, 2),
		NewPlayer("Charlie", 10, 1, 4),
	}
	sorted := goalsSort(players)
	expectedOrder := []string{"Bob", "Alice", "Charlie"}
	for i, player := range sorted {
		if player.Name != expectedOrder[i] {
			t.Errorf("Expected %s at position %d, got %s", expectedOrder[i], i, player.Name)
		}
	}
}

func TestRatingSort(t *testing.T) {
	players := []Player{
		NewPlayer("Alice", 10, 2, 5),
		NewPlayer("Bob", 15, 3, 2),
		NewPlayer("Charlie", 10, 1, 4),
	}
	sorted := ratingSort(players)
	expectedOrder := []string{"Charlie", "Alice", "Bob"}
	for i, player := range sorted {
		if player.Name != expectedOrder[i] {
			t.Errorf("Expected %s at position %d, got %s", expectedOrder[i], i, player.Name)
		}
	}
}
