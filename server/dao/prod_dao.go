package dao

import (
	"server/models"
	"xorm.io/xorm"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2024/1/18 16:36
//

type ProdDao struct {
	db *xorm.Engine
}

func NewProdDao(engine *xorm.Engine) *ProdDao {
	return &ProdDao{
		db: engine,
	}
}

func (dao *ProdDao) Get(id int64) (*models.TzProd, error) {
	return nil, nil
}
