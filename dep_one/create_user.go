package dep_one

import "github.com/angelcervera/go_monorepo_dependencies/dep_transitive"

type User struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

func CreateUser(userName string) User {
	return User{
		UserName: userName,
		Password: dep_transitive.GeneratePassword(),
	}
}
