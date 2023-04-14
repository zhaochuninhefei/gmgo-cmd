package version

import (
	"fmt"
	"github.com/spf13/cobra"
)

const (
	majorVersion string = "0" // 主版本号
	minorVersion string = "0" // 次版本号
	patchVersion string = "1" // 补丁版本号
)

// VersionCmd returns the Cobra Command for Version
func VersionCmd() *cobra.Command {
	return versionCommand
}

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "打印 gmgo-cmd 版本",
	Long:  `打印 gmgo-cmd 当前版本`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			return fmt.Errorf("version子命令不需要参数")
		}
		// Parsing of the command line is done so silence key usage
		cmd.SilenceUsage = true
		fmt.Print(GetVersionInfo())
		return nil
	},
}

// GetVersionInfo returns version information for the gmgo-key
func GetVersionInfo() string {
	return fmt.Sprintf("Version: v%s.%s.%s\n", majorVersion, minorVersion, patchVersion)
}
