package usecase

import "github.com/kiranraj27/sprint-planner/internal/domain"

type userUseCase struct {
    userRepo domain.UserRepository
}

func NewUserUseCase(userRepo domain.UserRepository) domain.UserUseCase {
    return &userUseCase{userRepo}
}

func (uc *userUseCase) Register(user *domain.User) error {
    return uc.userRepo.Create(user)
}

func (uc *userUseCase) GetUser(id int) (*domain.User, error) {
    return uc.userRepo.GetByID(id)
}

func (uc *userUseCase) ListUsers() ([]*domain.User, error) {
    return uc.userRepo.GetUsers()
}