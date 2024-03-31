package service

import (
	"GolangQuizlet/internal/domain"
	"context"
	"math/rand"
)

type QuizServiceImpl struct {
	repo domain.QuizRepository
}

func NewQuizService(repo domain.QuizRepository) domain.QuizService {
	return &QuizServiceImpl{repo}
}

func (s *QuizServiceImpl) GetNextQuestion() (domain.Question, error) {
	questions, err := s.repo.GetQuestions()
	if err != nil {
		return domain.Question{}, err
	}
	return questions[rand.Intn(len(questions))], nil
}

func (s *QuizServiceImpl) CheckAnswer(questionID, answer int) (bool, error) {
	questions, err := s.repo.GetQuestions()
	if err != nil {
		return false, err
	}

	for _, question := range questions {
		if question.ID == questionID {
			return question.Answer == answer, nil
		}
	}

	return false, nil
}
func (s *QuizServiceImpl) InsertQuestion(ctx context.Context, input domain.Question) error {
	return s.repo.InsertQuestion(ctx, input)
}
