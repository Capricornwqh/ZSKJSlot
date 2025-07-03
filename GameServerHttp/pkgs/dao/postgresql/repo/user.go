package pgsql_repo

import (
	pgsql_entity "SlotGameServer/pkgs/dao/postgresql/entity"
	"SlotGameServer/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type userDBRepo struct {
	db *gorm.DB
}

type UserDBRepo interface {
	// 新增用户
	AddUser(ctx *gin.Context, user *pgsql_entity.User) error
	// 通过邮箱查询
	FindUserByEmail(ctx *gin.Context, email string) (*pgsql_entity.User, error)
	// 通过ID查询
	FindUserByID(ctx *gin.Context, userId uint64) (*pgsql_entity.User, error)
	// 更新用户登录
	UpdateUserSignIn(ctx *gin.Context, user *pgsql_entity.User) error
	// 批量获取用户
	BatchUserByUserId(ctx *gin.Context, userId []uint64) ([]*pgsql_entity.User, error)
	// 用户名是否存在
	ExistsUsername(ctx *gin.Context, username string) (bool, error)
}

func NewUserDBRepo(db *gorm.DB) UserDBRepo {
	return &userDBRepo{
		db: db,
	}
}

// 新增用户
func (r *userDBRepo) AddUser(ctx *gin.Context, user *pgsql_entity.User) error {
	if r.db == nil || user == nil || user.UId <= 0 || len(user.Pass) <= 0 {
		return utils.ErrParameter
	}

	err := r.db.WithContext(ctx).Create(user).Error
	if err != nil {
		return err
	}

	return nil
}

// 通过邮箱查询
func (r *userDBRepo) FindUserByEmail(ctx *gin.Context, email string) (*pgsql_entity.User, error) {
	if r.db == nil || !utils.IsEmail(email) {
		return nil, utils.ErrParameter
	}

	userInfo := &pgsql_entity.User{}
	err := r.db.WithContext(ctx).Where("email=?", email).First(&userInfo).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, utils.ErrDBNotRecord
		}
		return nil, err
	}

	return userInfo, nil
}

// 通过ID查询
func (r *userDBRepo) FindUserByID(ctx *gin.Context, userId uint64) (*pgsql_entity.User, error) {
	if r.db == nil || userId <= 0 {
		return nil, utils.ErrParameter
	}

	userInfo := &pgsql_entity.User{}
	err := r.db.WithContext(ctx).Where("id=?", userId).First(&userInfo).Error
	if err != nil {
		return nil, err
	}

	return userInfo, nil
}

// 更新用户登录
func (r *userDBRepo) UpdateUserSignIn(ctx *gin.Context, user *pgsql_entity.User) error {
	if r.db == nil || user == nil || user.UId <= 0 {
		return utils.ErrParameter
	}

	err := r.db.WithContext(ctx).Model(&pgsql_entity.User{UId: user.UId}).
		UpdateColumns(pgsql_entity.User{
			LastLoginDate: user.LastLoginDate,
			UpdatedAt:     user.UpdatedAt,
			IPInfo:        user.IPInfo,
			Country:       user.Country,
			Language:      user.Language,
			Device:        user.Device,
			Terminal:      user.Terminal,
		}).Error
	if err != nil {
		return err
	}

	return nil
}

// 批量获取用户
func (r *userDBRepo) BatchUserByUserId(ctx *gin.Context, userId []uint64) ([]*pgsql_entity.User, error) {
	if r.db == nil || len(userId) <= 0 {
		return nil, utils.ErrParameter
	}

	sliceUser := []*pgsql_entity.User{}
	err := r.db.WithContext(ctx).Where("state<?", pgsql_entity.UFsuspended).
		Find(&sliceUser, userId).Error
	if err != nil {
		return nil, err
	}

	return sliceUser, nil
}

// 用户名是否存在
func (r *userDBRepo) ExistsUsername(ctx *gin.Context, username string) (bool, error) {
	if r.db == nil || len(username) <= 0 {
		return false, utils.ErrParameter
	}

	var count int64
	err := r.db.WithContext(ctx).Model(&pgsql_entity.User{}).
		Where("username=?", username).Count(&count).Error
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}
