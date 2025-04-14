package repository

import (
	"ecommerce-service/internal/domain"
	"log"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user domain.User) (domain.User, error)
	FindUser(email string) (domain.User, error)
	FindUserById(id uint) (domain.User, error)
	UpdateUser(id uint, user domain.User) (domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (repository userRepository) CreateUser(user domain.User) (domain.User, error) {
	err := repository.db.Create(&user).Error
	if err != nil {
		log.Printf("create user error. Details: %v\n", err)
		return domain.User{}, err
	}
	return domain.User{}, nil
}

func (repository userRepository) FindUser(email string) (domain.User, error) {
	return domain.User{}, nil
}

func (repository userRepository) FindUserById(id uint) (domain.User, error) {
	return domain.User{}, nil
}

func (repository userRepository) UpdateUser(id uint, user domain.User) (domain.User, error) {
	return domain.User{}, nil
}
