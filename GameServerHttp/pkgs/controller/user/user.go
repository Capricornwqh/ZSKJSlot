package controller_user

import (
	model "SlotGameServer/pkgs/model/user"
	service "SlotGameServer/pkgs/service/user"
	"SlotGameServer/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{
		userService: service,
	}
}

// 注册
func (c *UserController) SignUp(ctx *gin.Context) {
	req := &model.UserSignRequest{}
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}

	if len(req.Username) <= 0 {
		utils.HandleError(ctx, utils.ErrParameter)
		return
	}

	switch req.Typed {
	case model.ACCOUNTTYPE_EMAIL:
		// 邮箱
		resp, err := c.userService.EmailSignUp(ctx, req)
		if err != nil {
			utils.HandleError(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, resp)
		return
	case model.ACCOUNTTYPE_PHONE:
		// 手机
	default:
	}

	utils.HandleError(ctx, utils.ErrParameter)
}

// 登录
func (c *UserController) SignIn(ctx *gin.Context) {
	req := &model.UserSignRequest{}
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}

	switch req.Typed {
	case model.ACCOUNTTYPE_ACCOUNT:
		// 账号
		resp, err := c.userService.AccountSignIn(ctx, req)
		if err != nil {
			utils.HandleError(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, resp)
		return
	case model.ACCOUNTTYPE_EMAIL:
		// 邮箱
		resp, err := c.userService.EmailSignIn(ctx, req)
		if err != nil {
			utils.HandleError(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, resp)
		return
	case model.ACCOUNTTYPE_PHONE:
		// 手机
	default:
	}

	utils.HandleError(ctx, utils.ErrParameter)
}

// 验证码
func (c *UserController) VerifyCode(ctx *gin.Context) {
	req := &model.UserVerifyRequest{}
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}

	switch req.Typed {
	case model.ACCOUNTTYPE_EMAIL:
		//邮箱
		err = c.userService.VerifyEmail(ctx, req.Language, req.Object)
		if err != nil {
			utils.HandleError(ctx, err)
			return
		}
	case model.ACCOUNTTYPE_PHONE:
		//手机
	default:
		utils.HandleError(ctx, utils.ErrParameter)
		return

	}

	ctx.JSON(http.StatusOK, gin.H{})
}

// 设置location
func (c *UserController) SetLocation(ctx *gin.Context) {
	req := &model.UserProfile{}
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}

	if len(req.Country) <= 0 || len(req.Language) <= 0 {
		utils.HandleError(ctx, utils.ErrParameter)
		return
	}
}

// 获取profile
func (c *UserController) GetProfile(ctx *gin.Context) {
	resp, err := c.userService.GetProfile(ctx)
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
