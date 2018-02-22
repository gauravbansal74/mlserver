package cmd

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/gauravbansal74/mlserver/pkg/database"
	"github.com/gauravbansal74/mlserver/pkg/logger"
	"github.com/gauravbansal74/mlserver/pkg/rabbitmq"
	"github.com/gauravbansal74/mlserver/watcher"
)

// watcherCmd represents the watcher command
var watcherCmd = &cobra.Command{
	Use:   "watcher",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// startCmd represents the watcher start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		rabbitmq.Init(&rabbitmq.Configuration{
			ListenHost: viper.GetString("rabbitmq.host"),
			ListenPort: viper.GetInt("rabbitmq.port"),
			Username:   viper.GetString("rabbitmq.username"),
			Password:   viper.GetString("rabbitmq.password"),
			Exchange:   viper.GetString("rabbitmq.exchange"),
		})
		watcher.Init(viper.GetString("datasource.folder"))
	},
}

// processCmd represents the watcher process command
var processCmd = &cobra.Command{
	Use:   "process",
	Short: "Run activity Runner",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		dbI := database.LoadConfig(viper.GetString("mongo.url"), viper.GetString("mongo.database"))
		dbI.Init()
		rabbitmq.Init(&rabbitmq.Configuration{
			ListenHost: viper.GetString("rabbitmq.host"),
			ListenPort: viper.GetInt("rabbitmq.port"),
			Username:   viper.GetString("rabbitmq.username"),
			Password:   viper.GetString("rabbitmq.password"),
			Exchange:   viper.GetString("rabbitmq.exchange"),
		})
		ticker := time.NewTicker(60 * time.Second)
		defer ticker.Stop()
		for ; ; <-ticker.C {
			rmq, err := rabbitmq.GetDeliveryChannel()
			if err != nil {
				logger.Error(err, "Error while connecting with rabbitMQ. Runner will try to connect again after 10 seconds.")
			}
			defer rmq.Close()
			watcher.Consumer(rmq, viper.GetString("datasource.folder")+"/")
		}
	},
}

func init() {
	rootCmd.AddCommand(watcherCmd)
	watcherCmd.AddCommand(startCmd)
	watcherCmd.AddCommand(processCmd)
}
