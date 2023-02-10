package model

type (
	Customer struct {
		CustomerId int32  `json:"customer_id"`
		StoreId    int32  `json:"store_id"`
		FirstName  string `json:"first_name"`
		LastName   string `json:"last_name"`
		AddressId  int32  `json:"address_id"`
		Email      string `json:"email"`
		ActiveBool bool   `json:"activebool"`
		CreateDate string `json:"create_date"`
		LastUpdate string `json:"last_update"`
		Active     int8   `json:"active"`
	}
	CustomerTotalAverageSpend struct {
		Fullname         string  `json:"full_name"`
		TotalSpend       float32 `json:"total_spend"`
		AverageSpend     float32 `json:"average_spend"`
		TotalTransaction int16   `json:"total_transaction"`
	}
	CustomerCountry struct {
		Country       string `json:"country"`
		TotalCustomer int16  `json:"total_customer"`
	}
)
