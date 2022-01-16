package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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

// SetupRoutes - sets up all the routes application
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up routes")
	h.Router = mux.NewRouter()

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

// GetComment - retrieve a comment by id
func (h *Handler) GetCommentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		sendErrorResponse(w, "Unable to parse UINT from id", err)
		return
	}

	comment, err := h.Service.GetComment(uint(i))
	if err != nil {
		sendErrorResponse(w, "Error retrieve comment by id", err)
		return
	}

	if err := sendOkResponse(w, comment); err != nil {
		panic(err)
	}
}

// GetAllCommentsHandler - retrieve all comment from the comment service
func (h *Handler) GetAllCommentsHandler(w http.ResponseWriter, r *http.Request) {
	comments, err := h.Service.GetAllComments()
	if err != nil {
		sendErrorResponse(w, "Failed to retrieve all comments", err)
		return
	}
	if err := sendOkResponse(w, comments); err != nil {
		panic(err)
	}
}

func (h *Handler) PostCommentHandler(w http.ResponseWriter, r *http.Request) {
	var comment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		sendErrorResponse(w, "Failed to decode json body", err)
		return
	}

	comment, err := h.Service.PostComment(comment)
	if err != nil {
		sendErrorResponse(w, "Failed to post new comment", err)
		return
	}
	if err := sendOkResponse(w, comment); err != nil {
		panic(err)
	}
}

func (h *Handler) UpdateCommentHandler(w http.ResponseWriter, r *http.Request) {
	var comment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		sendErrorResponse(w, "Failed to decode json body", err)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	commentID, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		sendErrorResponse(w, "Unable to parse UINT from id", err)
		return
	}

	comment, err = h.Service.UpdateComment(uint(commentID), comment)
	if err != nil {
		sendErrorResponse(w, "Failed to update comment", err)
		return
	}

	if err := sendOkResponse(w, comment); err != nil {
		panic(err)
	}
}

func (h *Handler) DeleteCommentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	commentID, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		sendErrorResponse(w, "Unable to parse UINT from id", err)
		return
	}

	err = h.Service.DeleteComment(uint(commentID))
	if err != nil {
		sendErrorResponse(w, "Failed to delete comment by id", err)
		return
	}

	err = sendOkResponse(w, Response{
		Message: "Successfully deleted comment",
	})
	if err != nil {
		panic(err)
	}
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
