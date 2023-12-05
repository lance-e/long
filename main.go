package main

import (
	"flag"
	"log"
)

var name string

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) <= 0 {
		return
	}
	switch args[0] {
	case "go":
		cmd := flag.NewFlagSet("go", flag.ExitOnError)
		cmd.StringVar(&name, "name", "go语言", "帮助信息")
		_ = cmd.Parse(args[1:])
	case "java":
		cmd := flag.NewFlagSet("java", flag.ExitOnError)
		cmd.StringVar(&name, "n", "java语言", "帮助信息")
		_ = cmd.Parse(args[1:])

	}
	log.Printf("Name:%s", name)
}
