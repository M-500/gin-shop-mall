//@Author: wulinlin
//@Description:
//@File:  prod_repostory
//@Version: 1.0.0
//@Date: 2024/01/10 22:11

package prod_repositories

import (
	"backend/internal/models"
	databasenani "backend/pkg/database"
	"backend/pkg/utils"
	"gorm.io/gorm"
)

type IProdGroupRepository interface {
	Get(id int64) (*models.ProdTagModel, error)
	SearchProdTagPager(searchKey string, status int, pageSize, pageNum int) ([]*models.ProdTagModel, int64, error)
	Insert(*models.ProdTagModel) error
	Update(data *models.ProdTagModel, musColumns ...string) error
	Save(data *models.ProdTagModel, musColumns ...string) error
}

type ProdGroupRepository struct {
	DB *gorm.DB
}

func NewProdGroupRepository() IProdGroupRepository {
	return &ProdGroupRepository{
		DB: databasenani.GetDB(),
	}
}
func (repo *ProdGroupRepository) Get(id int64) (*models.ProdTagModel, error) {
	var query = models.ProdTagModel{}
	tx := repo.DB.Where("id = ?", id).First(&query)
	if tx.Error != nil && tx.Error != gorm.ErrRecordNotFound {
		return nil, tx.Error
	}
	if tx.RowsAffected < 1 {
		return nil, nil
	}
	return &query, nil
}

func (repo *ProdGroupRepository) SearchProdTagPager(searchKey string, status int, pageSize, page int) ([]*models.ProdTagModel, int64, error) {
	var queryData = []*models.ProdTagModel{}
	var total int64
	query := repo.DB.Model(&models.ProdTagModel{}).Limit(pageSize).Offset((page - 1) * pageSize)
	// 如果有关键字
	if utils.IsNotBlank(searchKey) {
		query = query.Where("title LIKE ?", "%"+searchKey+"%")
	}
	if status != 0 {
		query = query.Where(&models.ProdTagModel{Status: status})
	}
	query.Model(&models.ProdTagModel{}).Count(&total)
	// 分页
	tx := query.Scopes(utils.PaginateTemplate(int(page), int(pageSize))).Find(&queryData)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}
	return queryData, total, nil
}

func (repo *ProdGroupRepository) Insert(data *models.ProdTagModel) error {
	return repo.DB.Create(data).Error
}

func (repo *ProdGroupRepository) Update(data *models.ProdTagModel, musColumns ...string) error {
	return repo.DB.Model(data).Select(musColumns).Updates(data).Error
}

// 更高一层的封装
func (repo *ProdGroupRepository) Save(data *models.ProdTagModel, musColumns ...string) error {
	if data.ID > 0 {
		return repo.Update(data, musColumns...)
	}
	return repo.Insert(data)
}
