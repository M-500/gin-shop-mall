package repositories

import "gorm.io/gorm"

type IUserRepository interface {
	SelectByAccount(account string) error
	SelectByID(uid int64) error
	SelectByPhone(phone string) error
}

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (u *UserRepository) SelectByAccount(account string) error {
	return nil
}
func (u *UserRepository) SelectByID(uid int64) error {
	return nil
}
func (u *UserRepository) SelectByPhone(phone string) error {
	return nil
}
