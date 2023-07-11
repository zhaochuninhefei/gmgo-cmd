package pwd

import (
	"crypto/rand"
	"fmt"
	"github.com/spf13/cobra"
	"math/big"
	"strings"
)

// 用cobra定义一个命令行, 用于生成口令
//  gmgo-key pwd --length 16
//  gmgo-key pwd -l 16
//  gmgo-key pwd --strength 1
//  gmgo-key pwd -s 1
var pwdCommand = &cobra.Command{
	Use:   "pwd",
	Short: "口令生成器",
	Long:  `使用gmgo的口令生成器`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if length >= 4 && strength >= 0 {
			pwd := GeneratePassword(length, strength)
			println(pwd)
		} else {
			println("参数不正确, 请使用`gmgo-key pwd --help`查看帮助信息")
		}

		return nil
	},
}

var length int
var strength int

// PwdCmd returns the Cobra Command for pwd
//goland:noinspection GoNameStartsWithPackageName
func PwdCmd() *cobra.Command {
	pwdCommand.Flags().IntVarP(&length, "length", "l", 0, "口令长度(至少为4)")
	pwdCommand.Flags().IntVarP(&strength, "strength", "s", 0, "口令强度(1:大小写字母+数字, 2:大小写字母+数字+特殊符号, 默认:2)")
	return pwdCommand
}

const (
	uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	digits           = "0123456789"
	specialChars     = "!@#$%^&*()-=_+[]{}|;':,.<>?/~`"
)

func GeneratePassword(length int, strength int) string {
	if length < 4 {
		panic(fmt.Errorf("口令长度必须大于3"))
	}

	// 定义字符集合
	charSets := make([]string, 0)
	charSets = append(charSets, uppercaseLetters)
	charSets = append(charSets, lowercaseLetters)
	charSets = append(charSets, digits)
	if strength != 1 {
		charSets = append(charSets, specialChars)
	}

	// 生成初始口令，至少包含一个字符类型的字符
	initialPassword := strings.Builder{}
	// 遍历所有字符集合
	for _, charSet := range charSets {
		// 从当前字符集合中随机选取一个字符
		char, err := getRandomCharacter(charSet)
		if err != nil {
			panic(err)
		}
		initialPassword.WriteString(char)
	}

	// 生成剩余长度的随机字符
	remainingLength := length - len(charSets)
	for i := 0; i < remainingLength; i++ {
		// 随机选取一种字符集合
		setIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charSets))))
		charSet := charSets[setIndex.Int64()]
		// 从当前字符集合中随机选取一个字符
		char, err := getRandomCharacter(charSet)
		if err != nil {
			panic(err)
		}
		initialPassword.WriteString(char)
	}

	// 将初始口令随机打乱
	runeArray := []rune(initialPassword.String())
	shuffleRuneArray(runeArray)

	return string(runeArray)
}

// 从字符集合中获取一个随机字符
func getRandomCharacter(charSet string) (string, error) {
	charSetLength := len(charSet)
	charIndex, err := rand.Int(rand.Reader, big.NewInt(int64(charSetLength)))
	if err != nil {
		return "", err
	}
	return string(charSet[charIndex.Int64()]), nil
}

// 随机打乱字符数组
func shuffleRuneArray(arr []rune) {
	n := len(arr)
	// 从尾部开始遍历，每个位置都和前面的某个随机位置互换
	for i := n - 1; i > 0; i-- {
		j, _ := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		arr[i], arr[j.Int64()] = arr[j.Int64()], arr[i]
	}
}
