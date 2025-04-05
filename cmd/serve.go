/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"github.com/alynlin/myapi/pkg"
	"github.com/alynlin/myapi/pkg/controller"
	"github.com/alynlin/myapi/pkg/logging"
	sw "github.com/alynlin/myapi/pkg/model/v1"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap/zapcore"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called")

		log.Printf("Server started")
		router := createRouter()

		log.Fatal(router.Run(":8080"))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func createRouter() *gin.Engine {

	ctx := context.Background()

	apiLogger, err := logging.NewLogger().
		Field("pkg", "myapi").
		WithRequestId().
		Level(zapcore.DebugLevel.String()).
		Build(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't create configured logger: %v\n", err)
		os.Exit(1)
	}

	user_service := pkg.UserService(apiLogger)

	routes := sw.ApiHandleFunctions{
		UserAPI: &controller.UserAPIController{
			Service: user_service,
		},
	}

	router := sw.NewRouter(routes)

	return router
}
