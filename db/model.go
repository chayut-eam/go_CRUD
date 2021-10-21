package db

import (
	"time"

	"github.com/google/uuid"
	"crud/chrono"
)

type ProductMapping struct {
	Sku     *string `bson:"topvalue_sku"`
	Barcode *string `bson:"cj_barcode"`
}

type APILog struct {
	ID            *string      `bson:"_id"`
	TransactionID *string      `bson:"transaction_id"`
	OrderID       *string      `bson:"order_id,omitempty"`
	CreatedAt     *time.Time   `bson:"created_at"`
	Request       *APIRequest  `bson:"request"`
	Response      *APIResponse `bson:"response"`
}

type APIRequest struct {
	URL     string              `bson:"url"`
	Method  string              `bson:"method"`
	Headers map[string][]string `bson:"headers"`
	Body    interface{}         `bson:"body"`
}

type APIResponse struct {
	HTTPStatus int                 `bson:"http_status"`
	Headers    map[string][]string `bson:"header"`
	Body       interface{}         `bson:"body"`
	Latency    int64               `bson:"latency"`
}

func NewAPILog() *APILog {
	uuid := uuid.NewString()
	now := chrono.Now().Time
	return &APILog{
		ID:            &uuid,
		TransactionID: &uuid,
		CreatedAt:     &now,
	}
}

type Order struct {
	ID                 *string    `bson:"_id"`
	OrderID            *string    `bson:"order_id"`
	CJCode             *string    `bson:"cj_code"`
	CJBillNo           *string    `bson:"cj_bill_no"`
	GrandTotal         *float64   `bson:"grand_total"`
	Items              []*Item    `bson:"items"`
	CreatedAt          *time.Time `bson:"created_at"`
	IsPaymentConfirmed bool       `bson:"is_payment_confirmed"`
	PaymentConfirmedAt *time.Time `bson:"payment_confirmed_at"`
	ShouldRetry        bool       `bson:"should_retry"`
	IsPendingRetry     bool       `bson:"is_pending_retry"`
	LastAttemptAt      *time.Time `bson:"last_attempt_at"`
	RetryAttempt       int64      `bson:"retry_attempt"`
	DeletedAt          *time.Time `bson:"deleted_at"`
}

type Item struct {
	ItemID     *int64   `bson:"item_id"`
	Barcode    string   `bson:"barcode"`
	ItemName   *string  `bson:"item_name"`
	Price      *float64 `bson:"price"`
	QtyOrdered *int64   `bson:"qty_ordered"`
}
