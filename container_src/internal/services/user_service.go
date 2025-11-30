package services

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"server/internal/models"
	"server/internal/repositories"
)

type UserService interface {
	CreateUser(req *models.CreateUserRequest) (*models.User, error)
	GetUser(id uuid.UUID) (*models.User, error)
	GetUsers(limit, offset int) ([]models.User, error)
	UpdateUser(id uuid.UUID, req *models.UpdateUserRequest) (*models.User, error)
	DeleteUser(id uuid.UUID) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(req *models.CreateUserRequest) (*models.User, error) {
	// Check if user already exists
	existingUser, err := s.repo.FindByEmail(req.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}

	user := &models.User{
		Email: req.Email,
		Name:  req.Name,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) GetUser(id uuid.UUID) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) GetUsers(limit, offset int) ([]models.User, error) {
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}
	return s.repo.FindAll(limit, offset)
}

func (s *userService) UpdateUser(id uuid.UUID, req *models.UpdateUserRequest) (*models.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Name != "" {
		user.Name = req.Name
	}

	if err := s.repo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) DeleteUser(id uuid.UUID) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}
