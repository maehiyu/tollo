package repository

import (
	"context"
	"sync"

	"github.com/maehiyu/tollo/internal/userservice/domain/user"
)

// MockUserRepository is a mock implementation of UserRepository for testing.
// It stores users in memory.
type MockUserRepository struct {
	mu    sync.RWMutex
	users map[string]*user.User
}

// NewMockUserRepository creates a new mock user repository.
func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users: make(map[string]*user.User),
	}
}

// FindByID finds a user by ID.
func (r *MockUserRepository) FindByID(ctx context.Context, id string) (*user.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	u, ok := r.users[id]
	if !ok {
		return nil, user.ErrNotFound
	}
	return u, nil
}

// FindByEmail finds a user by email.
func (r *MockUserRepository) FindByEmail(ctx context.Context, email user.Email) (*user.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, u := range r.users {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, user.ErrNotFound
}

// Save saves a user. If the user already exists, it's updated. Otherwise, it's created.
// It also checks for email uniqueness on creation.
func (r *MockUserRepository) Save(ctx context.Context, u *user.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Check if it's a new user and if the email already exists
	_, err := r.findByIDInternal(u.ID)
	if err != nil { // User does not exist, so it's a creation
		if r.emailExistsInternal(u.Email) {
			return user.ErrEmailAlreadyExists
		}
	}

	r.users[u.ID] = u
	return nil
}

// Delete deletes a user by ID.
func (r *MockUserRepository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.users[id]; !ok {
		return user.ErrNotFound
	}
	delete(r.users, id)
	return nil
}

// --- Internal helpers without locks ---

func (r *MockUserRepository) findByIDInternal(id string) (*user.User, error) {
	u, ok := r.users[id]
	if !ok {
		return nil, user.ErrNotFound
	}
	return u, nil
}

func (r *MockUserRepository) emailExistsInternal(email user.Email) bool {
	for _, u := range r.users {
		if u.Email == email {
			return true
		}
	}
	return false
}
