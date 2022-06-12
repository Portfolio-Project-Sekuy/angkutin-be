package console

import (
	"fmt"
	"github.com/Portfolio-Project-Sekuy/angkutin-be/internal/delivery/graphqlsvc/graph/generated"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/Portfolio-Project-Sekuy/angkutin-be/internal/config"
	"github.com/Portfolio-Project-Sekuy/angkutin-be/internal/delivery/graphqlsvc/graph"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "server",
	Short: "run server",
	Long:  "this command is to start the server",
	Run:   run,
}

func init() {
	RootCmd.AddCommand(runCmd)
}

func run(cmd *cobra.Command, args []string) {
	errChan := make(chan error, 1)
	sigChan := make(chan os.Signal, 1)
	quitChan := make(chan bool, 1)

	go func() {
		for {
			select {
			case <-sigChan:
				quitChan <- true
			case e := <-errChan:
				logrus.Error(e)
				quitChan <- true
			}
		}
	}()

	server := echo.New()
	server.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	gqlConfig := generated.Config{Resolvers: &graph.Resolver{}}
	gqlHandler := handler.NewDefaultServer(generated.NewExecutableSchema(gqlConfig))
	server.POST("/query", func(c echo.Context) error {
		gqlHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	go func() {
		err := server.Start(fmt.Sprintf(":%s", config.HTTPPort()))
		if err != nil && err != http.ErrServerClosed {
			errChan <- err
		}
	}()

	<-quitChan
	logrus.Info("exiting...")
}
