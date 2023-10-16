package main

import (
	"fmt"
	"os"
)

func main() {
	name, found := os.LookupEnv("UTIL_NAME")
	if !found {
		name = "unknown"
	}

	fmt.Printf("Executing %s utility\n", name)
}
