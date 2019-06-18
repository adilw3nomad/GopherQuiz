package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type quizItem struct {
	question string
	answer   string
}

// CheckAnswer compares two strings and returns a different message
// if they match
func CheckAnswer(answer string, correctAnswer string) (string, int) {
	if answer == correctAnswer {
		return "Correct! Well done\n", 1
	}
	return ("WRONG! Answer is: " + correctAnswer + "\n"), 0
}

// ShowScore checks how many questiones were answered correctly and
// returns a string saying how many were correct and how many were incorrect
func ShowScore(numberOfCorrectAnswers int, numberOfQuestions int) string {
	wrongAnswers := numberOfQuestions - numberOfCorrectAnswers
	return ("Correct answers: " + strconv.Itoa(numberOfCorrectAnswers) + ". Incorrect answers: " + strconv.Itoa(wrongAnswers) + "\n")
}

func main() {
	csvFile, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Create user input reader
	inputReader := bufio.NewReader(os.Stdin)

	var numberOfCorrectAnswers int
	for i := range records {

		// Create quizItem object.
		quizItem := quizItem{records[i][0], records[i][1]}

		// Print out the question.
		fmt.Println("Question: ", quizItem.question)

		// Prompt user to input their answer.
		fmt.Print("Enter your answer now: ")

		// Expect answer to be given once they hit return.
		text, err := inputReader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Your answer is:", text)

		// Trim the newline suffix from the input.
		text = strings.TrimSuffix(text, "\n")

		// Compare the answer given by the user to the correct answer
		// Print a response accordingly.
		responseString, scoreIncrement := CheckAnswer(text, quizItem.answer)
		numberOfCorrectAnswers = numberOfCorrectAnswers + scoreIncrement
		fmt.Print(responseString)
	}
	fmt.Print(ShowScore(numberOfCorrectAnswers, len(records)))
}
