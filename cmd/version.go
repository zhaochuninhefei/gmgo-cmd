package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

const (
	majorVersion string = "0" // 主版本号
	minorVersion string = "0" // 次版本号
	patchVersion string = "0" // 补丁版本号
)

// VersionCmd returns the Cobra Command for Version
func VersionCmd() *cobra.Command {
	return versionCommand
}

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Print gmgo-cmd version.",
	Long:  `Print current version of the gmgo-cmd.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			return fmt.Errorf("version子命令不需要参数")
		}
		// Parsing of the command line is done so silence cmd usage
		cmd.SilenceUsage = true
		fmt.Print(GetVersionInfo())
		return nil
	},
}

// GetVersionInfo returns version information for the gmgo-cmd
func GetVersionInfo() string {
	return fmt.Sprintf("Version: v%s.%s.%s\n", majorVersion, minorVersion, patchVersion)
}
