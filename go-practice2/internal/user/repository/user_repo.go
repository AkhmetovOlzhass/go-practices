package repository

import (
    "go-prcatice2/internal/user/domain"
    "gorm.io/gorm"
)

type UserRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{db: db}
}

func (r *UserRepository) GetByID(id int) (*domain.User, error) {
    var user domain.User
    if err := r.db.First(&user, id).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *UserRepository) Create(user *domain.User) error {
    return r.db.Create(user).Error
}
