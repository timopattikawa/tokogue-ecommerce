package mock

import (
	"fmt"
	"github.com/stretchr/testify/mock"
)

type PasswordGeneratorMock struct {
	Mock mock.Mock
}

func (p *PasswordGeneratorMock) HashPassword(password string) (string, error) {
	arguments := p.Mock.Called(password)

	if arguments.Get(0) != nil && arguments.Get(1) == nil {
		return arguments.Get(0).(string), nil
	}

	return "", fmt.Errorf("Error generate password")
}

func (p *PasswordGeneratorMock) CheckPasswordHash(password, hash string) bool {
	arguments := p.Mock.Called(password, hash)

	return arguments.Get(0).(bool)
}
