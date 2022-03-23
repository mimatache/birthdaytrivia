package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mimatache/birthdaytrivia/trivia/question"
)

type trivia interface {
	Next() bool
	GetCurrentQuestion() (*question.Question, error)
	Reset()
}

type answer struct {
	Index int `json:"index"`
}

type GameStatus struct {
	IsAnswerCorrect bool `json:"isAnswerCorrect"`
	HasNext         bool `json:"hasNext"`
}

type questionToShow struct {
	Question string   `json:"question"`
	Answers  []string `json:"answers"`
}

type api struct {
	trivia trivia
}

func New(trivia trivia) *api {
	return &api{
		trivia: trivia,
	}
}

func (a *api) Register(router *mux.Router) {
	router.HandleFunc("/trivia/question", a.answerQuestion).Methods(http.MethodPost)
	router.HandleFunc("/trivia/question", a.getNextQuestion).Methods(http.MethodGet)
	router.HandleFunc("/trivia/reset", a.reset).Methods(http.MethodPost)
}

// answerQuestion checks whether the sent answer is correct
func (a *api) answerQuestion(w http.ResponseWriter, r *http.Request) {
	userAnswer := answer{}
	body := r.Body
	defer body.Close()
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&userAnswer)
	if err != nil {
		SendError(w, "could not decode the body", http.StatusInternalServerError)
		return
	}
	q, err := a.trivia.GetCurrentQuestion()
	if err != nil {
		SendError(w, err.Error(), http.StatusBadRequest)
		return
	}
	Send(w,
		GameStatus{
			IsAnswerCorrect: q.IsCorrect(userAnswer.Index),
			HasNext:         a.trivia.Next(),
		},
		http.StatusOK,
	)
}

func (a *api) getNextQuestion(w http.ResponseWriter, r *http.Request) {
	q, err := a.trivia.GetCurrentQuestion()
	if err != nil {
		SendError(w, err.Error(), http.StatusBadRequest)
		return
	}
	userQuestion := &questionToShow{
		Question: q.Text,
		Answers:  make([]string, len(q.Answers)),
	}
	for i, v := range q.Answers {
		userQuestion.Answers[i] = v.Text
	}
	Send(w, userQuestion, http.StatusOK)
}

func (a *api) reset(w http.ResponseWriter, r *http.Request) {
	a.trivia.Reset()
	w.WriteHeader(http.StatusOK)
}

type executionError struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

// SendError returns an error message with the given status code
func SendError(w http.ResponseWriter, message string, code int) {
	execError := &executionError{
		Error: message,
		Code:  code,
	}

	js, err := json.Marshal(execError)
	if err != nil {
		SendError(w, err.Error(), 400)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(js)
}

// Send writes a json body to the response writter
func Send(w http.ResponseWriter, value interface{}, code int) {
	js, err := json.Marshal(value)
	if err != nil {
		SendError(w, err.Error(), 400)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(js)
}
