package dep_transitive

import (
	"github.com/ozgio/strutil"
	"log"
)

func GeneratePassword() string {
	id, error := strutil.Random("abcdefghik", 5)
	if error != nil {
		log.Fatal(error)
	}

	return id
}
