package main

import (
	"fmt"
	"runtime"

	"github.com/lindell/go-burner-email-providers/burner"
)

func main() {
	_ = burner.IsBurnerEmail("test@test.com")

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d\n", m.Alloc)
}
