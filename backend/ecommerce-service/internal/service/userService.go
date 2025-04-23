package service

import (
	"ecommerce-service/internal/domain"
	"ecommerce-service/internal/dto"
	"ecommerce-service/internal/helper"
	"ecommerce-service/internal/repository"
	"errors"
	"log"
	"time"
)

type UserService struct {
	Repository repository.UserRepository
	Auth       helper.Auth
}

func (service UserService) findUserByEmail(email string) (*domain.User, error) {
	// perform some db operation
	user, err := service.Repository.FindUserByEmail(email)
	// business logic
	return &user, err
}

func (service UserService) Signup(input dto.UserSignup) (string, error) {
	hashedPassword, err := service.Auth.CreateHashedPassword(input.Password)
	if err != nil {
		log.Println(err)
		return "", err
	}

	user, err := service.Repository.CreateUser(domain.User{
		Email:    input.Email,
		Password: hashedPassword,
		Phone:    input.Phone,
	})

	if err != nil {
		log.Println(err)
	}

	// generate token
	return service.Auth.GenerateToken(user.ID, user.Email, user.UserType)
}

func (service UserService) Login(email string, password string) (string, error) {
	user, err := service.findUserByEmail(email)
	if err != nil {
		return "", errors.New("user does not exist with the provided email id")
	}

	// compare password and generate token
	err = service.Auth.VerifyPassword(password, user.Password)
	if err != nil {
		return "", err
	}

	// generate token
	return service.Auth.GenerateToken(user.ID, email, user.UserType)
}

func (s UserService) isVerifiedUser(id uint) bool {

	currentUser, err := s.Repository.FindUserById(id)

	return err == nil && currentUser.Verified
}

func (s UserService) GetVerificationCode(e domain.User) (int, error) {
	// if user already verified
	if s.isVerifiedUser(e.ID) {
		return 0, errors.New("user already verified")
	}

	// generate verification code
	code, err := s.Auth.GenerateCode()
	if err != nil {
		return 0, err
	}

	// update user
	user := domain.User{
		Expiry: time.Now().Add(30 * time.Minute),
		Code:   code,
	}

	_, err = s.Repository.UpdateUser(e.ID, user)

	if err != nil {
		return 0, errors.New("unable to update verification code")
	}

	// user, _ = s.Repository.FindUserById(e.ID)

	// send sms

	// return verification code
	return code, nil
}

func (s UserService) VerifyCode(id uint, code int) error {
	// if user already verified
	if s.isVerifiedUser(id) {
		log.Println("verified...")
		return errors.New("user already verified")
	}

	user, err := s.Repository.FindUserById(id)

	if err != nil {
		return err
	}

	if user.Code != code {
		return errors.New("verification code does not match")
	}

	if !time.Now().Before(user.Expiry) {
		return errors.New("verification code expired")
	}

	updateUser := domain.User{
		Verified: true,
	}

	_, err = s.Repository.UpdateUser(id, updateUser)

	if err != nil {
		return errors.New("unable to verify user")
	}

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
