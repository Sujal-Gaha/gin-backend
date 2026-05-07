package usecase

import (
	"context"
	"time"

	"gin-backend/internal/domain/user"
)

type CreateUserInput struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

type CreateUserUseCase struct {
	repo user.Repository
}

func NewCreateUserUseCase(repo user.Repository) *CreateUserUseCase {
	return &CreateUserUseCase{repo: repo}
}

func (u *CreateUserUseCase) Execute(ctx context.Context, input CreateUserInput) (*user.User, error) {
	usr := &user.User{
		ID:        "u_" + time.Now().Format("20060102150405"),
		Name:      input.Name,
		Email:     input.Email,
		CreatedAt: time.Now(),
	}

	if err := u.repo.Create(ctx, usr); err != nil {
		return nil, err
	}

	return usr, nil
}
