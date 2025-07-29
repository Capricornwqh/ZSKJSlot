package service_user

import (
	i18n_repo "SlotGameServer/pkgs/dao/i18n/repo"
	entity_pgsql "SlotGameServer/pkgs/dao/postgresql/entity"
	repo_pgsql "SlotGameServer/pkgs/dao/postgresql/repo"
	entity_redis "SlotGameServer/pkgs/dao/redis/entity"
	repo_redis "SlotGameServer/pkgs/dao/redis/repo"
	model "SlotGameServer/pkgs/model/user"
	"SlotGameServer/utils"
	utils_middleware "SlotGameServer/utils/middleware"
	"reflect"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserService struct {
	VerifyCodeRedisRepo repo_redis.VerifyCodeRedisRepo
	UserDBRepo          repo_pgsql.UserDBRepo
	UserRedisRepo       repo_redis.UserRedisRepo
	EmailI18NRepo       i18n_repo.EmailI18NRepo
}

// 检查所有属性是否已初始化
func (s *UserService) CheckInitialization() {
	v := reflect.ValueOf(s).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.IsNil() {
			logrus.Fatalf("%s field %s is not initialized", t.Name(), t.Field(i).Name)
		}
	}
}

// 获取用户的redis数据
func (s *UserService) GetUser(ctx *gin.Context) (*entity_pgsql.Account, error) {
	tmpUserId := ctx.GetUint64(utils_middleware.CONTEXT_USERID)
	if tmpUserId <= 0 {
		return nil, utils.ErrUserNotFound
	}

	// 从redis中获取
	tmpUser, err := s.UserRedisRepo.GetUser(ctx, tmpUserId)
	if err != nil {
		if err == utils.ErrRedisNotKey {
			//如果没有从db中获取
			tmpUser, err = s.UserDBRepo.FindUserByID(ctx, tmpUserId)
			if err != nil {
				logrus.WithContext(ctx).Error(err)
				return nil, utils.ErrUserNotFound
			}

			//设置reids
			err = s.UserRedisRepo.SetUser(ctx, tmpUser, entity_redis.TimeUserBase*time.Second)
			if err != nil {
				logrus.WithContext(ctx).Error(err)
				return nil, utils.ErrOperation
			}
		} else {
			logrus.WithContext(ctx).Error(err)
			return nil, err
		}
	}

	return tmpUser, nil
}

// 邮箱验证
func (s *UserService) VerifyEmail(ctx *gin.Context, lang, email string) error {
	// 获取验证码
	code, err := s.VerifyCodeRedisRepo.GetVerifyCode(ctx, email)
	if err != nil {
		//生成验证码
		result, err := utils.VerifyCode()
		if err != nil {
			logrus.WithContext(ctx).Error(err)
			return utils.ErrParse
		}
		code = strconv.Itoa(result)
		err = s.VerifyCodeRedisRepo.SetVerifyCode(ctx, email, code, entity_redis.TimeVerifyCode*time.Second)
		if err != nil {
			logrus.WithContext(ctx).Error(err)
			return utils.ErrOperation
		}
	}

	title, body, err := s.EmailI18NRepo.GetEmailVerifyCode(ctx, lang, "博戏", code)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return utils.ErrOperation
	}

	err = utils.SendEmail(utils.Conf.Email.UserName, []string{email}, title, body)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return utils.ErrOperation
	}

	return nil
}

// 邮箱注册
func (s *UserService) EmailSignUp(ctx *gin.Context, req *model.UserSignRequest) (*model.UserBaseInfo, error) {
	//校验验证码
	tmpCode, err := s.VerifyCodeRedisRepo.GetDelVerifyCode(ctx, req.Object)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return nil, utils.ErrVerifyCode
	}

	if tmpCode != req.Token {
		logrus.WithContext(ctx).Error(utils.ErrVerifyCode)
		return nil, utils.ErrVerifyCode
	}

	//用户名是否存在
	exists, err := s.UserDBRepo.ExistsUsername(ctx, req.Username)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return nil, utils.ErrOperation
	}
	if exists {
		logrus.WithContext(ctx).Error(utils.ErrUsernameExists)
		return nil, utils.ErrUsernameExists
	}

	nowTime := time.Now()
	tmpUser, err := s.UserDBRepo.FindUserByEmail(ctx, req.Object)
	if err != nil && err != utils.ErrDBNotRecord {
		logrus.WithContext(ctx).Error(err)
		return nil, utils.ErrParameter
	}

	tmpPassword, err := utils.GeneratePassword(12)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return nil, utils.ErrOperation
	}

	//如果没有账号，直接注册
	tmpPassword, err = utils.EncryptPassword(tmpPassword)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return nil, utils.ErrOperation
	}
	tmpUser = &entity_pgsql.Account{
		DeletedAt:     nowTime,
		CreatedAt:     nowTime,
		UpdatedAt:     nowTime,
		LastLoginDate: nowTime,
		Email:         req.Object,
		Username:      req.Username,
		Status:        entity_pgsql.UFactivated,
		Pass:          tmpPassword,
		IPInfo:        ctx.ClientIP(),
		Device:        req.Device,
		Terminal:      req.Terminal,
		Birthday:      nowTime,
		Country:       req.Country,
		Language:      req.Language,
	}

	//插入数据库
	err = s.UserDBRepo.AddUser(ctx, tmpUser)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return nil, utils.ErrOperation
	}

	tmpToken, err := utils_middleware.GenerateToken(tmpUser.Id, nowTime)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return nil, utils.ErrToken
	}

	//设置reids
	err = s.UserRedisRepo.SetUser(ctx, tmpUser, entity_redis.TimeUserBase*time.Second)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return nil, utils.ErrOperation
	}

	tmpUserLoginResponse := &model.UserBaseInfo{
		UserId:        tmpUser.Id,
		CreatedAt:     tmpUser.CreatedAt,
		LastLoginDate: tmpUser.LastLoginDate,
		Username:      tmpUser.Username,
		EMail:         tmpUser.Email,
		Pass:          tmpUser.Pass,
		Avatar:        tmpUser.Avatar,
		Mobile:        tmpUser.Mobile,
		Country:       tmpUser.Country,
		Language:      tmpUser.Language,
		Status:        uint32(tmpUser.Status),
		Token:         tmpToken,
	}
	return tmpUserLoginResponse, nil

}

// 邮箱登录
func (s *UserService) EmailSignIn(ctx *gin.Context, req *model.UserSignRequest) (*model.UserBaseInfo, error) {
	//校验验证码
	tmpCode, err := s.VerifyCodeRedisRepo.GetDelVerifyCode(ctx, req.Object)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return nil, utils.ErrVerifyCode
	}

	if tmpCode != req.Token {
		logrus.WithContext(ctx).Error(utils.ErrVerifyCode)
		return nil, utils.ErrVerifyCode
	}

	nowTime := time.Now()
	tmpUser, err := s.UserDBRepo.FindUserByEmail(ctx, req.Object)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return nil, utils.ErrParameter
	}

	if tmpUser.Status > entity_pgsql.UFsuspended {
		logrus.WithContext(ctx).Error(utils.ErrAccountAvailable)
		return nil, utils.ErrAccountAvailable
	}

	if req.Passwd != tmpUser.Pass {
		logrus.WithContext(ctx).Error(utils.ErrUserPassword)
		return nil, utils.ErrUserPassword
	}

	//更新
	tmpUser.UpdatedAt = nowTime
	tmpUser.LastLoginDate = nowTime
	tmpUser.IPInfo = ctx.ClientIP()
	tmpUser.Device = req.Device
	tmpUser.Terminal = req.Terminal
	tmpUser.Country = req.Country
	tmpUser.Language = req.Language
	err = s.UserDBRepo.UpdateUserSignIn(ctx, tmpUser)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return nil, utils.ErrOperation
	}

	tmpToken, err := utils_middleware.GenerateToken(tmpUser.Id, nowTime)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return nil, utils.ErrToken
	}

	//设置reids
	err = s.UserRedisRepo.SetUser(ctx, tmpUser, entity_redis.TimeUserBase*time.Second)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return nil, utils.ErrOperation
	}

	tmpUserLoginResponse := &model.UserBaseInfo{
		UserId:        tmpUser.Id,
		CreatedAt:     tmpUser.CreatedAt,
		LastLoginDate: tmpUser.LastLoginDate,
		Username:      tmpUser.Username,
		EMail:         tmpUser.Email,
		Pass:          tmpUser.Pass,
		Avatar:        tmpUser.Avatar,
		Mobile:        tmpUser.Mobile,
		Country:       tmpUser.Country,
		Language:      tmpUser.Language,
		Status:        uint32(tmpUser.Status),
		Token:         tmpToken,
	}
	return tmpUserLoginResponse, nil
}

// 账号登录
func (s *UserService) AccountSignIn(ctx *gin.Context, req *model.UserSignRequest) (*model.UserBaseInfo, error) {
	//校验token
	tokenId, err := utils_middleware.ParseToken(ctx.GetHeader("Authorization"))
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return nil, utils.ErrToken
	}

	nowTime := time.Now()
	tmpId, err := strconv.ParseUint(req.Object, 10, 64)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return nil, utils.ErrToken
	}

	if tmpId <= 0 || tmpId != tokenId {
		return nil, utils.ErrToken
	}

	tmpUser, err := s.UserDBRepo.FindUserByID(ctx, tmpId)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return nil, utils.ErrUserNotFound
	} else {
		if tmpUser.Status > entity_pgsql.UFsuspended {
			logrus.WithContext(ctx).Error(utils.ErrAccountAvailable)
			return nil, utils.ErrAccountAvailable
		}

		if req.Passwd != tmpUser.Pass {
			logrus.WithContext(ctx).Error(utils.ErrUserPassword)
			return nil, utils.ErrUserPassword
		}

		//更新
		tmpUser.UpdatedAt = nowTime
		tmpUser.LastLoginDate = nowTime
		tmpUser.IPInfo = ctx.ClientIP()
		tmpUser.Device = req.Device
		tmpUser.Terminal = req.Terminal
		tmpUser.Country = req.Country
		tmpUser.Language = req.Language
		err = s.UserDBRepo.UpdateUserSignIn(ctx, tmpUser)
		if err != nil {
			logrus.WithContext(ctx).Error(err)
			return nil, utils.ErrOperation
		}
	}

	tmpToken, err := utils_middleware.GenerateToken(tmpUser.Id, nowTime)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return nil, utils.ErrToken
	}

	//设置reids
	err = s.UserRedisRepo.SetUser(ctx, tmpUser, entity_redis.TimeUserBase*time.Second)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return nil, utils.ErrOperation
	}

	tmpUserLoginResponse := &model.UserBaseInfo{
		UserId:        tmpUser.Id,
		CreatedAt:     tmpUser.CreatedAt,
		LastLoginDate: tmpUser.LastLoginDate,
		Username:      tmpUser.Username,
		EMail:         tmpUser.Email,
		Avatar:        tmpUser.Avatar,
		Mobile:        tmpUser.Mobile,
		Country:       tmpUser.Country,
		Language:      tmpUser.Language,
		Status:        uint32(tmpUser.Status),
		Token:         tmpToken,
	}
	return tmpUserLoginResponse, nil
}

// 获取用户信息
func (s *UserService) GetProfile(ctx *gin.Context) (*model.UserBaseInfo, error) {
	tmpUser, err := s.GetUser(ctx)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return nil, utils.ErrUserNotFound
	}
	if tmpUser.Status > entity_pgsql.UFsuspended {
		logrus.WithContext(ctx).Error(utils.ErrAccountAvailable)
		return nil, utils.ErrAccountAvailable
	}

	tmpUserBaseInfo := &model.UserBaseInfo{
		UserId:        tmpUser.Id,
		CreatedAt:     tmpUser.CreatedAt,
		LastLoginDate: tmpUser.LastLoginDate,
		Username:      tmpUser.Username,
		EMail:         tmpUser.Email,
		Avatar:        tmpUser.Avatar,
		Mobile:        tmpUser.Mobile,
		Country:       tmpUser.Country,
		Language:      tmpUser.Language,
		Status:        uint32(tmpUser.Status),
		Birthday:      tmpUser.Birthday,
	}

	return tmpUserBaseInfo, nil
}

//
