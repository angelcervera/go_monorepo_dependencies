package dep_one

import "fmt"

func SayHi(name string) string {
	return fmt.Sprintf("Hi, %v", name)
}
