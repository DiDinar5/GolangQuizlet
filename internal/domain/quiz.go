package domain

type Question struct {
	ID      int
	Text    string
	Options []string
	Answer  int //
}

type QuizRepository interface {
	GetQuestions() ([]Question, error)
}

type QuizService interface {
	GetNextQuestion() (Question, error)
	CheckAnswer(questionID, answer int) (bool, error)
}
