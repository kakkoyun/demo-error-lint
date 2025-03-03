package main

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

// Custom error types for demonstration
type NotFoundError struct {
	Item string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s not found", e.Item)
}

// Sentinel errors
var (
	ErrInvalidInput = errors.New("invalid input")
	ErrTimeout      = errors.New("operation timed out")
)

// Function that returns an error
func fetchData() error {
	return ErrInvalidInput
}

// Function that returns a wrapped error
func processData() error {
	err := fetchData()
	if err != nil {
		// ISSUE: Using %v instead of %w in fmt.Errorf
		return fmt.Errorf("failed to process data: %w", err)
	}
	return nil
}

// Function that returns a custom error
func findItem(item string) error {
	return &NotFoundError{Item: item}
}

// Function demonstrating EOF handling (documented special case)
func readFullBuffer(r io.Reader, buf []byte) (int, error) {
	n, err := r.Read(buf)
	// This is actually allowed by the linter because io.EOF is documented
	// to be returned unwrapped
	if err == io.EOF {
		return n, nil
	}
	return n, err
}

func main() {
	// Demo 1: Error comparisons with ==
	err := fetchData()

	// ISSUE: Direct comparison instead of errors.Is
	if err == ErrInvalidInput {
		fmt.Println("Invalid input detected")
	}

	// Correct way
	if errors.Is(err, ErrInvalidInput) {
		fmt.Println("Invalid input detected (correctly checked)")
	}

	// Demo 2: Error type assertions
	err = findItem("document")

	// ISSUE: Type assertion instead of errors.As
	notFoundErr, ok := err.(*NotFoundError)
	if ok {
		fmt.Printf("Not found error: %s\n", notFoundErr.Item)
	}

	// Correct way - FIX: use *NotFoundError instead of NotFoundError
	var notFound *NotFoundError
	if errors.As(err, &notFound) {
		fmt.Printf("Not found error (correctly checked): %s\n", notFound.Item)
	}

	// Demo 3: Switch on error value
	err = processData()

	// ISSUE: Switch on error value
	switch err {
	case ErrInvalidInput:
		fmt.Println("Invalid input")
	case ErrTimeout:
		fmt.Println("Timeout")
	default:
		fmt.Println("Unknown error")
	}

	// Demo 4: Switch on error type
	// ISSUE: Type switch instead of errors.As
	switch e := err.(type) {
	case *NotFoundError:
		fmt.Printf("Not found: %s\n", e.Item)
	default:
		fmt.Println("Other error type")
	}

	// Demo 5: fmt.Errorf without %w
	inputErr := errors.New("input validation failed")

	// ISSUE: Using fmt.Errorf without %w
	wrappedErr := fmt.Errorf("operation failed: %w", inputErr)
	fmt.Println(wrappedErr)

	// Correct way
	properlyWrappedErr := fmt.Errorf("operation failed: %w", inputErr)
	fmt.Println(properlyWrappedErr)

	// Demo 6: Multiple errors in fmt.Errorf
	err1 := errors.New("first error")
	err2 := errors.New("second error")

	// ISSUE: Multiple %v in fmt.Errorf
	combinedErr := fmt.Errorf("multiple errors: %w and %w", err1, err2)
	fmt.Println(combinedErr)

	// Demo 7: Special case with sql.ErrNoRows
	if openDbErr() == sql.ErrNoRows {
		// This is actually allowed by the linter because sql.ErrNoRows is documented
		// to be returned unwrapped
		fmt.Println("No rows found")
	}

	// Demo 8: Custom functions returning errors
	if err := customOperation(); err != nil {
		fmt.Println("Custom operation failed:", err)
	}

	// Just to use all the variables
	_ = wrappedErr
	_ = properlyWrappedErr
	_ = combinedErr
}

func openDbErr() error {
	return sql.ErrNoRows
}

func customOperation() error {
	file, err := os.Open("nonexistent.txt")
	if err != nil {
		// ISSUE: Not using %w
		return fmt.Errorf("could not open file: %w", err)
	}
	defer file.Close()

	data := make([]byte, 100)
	_, err = file.Read(data)
	if err != nil && err != io.EOF {
		// ISSUE: Direct comparison with io.EOF (though this one is allowed)
		return fmt.Errorf("could not read file: %w", err)
	}

	// ISSUE: Using strings.Contains instead of errors.Is
	if strings.Contains(err.Error(), "permission denied") {
		return fmt.Errorf("permission issue: %w", err)
	}

	return nil
}
