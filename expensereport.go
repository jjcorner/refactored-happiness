package main

import (
	"fmt"
	"time"
)

type ExpenseType int

const (
	Dinner ExpenseType = iota + 1
	Breakfast
	CarRental
)

const (
	DinnerExpenseThreshold    = 5000
	BreakfastExpenseThreshold = 1000
)

type Expense struct {
	Type   ExpenseType
	Amount int
}

func (e Expense) IsMeal() bool {
	return e.Type == Dinner || e.Type == Breakfast
}

func PrintReport(expenses []Expense) {
	total := 0
	mealExpenses := 0

	fmt.Printf("Expenses %s\n", time.Now().Format("2006-01-02"))

	for _, expense := range expenses {
		if expense.IsMeal() {
			mealExpenses += expense.Amount
		}

		var expenseName string
		switch expense.Type {
		case Dinner:
			expenseName = "Dinner"
		case Breakfast:
			expenseName = "Breakfast"
		case CarRental:
			expenseName = "Car Rental"
		}

		var mealOverExpensesMarker string
		if expense.Type == Dinner && expense.Amount > DinnerExpenseThreshold || expense.Type == Breakfast && expense.Amount > BreakfastExpenseThreshold {
			mealOverExpensesMarker = "X"
		} else {
			mealOverExpensesMarker = " "
		}

		fmt.Printf("%s\t%d\t%s\n", expenseName, expense.Amount, mealOverExpensesMarker)
		total += expense.Amount
	}

	fmt.Printf("Meal expenses: %d\n", mealExpenses)
	fmt.Printf("Total expenses: %d\n", total)
}
