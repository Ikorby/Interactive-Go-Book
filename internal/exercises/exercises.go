package exercises

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Ikorby/Interactive-Go-Book/internal/progress"
)

// Exercise описывает одно упражнение
type Exercise struct {
	ID       string
	Question string
	Answer   string
}

// Набор упражнений (пример)
var ExercisesSet = []Exercise{
	{
		ID:       "1.1",
		Question: "What is the keyword to define a function in Go?",
		Answer:   "func",
	},
	{
		ID:       "1.2",
		Question: "How do you declare a variable of type int?",
		Answer:   "var x int",
	},
}

// ShowExercises показывает список упражнений и проверяет ответы
func ShowExercises(p *progress.Progress) {
	fmt.Println("\nAvailable exercises:")

	for i, ex := range ExercisesSet {
		fmt.Printf("%d. Exercise %s\n", i+1, ex.ID)
	}

	reader := bufio.NewReader(os.Stdin)

	for _, ex := range ExercisesSet {
		fmt.Printf("\nExercise %s: %s\nYour answer: ", ex.ID, ex.Question)
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)

		if strings.EqualFold(userInput, ex.Answer) {
			fmt.Println("Correct!")
			p.AddExercise(ex.ID)
		} else {
			fmt.Printf("Wrong! Correct answer is: %s\n", ex.Answer)
		}

		// Сохраняем прогресс после каждого упражнения
		if err := progress.SaveProgress(*p); err != nil {
			fmt.Println("Error saving progress:", err)
		}
	}
}

