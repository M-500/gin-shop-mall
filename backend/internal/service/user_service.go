package service

import "backend/internal/repositories/users_repositories"

type IUserService interface {
}

type UserService struct {
	repo users_repositories.IUserRepository
}

func NewUserService(r users_repositories.IUserRepository) IUserService {
	return &UserService{
		repo: r,
	}
}
