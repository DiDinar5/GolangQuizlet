package domain

import "context"

type Question struct {
	ID      int
	Text    string
	Options []string
	Answer  int
}

type QuizRepository interface {
	GetQuestions() ([]Question, error)
	InsertQuestion(ctx context.Context, input Question) error
}

type QuizService interface {
	GetNextQuestion() (Question, error)
	CheckAnswer(questionID, answer int) (bool, error)
	InsertQuestion(ctx context.Context, input Question) error
}
