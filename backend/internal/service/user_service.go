package service

import "backend/internal/repositories"

type IUserService interface {
}

type UserService struct {
	repo repositories.IUserRepository
}

func NewUserService(r repositories.IUserRepository) IUserService {
	return &UserService{
		repo: r,
	}
}
