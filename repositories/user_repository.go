// package repositories

// // repositories/user_repository.go

// import (
// 	"errors"
// )

// type User struct {
// 	ID       int
// 	Username string
// 	Password string
// }

// type UserRepository struct {
// 	users []User
// }

// func NewUserRepository() *UserRepository {
// 	return &UserRepository{
// 		users: []User{
// 			{ID: 1, Username: "user1", Password: "pass1"},
// 			{ID: 2, Username: "user2", Password: "pass2"},
// 			{ID: 3, Username: "user3", Password: "pass3"},

// 			// Thêm người dùng khác
// 		},
// 	}
// }

// func (repo *UserRepository) FindByUsername(username string) (*User, error) {
// 	for _, user := range repo.users {
// 		if user.Username == username {
// 			return &user, nil
// 		}
// 	}
// 	return nil, errors.New("user not found")
// }

package repositories

import (
	"JWT/models"
	"errors"
	"sync"
)

type UserRepository struct {
	users []models.User
	mu    sync.Mutex
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: []models.User{
			{ID: 1, Username: "user1", Password: "pass1"},
			{ID: 2, Username: "user2", Password: "pass2"},
			{ID: 3, Username: "user3", Password: "pass3"},
		},
	}
}

func (repo *UserRepository) AddUser(username, password string) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	for _, user := range repo.users {
		if user.Username == username {
			return errors.New("user already exists")
		}
	}

	newID := len(repo.users) + 1
	newUser := models.User{ID: newID, Username: username, Password: password}
	repo.users = append(repo.users, newUser)
	return nil
}

func (repo *UserRepository) FindByUsername(username string) (*models.User, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	for _, user := range repo.users {
		if user.Username == username {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (repo *UserRepository) ValidateCredentials(username, password string) bool {
	user, err := repo.FindByUsername(username)
	if err != nil {
		return false
	}
	return user.Password == password
}
