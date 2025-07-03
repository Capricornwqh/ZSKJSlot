package utils

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	mathRand "math/rand"
	"net"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"unicode"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// 字符串转int
func StringToInt(value string) int {
	result, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return 0
	}

	return int(result)
}

// 可用端口
func FindAvailablePort(startPort int) (int, error) {
	for port := startPort; port < startPort+200; port++ {
		conn, err := net.Dial("tcp", "localhost:"+strconv.Itoa(port))
		if err != nil {
			return port, nil
		}
		conn.Close()
	}
	return 0, ErrResource
}

// 获取content-type
func GetContentType(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return "", err
	}

	return http.DetectContentType(buffer), nil
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

// 计算aspect ration
func CalculateAspectRation(width, height int) (int, int) {
	g := gcd(width, height)
	return width / g, height / g
}

// 计算scale
func CalculateScale(tlWidth, tlHeight, screenWidth, screenHeight float64) float64 {
	scaleWidth := screenWidth / tlWidth
	scaleHeight := screenHeight / tlHeight

	return math.Min(scaleWidth, scaleHeight)
}

// 生成id
func GenerateId() string {
	return strings.ReplaceAll(uuid.NewString(), "-", "")
}

// 生成验证码
func VerifyCode() (int, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(900000))
	if err != nil {
		return 0, err
	}

	return int(n.Int64()) + 100000, nil
}

// 获取当前调用
func GetCaller() string {
	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		return details.Name()
	}

	return ""
}

// 生成随机密码
func GeneratePassword(length int) (string, error) {
	const (
		lowercase = "abcdefghijklmnopqrstuvwxyz"
		uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		digits    = "0123456789"
		special   = "!@#$%^&*()"
	)
	all := lowercase + uppercase + digits + special

	if length > len(all) {
		return "", fmt.Errorf("length must be less than or equal to %d", len(all))
	}

	password := make([]byte, length)
	used := make(map[byte]bool)

	// 确保密码包含至少一个小写字母、大写字母、数字和特殊符号
	password[0] = lowercase[randInt(len(lowercase))]
	used[password[0]] = true
	password[1] = uppercase[randInt(len(uppercase))]
	used[password[1]] = true
	password[2] = digits[randInt(len(digits))]
	used[password[2]] = true
	password[3] = special[randInt(len(special))]
	used[password[3]] = true

	for i := 4; i < length; i++ {
		for {
			char := all[randInt(len(all))]
			if !used[char] {
				password[i] = char
				used[char] = true
				break
			}
		}
	}

	// 打乱密码
	mathRand.Shuffle(len(password), func(i, j int) {
		password[i], password[j] = password[j], password[i]
	})

	return string(password), nil
}

func randInt(max int) int {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		panic(err) // 处理错误
	}
	return int(n.Int64())
}

// 验证密码
func IsPassword(loginPasswd, userPasswd string) bool {
	if len(loginPasswd) == 0 && len(userPasswd) == 0 {
		return false
	}
	err := bcrypt.CompareHashAndPassword([]byte(userPasswd), []byte(loginPasswd))
	return err == nil
}

// 加密密码
func EncryptPassword(passwd string) (string, error) {
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	return string(hashPwd), err
}

// 判断是否为中文
func IsChinese(str string) bool {
	for _, v := range str {
		if unicode.Is(unicode.Han, v) {
			return true
		}
	}
	return false
}

func ArrayNotInArray(original []string, search []string) []string {
	var result []string
	originalMap := make(map[string]bool)
	for _, v := range original {
		originalMap[v] = true
	}
	for _, v := range search {
		if _, ok := originalMap[v]; !ok {
			result = append(result, v)
		}
	}
	return result
}

func UniqueArray[T comparable](input []T) []T {
	result := make([]T, 0, len(input))
	seen := make(map[T]bool, len(input))
	for _, element := range input {
		if !seen[element] {
			result = append(result, element)
			seen[element] = true
		}
	}
	return result
}

// uuid验证
func IsValidUUID(uuidStr string) bool {
	_, err := uuid.Parse(uuidStr)
	return err == nil
}
