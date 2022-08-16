package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	var name string                                                   // [1]
	flag.StringVar(&name, "name", "everyone", "The greeting object.") // [2]
	flag.Parse()

	fmt.Printf("Hello, %v!\n", name)

	var err error
	n, err := io.WriteString(os.Stdout, "Hello, everyone!\n")
	fmt.Println(n, err)

}
