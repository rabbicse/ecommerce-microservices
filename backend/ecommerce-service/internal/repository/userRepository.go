package repository

import (
	"ecommerce-service/internal/domain"
	"errors"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	CreateUser(user domain.User) (domain.User, error)
	FindUserByEmail(email string) (domain.User, error)
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
	return user, nil
}

func (repository userRepository) FindUserByEmail(email string) (domain.User, error) {
	var user domain.User

	err := repository.db.First(&user, "email=?", email).Error
	if err != nil {
		log.Printf("find user error. Details: %v\n", err)
		return domain.User{}, err
	}
	return user, nil
}

func (repository userRepository) FindUserById(id uint) (domain.User, error) {
	var user domain.User

	err := repository.db.First(&user, id).Error

	if err != nil {
		log.Printf("find user error %v", err)
		return domain.User{}, errors.New("user does not exist")
	}

	return user, nil
}

func (repository userRepository) UpdateUser(id uint, u domain.User) (domain.User, error) {
	var user domain.User

	err := repository.db.Model(&user).Clauses(clause.Returning{}).Where("id=?", id).Updates(u).Error

	if err != nil {
		log.Printf("error on update %v", err)
		return domain.User{}, errors.New("failed update user")
	}

	return user, nil
}
