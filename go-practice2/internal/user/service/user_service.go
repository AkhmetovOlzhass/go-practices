package service

import (
    "errors"
    "go-prcatice2/internal/user/domain"
    "go-prcatice2/internal/user/repository"
)

type UserService struct {
    repo *repository.UserRepository
}

func NewUserService(r *repository.UserRepository) *UserService {
    return &UserService{repo: r}
}

func (s *UserService) GetUser(id int) (*domain.User, error) {
    if id <= 0 {
        return nil, errors.New("invalid id")
    }
    return s.repo.GetByID(id)
}

func (s *UserService) CreateUser(name string) (*domain.User, error) {
    if name == "" {
        return nil, errors.New("invalid name")
    }
    user := &domain.User{Name: name}
    return user, s.repo.Create(user)
}
