# Go Error Linting Demo

A demo program to check/try out error linters. This repository demonstrates various error handling patterns in Go—both problematic ones and their correct alternatives—to test the capabilities of error linters like [go-errorlint](https://github.com/polyfloyd/go-errorlint).

## Goals

- Demonstrate common error handling anti-patterns in Go
- Show how to properly handle errors using the errors package introduced in Go 1.13
- Provide a test bed for error linting tools
- Illustrate how automatic fixes work with error linters

## Error Patterns Demonstrated

This demo includes examples of:

1. **Direct error comparisons** (`err == ErrSomething`) instead of `errors.Is()`
2. **Error type assertions** (`err.(*CustomError)`) instead of `errors.As()`
3. **Type switches** on errors instead of using `errors.As()`
4. **Missing error wrapping** with `fmt.Errorf()` using `%v` instead of `%w`
5. **String matching** (`strings.Contains(err.Error(), "text")`) instead of proper error handling
6. **Special cases** like handling of documented errors like `io.EOF` and `sql.ErrNoRows`

## Usage

### Building and Running

```bash
# Build the demo program
make build

# Run the demo program
make run
```

### Using the Error Linter

```bash
# Run the linter to check for issues
make lint

# Auto-fix issues where possible
make lint-fix
```

## What the Linter Will Find

The linter will detect issues like:

```go
// ISSUE: Direct comparison instead of errors.Is
if err == ErrInvalidInput {
    // ...
}

// ISSUE: Type assertion instead of errors.As
notFoundErr, ok := err.(*NotFoundError)

// ISSUE: Using %v instead of %w in fmt.Errorf
return fmt.Errorf("failed to process data: %v", err)
```

## Correct Patterns

The demo also shows correct error handling patterns:

```go
// Correct: Using errors.Is
if errors.Is(err, ErrInvalidInput) {
    // ...
}

// Correct: Using errors.As
var notFound *NotFoundError
if errors.As(err, &notFound) {
    // ...
}

// Correct: Using %w to wrap errors
return fmt.Errorf("operation failed: %w", err)
```

## License

This project is open source and available under the [MIT License](LICENSE).