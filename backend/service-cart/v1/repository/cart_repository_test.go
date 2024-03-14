package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"log"
	"regexp"
	"service-cart/internal/domain"
	"testing"
	"time"
)

func PrepSqlMock() (*sqlx.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal("Fail to prepare sql mock_repo")
	}
	return sqlx.NewDb(db, "postgres"), mock
}

var dummyCart = domain.Cart{
	Id:          uuid.New(),
	UserId:      1,
	TotalAmount: 1000,
	StatusCart:  domain.InBasket,
	CreateAt:    time.Now(),
}

func TestCreateCart_Success(t *testing.T) {
	mockConnDb, mockSql := PrepSqlMock()
	fakeMemberRepo := NewCartRepository(mockConnDb)
	prepare := mockSql.ExpectPrepare(regexp.QuoteMeta(`INSERT INTO public.tb_cart 
	    (user_id, total_amount, status_cart, created_at) VALUES($1, $2, $3, $4)`))
	prepare.ExpectExec().WithArgs().WillReturnResult(sqlmock.NewResult(0, 1))

	err := fakeMemberRepo.CreateNewCart(dummyCart)
	assert.NoError(t, err)
}

func TestCreateCart_Fail(t *testing.T) {
	mockConnDb, mockSql := PrepSqlMock()
	fakeMemberRepo := NewCartRepository(mockConnDb)
	prepare := mockSql.ExpectPrepare(regexp.QuoteMeta(`INSERT INTO public.tb_cart 
	    (user_id, total_amount, status_cart, created_at) VALUES($1, $2, $3, $4)`))
	prepare.ExpectExec().WithArgs().WillReturnResult(sqlmock.NewResult(0, 0))

	err := fakeMemberRepo.CreateNewCart(dummyCart)
	assert.NoError(t, err)
}

func TestFindCurrentCart_CurrentCartExist(t *testing.T) {
	mockConnDb, mockSql := PrepSqlMock()
	fakeMemberRepo := NewCartRepository(mockConnDb)
	row := mockSql.NewRows([]string{"id", "user_id", "total_amount", "status_cart", "create_at"}).
		AddRow(dummyCart.Id, dummyCart.UserId, dummyCart.TotalAmount, dummyCart.StatusCart, dummyCart.CreateAt)

	prepare := mockSql.ExpectPrepare(regexp.QuoteMeta(`SELECT 
    id, user_id, total_amount, status_cart, created_at 
	FROM public.tb_cart WHERE user_id = $1 AND status_cart = 'IN_BASKET'`))
	prepare.ExpectQuery().WithArgs(1).WillReturnRows(row)

	cart := fakeMemberRepo.FindCurrentCart(1)
	assert.NotEqual(t, cart, nil)
	assert.Equal(t, cart.Id, dummyCart.Id)
	assert.Equal(t, cart.UserId, 1)
}

func TestFindCurrentCart_CurrentCartNotExist(t *testing.T) {
	mockConnDb, mockSql := PrepSqlMock()
	fakeMemberRepo := NewCartRepository(mockConnDb)

	prepare := mockSql.ExpectPrepare(regexp.QuoteMeta(`SELECT 
    id, user_id, total_amount, status_cart, created_at 
	FROM public.tb_cart WHERE user_id = $1 AND status_cart = 'IN_BASKET'`))
	prepare.ExpectQuery().WithArgs(1).WillReturnRows(mockSql.NewRows(nil))

	cart := fakeMemberRepo.FindCurrentCart(1)
	assert.Nil(t, cart)
}

func TestUpdateCart_UpdateSuccessful(t *testing.T) {
	mockConnDb, mockSql := PrepSqlMock()
	fakeMemberRepo := NewCartRepository(mockConnDb)

	prepare := mockSql.ExpectPrepare(regexp.QuoteMeta(`UPDATE public.tb_cart
	SET user_id=$2, total_amount=$3, status_cart=$4, created_at=$5
	WHERE id=$1`))
	prepare.ExpectExec().
		WithArgs(dummyCart.Id, dummyCart.UserId, 10000, dummyCart.StatusCart, dummyCart.CreateAt).
		WillReturnResult(sqlmock.NewResult(0, 1))

	var dummyCartUpdate = domain.Cart{
		Id:          dummyCart.Id,
		UserId:      1,
		TotalAmount: 10000,
		StatusCart:  domain.InBasket,
		CreateAt:    dummyCart.CreateAt,
	}

	err := fakeMemberRepo.UpdateCart(dummyCartUpdate)
	assert.NoError(t, err)
}
