go-burner-email-providers
----
[![Go Reference](https://pkg.go.dev/badge/github.com/lindell/go-burner-email-providers/burner.svg)](https://pkg.go.dev/github.com/lindell/go-burner-email-providers/burner)
[![Daily list sync](https://github.com/lindell/go-burner-email-providers/workflows/Daily%20list%20sync/badge.svg)](https://github.com/lindell/go-burner-email-providers/actions?query=workflow%3A%22Daily+list+sync%22)
[![Go Report Card](https://goreportcard.com/badge/github.com/lindell/go-burner-email-providers)](https://goreportcard.com/report/github.com/lindell/go-burner-email-providers)


Go package that detects burner (temporary) emails based on the community maintained [wesbos/burner-email-providers](https://github.com/wesbos/burner-email-providers) list. This repository is synced daily against that list.

It does currently contain 27,205 domains and the lookup is done with a hash set for instant results.

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
| On Disc | 0.77 Mb |
| Memory | 1.20 Mb |

