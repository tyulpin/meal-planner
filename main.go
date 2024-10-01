package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"meal-planner/mealplan"
	"meal-planner/ollama"
)

func main() {
	mp := mealplan.New()

	for {
		printMenu()
		choice := getMenuChoice()

		switch choice {
		case 1:
			addMeal(mp)
		case 2:
			mp.View()
		case 3:
			mealplan.Save(mp)
		case 4:
			mp = mealplan.Load()
		case 5:
			getMealSuggestion()
		case 6:
			chatWithLLM()
		case 7:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func printMenu() {
	fmt.Println("\n1. Add meal")
	fmt.Println("2. View meal plan")
	fmt.Println("3. Save meal plan")
	fmt.Println("4. Load meal plan")
	fmt.Println("5. Get meal suggestion from LLM")
	fmt.Println("6. Chat with LLM")
	fmt.Println("7. Exit")
}

func getMenuChoice() int {
	var choice int
	fmt.Print("Choose an option: ")
	fmt.Scanln(&choice)
	return choice
}

func addMeal(mp *mealplan.MealPlan) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter day of the week: ")
	day, _ := reader.ReadString('\n')
	day = strings.TrimSpace(day)

	fmt.Print("Enter meal type (breakfast/lunch/dinner): ")
	mealType, _ := reader.ReadString('\n')
	mealType = strings.TrimSpace(mealType)

	fmt.Print("Enter meal name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	mp.AddMeal(day, mealType, name)
	fmt.Println("Meal added successfully!")
}

func getMealSuggestion() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter meal type (breakfast/lunch/dinner): ")
	mealType, _ := reader.ReadString('\n')
	mealType = strings.TrimSpace(mealType)

	prompt := fmt.Sprintf("Suggest a healthy %s meal:", mealType)
	response, err := ollama.Query(prompt)
	if err != nil {
		fmt.Printf("Error getting suggestion: %v\n", err)
		return
	}

	fmt.Printf("LLM Suggestion: %s\n", response)
}

func chatWithLLM() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("You: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			break
		}

		response, err := ollama.Query(input)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		fmt.Printf("LLM: %s\n", response)
	}
}
