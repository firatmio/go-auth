package models

import "errors"

type Permission int64

const (
	PermRead   Permission = 1 << 0 // 1
	PermWrite  Permission = 1 << 1 // 2
	PermDelete Permission = 1 << 2 // 4
	PermAdmin  Permission = 1 << 3 // 8
)

type User struct {
	ID          int        `json:"id"`
	Username    string     `json:"username"`
	Password    string     `json:"password"`
	Permissions Permission `json:"permissions"`
}

var (
	Users  = make(map[string]User)
	NextID = 1
)

func CreateUser(username, password string, permissions Permission) error {
	if _, exists := Users[username]; exists {
		return errors.New("user already exists")
	}

	user := User{
		ID:          NextID,
		Username:    username,
		Password:    password,
		Permissions: permissions,
	}

	Users[username] = user
	NextID++
	return nil
}

func GetUser(username string) (User, error) {
	user, exists := Users[username]
	if !exists {
		return User{}, errors.New("user not found")
	}
	return user, nil
}

func GetAllUsers() []User {
	users := make([]User, 0, len(Users))
	for _, u := range Users {
		users = append(users, u)
	}
	return users
}
