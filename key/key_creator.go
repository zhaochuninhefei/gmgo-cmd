package key

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/spf13/cobra"
)

// 用cobra定义一个命令行,用于生成密钥
var keyCommand = &cobra.Command{
	Use:   "key",
	Short: "随机密钥生成器",
	Long:  `使用gmgo的随机密钥生成器`,
	Run: func(cmd *cobra.Command, args []string) {
		if length > 0 {
			key := generateKey(length)
			println(key)
		} else {
			println("缺少参数, 请使用`gmgo-key key --help`查看帮助信息")
		}

	},
}

// 定义length并从cobra命令行参数中读取赋值

var length int

// KeyCmd returns the Cobra Command for key
func KeyCmd() *cobra.Command {
	keyCommand.Flags().IntVarP(&length, "length", "l", 32, "Key length")
	return keyCommand
}

// 定义generateKey函数，根据传入的长度生成一个随机密钥
func generateKey(length int) string {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(randomBytes)[:length]
}
