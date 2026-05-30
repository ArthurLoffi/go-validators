# go-validators

A lightweight Go package for common input validation using regular expressions.

## Installation

```bash
go get github.com/ArthurLoffi/go-validators
```

> Requires Go 1.22.2 or higher.

## Usage

```go
package main

import (
	"fmt"
	govalidators "github.com/ArthurLoffi/go-validators"
)

func main() {
	fmt.Println(govalidators.ValidateName("Arthur Loffi")) // true
	fmt.Println(govalidators.ValidateName("João Silva"))   // true
	fmt.Println(govalidators.ValidateName("John123"))      // false
	fmt.Println(govalidators.ValidateName(""))             // false
}
```

## Available Validators

| Function | Description | Accepts |
|---|---|---|
| `ValidateName(s string) bool` | Validates that a name contains only letters (including accented characters) and spaces | `"Arthur"`, `"João Silva"` |

### `ValidateName`

Validates that a string contains only alphabetic characters — including Latin accented characters (`À–ÿ`) — and spaces. Numbers, symbols, and special characters are rejected.

```go
govalidators.ValidateName("Maria")       // true
govalidators.ValidateName("José Santos") // true
govalidators.ValidateName("John_Doe")   // false
govalidators.ValidateName("123")         // false
```

## Project Structure

```
go-validators/
├── go.mod
└── validate.go
```

## Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request with new validators or improvements.

1. Fork the project
2. Create your feature branch (`git checkout -b feat/my-validator`)
3. Commit your changes (`git commit -m 'feat: add my-validator'`)
4. Push to the branch (`git push origin feat/my-validator`)
5. Open a Pull Request

## License

This project is licensed under the [MIT License](./LICENSE).

---

Made by [ArthurLoffi](https://github.com/ArthurLoffi)