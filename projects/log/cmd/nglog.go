package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"log/internal/nglog"
)

var serviceName string

var logListCmd = &cobra.Command{
	Use: "list",
	Short: "获取服务列表",
	Long: "获取服务列表",
	Run: func(cmd *cobra.Command, args []string) {
		nglog.List()
	},
}

var downloadLogCmd = &cobra.Command{
	Use: "download",
	Short: "获取日志",
	Long: "获取指定服务的日志",
	Run: func(cmd *cobra.Command, args []string) {
		nglog.Download(serviceName)
		log.Println("write file")
	},
}

func init()  {
	downloadLogCmd.Flags().StringVarP(&serviceName, "service", "s", "", "要下载的服务的服务名称" )
}