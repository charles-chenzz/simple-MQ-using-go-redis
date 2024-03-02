package types

// all mysql related struct and connect
const (
	BizError        int64 = 1000
	ValidatingError int64 = 1001
	DataCreateError int64 = 1002
)

const (
	DSN = "127.0.0.1"
)

// User is database table
type User struct {
	UserId   int64  `json:"user_id"`
	UserName string `json:"user_name"`
}

type Order struct {
	OrderId       uint64 `json:"order_id"`
	TransactionId uint64 `json:"transaction_id"`
	ProductId     int64  `json:"product_id,omitempty"`
	ProductType   int64  `json:"product_type,omitempty"`
	Quantity      int64  `json:"quantity,omitempty"`
	Size          string `json:"size,omitempty"`
	Color         string `json:"color,omitempty"`
	Status        int64  `json:"status,omitempty"`
	RetryTime     int64  `json:"retry_time,omitempty"`
}

type Shipping struct {
	Email             string `json:"email,omitempty"`
	Address           string `json:"address,omitempty"`
	FirstName         string `json:"first_name,omitempty"`
	LastName          string `json:"last_name,omitempty"`
	ApartmentSuiteEtc string `json:"apartment_suite_etc,omitempty"`
	City              string `json:"city,omitempty"`
	State             int64  `json:"state,omitempty"`
	ZipCode           int64  `json:"zip_code,omitempty"`
	Phone             int64  `json:"phone,omitempty"`
}
