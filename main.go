package main

import (
	"flag"
	"log"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "lance", "用户名")
	flag.StringVar(&name, "n", "lance", "用户名")
	flag.Parse()
	log.Printf("name: %s", name)
}
