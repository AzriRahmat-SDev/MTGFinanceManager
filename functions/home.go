package functions

import (
	"fmt"
	"net/http"
)

//invokes the client side FEnd
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to home page on the client side")
}
