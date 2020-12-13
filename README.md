go-burner-email-providers
----
[![Go Reference](https://pkg.go.dev/badge/github.com/lindell/go-burner-email-providers/burner.svg)](https://pkg.go.dev/github.com/lindell/go-burner-email-providers/burner)


Go package that detects burner (temporary) emails based on the community maintained [wesbos/burner-email-providers](https://github.com/wesbos/burner-email-providers) list.

This repository is synced daily against that list.

## Installation

```
go get github.com/lindell/go-burner-email-providers
```

## Usage

```go
import (
    "github.com/lindell/go-burner-email-providers/burner"
)

func main() {
	isBurnerEmail := burner.IsBurnerEmail("test@temp-mail.org")
	fmt.Println(isBurnerEmail) // true

	isBurnerEmail = burner.IsBurnerEmail("johan@gmail.com")
	fmt.Println(isBurnerEmail) // false

	isBurnerDomain := burner.IsBurnerDomain("temp-mail.org")
	fmt.Println(isBurnerDomain) // true

	isBurnerEmail = burner.IsBurnerDomain("gmail.com")
	fmt.Println(isBurnerEmail) // false
}
```

## Size

Since the list of domains is quite large, the binary size and memory usage is not insignificant.

The increase of using this package is:
| Where | Size diff |
| -| -|
| On Disc | 3.22 Mb |
| Memory | 4.77 Mb |

