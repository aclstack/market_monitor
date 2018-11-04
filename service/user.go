package service

import (
	"github.com/JiangInk/market_monitor/extend/utils"
	"github.com/JiangInk/market_monitor/models"
	"github.com/rs/zerolog/log"
)

// UserService 用户服务层逻辑
type UserService struct{}

// QueryUser 查询用户
func (us UserService) QueryUser() (user models.User, err error) {
	return
}

// StoreUser 添加用户
func (us UserService) StoreUser(email string, pass string) (userID uint, err error) {
	log.Info().Msg("enter storeUser service")

	user := &models.User{
		Email:    email,
		UserName: email,
		Password: pass,
		Status:   true,
	}
	user.Password = utils.Md5(user.Email + user.Password)
	log.Debug().Msgf("user password: %s", user.Password)
	userID, err = user.Insert()
	return
}

// UpdateUser 更新用户
func (us UserService) UpdateUser(userID uint) {
	return
}

// DestroyUser 删除用户
func (us UserService) DestroyUser(userID uint) error {
	// log.Info().Msg("enter removeUser service.")
	return nil
}
