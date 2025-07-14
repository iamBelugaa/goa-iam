package userstore

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/iamBelugaa/goa-iam/gen/user/gen/user"
	userdomain "github.com/iamBelugaa/goa-iam/internal/domain/user"
)

type memory struct {
	mu           sync.RWMutex
	emailToIdMap map[string]string
	users        map[string]*user.User
}

func NewMemoryStore() *memory {
	return &memory{}
}

func (m *memory) QueryUserById(context context.Context, userID string) (*user.User, error) {
	user, ok := m.users[userID]
	if !ok {
		return nil, fmt.Errorf("user with id %s doesn't exists", userID)
	}
	return user, nil
}

func (m *memory) QueryUserByEmail(context context.Context, email string) (*user.User, error) {
	userID, ok := m.emailToIdMap[email]
	if !ok {
		return nil, fmt.Errorf("user with email %s doesn't exists", email)
	}

	user, ok := m.users[userID]
	if !ok {
		return nil, fmt.Errorf("user with email %s doesn't exists", email)
	}

	return user, nil
}

func (m *memory) Create(context context.Context, cmd user.CreateUserRequest) (*user.User, error) {
	m.mu.RLock()
	if _, ok := m.emailToIdMap[cmd.Email]; ok {
		m.mu.RUnlock()
		return nil, fmt.Errorf("user with email %s already exists", cmd.Email)
	}
	m.mu.RUnlock()

	user := user.User{
		ID:        uuid.New().String(),
		FirstName: cmd.FirstName,
		LastName:  cmd.LastName,
		Email:     cmd.Email,
		Status:    userdomain.UserStatusActive,
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.emailToIdMap[cmd.Email]; !ok {
		m.emailToIdMap[cmd.Email] = user.ID
		m.users[user.ID] = &user
	} else {
		return nil, fmt.Errorf("user with email %s already exists", cmd.Email)
	}

	return &user, nil
}
