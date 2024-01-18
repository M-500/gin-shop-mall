package service

import (
	"backend/internal/entity"
	"backend/internal/repositories/prod_repositories"
)

type ProdService struct {
	repo prod_repositories.IProdGroupRepository
}

func NewProdService() *ProdService {
	return &ProdService{
		repo: prod_repositories.NewProdGroupRepository(),
	}
}

func (s *ProdService) QueryProdTagList(searchKey string, status int, pageSize, pageNum int) ([]*entity.ProdTag, int64, error) {
	return s.repo.SearchProdTagPager(searchKey, status, pageSize, pageNum)
}

func (s *ProdService) AddProdTag(data *entity.ProdTag, musColumns ...string) error {
	return s.repo.Save(data, musColumns...)
}
func (s *ProdService) Get(id int64) (*entity.ProdTag, error) {
	return s.repo.Get(id)
}

func (s *ProdService) Delete(id int64) error {
	return s.repo.Delete(id)
}
