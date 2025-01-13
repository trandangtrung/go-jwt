package repositories

// repositories/user_repository.go

import (
	"errors"
)

type User struct {
	ID       int
	Username string
	Password string
}

type UserRepository struct {
	users []User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: []User{
			{ID: 1, Username: "user1", Password: "pass1"},
			{ID: 2, Username: "user2", Password: "pass2"},
			{ID: 3, Username: "user3", Password: "pass3"},

			// Thêm người dùng khác
		},
	}
}

func (repo *UserRepository) FindByUsername(username string) (*User, error) {
	for _, user := range repo.users {
		if user.Username == username {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}
