package main

import "fmt"

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
)

func main() {
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
}
