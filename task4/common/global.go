package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Result 统一响应体
type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// UnauthorizedError 未授权异常
type UnauthorizedError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e *UnauthorizedError) Error() string {
	return fmt.Sprintf("%d %s", e.Code, e.Msg)
}

// BusinessError 业务异常
type BusinessError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e *BusinessError) Error() string {
	return fmt.Sprintf("%d %s", e.Code, e.Msg)
}

type GlobalError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Err  string `json:"error"`
}

func (e *GlobalError) Error() string {
	return fmt.Sprintf("%d %s : %v", e.Code, e.Msg, e.Err)
}

func OK(data interface{}) *Result {
	return &Result{
		Code: http.StatusOK,
		Msg:  "成功",
		Data: data,
	}
}

func Failed(msg string) *BusinessError {
	return &BusinessError{
		Code: http.StatusBadRequest,
		Msg:  msg,
	}
}

func Unauthorized(msg string) *UnauthorizedError {
	return &UnauthorizedError{
		Code: http.StatusUnauthorized,
		Msg:  msg,
	}
}

func GError(err string) *GlobalError {
	return &GlobalError{
		Code: http.StatusInternalServerError,
		Msg:  "Internal Server Error",
		Err:  err,
	}
}

type Config struct {
	JWTSecret string
	Duration  time.Duration
}

var AppConfig = Config{
	JWTSecret: "2025JWT_2290", // 从环境变量获取
	Duration:  24 * time.Hour, // JWT 有效期
}

type UserRealm struct {
	UserId   uint
	UserName string
}

func GetCurrentUser(c *gin.Context) *UserRealm {
	if userReal, exists := c.Get("USER_REAL"); exists {
		if realm, ok := userReal.(*UserRealm); ok {
			return realm
		}
	}
	return nil
}
