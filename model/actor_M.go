package model

import (
	"dvdrentals_backend/database"
	"net/http"
	"strconv"
)

type (
	ActorWMostFilm struct {
		Fullname   string `json:"full_name"`
		LastUpdate string `json:"last_update"`
		Total      int16  `json:"total_film"`
	}
	TotalAverageActorFilmCategory struct {
		Name    string  `json:"name"`
		Total   int16   `json:"total_actor"`
		Average float32 `json:"average_actor_per_film"`
	}
)

func GetActorWMostFilm() (Response, error) {
	var res Response
	var obj ActorWMostFilm
	var arrobj []ActorWMostFilm

	db := database.GetDBInstance()

	rows, err := db.Raw("select coalesce (a.first_name , '') || ' ' || coalesce (a.last_name , '') as full_name, a.last_update, " +
		"count(1) as total_film from actors a join film_actors fa on fa.actor_id = a.actor_id group by a.actor_id " +
		"order by 3 desc limit 7").Rows()

	if err != nil {
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&obj.Fullname,
			&obj.LastUpdate,
			&obj.Total)

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

func GetTotalAverageActorFilmCategory() (Response, error) {
	var res Response
	var obj TotalAverageActorFilmCategory
	var arrobj []TotalAverageActorFilmCategory

	db := database.GetDBInstance()

	rows, err := db.Raw("select name, sum(total) as total_actor, cast(avg(total) as decimal(10,2)) as average_actor_per_film " +
		"from(select f.film_id, count(1) as total, c.name from films f join film_actors fa on fa.film_id = f.film_id join film_categorys fc " +
		"on f.film_id = fc.film_id join categorys c on c.category_id = fc.category_id group by 1, 3) q group by 1").Rows()

	if err != nil {
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&obj.Name,
			&obj.Total,
			&obj.Average)

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
