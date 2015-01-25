package main

import (
	"fmt"
	"github.com/caarlos0/go-playground/essential-go/pages"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Not enough args")
	}
	filename := os.Args[1]
	p, err := pages.NewPage(filename)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(p)
}
