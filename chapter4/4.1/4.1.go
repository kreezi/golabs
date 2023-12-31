package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "ch03/ex12: must have 2 arguments.")
		os.Exit(1)
	}
	fmt.Printf("%d\n", sha256PopCount(os.Args[1], os.Args[2]))
}

func sha256PopCount(a, b string) int {
	digesta := sha256.Sum256([]byte(a))
	digestb := sha256.Sum256([]byte(b))
	return popCount(digesta, digestb)
}

func popCount(a, b [32]byte) int {
	pop := 0
	for i := range a {
		pop += int(pc[a[i]^b[i]])
	}
	return pop
}

//go run 4.1.go str rts
