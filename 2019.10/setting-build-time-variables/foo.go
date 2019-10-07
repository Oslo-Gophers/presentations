// Code use in the slides.
// Run with:
//   go run -ldflags "-X main.version=1.45" foo.go

package main

import "fmt"

var version string // empty, should be assigned wit go build -ldflags ..

func main() {
	fmt.Printf("Foobar version %q starting!\n", version)
}
