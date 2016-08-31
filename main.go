package main

import (
	"github.com/faerulsalamun/golang-boilerplate/server"
	"flag"
	"fmt"
	"os"
	"github.com/faerulsalamun/golang-boilerplate/config"
)

func main() {
	enviroment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*enviroment)
	server.Init()
}