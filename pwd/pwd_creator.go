package pwd

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/spf13/cobra"
)

// 用cobra定义一个命令行, 用于生成口令
// gmgo-key pwd --length 16
// gmgo-key pwd -l 16
var pwdCommand = &cobra.Command{
	Use:   "pwd",
	Short: "口令生成器",
	Long:  `使用gmgo的口令生成器`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if length > 0 {
			pwd := GeneratePassword(length)
			println(pwd)
		} else {
			println("缺少参数, 请使用`gmgo-key pwd --help`查看帮助信息")
		}

		return nil
	},
}

var length int

// PwdCmd returns the Cobra Command for pwd
func PwdCmd() *cobra.Command {
	pwdCommand.Flags().IntVarP(&length, "length", "l", 0, "口令长度")
	return pwdCommand
}

func GeneratePassword(length int) string {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(randomBytes)[:length]
}
