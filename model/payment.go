package model

type (
	Payment struct {
		PaymentId   int32   `json:"payment_id"`
		CustomerId  int32   `json:"customer_id"`
		StaffId     int32   `json:"staff_id"`
		RentalId    int32   `json:"rental_id"`
		Amount      float32 `json:"amount"`
		PaymentDate string  `json:"payment_date"`
	}
	TotalTransaction struct {
		Criteria         string `json:"criteria"`
		TotalTransaction string `json:"total_transaction"`
	}
)
