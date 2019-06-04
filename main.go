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

// Rubbish ....
func CheckAnswer(answer string, correctAnswer string) string {
	if answer == correctAnswer {
		return "Correct! Well done"
	} else {
		return ("WRONG! Answer is: " + correctAnswer)
	}
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
	for i, _ := range records {
		// Create a quizItem object.
		quizItem := quizItem{records[i][0], records[i][1]}

		// Print out the question to the user.
		fmt.Println("Question:", quizItem.question)

		// Create a reader and allow the user to input their answer.
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter your answer now: ")

		// Expect an answer to be given once they hit return.
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Your answer is:", text)

		// Trim the newline suffix from the input
		text = strings.TrimSuffix(text, "\n")

		fmt.Println(CheckAnswer(text, quizItem.answer))
	}

}
