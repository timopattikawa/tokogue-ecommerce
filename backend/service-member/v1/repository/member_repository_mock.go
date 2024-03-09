package repository

import (
	"github.com/stretchr/testify/mock"
	"log"
	"service-member/internal/domain"
)

type MemberRepositoryMock struct {
	Mock mock.Mock
}

func (m *MemberRepositoryMock) CreateMember(member domain.Member) error {
	arguments := m.Mock.Called(member)
	log.Println(arguments.Get(0))
	if arguments.Get(0) != nil {
		return arguments.Get(0).(error)
	}

	return nil
}

func (m *MemberRepositoryMock) FindMemberById(id int) (domain.Member, error) {
	arguments := m.Mock.Called(id)
	if arguments.Get(0) == nil && arguments.Get(1) != nil {
		return domain.Member{}, arguments.Get(1).(error)
	}
	member := arguments.Get(0).(domain.Member)
	return member, nil
}

func (m *MemberRepositoryMock) DeleteMemberById(id domain.Member) (domain.Member, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MemberRepositoryMock) FindMemberByEmail(email string) (domain.Member, error) {
	arguments := m.Mock.Called(email)
	if arguments.Get(0) == nil && arguments.Get(1) != nil {
		return domain.Member{}, arguments.Get(1).(error)
	}
	member := arguments.Get(0).(domain.Member)
	return member, nil
}
