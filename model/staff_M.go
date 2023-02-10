package model

type (
	Staff struct {
		StaffId    int32  `json:"staff_id"`
		FirstName  string `json:"first_name"`
		LastName   string `json:"last_name"`
		AddressId  int32  `json:"address_id"`
		Email      string `json:"email"`
		StoreId    int32  `json:"store_id"`
		Active     bool   `json:"active"`
		Username   string `json:"username"`
		Password   string `json:"password"`
		LastUpdate string `json:"last_update"`
	}
	StaffPerformance struct {
		Fullname         string  `json:"full_name"`
		TotalReceived    float32 `json:"total_received"`
		AverageReceived  float32 `json:"average_received"`
		TotalTransaction int16   `json:"total_transaction"`
		TotalCustomer    int16   `json:"total_customer"`
	}
)