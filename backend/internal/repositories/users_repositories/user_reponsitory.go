package users_repositories

import (
	"backend/internal/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Get(uid int64) (*models.SysUserModel, error)
	FindByAccount(account string) (*models.SysUserModel, error)
	FindByPhone(phone string) (*models.SysUserModel, error)
	FindAllPager(searchKey string, page, size int) ([]models.SysUserModel, int64, error)
	Insert(data *models.SysUserModel) error
	Update(data *models.SysUserModel, musColumns ...string) error
	Save(data *models.SysUserModel, musColumns ...string) error
}

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		DB: db,
	}
}
func (u *UserRepository) Get(uid int64) (*models.SysUserModel, error) {
	user := models.SysUserModel{}
	tx := u.DB.First(&user, uid)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected < 1 {
		return nil, nil
	}
	return &user, nil
}
func (u *UserRepository) FindByAccount(account string) (*models.SysUserModel, error) {
	queryData := models.SysUserModel{}
	u.DB.Where("username = ?", account).First(&queryData)

	return nil, nil
}

func (u *UserRepository) FindByPhone(phone string) (*models.SysUserModel, error) {
	queryData := models.SysUserModel{}
	u.DB.Where("phone = ?", phone).First(&queryData)

	return nil, nil
}

func (u *UserRepository) FindAllPager(searchKey string, page, size int) ([]models.SysUserModel, int64, error) {
	var users []models.SysUserModel
	var total int64
	query := u.DB.Model(&models.SysUserModel{})
	if searchKey != "" {
		query = query.Where("username LIKE ?", "%"+searchKey+"%")
	}
	query.Count(&total)
	err := query.Offset((page - 1) * size).Limit(size).Find(&users).Error
	return users, total, err
}

func (u *UserRepository) Insert(data *models.SysUserModel) error {
	return u.DB.Create(data).Error
}

func (u *UserRepository) Update(data *models.SysUserModel, musColumns ...string) error {
	return u.DB.Model(data).Select(musColumns).Updates(data).Error
}

func (u *UserRepository) Save(data *models.SysUserModel, musColumns ...string) error {
	return nil
}
