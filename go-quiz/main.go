package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type Question struct {
	question string
	answer   string
}

func newQuestion(data []string) *Question {
	q := Question{
		question: data[0],
		answer:   data[1],
	}

	return &q
}

func (q Question) ask() bool {
	fmt.Println(q.question, "?")

	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')

	text = strings.Replace(text, "\n", "", 1)

	if text == q.answer {
		return true
	}

	return false
}

func main() {
	file, err := os.Open(
		"problems.csv",
	)

	if err != nil {
		fmt.Println("error opening problems.csv file")
		panic(1)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	csv_data, err := csvReader.ReadAll()

	if err != nil {
		fmt.Println("error parsing problems.csv file")
		panic(1)
	}

	questions := make([]Question, 0)

	for _, record := range csv_data {
		q := newQuestion(record)
		questions = append(questions, *q)
	}

	answersCorrect := 0
	for _, q := range questions {
		answer := q.ask()

		if answer == true {
			answersCorrect += 1
		}
	}

	fmt.Println("number of correct answers is:", answersCorrect)

	os.Exit(0)
}
