package x509

import (
	"fmt"
	"gitee.com/zhaochuninhefei/gmgo/x509"
	"github.com/spf13/cobra"
)

// gmgo-key x509 --text --in cert.pem
// gmgo-key x509 -t -i cert.pem
var x509Command = &cobra.Command{
	Use:   "x509",
	Short: "x509相关指令",
	Long:  `使用gmgo的x509相关指令`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if isText && certPath != "" {
			certificate, err := x509.ReadCertificateFromPemFile(certPath)
			if err != nil {
				return err
			}
			text, err := x509.CertificateText(certificate)
			if err != nil {
				return err
			}
			fmt.Println(text)
		} else {
			fmt.Println("缺少参数, 请使用`gmgo-key x509 --help`查看帮助信息")
		}

		return nil
	},
}

var isText bool
var certPath string

// X509Cmd returns the Cobra Command for x509
func X509Cmd() *cobra.Command {
	x509Command.Flags().BoolVarP(&isText, "text", "t", false, "输出证书的文本信息")
	x509Command.Flags().StringVarP(&certPath, "in", "i", "", "证书文件路径")
	return x509Command
}
