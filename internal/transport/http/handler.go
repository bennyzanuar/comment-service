package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler - stores pointer to comments service
type Handler struct {
	Router *mux.Router
}

// NewHandler - return a pointer to a Handler
func NewHandler() *Handler {
	return &Handler{}
}

// SetupRoutes - sets up all the routes application
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "I am alive")
	})
}
