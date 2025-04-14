package service

import (
	"ecommerce-service/internal/domain"
	"ecommerce-service/internal/dto"
	"ecommerce-service/internal/repository"
	"fmt"
	"log"
)

type UserService struct {
	Repository repository.UserRepository
}

func (s UserService) FindUserByEmail(email string) (*domain.User, error) {
	// perform some db operation
	// business logic
	return nil, nil
}

func (service UserService) Signup(input dto.UserSignup) (string, error) {
	log.Println(input)

	user, err := service.Repository.CreateUser(domain.User{
		Email:    input.Email,
		Password: input.Password,
		Phone:    input.Phone,
	})

	// generate token
	log.Println(user)
	userInfo := fmt.Sprintf("%v,%v,%v", user.ID, user.Email, user.UserType)
	return userInfo, err
}

func (s UserService) Login(input any) (string, error) {
	return "", nil
}

func (s UserService) GetVerificationCode(e domain.User) (int, error) {
	return 0, nil
}

func (s UserService) VerifyCode(id uint, code int) error {
	return nil
}

func (s UserService) CreateProfile(id uint, input any) error {
	return nil
}

func (s UserService) GetProfile(id uint, input any) (*domain.User, error) {
	return nil, nil
}

func (s UserService) BecomeSeller(id uint, input any) (string, error) {
	return "", nil
}

func (s UserService) FindCart(id uint) ([]interface{}, error) {
	return nil, nil
}

func (s UserService) CreateCart(input any, user domain.User) ([]interface{}, error) {
	return nil, nil
}

func (s UserService) CreateOrder(user domain.User) (int, error) {
	return 0, nil
}

func (s UserService) GetOrders(user domain.User) ([]interface{}, error) {
	return nil, nil
}

func (s UserService) GetOrderById(id uint, uId uint) (interface{}, error) {
	return nil, nil
}
