package question_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mimatache/birthdaytrivia/trivia/question"
)

var (
	validAnswer   = question.Answer{Text: "good answer", IsTrue: true}
	invalidAnswer = question.Answer{Text: "bad answer", IsTrue: false}
	noTextAnswer  = question.Answer{}
)

func Test_Question_IsCorrect(t *testing.T) {
	tests := []struct {
		name    string
		answer  int
		correct bool
	}{
		{
			name:    "correct",
			answer:  0,
			correct: true,
		},
		{
			name:    "incorrect",
			answer:  1,
			correct: false,
		},
		{
			name:    "out of range",
			answer:  2,
			correct: false,
		},
		{
			name:    "out of range negative",
			answer:  -1,
			correct: false,
		},
	}
	q := &question.Question{
		Text:    "Your question here?",
		Answers: []question.Answer{validAnswer, invalidAnswer},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.correct, q.IsCorrect(test.answer))
		})
	}
}

func Test_Question_Validate(t *testing.T) {
	tests := []struct {
		name        string
		question    question.Question
		expectedErr error
	}{
		{
			name: "correct",
			question: question.Question{
				Text:    "Your question here?",
				Answers: []question.Answer{validAnswer, invalidAnswer},
			},
			expectedErr: nil,
		},
		{
			name: "no enough answers",
			question: question.Question{
				Text:    "Your question here?",
				Answers: []question.Answer{validAnswer},
			},
			expectedErr: question.ErrNotEnoughAnswers,
		},
		{
			name: "no text",
			question: question.Question{
				Text:    "",
				Answers: []question.Answer{validAnswer, invalidAnswer},
			},
			expectedErr: question.ErrNoTextProvided,
		},
		{
			name: "no text answer",
			question: question.Question{
				Text:    "Your question here?",
				Answers: []question.Answer{validAnswer, invalidAnswer, noTextAnswer},
			},
			expectedErr: question.ErrNoTextProvided,
		},
		{
			name: "no correct answer",
			question: question.Question{
				Text:    "Your question here?",
				Answers: []question.Answer{invalidAnswer, invalidAnswer},
			},
			expectedErr: question.ErrNoValidAnswer,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.question.Validate()
			assert.ErrorIs(t, err, test.expectedErr)
		})
	}
}
