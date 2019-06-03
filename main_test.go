package main

import "testing"

func TestCheckCorrectAnswer(t *testing.T) {
	got := CheckAnswer("3", "3")
	if got != "Correct! Well done" {
		t.Errorf("CheckCorrectAnswer = %v; want 'Correct! Well Done'", got)
	}
}

func TestCheckIncorrectAnswer(t *testing.T) {
	got := CheckAnswer("3", "2")
	if got != "WRONG! Answer is: 2" {
		t.Errorf("CheckIncorrectAnswer = %v; want 'WRONG! Answer is: 2'", got)
	}
}
