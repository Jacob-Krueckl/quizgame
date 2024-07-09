package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Question struct {
	question string
	answer   int
	options  []string
}

type Quiz struct {
	questions []Question
}

func (q *Quiz) shuffle() {
	for i := range q.questions {
		j := rand.Intn(i + 1)
		q.questions[i], q.questions[j] = q.questions[j], q.questions[i]
	}
}

// askQuestion asks the user a question and returns the user's answer
func (q *Quiz) askQuestion(i int) int {
	var userAnswer int
	fmt.Printf("Question: %s\n", q.questions[i].question)
	for i, option := range q.questions[i].options {
		if option == "" {
			continue
		}
		fmt.Printf("  %v: %v\n", i+1, option)
	}
	fmt.Print("Your answer: ")
	fmt.Scanf("%d\n", &userAnswer)
	return userAnswer
}

// addQuestion adds a question to the quiz
func (q *Quiz) addQuestion(question string, answer int, options []string) {
	q.questions = append(q.questions, Question{question, answer, options})
}

// createQuiz creates a quiz from multiple lines of CSV data
func createQuiz(lines *[][]string) Quiz {

	quiz := Quiz{}

	for i, line := range *lines {

		// ignore the header
		if i == 0 {
			continue
		}

		questionFromFile := line[0]
		answerFromFile, err := strconv.Atoi(line[1])
		if err != nil {
			fmt.Println("Error:", err)
			return quiz
		}
		optionsFromFile := line[2:]

		quiz.addQuestion(questionFromFile, answerFromFile, optionsFromFile)
	}

	return quiz

}

// quizGame runs the quiz game
func quizGame(quiz Quiz, correctWriteChan chan int) {

	for _, currentQuestion := range quiz.questions {

		correct := 0
		fmt.Printf("Question: %s\n", currentQuestion.question)
		for i, option := range currentQuestion.options {
			if option == "" {
				continue
			}
			fmt.Printf("  %v: %v\n", i+1, option)
		}

		var userAnswer int
		fmt.Print("Your answer: ")
		fmt.Scanf("%d\n", &userAnswer)

		if userAnswer == currentQuestion.answer {
			correct++
		} else {
			fmt.Printf("Wrong answer. Correct answer is %v\n", currentQuestion.answer)
		}

		correctWriteChan <- correct
	}

	// signal that the game is over
	correctWriteChan <- -1
}

// wrapUp prints the final score and waits for user input to exit
func wrapUp(correct int, total int) {
	fmt.Printf("You got %v out of %v correct\n", correct, total)
	fmt.Print("Press anything to exit")
	fmt.Scanf("%v\n")
}

func main() {

	quizLength := 10 * time.Second

	fmt.Println("Quiz Game")
	fmt.Printf("  You have %v to answer all the questions\n", quizLength)
	fmt.Printf("  Enter the number of the correct answer for each question\n\n")

	file, err := os.Open("problems.csv")
	if err != nil {
		fmt.Println("Error when opening CSV:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error when reading all lines:", err)
		return
	}

	correctWriteChan := make(chan int, len(lines)-1)

	quiz := createQuiz(&lines)
	quiz.shuffle()

	go quizGame(quiz, correctWriteChan)

	correct := 0
	for {
		// select will block until one of the cases is ready
		select {
		// time.After will send a message on the channel after the specified time
		case <-time.After(quizLength):
			fmt.Println("\nTime's up!")
			wrapUp(correct, len(lines)-1)
			return
		// correctWriteChan will send either 1 for a correct answer, 0 for an incorrect answer, or -1 for the end of the game
		case currentResult := <-correctWriteChan:
			if currentResult == -1 {
				wrapUp(correct, len(lines)-1)
				return
			}
			correct += currentResult
		}
	}

}
