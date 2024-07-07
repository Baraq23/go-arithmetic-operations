
## Go Arithmetic Operations

This is a simple Go program that performs basic arithmetic operations (addition, multiplication, and division) on two floating-point numbers provided as command-line arguments.
Requirements

```bash
    Go 1.16 or higher
```

# Installation

   - Clone the repository:

```bash
git clone https://github.com/yourusername/go-arithmetic-operations.git

cd go-arithmetic-operations
```

   - Ensure you have Go installed. If not, download and install it from the official website.

# Usage

- To run the program, use the following command:

```bash
go run . <num1> <num2>
```

- Replace <num1> and <num2> with the two numbers you want to perform arithmetic operations on. Both numbers must be valid floating-point numbers.

_Example_

```bash
go run . 10.5 2.3
```

_Output:_

```bash

The sum is: 12.800
The product is: 24.150
The quotient is: 4.565

```

# Functions

    Add(a, b float64) float64   - Returns the sum of a and b.

    Mult(a, b float64) float64  - Returns the product of a and b.

    Div(a, b float64) float64   - Returns the quotient of a divided by b.

    checkNum(a, b string) bool  - Checks if both a and b are valid numeric strings.