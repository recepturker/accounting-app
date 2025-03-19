package services

import (
	"errors"

	"github.com/recepturker/accounting-app/internal/models"
	"github.com/recepturker/accounting-app/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

// Kullanıcıyı doğrula
func (s *UserService) AuthenticateUser(email, password string) (*models.User, error) {
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("kullanıcı bulunamadı")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("hatalı şifre")
	}

	return user, nil
}

func (s *UserService) RegisterUser(user *models.User) error {
	return s.userRepo.CreateUser(user)
}
