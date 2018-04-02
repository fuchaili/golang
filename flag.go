package main

import (
	"flag"
	"fmt"
)

func main() {
	fileName := flag.String("file", "fileName", "input a filename")
	flag.Parse()
	fmt.Printf("%s\n", *fileName)
	fmt.Printf("%T\n", fileName)
	fmt.Printf("%T\n", *fileName)
	fmt.Println("filename", fileName)

}
