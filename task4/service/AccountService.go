package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"task4/common"
	"task4/dao"
	"task4/domain"
	"time"
)

type UserDTO struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	Email    string `form:"email"`
}

// Register go docs
// @Summary      注册新用户
// @Description  注册新用户
// @Tags		 公开API
// @Accept       multipart/form-data
// @Produce      json
// @Param		 username formData  string true "用户名" default(admin)
// @Param		 password formData  string true "密码" default(123456)
// @Param		 email formData  string false "邮箱" default(119@qq.com)
// @Success      200  {object}  instance.SuccessResponse
// @Failure      400  {object}  instance.FailedResponse
// @Failure      500  {object}  instance.ErrorResponse
// @Router       /public/register [post]
func Register(c *gin.Context) {
	userDTO := &UserDTO{}
	err := c.ShouldBind(userDTO)
	if err != nil {
		panic(common.Failed(err.Error()))
	}
	// 加密密码
	hashedPassword, err1 := bcrypt.GenerateFromPassword([]byte(userDTO.Password), bcrypt.DefaultCost)
	if err1 != nil {
		panic(common.Failed(err1.Error()))
	}
	userDTO.Password = string(hashedPassword)

	//查询用户是否存在
	if exists := dao.QueryUserCount(userDTO.Username); exists {
		panic(common.Failed("用户已存在"))
	}

	//注册新用户
	user := &domain.User{
		Username: userDTO.Username,
		Password: userDTO.Password,
		Email:    userDTO.Email,
	}
	if errs := dao.CreateUser(user); errs != nil {
		panic(errs)
	}
	c.JSON(http.StatusOK, common.OK("注册成功"))
}

// Login go docs
// @Summary      用户登录
// @Description  用户登录
// @Tags		 公开API
// @Accept       multipart/form-data
// @Produce      json
// @Param		 username formData  string true "用户名" default(admin)
// @Param		 password formData  string true "密码" default(123456)
// @Success      200  {object}  instance.SuccessResponse
// @Failure      400  {object}  instance.FailedResponse
// @Failure      401  {object}  instance.UnauthorizedResponse
// @Failure      500  {object}  instance.ErrorResponse
// @Router       /public/login [post]
func Login(c *gin.Context) {
	userDTO := &UserDTO{}
	err := c.ShouldBind(userDTO)
	if err != nil {
		panic(common.Failed(err.Error()))
	}
	//查询用户信息
	var storeUser *domain.User
	var err1 error
	if storeUser, err1 = dao.QueryUserByName(userDTO.Username); err1 != nil {
		if errors.Is(err1, gorm.ErrRecordNotFound) {
			panic(common.Failed("用户不存在"))
		}
		panic(err)
	}
	// 验证密码
	if err2 := bcrypt.CompareHashAndPassword([]byte(storeUser.Password), []byte(userDTO.Password)); err2 != nil {
		panic(common.Unauthorized(err2.Error()))
	}

	// 生成 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       storeUser.ID,
		"username": storeUser.Username,
		"exp":      jwt.NewNumericDate(time.Now().Add(common.AppConfig.Duration)),
	})

	tokenString, err3 := token.SignedString([]byte(common.AppConfig.JWTSecret))
	if err3 != nil {
		panic(common.Failed(err3.Error()))
	}
	c.JSON(http.StatusOK, common.OK(tokenString))
}

// Logout 退出登录
func Logout(c *gin.Context) {
	//注销时,利用redis存储黑名单
}
