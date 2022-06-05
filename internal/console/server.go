package console

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"

	"github.com/Portfolio-Project-Sekuy/angkutin-be/internal/config"
	"github.com/labstack/echo/v4"
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

	go func() {
		err := server.Start(fmt.Sprintf(":%s", config.HTTPPort()))
		if err != nil && err != http.ErrServerClosed {
			errChan <- err
		}
	}()

	<-quitChan
	logrus.Info("exiting...")
}
