package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestPrintReport_PreRefactor(t *testing.T) {

	t.Run("Expense in range", func(t *testing.T) {
		// Running the function
		expenses := []Expense{
			{Type: Dinner, Amount: 4500},
			{Type: Breakfast, Amount: 900},
			{Type: CarRental, Amount: 4000},
		}

		// Capture the output of the PrintReport function
		capturedOutput, err := captureOutput(expenses, PrintReport)
		if err != nil {
			t.Fatalf("Failed to capture output: %v", err)
		}

		// Expected output
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
		// Running the function
		expenses := []Expense{
			{Type: Dinner, Amount: 8000},
			{Type: Breakfast, Amount: 900},
			{Type: CarRental, Amount: 4000},
		}

		// Capture the output of the PrintReport function
		capturedOutput, err := captureOutput(expenses, PrintReport)
		if err != nil {
			t.Fatalf("Failed to capture output: %v", err)
		}

		// Expected output
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
		// Running the function
		var expenses []Expense

		// Capture the output of the PrintReport function
		capturedOutput, err := captureOutput(expenses, PrintReport)
		if err != nil {
			t.Fatalf("Failed to capture output: %v", err)
		}

		// Expected output
		expectedDate := time.Now().Format("2006-01-02")
		expectedOutput := "Expenses " + expectedDate + "\n" +
			"Meal expenses: 0\n" +
			"Total expenses: 0\n"

		assert.Equal(t, expectedOutput, capturedOutput)
	})
}

/*func TestPrintReportAddNewExpense_PostRefactor(t *testing.T) {

	t.Run("Expense in range", func(t *testing.T) {
		// Running the function
		expenses := []Expense{
			{Type: Dinner, Amount: 4500},
			{Type: Breakfast, Amount: 900},
			{Type: CarRental, Amount: 4000},
			{Type: 5, Amount: 20000},
		}

		// Capture the output of the PrintReport function
		capturedOutput, err := captureOutput(expenses, PrintReport)
		if err != nil {
			t.Fatalf("Failed to capture output: %v", err)
		}

		// Expected output
		expectedDate := time.Now().Format("2006-01-02")
		expectedOutput := "Expenses " + expectedDate + "\n" +
			"Dinner\t4500\t \n" +
			"Breakfast\t900\t \n" +
			"Car Rental\t4000\t \n" +
			"Meal expenses: 5400\n" +
			"Total expenses: 9400\n"

		assert.Equal(t, expectedOutput, capturedOutput)
	})
}*/

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
