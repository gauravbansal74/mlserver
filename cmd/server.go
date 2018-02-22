package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/gauravbansal74/mlserver/server"
	"github.com/gauravbansal74/mlserver/server/handlers"
	"github.com/gauravbansal74/mlserver/server/route"

	"github.com/gauravbansal74/mlserver/pkg/database"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "to start mlserver",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// init database
		dbI := database.LoadConfig(viper.GetString("mongo.url"), viper.GetString("mongo.database"))
		dbI.Init()
		handlers.Load()
		server.Run(route.LoadHTTP(), route.LoadHTTPS(), server.Config{
			Hostname: viper.GetString("server.host"),
			HTTPPort: viper.GetInt("server.port"),
			UseHTTP:  true,
			UseHTTPS: false,
		})
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.PersistentFlags().IntP("port", "p", 8080, "Port on which the server should listen")
	serverCmd.PersistentFlags().StringP("host", "", "127.0.0.1", "Host/interface on which the server should listen")
	serverCmd.PersistentFlags().BoolP("debug", "", true, "Enable/Disable debugger")

	viper.BindPFlag("Server.port", serverCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("Server.host", serverCmd.PersistentFlags().Lookup("host"))
	viper.BindPFlag("Server.debug", serverCmd.PersistentFlags().Lookup("debug"))
}
