package repository

import (
	"github.com/julianlee107/user/domain/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	// InitTable 初始化表格
	InitTable() error
	// FindUserByName 根据用户名称查找用户信息
	FindUserByName(string) (*model.User, error)
	// FindUserByID 根据ID查找用户信息
	FindUserByID(int64) (*model.User, error)
	// DeleteUserByID 删除用户
	DeleteUserByID(int64) error
	CreateUser(user *model.User) (int64, error)
	UpdateUser(user *model.User) error
	FindAll() ([]model.User, error)
}
type UserRepository struct {
	mysqlDB *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{mysqlDB: db}
}

// InitTable 创建表格
func (u *UserRepository) InitTable() error {
	return u.mysqlDB.Migrator().CreateTable(&model.User{})
}

// FindUserByName 根据名称查找
func (u *UserRepository) FindUserByName(name string) (*model.User, error) {
	user := &model.User{}
	return user, u.mysqlDB.Where("user_name = ?", name).Take(&user).Error
}

// FindUserByID 根据ID查找
func (u *UserRepository) FindUserByID(id int64) (user *model.User, err error) {
	user = &model.User{}
	return user, u.mysqlDB.Where("id = ?", id).Take(&user).Error
}

// CreateUser 创建用户
func (u *UserRepository) CreateUser(user *model.User) (userID int64, err error) {
	return user.ID, u.mysqlDB.Create(user).Error
}

// DeleteUserByID 删除用户
func (u *UserRepository) DeleteUserByID(id int64) error {
	return u.mysqlDB.Where("id = ?", id).Delete(&model.User{}).Error
}

// UpdateUser 更新用户
func (u *UserRepository) UpdateUser(user *model.User) error {
	return u.mysqlDB.Model(user).Updates(&user).Error
}

func (u *UserRepository) FindAll() ([]model.User, error) {
	var users []model.User
	return users, u.mysqlDB.Find(&users).Error
}
