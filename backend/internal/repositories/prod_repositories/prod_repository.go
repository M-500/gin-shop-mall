//@Author: wulinlin
//@Description:
//@File:  prod_repostory
//@Version: 1.0.0
//@Date: 2024/01/10 22:11

package prod_repositories

import (
	"backend/internal/models"
	databasenani "backend/pkg/database"
	"gorm.io/gorm"
)

type IProdGroupRepository interface {
	Get(id int64) (*models.ProdTagModel, error)
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
