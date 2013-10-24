package main

import (
	"bytes"
	"fmt"
	"github.com/polds/imgbase64"
	"image/png"
	"os"
)

// Read from a bytes.Buffer.
// For simplicity I am skipping all error checks (bad idea)
func main() {
	file, _ := os.Open("test.png")

	m, _ := png.Decode(file)
	file.Close()

	// Do any adjustments to the image you need to
	// here such as resize and any filters
	// you might apply to the image

	var b bytes.Buffer
	png.Encode(&b, m)

	img := imgbase64.FromBuffer(b)
	fmt.Println(img)
}
