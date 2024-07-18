# Excel to Struct Converter

This project automates the conversion of Excel data into Go structs without the need to manually hardcode each field.

## Features

- Automatic mapping of CSV columns to struct fields using struct tags.
- Utilizes Go's reflection capabilities to dynamically assign values.

## Example Struct Definition

```go

type Item struct {
	ID       int64    `excel:"id"`
	Name     string   `excel:"name"`
	Price    float64  `excel:"price"`
	Tags     []string `excel:"tags"`
	Optional *string  `excel:"optional"`
}
