package console

import (
	"github.com/Portfolio-Project-Sekuy/angkutin-be/internal/config"
	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// RootCmd root command
var RootCmd = &cobra.Command{
	Use:   "root",
	Short: "root command",
	Long:  "This is the root command",
}

// Execute execute root command
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}

func init() {
	config.GetConfig()
	setupLogger()
}

// setupLogger setup logrus logger
func setupLogger() {
	formatter := runtime.Formatter{
		ChildFormatter: &logrus.JSONFormatter{},
		Line:           true,
		File:           true,
	}

	if config.Env() == "development" {
		formatter = runtime.Formatter{
			ChildFormatter: &logrus.TextFormatter{
				ForceColors:   true,
				FullTimestamp: true,
			},
			Line: true,
			File: true,
		}
	}

	logrus.SetFormatter(&formatter)
	logrus.SetOutput(os.Stdout)

	logLevel, err := logrus.ParseLevel(config.LogLevel())
	if err != nil {
		logLevel = logrus.DebugLevel
	}

	logrus.SetLevel(logLevel)
}
