package db

import (
	"context"
	"fmt"
	"phone-number-manager/internal/models"
	"sync"

	"github.com/google/uuid"
)

type memPhoneBookRepo struct {
	data map[string]*models.PhoneBook
	mu   sync.RWMutex
}

func NewInMemoryPhoneBookRepository(ctx context.Context) PhoneBookRepository {
	return &memPhoneBookRepo{
		data: make(map[string]*models.PhoneBook),
	}
}
func (m *memPhoneBookRepo) CreateEntry(ctx context.Context, entry *models.PhoneBook) error {
	m.mu.Lock()
	uuid := uuid.New().String()
	entry.ID = uuid
	defer m.mu.Unlock()
	m.data[entry.ID] = entry
	return nil
}

func (m *memPhoneBookRepo) GetEntry(ctx context.Context, uuid string) (*models.PhoneBook, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	entry, exists := m.data[uuid]
	if !exists {
		return nil, fmt.Errorf("entry not found")
	}
	return entry, nil
}
