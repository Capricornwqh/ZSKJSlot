package utils_middleware

import (
	"SlotGameServer/utils"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

const CONTEXT_USERID = "userId"

const TOKEN_EXPIRES = 7 * 24 * time.Hour

const TOKEN_ISSUER = "KongMing"

var jwtSecret = []byte("qHS~DkG6P}x3uZ")

type TokenClaims struct {
	UserId uint64 `json:"userId"`
	jwt.RegisteredClaims
}

// 认证，对以authorization为头部，形式为`bearer token`的Token进行验证
func WeakAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 从请求头中获取token
		token := ctx.GetHeader("Authorization")
		if token != "" {
			// 验证token
			userId, err := ParseToken(token)
			if err == nil {
				ctx.Set(CONTEXT_USERID, userId)
			}
		}

		ctx.Next()
	}
}

// 强制认证，对以authorization为头部，形式为`bearer token`的Token进行验证
func ForceAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 从请求头中获取token
		token := ctx.GetHeader("Authorization")
		if token == "" {
			logrus.WithContext(ctx).Error("token is missing")
			utils.HandleError(ctx, utils.ErrToken)
			return
		}

		// 验证token
		userId, err := ParseToken(token)
		if err != nil {
			logrus.WithContext(ctx).Error(err)
			utils.HandleError(ctx, utils.ErrToken)
			return
		}
		ctx.Set(CONTEXT_USERID, userId)

		ctx.Next()
	}
}

// 生成token
func GenerateToken(userId uint64, times time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, TokenClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(times.Add(TOKEN_EXPIRES)),
			Issuer:    TOKEN_ISSUER,
		},
	})

	tmpToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", utils.ErrParse
	}

	return tmpToken, nil
}

// 解析token
func ParseToken(strToken string) (uint64, error) {
	// 检查Bearer token
	parts := strings.SplitN(strToken, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		return 0, utils.ErrToken
	}

	// 解析JWT token
	token, err := jwt.Parse(parts[1], func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, utils.ErrToken
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return 0, utils.ErrToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, utils.ErrToken
	}

	userId, ok := claims["userId"].(float64)
	if !ok {
		return 0, utils.ErrToken
	}

	return uint64(userId), nil
}
