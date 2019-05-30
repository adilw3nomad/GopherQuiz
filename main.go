package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type quizItem struct {
	question string
	answer   string
	correct  bool
}

func main() {
	// Open up the csv file. TODO Use flag here to specify file path
	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	// Keep file in memory until main() has finished
	defer file.Close()

	// we need to find the number of lines in the file to use as len argument for making the slice
	quiz := make([]quizItem, 10)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// we then create a quiz item for each of the lines in the quiz
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
