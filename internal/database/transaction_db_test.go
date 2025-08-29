package database

import (
	"database/sql"
	"testing"

	"github.com/SamuelDevMobile/ms-wallet/internal/entity"
	"github.com/stretchr/testify/suite"
)

type TransactionDBTestSuite struct {
	suite.Suite
	DB            *sql.DB
	Client1       *entity.Client
	Client2       *entity.Client
	AccountFrom   *entity.Account
	AccountTo     *entity.Account
	TransactionDB *TransactionDB
}

func (s *TransactionDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.DB = db
	db.Exec("Create table clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("Create table accounts (id varchar(255), client_id varchar(255), balance int, created_at date)")
	db.Exec("Create table transactions (id varchar(255), account_id_from varchar(255), account_id_to varchar(255), amount int, created_at date)")

	client1, err := entity.NewClient("Samuel", "s@s.com")
	s.Nil(err)
	s.Client1 = client1
	client2, err := entity.NewClient("Isac", "i@i.com")
	s.Nil(err)
	s.Client2 = client2

	// creating account
	AccountFrom := entity.NewAccount(s.Client1)
	AccountFrom.Balance = 1000
	s.AccountFrom = AccountFrom
	AccountTo := entity.NewAccount(s.Client2)
	AccountTo.Balance = 1000
	s.AccountTo = AccountTo

	s.TransactionDB = NewTransactionDB(db)
}

func (s *TransactionDBTestSuite) TearDownSuite() {
	s.DB.Close()
	s.DB.Exec("DROP TABLE clients")
	s.DB.Exec("DROP TABLE accounts")
	s.DB.Exec("DROP TABLE transactions")
}

func TestTransactionTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDBTestSuite))
}

func (s *TransactionDBTestSuite) TestCreate() {
	transaction, err := entity.NewTransaction(s.AccountFrom, s.AccountTo, 100)
	s.Nil(err)
	err = s.TransactionDB.Create(transaction)
	s.Nil(err)
}
