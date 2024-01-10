package expensereport

import (
	"fmt"
	"time"
)

type ExpenseType int

const (
	Dinner ExpenseType = iota + 1
	Breakfast
	CarRental
	Lunch
)

const (
	DinnerExpenseThreshold    = 5000
	BreakfastExpenseThreshold = 1000
	LunchExpenseThreshold     = 2000
)

type Expense struct {
	Type   ExpenseType
	Amount int
}

func (e Expense) IsMeal() bool {
	return e.Type == Dinner || e.Type == Breakfast || e.Type == Lunch
}

func (e Expense) OverThreshold() bool {
	switch e.Type {
	case Dinner:
		return e.Amount > DinnerExpenseThreshold
	case Breakfast:
		return e.Amount > BreakfastExpenseThreshold
	case Lunch:
		return e.Amount > LunchExpenseThreshold
	default:
		return false
	}
}

func (e Expense) Name() string {
	switch e.Type {
	case Dinner:
		return "Dinner"
	case Breakfast:
		return "Breakfast"
	case Lunch:
		return "Lunch"
	case CarRental:
		return "Car Rental"
	default:
		return "Unknown"
	}
}

func printExpense(expenses []Expense) {
	fmt.Printf("Expenses %s\n", time.Now().Format("2006-01-02"))
	for _, expense := range expenses {
		marker := " "
		if expense.OverThreshold() {
			marker = "X"
		}
		fmt.Printf("%s\t%d\t%s\n", expense.Name(), expense.Amount, marker)
	}
}
func calculateTotal(expenses []Expense) (total int, mealExpenses int) {
	for _, expense := range expenses {
		if expense.IsMeal() {
			mealExpenses += expense.Amount
		}
		total += expense.Amount
	}
	return
}

func PrintReport(expenses []Expense) {
	total, mealExpense := calculateTotal(expenses)
	printExpense(expenses)
	fmt.Printf("Meal expenses: %d\n", mealExpense)
	fmt.Printf("Total expenses: %d\n", total)
}
