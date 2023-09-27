package function

import (
	"encoding/json"
	"fmt"
	"github.com/angelcervera/go_monorepo_dependencies/dep_one"
	"log"
	"net/http"

	_ "github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("CreateUser", createUser)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	userName := r.URL.Query().Get("user_name")
	user := dep_one.CreateUser(userName)
	body, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, string(body))
}
