package utils

import (
	"BusinessServer/env"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/golang-jwt/jwt/v5"
	"io"
	"strconv"
	"strings"
	"time"
)

var node, _ = snowflake.NewNode(env.GetConfig().ClusterID)

// GenerateId 使用雪花算法生成全局唯一ID
func GenerateId() string {
	return node.Generate().String()
}

func GenerateCode() string {
	return "swust" + node.Generate().String()[11:]
}

func IsBlank(str *string) bool {
	if str == nil || strings.TrimSpace(*str) == "" {
		return true
	}
	return false
}

func ConvertStr2Int64(str string) int64 {
	atoi, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return int64(atoi)
}

func ConvertStr2Float32(str string) float32 {
	value, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return 0.0
	}
	return float32(value)
}

func GenerateSalt(length int) string {
	if length < 0 {
		length = 16
	}
	if length%2 != 0 {
		length += 1
	}
	b := make([]byte, length/2)
	dst := make([]byte, length)
	rand.Read(b)
	hex.Encode(dst, b)
	return string(dst)
}

func Md5Digest(password, salt string) string {
	h := md5.New()
	io.WriteString(h, password+salt)
	return fmt.Sprintf("%x", h.Sum(nil))
}

type CurrentUser struct {
	ID   string
	Name string
}
type MyClaims struct {
	UserName string `json:"userName"`
	UserId   string `json:"userId"`
	jwt.RegisteredClaims
}

func GenerateJWT(username string, userId string) (string, error) {
	expireTime := time.Duration(env.GetConfig().ExpireTime)
	claim := MyClaims{
		username,
		userId,
		jwt.RegisteredClaims{
			Issuer:    env.GetConfig().SignUser,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireTime * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()), // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()), // 生效时间
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	fmt.Println("密钥", env.GetConfig().Secrect)
	token, err := t.SignedString([]byte(env.GetConfig().Secrect))
	return token, err
}

func ParseToken(tokenString string) (*MyClaims, error) {
	t, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(env.GetConfig().Secrect), nil
	})
	if claims, ok := t.Claims.(*MyClaims); ok && t.Valid {
		return claims, nil
	}
	return nil, err
}
