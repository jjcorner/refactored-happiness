package expensereport

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func captureOutput(expenses []Expense, f func([]Expense)) (string, error) {
	// Saving the original stdout
	originalStdout := os.Stdout

	// Creating a pipe to capture stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Capturing the output
	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		_, _ = buf.ReadFrom(r)
		outC <- buf.String()
	}()

	f(expenses)

	// Restoring stdout and closing the write end of the pipe
	w.Close()
	os.Stdout = originalStdout

	return <-outC, nil
}

func TestPrintReport_PreRefactor(t *testing.T) {

	t.Run("Expense in range", func(t *testing.T) {
		expenses := []Expense{
			{Type: Dinner, Amount: 4500},
			{Type: Breakfast, Amount: 900},
			{Type: CarRental, Amount: 4000},
		}

		capturedOutput, err := captureOutput(expenses, PrintReport)
		if err != nil {
			t.Fatalf("Failed to capture output: %v", err)
		}

		expectedDate := time.Now().Format("2006-01-02")
		expectedOutput := "Expenses " + expectedDate + "\n" +
			"Dinner\t4500\t \n" +
			"Breakfast\t900\t \n" +
			"Car Rental\t4000\t \n" +
			"Meal expenses: 5400\n" +
			"Total expenses: 9400\n"

		assert.Equal(t, expectedOutput, capturedOutput)
	})

	t.Run("Over Expense", func(t *testing.T) {
		expenses := []Expense{
			{Type: Dinner, Amount: 8000},
			{Type: Breakfast, Amount: 900},
			{Type: CarRental, Amount: 4000},
		}

		capturedOutput, err := captureOutput(expenses, PrintReport)
		if err != nil {
			t.Fatalf("Failed to capture output: %v", err)
		}

		expectedDate := time.Now().Format("2006-01-02")
		expectedOutput := "Expenses " + expectedDate + "\n" +
			"Dinner\t8000\tX\n" +
			"Breakfast\t900\t \n" +
			"Car Rental\t4000\t \n" +
			"Meal expenses: 8900\n" +
			"Total expenses: 12900\n"

		assert.Equal(t, expectedOutput, capturedOutput)
	})

	t.Run("No expenses", func(t *testing.T) {
		var expenses []Expense

		capturedOutput, err := captureOutput(expenses, PrintReport)
		if err != nil {
			t.Fatalf("Failed to capture output: %v", err)
		}

		expectedDate := time.Now().Format("2006-01-02")
		expectedOutput := "Expenses " + expectedDate + "\n" +
			"Meal expenses: 0\n" +
			"Total expenses: 0\n"

		assert.Equal(t, expectedOutput, capturedOutput)
	})
}

func TestPrintReport_AddLunchExpense(t *testing.T) {
	t.Run("Expense in range", func(t *testing.T) {
		expenses := []Expense{
			{Type: Dinner, Amount: 4500},
			{Type: Breakfast, Amount: 900},
			{Type: Lunch, Amount: 1000},
			{Type: CarRental, Amount: 4000},
		}

		capturedOutput, err := captureOutput(expenses, PrintReport)
		if err != nil {
			t.Fatalf("Failed to capture output: %v", err)
		}

		expectedDate := time.Now().Format("2006-01-02")
		expectedOutput := "Expenses " + expectedDate + "\n" +
			"Dinner\t4500\t \n" +
			"Breakfast\t900\t \n" +
			"Lunch\t1000\t \n" +
			"Car Rental\t4000\t \n" +
			"Meal expenses: 6400\n" +
			"Total expenses: 10400\n"

		assert.Equal(t, expectedOutput, capturedOutput)
	})

	t.Run("Over Expense", func(t *testing.T) {
		expenses := []Expense{
			{Type: Dinner, Amount: 8000},
			{Type: Breakfast, Amount: 900},
			{Type: Lunch, Amount: 4000},
			{Type: CarRental, Amount: 4000},
		}

		capturedOutput, err := captureOutput(expenses, PrintReport)
		if err != nil {
			t.Fatalf("Failed to capture output: %v", err)
		}

		expectedDate := time.Now().Format("2006-01-02")
		expectedOutput := "Expenses " + expectedDate + "\n" +
			"Dinner\t8000\tX\n" +
			"Breakfast\t900\t \n" +
			"Lunch\t4000\tX\n" +
			"Car Rental\t4000\t \n" +
			"Meal expenses: 12900\n" +
			"Total expenses: 16900\n"

		assert.Equal(t, expectedOutput, capturedOutput)
	})

}
