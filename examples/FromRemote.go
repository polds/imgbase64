package main

import (
	"fmt"
	"github.com/polds/imgbase64"
)

// For this example a DefaultImage will not be set
func main() {
	img := imgbase64.FromRemote("https://s3.amazonaws.com/uploads.hipchat.com/46533/311662/7rcjJ1KHy128Dxv/transparent.png")

	fmt.Println(img)
}
