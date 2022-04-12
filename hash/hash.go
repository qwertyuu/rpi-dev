package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

var (
	buildTime string
)

func main() {
	fmt.Printf("Build Time: %s\n", buildTime)
	h := md5.New()
	io.WriteString(h, "Francis")
	fmt.Printf("%x", h.Sum(nil))
	fmt.Printf("\n")
	fmt.Printf("233168")
}
