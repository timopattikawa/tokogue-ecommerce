package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"regexp"
	"service-cart/internal/domain"
	"testing"
)

func prepareMock() (*sqlx.DB, sqlmock.Sqlmock) {
	db, s, err := sqlmock.New()
	if err != nil {
		return nil, nil
	}

	return sqlx.NewDb(db, "postgres"), s
}

var dummyCartDetail = domain.CartDetail{
	Id:        1,
	CartId:    uuid.New(),
	ProductId: 1,
	Qty:       10,
}

func TestInsertCartDetail_SuccessToInsert(t *testing.T) {
	mockConnDb, mockSql := prepareMock()

	repository := NewCartDetailRepository(mockConnDb)

	prepare := mockSql.ExpectPrepare(regexp.QuoteMeta(`INSERT INTO public.tb_cart_detail 
    (product_id, qty, cart_id) VALUES($1, $2, $3)`))
	prepare.ExpectExec().WithArgs(
		dummyCartDetail.ProductId, dummyCartDetail.Qty, dummyCartDetail.CartId).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := repository.InsertCartDetail(dummyCartDetail)
	assert.NoError(t, err)
}

func TestInsertCartDetail_FailToUpdate(t *testing.T) {
	mockConnDb, mockSql := prepareMock()

	repository := NewCartDetailRepository(mockConnDb)

	prepare := mockSql.ExpectPrepare(regexp.QuoteMeta(`INSERT INTO public.tb_cart_detail 
    (id, product_id, qty, cart_id) VALUES(nextval('tb_cart_detail_id_seq'::regclass), $1, $2, $3)`))
	prepare.ExpectExec().WithArgs(
		dummyCartDetail.ProductId, dummyCartDetail.Qty, dummyCartDetail.CartId).
		WillReturnResult(sqlmock.NewResult(0, 0))

	err := repository.InsertCartDetail(dummyCartDetail)
	assert.NotNil(t, err)
}

func TestUpdateDetailCart_SuccessToUpdate(t *testing.T) {
	mockConnDb, mockSql := prepareMock()

	repository := NewCartDetailRepository(mockConnDb)

	prepare := mockSql.ExpectPrepare(regexp.QuoteMeta(`UPDATE public.tb_cart_detail
		SET product_id=$2, qty=$3, cart_id=$4
		WHERE id=$1`))
	prepare.ExpectExec().WithArgs(
		dummyCartDetail.Id, dummyCartDetail.ProductId, dummyCartDetail.Qty, dummyCartDetail.CartId).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := repository.UpdateCartDetail(dummyCartDetail)
	assert.NoError(t, err)
}

func TestFindByCartIdAndProductId_SuccessToFind(t *testing.T) {
	mockConnDb, mockSql := prepareMock()

	repository := NewCartDetailRepository(mockConnDb)

	rows := mockSql.NewRows([]string{"id", "product_id", "qty", "cart_id"}).AddRow(
		dummyCartDetail.Id, dummyCartDetail.ProductId, dummyCartDetail.Qty, dummyCartDetail.CartId,
	)

	prepare := mockSql.ExpectPrepare(regexp.QuoteMeta(`SELECT id, product_id, qty, cart_id
		FROM public.tb_cart_detail WHERE cart_id = $1 AND product_id = $2`))
	prepare.ExpectQuery().WithArgs(
		dummyCartDetail.CartId, dummyCartDetail.ProductId).
		WillReturnRows(rows)

	res := repository.FindByCartIdAndProductId(dummyCartDetail.CartId, dummyCartDetail.ProductId)
	assert.NotNil(t, res)
}
