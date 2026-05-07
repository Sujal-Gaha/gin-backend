package mem_repo

import (
	"context"
	"sync"

	"gin-backend/internal/domain"
	"gin-backend/internal/domain/user"
)

type UserRepository struct {
	mu    sync.RWMutex
	users map[string]*user.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make(map[string]*user.User),
	}
}

func (r *UserRepository) Create(ctx context.Context, u *user.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.users[u.ID] = u
	return nil
}

func (r *UserRepository) FindByID(ctx context.Context, id string) (*user.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	u, ok := r.users[id]
	if !ok {
		return nil, domain.ErrNotFound
	}
	return u, nil
}

func (r *UserRepository) FindAll(ctx context.Context) ([]*user.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	list := make([]*user.User, 0, len(r.users))
	for _, u := range r.users {
		list = append(list, u)
	}
	return list, nil
}

func (r *UserRepository) Update(ctx context.Context, u *user.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.users[u.ID]; !ok {
		return domain.ErrNotFound
	}
	r.users[u.ID] = u
	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.users[id]; !ok {
		return domain.ErrNotFound
	}
	delete(r.users, id)
	return nil
}
