package dep_one

type User struct {
	UserName string
	Password string
}

func CreateUser(userName string) User {
	return User{
		UserName: userName,
		Password: "",
	}
}
