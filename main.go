package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
)

type Name string

func (n *Name) String() string {
	return fmt.Sprint(*n)
}
func (n *Name) Set(value string) error {
	if len(*n) > 0 {
		return errors.New("name flag already set")
	}
	*n = Name("lance:" + value)
	return nil
}
func main() {
	var name Name
	flag.Var(&name, "name", "帮助信息")
	flag.Parse()

	log.Printf("Name:%s", name)
}
