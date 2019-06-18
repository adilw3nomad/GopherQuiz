package main

import "testing"

func TestCheckCorrectAnswer(t *testing.T) {
	responseString, scoreIncrement := CheckAnswer("3", "3")
	if responseString != "Correct! Well done\n" {
		t.Errorf("CheckCorrectAnswer, responseString = %v; want 'Correct! Well done\n'", responseString)
	}
	if scoreIncrement != 1 {
		t.Errorf("CheckCorrectAnswer, scoreIncrementt = %v; want true", scoreIncrement)
	}
}

func TestCheckIncorrectAnswer(t *testing.T) {
	responseString, scoreIncrement := CheckAnswer("3", "2")
	if responseString != "WRONG! Answer is: 2\n" {
		t.Errorf("CheckIncorrectAnswer responseString = %v; want 'WRONG! Answer is: 2\n'", responseString)
	}
	if scoreIncrement != 0 {
		t.Errorf("CheckCorrectAnswer, scoreIncrement = %v; want false", scoreIncrement)
	}
}

func TestShowScore(t *testing.T) {
	score, quizItems := 6, 10
	responseString := ShowScore(score, quizItems)
	if responseString != "Correct answers: 6. Incorrect answers: 4\n" {
		t.Errorf("ShowScore, responseString = %v; want 'Correct answers: 6. Incorrect answers: 4\n", responseString)
	}
}
