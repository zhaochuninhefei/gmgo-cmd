package main

import (
	"fmt"
	"os"

	"gitee.com/zhaochuninhefei/gmgo-cmd/cmd"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "gmgo-cmd",
		Short: "gmgo的命令行工具",
		Long:  "gmgo的命令行工具, 用于提供gmgo各种功能的命令行接口",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello from go-cmd!")
		},
	}

	// Add subcommand to root command
	rootCmd.AddCommand(cmd.VersionCmd())
	rootCmd.AddCommand(cmd.X509Cmd())

	// Parse command line arguments
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
