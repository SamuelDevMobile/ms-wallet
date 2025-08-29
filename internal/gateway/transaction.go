package gateway

import "github.com/SamuelDevMobile/ms-wallet/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
