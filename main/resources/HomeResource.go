package resources

import (
	"fmt"
	"net/http"
)

func ServerHealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Looks good!")
}
