package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func SetUp() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/health", health)
	return router
}

// health check endpoint
func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "We healthly")
}
