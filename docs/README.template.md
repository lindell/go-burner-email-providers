go-burner-email-providers
----

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
    fmt.Println(isBurnerEmail)

    isBurnerDomain := burner.IsBurnerDomain("temp-mail.org")
    fmt.Println(isBurnerDomain)
}
```

## Size

Since the list of domains is quite large, the binary size and memory usage is not insignificant.

The increase of using this package is:
| Where | Size diff |
| -| -|
| On Disc | {{.SizeDiff}} |
| Memory | {{.MemDiff}} |

