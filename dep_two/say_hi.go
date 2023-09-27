package dep_two

import (
	"log"
)

func SayHi(name string, id string) {
	log.Printf("Hi, %v with id [%v]", name, id)
}
