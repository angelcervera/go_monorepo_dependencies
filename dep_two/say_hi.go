package dep_two

import (
	"log"
)

func SayHi(name string, password string) {
	log.Printf("Hi, %v with passwod [%v]", name, password)
}
