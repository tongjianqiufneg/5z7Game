package dao

import (
	"5z7Game/pkg/dto/request"
	"github.com/jinzhu/gorm"
)

// UserDao
type UserDao struct {
	BaseDao
}

// User return UserDao pointer with db connection
func User(db *gorm.DB) *UserDao {
	dao := &UserDao{}
	dao.db = db
	return dao
}

// Auth 校验用户
func (dao *UserDao) Auth(req request.UserAuthRequest) error {
	return nil
}

// Register 注册用户
func (dao *UserDao) Register(req request.UserRegisterRequest) error {
	return nil
}
