package usecase

import (
	"context"
	"time"

	"gin-backend/internal/domain/todo"
	"gin-backend/internal/domain/user"
)

type CreateTodoInput struct {
	UserID string `json:"user_id" binding:"required"`
	Title  string `json:"title" binding:"required"`
}

type CreateTodoUseCase struct {
	repo     todo.Repository
	userRepo user.Repository
}

func NewCreateTodoUseCase(repo todo.Repository, userRepo user.Repository) *CreateTodoUseCase {
	return &CreateTodoUseCase{repo: repo, userRepo: userRepo}
}

func (u *CreateTodoUseCase) Execute(ctx context.Context, input CreateTodoInput) (*todo.Todo, error) {
	if _, err := u.userRepo.FindByID(ctx, input.UserID); err != nil {
		return nil, err
	}

	t := &todo.Todo{
		ID:        time.Now().Format("20060102150405"),
		UserID:    input.UserID,
		Title:     input.Title,
		Completed: false,
		CreatedAt: time.Now(),
	}

	if err := u.repo.Create(ctx, t); err != nil {
		return nil, err
	}

	return t, nil
}
