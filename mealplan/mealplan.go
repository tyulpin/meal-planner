package mealplan

import "fmt"

type Meal struct {
	Name string
	Type string
}

type MealPlan struct {
	Days map[string][]Meal
}

func New() *MealPlan {
	return &MealPlan{
		Days: make(map[string][]Meal),
	}
}

func (mp *MealPlan) AddMeal(day, mealType, name string) {
	meal := Meal{Name: name, Type: mealType}
	mp.Days[day] = append(mp.Days[day], meal)
}

func (mp *MealPlan) View() {
	for day, meals := range mp.Days {
		fmt.Printf("\n%s:\n", day)
		for _, meal := range meals {
			fmt.Printf("  %s: %s\n", meal.Type, meal.Name)
		}
	}
}

