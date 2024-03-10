package usecase

import (
	"encoding/json"
	"log"
	"service-member/internal/domain"
	"service-member/internal/dto"
	"service-member/internal/exception"
	"service-member/internal/util"
)

type MemberUsecaseImpl struct {
	memberRepo   domain.MemberRepository
	jwtGenerator util.JwtGenerator
	hashPassword util.PasswordGenerator
	kafkaMessage util.KafkaUsecase
}

func (m MemberUsecaseImpl) SignUpMember(request dto.MemberSignUpRequest) (dto.MemberAuthResponse, error) {

	_, err := m.memberRepo.FindMemberByEmail(request.Email)
	if err == nil {
		return dto.MemberAuthResponse{}, exception.BadRequestError{Message: "Member has been signup"}
	}

	token, err := m.jwtGenerator.NewAccessToken(request.Email, request.Name)
	if err != nil {
		log.Printf("Error generate jwt {%s}", err)
	}

	password, err := m.hashPassword.HashPassword(request.Password)
	if err != nil {
		log.Printf("Error hash password {%s}", err)
		return dto.MemberAuthResponse{}, err

	}
	newMember := domain.Member{
		Name:      request.Name,
		Password:  password,
		AccessKey: token,
		Email:     request.Email,
	}

	err = m.memberRepo.CreateMember(newMember)
	if err != nil {
		log.Printf("Fail to create member {%s}", err)
		return dto.MemberAuthResponse{}, err
	}

	messageToKafka := dto.SendOtpKafka{
		Type:  "REGISTRATION",
		Email: request.Email,
	}
	log.Println(messageToKafka)

	marshal, err := json.Marshal(messageToKafka)
	if err != nil {
		log.Printf("Potential fail send message notification becase error marshal json {%s}", err)
	}

	sendToKafka := m.kafkaMessage.ProduceMessage(marshal)
	if !sendToKafka {
		log.Printf("Fail send to kafka with message : {%s}", messageToKafka)
	}
	log.Printf("Send kafka with message : {%s}", messageToKafka)

	return dto.MemberAuthResponse{Token: token}, nil
}

func (m MemberUsecaseImpl) SignInMember(request dto.MemberSignInRequest) (dto.MemberAuthResponse, error) {
	member, err := m.memberRepo.FindMemberByEmail(request.Email)
	if err != nil {
		return dto.MemberAuthResponse{}, exception.NotFoundError{Message: "Member not found"}
	}

	isValidPassword := m.hashPassword.CheckPasswordHash(request.Password, member.Password)
	if !isValidPassword {
		return dto.MemberAuthResponse{}, exception.BadRequestError{Message: "Invalid Password please check again!!"}
	}

	messageToKafka := dto.SendOtpKafka{
		Type:  "LOGIN",
		Email: request.Email,
	}
	log.Println(messageToKafka)

	marshal, err := json.Marshal(messageToKafka)
	if err != nil {
		log.Printf("Potential fail send message notification becase error marshal json {%s}", err)
	}

	sendToKafka := m.kafkaMessage.ProduceMessage(marshal)
	if !sendToKafka {
		log.Printf("Fail send to kafka with message : {%s}", messageToKafka)
	}
	log.Printf("Send kafka with message : {%s}", messageToKafka)

	return dto.MemberAuthResponse{Token: member.AccessKey}, nil
}

func (m MemberUsecaseImpl) DeleteMember(header dto.Header) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m MemberUsecaseImpl) GetMemberById(id int) (dto.MemberResponse, error) {
	member, err := m.memberRepo.FindMemberById(id)
	if err != nil {
		return dto.MemberResponse{}, exception.NotFoundError{Message: "Member not found"}
	}

	return dto.MemberResponse{
		Id:       member.Id,
		Name:     member.Name,
		Email:    member.Email,
		CreateAt: member.CreateAt,
	}, nil
}

func NewUsecaseMember(member domain.MemberRepository,
	generator util.JwtGenerator,
	passwordGenerator util.PasswordGenerator,
	kafkaMessage util.KafkaUsecase) domain.MemberService {
	return &MemberUsecaseImpl{
		memberRepo:   member,
		jwtGenerator: generator,
		hashPassword: passwordGenerator,
		kafkaMessage: kafkaMessage,
	}
}
