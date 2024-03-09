package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"service-member/config"
)

func NewConnectionPostgres(config *config.Config) *sqlx.DB {
	dns := fmt.Sprintf(
		"dbname=%s host=%s port=%d user=%s password=%s sslmode=disable",
		config.Database.Name,
		config.Database.Host,
		config.Database.Port,
		config.Database.User,
		config.Database.Password)

	sql, err := sqlx.Open("postgres", dns)
	if err != nil {
		log.Fatalf("Fail open Database {%s}", err.Error())
	}

	if err := sql.Ping(); err != nil {
		log.Fatalf("Fail ping Database {%s}", err.Error())
	}

	return sql
}
