package repository

import (
	"GolangQuizlet/internal/domain"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type QuizRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewQuizRepository(db *pgxpool.Pool) domain.QuizRepository {
	return &QuizRepositoryImpl{db: db}
}

func (repo *QuizRepositoryImpl) GetQuestions() ([]domain.Question, error) {
	rows, err := repo.db.Query(context.Background(), "SELECT id, text, options, answer FROM questions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var questions []domain.Question
	for rows.Next() {
		var q domain.Question
		var options []string
		err := rows.Scan(
			&q.ID,
			&q.Text,
			&q.Options,
			&q.Answer,
		)
		if err != nil {
			return nil, err
		}
		q.Options = options
		questions = append(questions, q)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return questions, nil

	/*return []domain.Question{
		{ID: 1, Text: "What does 'Go' stand for?", Options: []string{"Gopher", "GoLang", "Go Programming", "None of the above"}, Answer: 1},
	}, nil*/
}
func (repo *QuizRepositoryImpl) InsertQuestion(ctx context.Context, input domain.Question) error {
	sql := `INSERT INTO questions (text, options, answer) 
	VALUES ($1,$2,$3)`
	_, err := repo.db.Exec(ctx, sql, input.Text, input.Options, input.Answer)
	if err != nil {
		return fmt.Errorf("error saving question to the database: %w", err)
	}

	return nil
	/*args := []interface{}{
		input.ID,
		input.Text,
		input.Options,
		input.Answer,
	}
	db, err := repo.db.Exec(q)*/
}
