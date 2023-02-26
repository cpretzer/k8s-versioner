package main

import (
	"fmt"
	"os"
)

func main() {
	message, exists := os.LookupEnv("K8S_VERSIONER_MESSAGE")

	if !exists {
		message = "XO ReiRay"
	}

	fmt.Fprintln(os.Stderr, fmt.Sprintf("Your message is %s", message))
}
