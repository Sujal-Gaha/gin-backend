package mem_repo

import (
	"context"
	"sync"

	"gin-backend/internal/domain"
	"gin-backend/internal/domain/todo"
)

type TodoRepository struct {
	mu    sync.RWMutex
	todos map[string]*todo.Todo
}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{
		todos: make(map[string]*todo.Todo),
	}
}

func (r *TodoRepository) Create(ctx context.Context, t *todo.Todo) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.todos[t.ID] = t
	return nil
}

func (r *TodoRepository) FindByID(ctx context.Context, id string) (*todo.Todo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	t, ok := r.todos[id]
	if !ok {
		return nil, domain.ErrNotFound
	}
	return t, nil
}

func (r *TodoRepository) FindAll(ctx context.Context) ([]*todo.Todo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	list := make([]*todo.Todo, 0, len(r.todos))
	for _, t := range r.todos {
		list = append(list, t)
	}
	return list, nil
}

func (r *TodoRepository) FindByUserID(ctx context.Context, userID string) ([]*todo.Todo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	list := make([]*todo.Todo, 0)
	for _, t := range r.todos {
		if t.UserID == userID {
			list = append(list, t)
		}
	}
	return list, nil
}

func (r *TodoRepository) Update(ctx context.Context, t *todo.Todo) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.todos[t.ID]; !ok {
		return domain.ErrNotFound
	}
	r.todos[t.ID] = t
	return nil
}

func (r *TodoRepository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.todos[id]; !ok {
		return domain.ErrNotFound
	}
	delete(r.todos, id)
	return nil
}
