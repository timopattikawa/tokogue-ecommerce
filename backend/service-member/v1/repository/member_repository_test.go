package repository

import (
	"bou.ke/monkey"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"log"
	"regexp"
	"service-member/internal/domain"
	"testing"
	"time"
)

func DBMockPrep() (*sqlx.DB, sqlmock.Sqlmock) {
	sqlDB, mockDb, err := sqlmock.New()
	if err != nil {
		log.Fatalf("Someting error %s expected when mocking database", err)
	}
	return sqlx.NewDb(sqlDB, "pq"), mockDb
}

func TestCreateMember_ResultSuccess(t *testing.T) {
	sqlM, mock := DBMockPrep()
	fakeMemberRepo := NewMemberRepository(sqlM)
	defer sqlM.Close()

	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	})
	var memberDummy = &domain.Member{
		Id:        101,
		Name:      "Timo Pattikawa",
		Password:  "asdfasdf",
		AccessKey: "TMPTMPTMPTMP",
		Email:     "timo@gmail.com",
		CreateAt:  time.Now(),
	}

	query := `INSERT INTO member ("name", "password", access_key, email, create_at) values ($1, $2, $3, $4, $5)`

	prepare := mock.ExpectPrepare(regexp.QuoteMeta(query))
	prepare.ExpectExec().
		WithArgs(memberDummy.Name, memberDummy.Password, memberDummy.AccessKey, memberDummy.Email, memberDummy.CreateAt).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := fakeMemberRepo.CreateMember(*memberDummy)
	assert.NoError(t, err)
}

func TestCreateMember_ResultFailCauseError(t *testing.T) {
	sqlM, mock := DBMockPrep()
	defer sqlM.Close()

	fakeMemberRepo := NewMemberRepository(sqlM)
	var memberDummy = &domain.Member{
		Id:        101,
		Name:      "Timo Pattikawa",
		Password:  "asdfasdf",
		AccessKey: "TMPTMPTMPTMP",
		Email:     "timo@gmail.com",
		CreateAt:  time.Now(),
	}
	query := `INSERT INTO member ("name", "password", access_key, email, create_at) values ($1, $2, $3, $4, $5)`

	prepare := mock.ExpectPrepare(regexp.QuoteMeta(query))
	prepare.ExpectExec().
		WithArgs(memberDummy.Name, memberDummy.Password, memberDummy.AccessKey, memberDummy.CreateAt, memberDummy.CreateAt).
		WillReturnResult(sqlmock.NewResult(0, 0))

	err := fakeMemberRepo.CreateMember(*memberDummy)
	assert.Error(t, err)
}

func TestFindMemberById_ResultSuccessToFind(t *testing.T) {
	sqlMck, mock := DBMockPrep()
	defer sqlMck.Close()

	fakeMemberRepo := NewMemberRepository(sqlMck)
	var memberDummy = &domain.Member{
		Id:        101,
		Name:      "Timo Pattikawa",
		Password:  "asdfasdf",
		AccessKey: "TMPTMPTMPTMP",
		Email:     "timo@gmail.com",
		CreateAt:  time.Now(),
	}
	rows := mock.NewRows([]string{"id", "name", "password", "access_key", "email", "create_at"}).
		AddRow(memberDummy.Id, memberDummy.Name, memberDummy.Password, memberDummy.AccessKey, memberDummy.Email, memberDummy.CreateAt)

	query := `SELECT id, "name", "password", access_key, email, create_at 
		FROM member WHERE id = $1`
	prepare := mock.ExpectPrepare(regexp.QuoteMeta(query))
	prepare.ExpectQuery().WithArgs(101).WillReturnRows(rows)

	memberById, err := fakeMemberRepo.FindMemberById(101)
	assert.Equal(t, memberById.Email, memberDummy.Email)
	assert.Nil(t, err)
}

func TestFindMemberById_ResultFailToFind404(t *testing.T) {
	sqlMck, mock := DBMockPrep()
	defer sqlMck.Close()

	fakeMemberRepo := NewMemberRepository(sqlMck)

	query := `SELECT id, "name", "password", access_key, email, create_at 
		FROM member WHERE id = $1`
	prepare := mock.ExpectPrepare(regexp.QuoteMeta(query))
	prepare.ExpectQuery().WithArgs(101).WillReturnRows(mock.NewRows([]string{}))

	memberById, err := fakeMemberRepo.FindMemberById(101)
	assert.NotNil(t, err)
	assert.Equal(t, domain.Member{}, memberById)
}

func TestFindMemberByEmail_ResultSuccess(t *testing.T) {
	sqlMck, mock := DBMockPrep()
	defer sqlMck.Close()

	fakeMemberRepo := NewMemberRepository(sqlMck)
	var memberDummy = &domain.Member{
		Id:        101,
		Name:      "Timo Pattikawa",
		Password:  "asdfasdf",
		AccessKey: "TMPTMPTMPTMP",
		Email:     "timo@gmail.com",
		CreateAt:  time.Now(),
	}
	rows := mock.NewRows([]string{"id", "name", "password", "access_key", "email", "create_at"}).
		AddRow(memberDummy.Id, memberDummy.Name, memberDummy.Password, memberDummy.AccessKey, memberDummy.Email, memberDummy.CreateAt)

	query := `SELECT id, "name", "password", access_key, email, create_at 
		FROM member WHERE email = $1`
	prepare := mock.ExpectPrepare(regexp.QuoteMeta(query))
	prepare.ExpectQuery().WithArgs("timo@gmail.com").WillReturnRows(rows)

	memberById, err := fakeMemberRepo.FindMemberByEmail("timo@gmail.com")
	assert.Equal(t, memberById.Email, memberDummy.Email)
	assert.Nil(t, err)
}
