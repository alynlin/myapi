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
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

		srv := &http.Server{
			Addr:    ":8080",
			Handler: router,
		}

		// 在协程中启动服务器
		go func() {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("Server error: %v", err)
			}
		}()

		// 监听终止信号
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		log.Println("Shutting down server...")

		// 设置 5 秒超时上下文
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// 优雅关闭服务器
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal("Server forced to shutdown:", err)
		}

		select {
		case <-ctx.Done():
			//todo
			log.Println("Server exited")
		}
	},
}

func init() {
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
