package database

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"crud/internal/models"

	"github.com/google/uuid"
)

var (
	jsonFile = "users.json"
	mutex    sync.RWMutex // Read-Write Mutex for concurrency safety
	users    map[string]models.User
)

func InitJSONStorage(filename string) {
	mutex.Lock()
	defer mutex.Unlock()
	jsonFile = filename

	if _, err := os.Stat(jsonFile); os.IsNotExist(err) {
		os.WriteFile(jsonFile, []byte("{}"), 0644)
	}
	loadUsers()
}

func loadUsers() {
	data, err := os.ReadFile(jsonFile)
	if err != nil {
		fmt.Println("❌ Error reading JSON file:", err)
		users = make(map[string]models.User)
		return
	}

	if err := json.Unmarshal(data, &users); err != nil {
		fmt.Println("❌ Error unmarshaling JSON:", err)
		users = make(map[string]models.User) // Reset on failure
	}
}

func saveUsers() error {

	data, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return fmt.Errorf("❌ Error marshaling JSON: %v", err)
	}

	if err := os.WriteFile(jsonFile, data, 0644); err != nil {
		return fmt.Errorf("❌ Error writing to JSON file: %v", err)
	}

	return nil
}

func GetJSONUsers() ([]models.User, error) {
	mutex.RLock()
	defer mutex.RUnlock()

	var userList []models.User
	for _, user := range users {
		userList = append(userList, user)
	}
	return userList, nil
}

func GetJSONUserByID(id string) (models.User, bool) {
	mutex.RLock()
	defer mutex.RUnlock()

	user, exists := users[id]
	return user, exists
}

func CreateJSONUser(user models.User) error {
	mutex.Lock()
	defer mutex.Unlock()

	id := uuid.New().String()
	users[id] = user
	return saveUsers()
}

func UpdateJSONUser(id string, updatedUser models.User) error {
	mutex.Lock()
	defer mutex.Unlock()

	if _, exists := users[id]; !exists {
		return fmt.Errorf("❌ User with ID %s not found", id)
	}

	updatedUser.ID = id // Preserve the ID
	users[id] = updatedUser
	return saveUsers()
}

func DeleteJSONUser(id string) error {
	mutex.Lock()
	defer mutex.Unlock()

	if _, exists := users[id]; !exists {
		return fmt.Errorf("❌ User with ID %s not found", id)
	}

	delete(users, id)
	return saveUsers()
}
