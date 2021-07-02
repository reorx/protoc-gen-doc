package main

import (
	"flag"
	"fmt"
	"os"
)

type Flags struct {
	configFile string
}

var f = Flags{}

func init() {
	flag.StringVar(&f.configFile, "f", "", "config file path")
}

func main() {
	flag.Parse()

	if f.configFile == "" {
		fmt.Printf("error: missing argument '-f'\n")
		os.Exit(1)
	}
	fmt.Println(f.configFile)
}
