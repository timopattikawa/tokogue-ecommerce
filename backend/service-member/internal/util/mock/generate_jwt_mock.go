package mock

import (
	"fmt"
	"github.com/stretchr/testify/mock"
)

type JwtGeneratorImplMock struct {
	Mock mock.Mock
}

func (j *JwtGeneratorImplMock) NewAccessToken(email string, name string) (string, error) {
	arguments := j.Mock.Called(email, name)

	if arguments.Get(0) != nil && arguments.Get(1) == nil {
		return arguments.Get(0).(string), nil
	}

	return "", fmt.Errorf("error get jwt token")
}

func (j *JwtGeneratorImplMock) VerifyToken(token string) error {
	//TODO implement me
	panic("implement me")
}
