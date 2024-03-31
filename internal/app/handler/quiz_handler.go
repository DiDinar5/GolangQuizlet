package handler

import (
	"GolangQuizlet/internal/domain"
	"encoding/json"
	"net/http"
)

type QuizHandler struct {
	service domain.QuizService
}

func NewQuizHandler(service domain.QuizService) *QuizHandler {
	return &QuizHandler{service}
}

func (h *QuizHandler) GetNextQuestion(w http.ResponseWriter, r *http.Request) {
	question, err := h.service.GetNextQuestion()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(question)
}

func (h *QuizHandler) CheckAnswer(w http.ResponseWriter, r *http.Request) {

}

func (h *QuizHandler) InsertQuestion(w http.ResponseWriter, r *http.Request) {
	var question domain.Question
	if err := json.NewDecoder(r.Body).Decode(&question); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.InsertQuestion(r.Context(), question); err != nil {
		http.Error(w, "Failed to save the question", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
