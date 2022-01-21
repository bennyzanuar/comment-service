package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/bennyzanuar/comment-service/internal/comment"
	"github.com/gorilla/mux"
)

// Handler - stores pointer to comments service
type Handler struct {
	Router  *mux.Router
	Service *comment.Service
}

// Response - an object to store respose from this API
type Response struct {
	Message string
	Error   string
}

// NewHandler - return a pointer to a Handler
func NewHandler(service *comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"Method": r.Method,
			"Path":   r.URL.Path,
		}).Info("Handle request")
		next.ServeHTTP(rw, r)
	})
}

// SetupRoutes - sets up all the routes application
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up routes")
	h.Router = mux.NewRouter()
	h.Router.Use(LoggingMiddleware)

	h.Router.HandleFunc("/api/comment", h.GetAllCommentsHandler).Methods("GET")
	h.Router.HandleFunc("/api/comment", h.PostCommentHandler).Methods("POST")
	h.Router.HandleFunc("/api/comment/{id}", h.GetCommentHandler).Methods("GET")
	h.Router.HandleFunc("/api/comment/{id}", h.UpdateCommentHandler).Methods("PUT")
	h.Router.HandleFunc("/api/comment/{id}", h.DeleteCommentHandler).Methods("DELETE")

	h.Router.HandleFunc("/api/health", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
		rw.WriteHeader(http.StatusOK)
		err := json.NewEncoder(rw).Encode(Response{
			Message: "I am alive ",
		})
		if err != nil {
			panic(err)
		}
	})
}

func sendOkResponse(w http.ResponseWriter, resp interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(resp)
}

func sendErrorResponse(w http.ResponseWriter, message string, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(Response{Message: message, Error: err.Error()}); err != nil {
		panic(err)
	}
}
