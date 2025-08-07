package common

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"strings"
)

// CustomRecovery  返回统一错误响应
func CustomRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 记录错误日志
				log.Printf("请求路径[%s] >>> %v\n", c.Request.URL, err)

				switch objErr := err.(type) {
				case *BusinessError:
					c.AbortWithStatusJSON(http.StatusBadRequest, objErr)
				case *UnauthorizedError:
					c.AbortWithStatusJSON(http.StatusUnauthorized, objErr)
				default:
					errors := GError(err.(error).Error())
					c.AbortWithStatusJSON(http.StatusInternalServerError, errors)
				}
			}
		}()
		c.Next()
	}
}

// JWTAuth 统一鉴权
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求头中的 Authorization 字段
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, Unauthorized("authorization header required"))
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
		// ValidateToken 验证 JWT Token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// 检查签名方法是否正确
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(AppConfig.JWTSecret), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, Unauthorized(err.Error()))
			c.Abort()
			return
		}

		// ExtractClaims 提取 JWT 声明信息
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userReal := &UserRealm{
				UserId:   uint(claims["id"].(float64)),
				UserName: claims["username"].(string),
			}
			c.Set("USER_REAL", userReal)
		} else {
			c.JSON(http.StatusUnauthorized, Unauthorized("invalid token"))
			c.Abort()
			return
		}
		c.Next()
	}
}
