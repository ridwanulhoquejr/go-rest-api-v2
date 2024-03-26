package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ridwanulhoquejr/go-rest-api-v2/cmd/internal/comment"
)

type ApiError struct {
	Error      string `json:"error"`
	Details    string `json:"details"`
	StatusCode int    `json:"status_code"`
}

var (
	ErrInernalServer = "internal server error"
	ErrPathNotFound  = "path not found"
	ErrNotFound      = map[string]string{"error": "not found for the given id"}
	ErrUnprocessable = ApiError{
		Error:   "unprocessable entity",
		Details: "unable to process the entity",
	}
)

type CommentService interface {
	GetComment(ctx context.Context, ID string) (comment.Comment, error)
	PostComment(context.Context, comment.Comment) (comment.Comment, error)
	UpdateComment(ctx context.Context, ID string, newCmt comment.Comment) (comment.Comment, error)
	DeleteComment(ctx context.Context, ID string) error
	GetMultipleComment(ctx context.Context) ([]comment.Comment, error)
}

func (h *Handler) PostComment(
	w http.ResponseWriter,
	r *http.Request,
) {
	var cmt comment.Comment

	if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cmt, err := h.Service.PostComment(r.Context(), cmt)
	if err != nil {
		log.Println(err)
		return
	}

	if err := WriteJson(w, http.StatusOK, cmt); err != nil {
		panic(err)
	}

}

//? handler func for our route

// 1. getmultiple comments
func (h *Handler) GetMultipleComment(
	w http.ResponseWriter, r *http.Request) {

	cmts, err := h.Service.GetMultipleComment(r.Context())

	if err != nil {
		log.Println("failed to get multiple comments from service layer", err)
		return
	}

	if err := WriteJson(w, http.StatusOK, cmts); err != nil {
		log.Println("failed to write json response", err)
		return
	}
}

// get single comment
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		WriteJson(w, http.StatusNotFound, ErrPathNotFound)
		return
	}

	cmt, err := h.Service.GetComment(r.Context(), id)
	if err != nil {
		log.Println("failed to get comment from service layer", err)
		WriteJson(w, http.StatusInternalServerError, ErrNotFound)
		return
	}

	if err := WriteJson(w, http.StatusOK, cmt); err != nil {
		WriteJson(w, http.StatusInternalServerError, ErrInernalServer)
		return
	}
}

func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	var updatedCmt comment.Comment

	if err := json.NewDecoder(r.Body).Decode(&updatedCmt); err != nil {
		WriteJson(w, http.StatusBadRequest, ApiError{
			Error:      "unprocessable entity",
			Details:    "Server could not process the entity, check the request body",
			StatusCode: http.StatusUnprocessableEntity,
		})
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		WriteJson(w, http.StatusNotFound, ErrInernalServer)
		return
	}

	cmt, err := h.Service.UpdateComment(r.Context(), id, updatedCmt)

	if err != nil {
		WriteJson(w, http.StatusInternalServerError, ErrInernalServer)
		return
	}

	if err := WriteJson(w, http.StatusOK, cmt); err != nil {
		WriteJson(w, http.StatusInternalServerError, ErrInernalServer)
		return
	}

}

// delete a comment
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		WriteJson(w, http.StatusNotFound, ErrNotFound)
		return
	}

	if err := h.Service.DeleteComment(r.Context(), id); err != nil {
		WriteJson(w, http.StatusInternalServerError, ErrInernalServer)
		return
	}

	err := WriteJson(w, http.StatusOK, map[string]string{"result": "deleted comment successfully"})

	if err != nil {
		WriteJson(w, http.StatusInternalServerError, ErrInernalServer)
		return
	}
}

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
