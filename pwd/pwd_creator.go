package pwd

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"math/big"
	"os"
	"path"
	"strings"
)

// 用cobra定义一个命令行, 用于生成口令
//  gmgo-key pwd --length 16
//  gmgo-key pwd -l 16
//  gmgo-key pwd --strength 1
//  gmgo-key pwd -s 1
//  gmgo-key pwd --display test.com|testuser
//  gmgo-key pwd -d test.com|testuser
var pwdCommand = &cobra.Command{
	Use:   "pwd",
	Short: "口令生成器",
	Long:  `使用gmgo的口令生成器,支持大小写字母、数字和部分特殊符号(~!@#$%^&_-+=|:;)`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if length >= 4 && strength >= 0 {
			pwd := GeneratePassword(length, strength)
			println(pwd)
		} else if display != "" {
			// 查看口令
			pwd := displayPwd(display)
			println(pwd)
		} else {
			println("参数不正确, 请使用`gmgo-key pwd --help`查看帮助信息")
		}

		return nil
	},
}

var length int
var strength int
var saveKey string
var display string

// PwdCmd returns the Cobra Command for pwd
//goland:noinspection GoNameStartsWithPackageName
func PwdCmd() *cobra.Command {
	pwdCommand.Flags().IntVarP(&length, "length", "l", 0, "口令长度(至少为4)")
	pwdCommand.Flags().IntVarP(&strength, "strength", "s", 0, "口令强度(1:大小写字母+数字, 2:大小写字母+数字+特殊符号, 默认:2)")
	pwdCommand.Flags().StringVarP(&saveKey, "key", "k", "", "口令保存键值，格式: `用户名@目标域名`，如`testuser@test.com`")
	pwdCommand.Flags().StringVarP(&display, "display", "d", "", "显示口令，包含`@`时严格匹配口令键值，不包含`@`时作为口令键值的后缀查找")
	return pwdCommand
}

const (
	uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	digits           = "0123456789"
	specialChars     = "~!@#$%^&_-+=|:;"    // 去除了一些容易被误解或错误格式化的字符，比较全面的是 ~!@#$%^&*_-+=\|(){}[]:;"'<>,.?/`
	pwdFileDir       = ".gmgo-cmd/pwd.json" // 口令保存文件相对路径
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

	pwdStr := string(runeArray)

	// 判断是否传入了 key 参数
	if saveKey != "" {
		// 检查saveKey是否满足格式: `用户名@目标域名`
		if !strings.Contains(saveKey, "@") {
			panic(errors.New("口令保存键值需要满足格式: 用户名@目标域名"))
		}

		// 读取 PwdFileDir json文件内容，转为 map
		pwdMap, err := readPwdFile()
		if err != nil {
			panic(err)
		}
		// 将当前口令保存到 pwdMap 中
		pwdMap[saveKey] = pwdStr
		// 将pwdMap写入PwdFileDir文件
		err = writePwdFile(pwdMap)
		if err != nil {
			panic(err)
		}
	}
	return pwdStr
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

// 读取PwdFileDir json文件内容，转为 map
func readPwdFile() (map[string]string, error) {
	pwdMap := make(map[string]string)

	// 获取当前用户根目录
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	filePath := path.Join(homeDir, pwdFileDir)
	fileDir := path.Dir(filePath)

	// 如果fileDir不存在，或不是目录，则创建对应目录
	dirInfo, err := os.Stat(fileDir)
	if os.IsNotExist(err) || !dirInfo.IsDir() {
		if err = os.MkdirAll(fileDir, os.ModePerm); err != nil {
			return nil, err
		}
	}

	// 判断filePath是否存在且是文件
	fileInfo, err := os.Stat(filePath)
	if os.IsNotExist(err) || !fileInfo.Mode().IsRegular() {
		return pwdMap, nil
	}

	// 读取PwdFileDir文件内容(json格式)
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	// 将文件内容反序列化到map
	err = json.Unmarshal(content, &pwdMap)
	if err != nil {
		return nil, err
	}
	return pwdMap, nil
}

//  将map内容写入PwdFileDir文件
func writePwdFile(pwdMap map[string]string) error {
	// 将map序列化成json格式
	pwdJson, err := json.Marshal(pwdMap)
	if err != nil {
		return err
	}
	// 获取当前用户根目录
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	// 将序列化后的json写入PwdFileDir文件
	err = ioutil.WriteFile(path.Join(homeDir, pwdFileDir), pwdJson, 0644)
	if err != nil {
		return err
	}
	return nil
}

func displayPwd(dispaly string) string {
	// 获取当前用户根目录
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	filePath := path.Join(homeDir, pwdFileDir)
	// 读取PwdFileDir文件内容(json格式)
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	// 将文件内容反序列化到map
	var pwdMap map[string]string
	err = json.Unmarshal(content, &pwdMap)
	if err != nil {
		panic(err)
	}

	// 判断 dispaly 是否包含"@"
	if strings.Contains(dispaly, "@") {
		// 直接尝试从 pwdMap 中获取口令
		pwd, ok := pwdMap[dispaly]
		if ok {
			return pwd
		} else {
			panic(errors.New("口令不存在"))
		}
	} else {
		// 将 dispaly 作为口令键值的后缀，遍历pwdMap查找所有匹配的口令
		var pwdStrArr []string
		for key, value := range pwdMap {
			if strings.HasSuffix(key, dispaly) {
				pwdStrArr = append(pwdStrArr, key+":"+value)
			}
		}
		if len(pwdStrArr) == 0 {
			panic(errors.New("口令不存在"))
		} else {
			// 将 pwdStrArr 拼接为字符串返回
			return strings.Join(pwdStrArr, "\n")
		}
	}
}
