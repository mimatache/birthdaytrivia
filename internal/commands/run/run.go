package run

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"

	"github.com/mimatache/birthdaytrivia/trivia"
	"github.com/mimatache/birthdaytrivia/trivia/api"
)

func Command() *cobra.Command {
	var questionsFiles []string
	var port int

	cmd := &cobra.Command{
		Use:   "run",
		Short: "run the trivia game",
		RunE: func(cmd *cobra.Command, args []string) error {
			opts := []trivia.Option{}
			for _, questionsFile := range questionsFiles {
				qs, err := os.ReadFile(questionsFile)
				if err != nil {
					return err
				}
				opts = append(opts, trivia.WithQuestions(qs))
			}
			tg, err := trivia.NewGame(opts...)
			if err != nil {
				return err
			}
			game := api.New(tg)

			router := mux.NewRouter()
			game.Register(router.PathPrefix("/api/v1").Subrouter())
			if err := http.ListenAndServe(fmt.Sprintf(":%d", port), router); err != nil && err != http.ErrServerClosed {
				return err
			}
			fmt.Println("stop received")
			fmt.Println("wg done")
			return err
		},
	}

	cmd.Flags().IntVarP(&port, "port", "p", 8080, "api server port")
	cmd.Flags().StringSliceVarP(&questionsFiles, "file", "f", []string{"questions/_template.yaml"}, "file containing the questions")
	return cmd
}
