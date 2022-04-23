package commands

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"

	"github.com/mimatache/birthdaytrivia/internal/commands/about"
	"github.com/mimatache/birthdaytrivia/internal/logger"
	"github.com/mimatache/birthdaytrivia/internal/web"
	"github.com/mimatache/birthdaytrivia/photos"
	"github.com/mimatache/birthdaytrivia/trivia"
	"github.com/mimatache/birthdaytrivia/trivia/api"
)

// Root is the start of the application. All other commands should either be added to this command, or to a command that is in intself added to this command
func Root() *cobra.Command {
	var port int

	rootCommand := &cobra.Command{
		Use:   "birthdaytrivia",
		Short: "Hello Mariashi, would you like to play a game?",
		RunE: func(cmd *cobra.Command, args []string) error {
			log, sync, err := logger.New(true)
			if err != nil {
				return err
			}
			defer func() {
				_ = sync()
			}()
			tg, err := trivia.DefaultGame()
			if err != nil {
				return err
			}
			game := api.New(tg)

			router := mux.NewRouter()
			router.Use(logger.HTTPLogging(log))
			subr := router.PathPrefix("/api/v1").Subrouter()
			game.Register(subr)
			photos.Register(subr)

			webApp, err := web.GetUIHandler()
			if err != nil {
				return fmt.Errorf("start: %w", err)
			}
			fmt.Printf(
				`
🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂
🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂
🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂
Hello si bine ai venit la un mic joc de trivia pe care l-am facut noi pentru tine, cado de ziua ta normal !!!!!!!!!!!!
Ca sa incepi sa joci, deschide un broswer (te rog nu internet explorer, nu cred ca merge acolo 😬) pe http://localhost:%d si urmareste instructiunile
Spor la treaba!
🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂
🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂
🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂🎂
`, port,
			)
			router.PathPrefix("/").Handler(webApp)
			if err := http.ListenAndServe(fmt.Sprintf(":%d", port), router); err != nil && err != http.ErrServerClosed {
				return err
			}
			fmt.Println("stop received")
			fmt.Println("wg done")
			return err
		},
	}

	rootCommand.Flags().IntVarP(&port, "port", "p", 8080, "api server port")

	rootCommand.AddCommand(
		about.Command(),
	)

	return rootCommand
}
