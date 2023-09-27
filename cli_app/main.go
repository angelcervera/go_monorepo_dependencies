package main

import (
	"github.com/angelcervera/go_monorepo_dependencies/dep_one"
	"github.com/angelcervera/go_monorepo_dependencies/dep_two"
)

func main() {
	user := dep_one.CreateUser("angelcervera")
	dep_two.SayHi(user.UserName, user.Password)
}
