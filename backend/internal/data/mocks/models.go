package mocks

import "github.org/tawanr/sensorly/internal/data"

func NewMockModels() *data.Models {
	return &data.Models{
		Users:  UserModel{},
		Tokens: data.TokenModel{},
	}
}
