package database

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"crud/internal/models"

	"github.com/google/uuid"
)

type JSONStorage struct {
	filename string
	mutex    sync.RWMutex
	users    map[string]models.User
}

func NewJSONStorage(filename string) *JSONStorage {
	storage := &JSONStorage{filename: filename}
	storage.init()
	return storage
}

func (s *JSONStorage) init() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, err := os.Stat(s.filename); os.IsNotExist(err) {
		os.WriteFile(s.filename, []byte("{}"), 0644)
	}
	s.loadUsers()
}

func (s *JSONStorage) loadUsers() {
	data, err := os.ReadFile(s.filename)
	if err != nil {
		fmt.Println("❌ Error reading JSON file:", err)
		s.users = make(map[string]models.User)
		return
	}

	if err := json.Unmarshal(data, &s.users); err != nil {
		fmt.Println("❌ Error unmarshaling JSON:", err)
		s.users = make(map[string]models.User)
	}
}

func (s *JSONStorage) saveUsers() error {
	data, err := json.MarshalIndent(s.users, "", "  ")
	if err != nil {
		return fmt.Errorf("❌ Error marshaling JSON: %v", err)
	}

	if err := os.WriteFile(s.filename, data, 0644); err != nil {
		return fmt.Errorf("❌ Error writing to JSON file: %v", err)
	}

	return nil
}

func (s *JSONStorage) GetUsers() ([]models.User, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	var userList []models.User
	for _, user := range s.users {
		userList = append(userList, user)
	}
	return userList, nil
}

func (s *JSONStorage) GetUserByID(id string) (models.User, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	user, exists := s.users[id]
	if !exists {
		return models.User{}, fmt.Errorf("user not found")
	}
	return user, nil
}

func (s *JSONStorage) CreateUser(user models.User) (string, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	id := uuid.New().String()
	user.ID = id
	s.users[id] = user
	if err := s.saveUsers(); err != nil {
		return "", err
	}
	return id, nil
}

func (s *JSONStorage) UpdateUser(id string, updatedUser models.User) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, exists := s.users[id]; !exists {
		return fmt.Errorf("user not found")
	}

	updatedUser.ID = id
	s.users[id] = updatedUser
	return s.saveUsers()
}

func (s *JSONStorage) DeleteUser(id string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, exists := s.users[id]; !exists {
		return fmt.Errorf("user not found")
	}

	delete(s.users, id)
	return s.saveUsers()
}
