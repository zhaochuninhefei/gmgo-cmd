package main

import (
	"fmt"
	"gitee.com/zhaochuninhefei/gmgo-cmd/key"
	"gitee.com/zhaochuninhefei/gmgo-cmd/pwd"
	"gitee.com/zhaochuninhefei/gmgo-cmd/version"
	"gitee.com/zhaochuninhefei/gmgo-cmd/x509"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "gmgo-cmd",
		Short: "gmgo的命令行工具",
		Long:  "gmgo的命令行工具, 用于提供gmgo各种功能的命令行接口",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello from go-key!")
		},
	}

	// Add subcommand to root command
	rootCmd.AddCommand(version.VersionCmd())
	rootCmd.AddCommand(x509.X509Cmd())
	rootCmd.AddCommand(pwd.PwdCmd())
	rootCmd.AddCommand(key.KeyCmd())

	// Parse command line arguments
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
