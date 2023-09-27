package dep_transitive

import (
	"github.com/ozgio/strutil"
	"log"
)

func GeneratePassword() string {
	id, err := strutil.Random("abcdefghik", 5)
	if err != nil {
		log.Fatal(err)
	}

	return id
}
