package main

import (
	"fmt"
	"github.com/polds/imgbase64"
)

// Open and convert a local file.
func main() {
	img := imgbase64.FromLocal("test.png")

	fmt.Println(img)
}
