package question

import (
	"errors"
	"fmt"
)

var (
	ErrNoTextProvided   = errors.New("no text provided")
	ErrNotEnoughAnswers = errors.New("at least two answers must be provided")
	ErrNoValidAnswer    = errors.New("at least one answer must be marked as correct")
)

type Question struct {
	Text    string   `yaml:"question"`
	Answers []Answer `yaml:"answers"`
	Image   string   `yaml:"image"`
}

func (q *Question) IsCorrect(answer int) bool {
	if len(q.Answers) <= answer || answer < 0 {
		return false
	}
	if !q.Answers[answer].IsTrue {
		return false
	}
	return true
}

func (q *Question) Validate() error {
	if q.Text == "" {
		return ErrNoTextProvided
	}
	if len(q.Answers) < 2 {
		return ErrNotEnoughAnswers
	}
	hasCorrectAnswer := false
	for _, a := range q.Answers {
		if err := a.Validate(); err != nil {
			return fmt.Errorf("invalid answer: %w", err)
		}
		if a.IsTrue {
			hasCorrectAnswer = a.IsTrue
		}
	}
	if !hasCorrectAnswer {
		return fmt.Errorf("%w; %s", ErrNoValidAnswer, q.Text)
	}
	return nil
}

type Answer struct {
	Text   string `yaml:"answer"`
	IsTrue bool   `yaml:"isTrue"`
}

func (a Answer) Validate() error {
	if a.Text == "" {
		return ErrNoTextProvided
	}
	return nil
}
