package run

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/mimatache/birthdaytrivia/trivia"
)

func Command() *cobra.Command {
	var questionsFiles []string

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
			q, hasNext := tg.Next()
			for hasNext {
				fmt.Println(q.Text)
				for i, v := range q.Answers {
					fmt.Printf("\t %d) %s \n", i+1, v.Text)
				}
				fmt.Print("Answer:  ")
				reader := bufio.NewReader(os.Stdin)
				text, err := reader.ReadString('\n')
				if err != nil {
					return err
				}
				// convert CRLF to LF
				text = strings.ReplaceAll(text, "\n", "")
				a, err := strconv.Atoi(text)
				if err != nil {
					return err
				}
				if tg.IsAnswerCorrect(a - 1) {
					fmt.Println("Correct")
					q, hasNext = tg.Next()
					continue
				}
				fmt.Println("Incorrect")
			}

			return nil
		},
	}

	cmd.Flags().StringSliceVarP(&questionsFiles, "file", "f", []string{"questions/_template.yaml"}, "file containing the questions")
	return cmd
}
