package model

import (
	"dvdrentals_backend/database"
	"net/http"
	"strconv"
)

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

func GetAllStaff() (Response, error) {
	var res Response
	db := database.GetDBInstance()

	staff := []Staff{}

	if err := db.Find(&staff).Error; err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Successfully collect " + strconv.Itoa(len(staff)) + " data"
	res.Data = staff

	return res, nil
}

func GetStaffPerformance() (Response, error) {
	var res Response
	var obj StaffPerformance
	var arrobj []StaffPerformance

	db := database.GetDBInstance()

	rows, err := db.Raw("select coalesce (s.first_name , '') || ' ' || coalesce (s.last_name , '') as full_name, " +
		"cast(sum(p.amount) as decimal(10,2)) as total_received, cast(avg(p.amount) as decimal(10,2)) as average_received, " +
		"count(1) as total_transaction, count(distinct p.customer_id) as total_customer from staffs s " +
		"join payments p on p.staff_id = s.staff_id group by s.staff_id order by 2 desc").Rows()

	if err != nil {
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&obj.Fullname,
			&obj.TotalReceived,
			&obj.AverageReceived,
			&obj.TotalTransaction,
			&obj.TotalCustomer)

		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Successfully collect " + strconv.Itoa(len(arrobj)) + " data"
	res.Data = arrobj

	return res, nil
}
