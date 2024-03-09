package repository

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"log"
	"service-member/internal/domain"
	"time"
)

type MemberRepositoryImpl struct {
	db *sqlx.DB
}

func (m MemberRepositoryImpl) FindMemberByEmail(email string) (domain.Member, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30000)
	defer cancel()
	query := `SELECT id, "name", "password", access_key, email, create_at 
		FROM member WHERE email = $1`
	stmt, err := m.db.Prepare(query)
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Printf("Fail to close prepare query")
		}
	}(stmt)

	if err != nil {
		log.Printf("Error to prepare query {%s}, err: {%s}", query, err.Error())
		return domain.Member{}, err
	}

	var member domain.Member
	err = stmt.QueryRowContext(ctx, email).Scan(
		&member.Id,
		&member.Name,
		&member.Password,
		&member.AccessKey,
		&member.Email,
		&member.CreateAt)
	if err != nil {
		log.Printf("Error to query row: {%s}", err.Error())
		return domain.Member{}, err
	}

	return member, nil

}

func (m MemberRepositoryImpl) CreateMember(member domain.Member) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30000)
	defer cancel()

	query := `INSERT INTO member ("name", "password", access_key, email, create_at) values ($1, $2, $3, $4, $5)`
	prep, err := m.db.Prepare(query)

	defer func(prep *sql.Stmt) {
		err := prep.Close()
		if err != nil {
			log.Printf("Fail to close prepare query")
		}
	}(prep)

	if err != nil {
		log.Printf("Error to prepare query {%s}, err: {%s}", query, err.Error())
		return err
	}

	member.CreateAt = time.Now()
	_, err = prep.ExecContext(ctx,
		&member.Name,
		&member.Password,
		&member.AccessKey,
		&member.Email,
		&member.CreateAt,
	)
	if err != nil {
		log.Printf("Error to Exec: {%s}", err.Error())
		return err
	}

	return nil

}

func (m MemberRepositoryImpl) FindMemberById(id int) (domain.Member, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30000)
	defer cancel()

	query := `SELECT id, "name", "password", access_key, email, create_at 
		FROM member WHERE id = $1`
	prepare, err := m.db.Prepare(query)
	defer func(prepare *sql.Stmt) {
		err := prepare.Close()
		if err != nil {
			log.Printf("Fail to close prepare query")
		}
	}(prepare)

	if err != nil {
		log.Printf("Error to prepare query {%s}, err: {%s}", query, err.Error())
		return domain.Member{}, err
	}

	var member domain.Member
	err = prepare.QueryRowContext(ctx, id).Scan(
		&member.Id,
		&member.Name,
		&member.Password,
		&member.AccessKey,
		&member.Email,
		&member.CreateAt)
	if err != nil {
		log.Printf("Error to query row: {%s}", err.Error())
		return domain.Member{}, err
	}

	return member, nil
}

func (m MemberRepositoryImpl) DeleteMemberById(id domain.Member) (domain.Member, error) {
	//TODO implement me
	panic("implement me")
}

func NewMemberRepository(Db *sqlx.DB) domain.MemberRepository {
	return &MemberRepositoryImpl{
		db: Db,
	}
}
