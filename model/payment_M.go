package model

import (
	"dvdrentals_backend/database"
	"net/http"
	"strconv"
)

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

func GetAllPayment() (Response, error) {
	var res Response
	db := database.GetDBInstance()

	payment := []Payment{}

	if err := db.Find(&payment).Error; err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Successfully collect " + strconv.Itoa(len(payment)) + " data"
	res.Data = payment

	return res, nil
}

func GetTotalTransactionByAmount() (Response, error) {
	var res Response
	var obj TotalTransaction
	var arrobj []TotalTransaction

	db := database.GetDBInstance()

	rows, err := db.Raw("select criteria, count(1) as total_transaction from (select amount, case when amount > 8 then 'Expensive' " +
		"when amount > 4 then 'Moderate' when amount > 0 then 'Cheap' else 'Free' end as criteria from payments p ) q " +
		"group by 1 order by 2 desc").Rows()

	if err != nil {
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&obj.Criteria,
			&obj.TotalTransaction)

		if err != nil {
			return res, nil
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Successfully collect " + strconv.Itoa(len(arrobj)) + " data"
	res.Data = arrobj

	return res, nil
}
