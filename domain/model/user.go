package model

type User struct {
	// 主键
	ID           int64  `gorm:"primary)key;not_null;auto_increment"`
	UserName     string `gorm:"unique_index;not_null"`
	FirstName    string
	HashPassword string
}
