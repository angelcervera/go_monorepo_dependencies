package main

import (
	"fmt"

	"github.com/angelcervera/go_monorepo_dependencies/dep_one"
)

func main() {
	message := dep_one.SayHi("Angel")
	fmt.Println(message)
}
