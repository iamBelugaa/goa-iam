// Package userstore provides an in-memory implementation of the UserStorer interface.
package userstore

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"

	"github.com/iamBelugaa/goa-iam/gen/user"
	userdomain "github.com/iamBelugaa/goa-iam/internal/domain/user"
)

// memory implements the UserStorer interface using in-memory maps.
type memory struct {
	mu           sync.RWMutex          // protects access to users and emailToIdMap
	emailToIdMap map[string]string     // maps email addresses to user IDs
	users        map[string]*user.User // stores user data by ID
}

// NewMemoryStore creates and returns a new instance of the in-memory user store.
func NewMemoryStore() *memory {
	return &memory{
		emailToIdMap: make(map[string]string),
		users:        make(map[string]*user.User),
	}
}

// QueryById retrieves a user from memory by their user ID.
func (m *memory) QueryById(ctx context.Context, userID string) (*user.User, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	user, ok := m.users[userID]
	if !ok {
		return nil, fmt.Errorf("user with id %s doesn't exist", userID)
	}
	return user, nil
}

// QueryByEmail retrieves a user from memory by their email address.
func (m *memory) QueryByEmail(ctx context.Context, email string) (*user.User, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	userID, ok := m.emailToIdMap[email]
	if !ok {
		return nil, fmt.Errorf("user with email %s doesn't exist", email)
	}

	user, ok := m.users[userID]
	if !ok {
		return nil, fmt.Errorf("user with email %s doesn't exist", email)
	}

	return user, nil
}

// Create adds a new user to the in-memory store.
func (m *memory) Create(ctx context.Context, cmd *user.CreateUserRequest) (*user.User, error) {
	// Check for duplicate email with read lock.
	m.mu.RLock()
	if _, exists := m.emailToIdMap[cmd.Email]; exists {
		m.mu.RUnlock()
		return nil, fmt.Errorf("user with email %s already exists", cmd.Email)
	}
	m.mu.RUnlock()

	// Create new user object.
	newUser := &user.User{
		ID:        uuid.New().String(),
		FirstName: cmd.FirstName,
		LastName:  cmd.LastName,
		Email:     cmd.Email,
		Status:    userdomain.UserStatusActive,
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
	}

	// Store the user with write lock.
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.emailToIdMap[cmd.Email]; exists {
		return nil, fmt.Errorf("user with email %s already exists", cmd.Email)
	}

	m.emailToIdMap[cmd.Email] = newUser.ID
	m.users[newUser.ID] = newUser

	return newUser, nil
}

// List returns a slice containing all users currently stored in memory.
func (s *memory) List(ctx context.Context) ([]*user.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// Prepare a slice with the exact length of the stored users map.
	users := make([]*user.User, len(s.users))
	var index int

	for _, user := range s.users {
		users[index] = user
		index++
	}

	return users, nil
}
