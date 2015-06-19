package main

import "fmt"

const (
	// Version
	MAJOR = 0
	MINOR = 1
	PATCH = 0
)

var version = fmt.Sprintf("%d.%d.%d", MAJOR, MINOR, PATCH)

func main() {
	fmt.Printf("Triage, v%s\n", version)
}
