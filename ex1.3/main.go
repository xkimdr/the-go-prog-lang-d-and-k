package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	s1 := time.Now()
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Println(time.Since(s1))
	s2 := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(time.Since(s2))
}
