package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

type quizItem struct {
	question string
	answer   string
}

// CheckAnswer compares two strings and returns a different message
// if they match
func CheckAnswer(answer string, correctAnswer string) string {
	if answer == correctAnswer {
		return "Correct! Well done\n"
	}
	return ("WRONG! Answer is: " + correctAnswer + "\n")
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
		fmt.Print(CheckAnswer(text, quizItem.answer))
	}
}
