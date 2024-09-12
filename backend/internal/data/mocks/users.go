package mocks

import (
	"fmt"
	"time"

	"github.org/tawanr/sensorly/internal/data"
)

var mockUser = data.User{
	ID:        1,
	Email:     "test@example.com",
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

type UserModel struct{}

func (m UserModel) Insert(user *data.User) error {
	user.ID = 1
	fmt.Println("id:", user.ID)
	return nil
}

func (m UserModel) GetByEmail(email string) (*data.User, error) {
	if email == mockUser.Email {
		return &mockUser, nil
	}
	return &data.User{}, nil
}

func (m UserModel) Update(user *data.User) error {
	return nil
}

func (m UserModel) Delete(id int64) error {
	switch id {
	case mockUser.ID:
		return nil
	default:
		return data.ErrRecordNotFound
	}
}

func (m UserModel) GetForToken(tokenScope, tokenPlaintext string) (*data.User, error) {
	return &mockUser, nil
}
