package trivia

import (
	"errors"

	"gopkg.in/yaml.v3"

	"github.com/hashicorp/go-multierror"

	"github.com/mimatache/birthdaytrivia/trivia/question"
)

var ErrNoMoreQuestions = errors.New("no more questions")

type Game struct {
	Questions       []*question.Question `json:"questions"`
	currentQuestion int
}

type Option func(g *Game) error

func WithQuestions(questions []byte) Option {
	return func(g *Game) error {
		var errs error
		q := []*question.Question{}
		if err := yaml.Unmarshal(questions, &q); err != nil {
			return err
		}
		for _, v := range q {
			if err := v.Validate(); err != nil {
				errs = multierror.Append(errs, err)
			}
		}
		g.Questions = append(g.Questions, q...)
		return errs
	}
}

func NewGame(options ...Option) (*Game, error) {
	game := &Game{
		currentQuestion: 0,
	}
	var errs error
	for _, o := range options {
		if err := o(game); err != nil {
			errs = multierror.Append(errs, err)
		}
	}
	return game, errs
}

func (g *Game) GetCurrentQuestion() (*question.Question, error) {
	if len(g.Questions) > g.currentQuestion {
		return g.Questions[g.currentQuestion], nil
	}
	return nil, ErrNoMoreQuestions
}

// Next returns the next question in the list if it exists and a true. Otherwise returns and empty question and false
func (g *Game) Next() bool {
	ok := len(g.Questions)-1 > g.currentQuestion
	g.currentQuestion += 1
	return ok
}

func (g *Game) Reset() {
	g.currentQuestion = 0
}
