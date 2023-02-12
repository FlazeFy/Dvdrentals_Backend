package model

import (
	"dvdrentals_backend/database"
	"net/http"
	"strconv"
)

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

func GetAllCustomer() (Response, error) {
	var res Response
	db := database.GetDBInstance()

	customer := []Customer{}

	if err := db.Find(&customer).Error; err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Successfully collect " + strconv.Itoa(len(customer)) + " data"
	res.Data = customer

	return res, nil
}

func GetAllCustomerTotalAvgSpending() (Response, error) {
	var res Response
	var obj CustomerTotalAverageSpend
	var arrobj []CustomerTotalAverageSpend

	db := database.GetDBInstance()

	rows, err := db.Raw("select coalesce (c.first_name , '') || ' ' || coalesce (c.last_name , '') as full_name, " +
		"cast(sum(p.amount) as decimal(10,2)) as total_spend, cast(avg(p.amount) as decimal(10,2)) as average_spend, " +
		"count(c.customer_id) as total_transaction from customers c join payments p on p.customer_id = c.customer_id " +
		"group by c.customer_id order by 2 desc").Rows()

	if err != nil {
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&obj.Fullname,
			&obj.TotalSpend,
			&obj.AverageSpend,
			&obj.TotalTransaction)

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

func GetCountryWMostCustomer() (Response, error) {
	var res Response
	var obj CustomerCountry
	var arrobj []CustomerCountry

	db := database.GetDBInstance()

	rows, err := db.Raw("select c.country, count(1) as total_customer from countrys c join citys c2 on c2.country_id = c.country_id " +
		"join address a on a.city_id  = c2.city_id join customers c3 on c3.address_id = a.address_id group by 1 order by 2 desc " +
		"limit 7").Rows()

	if err != nil {
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&obj.Country,
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
