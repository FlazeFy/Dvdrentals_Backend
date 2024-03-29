package model

import (
	"dvdrentals_backend/database"
	"net/http"
	"strconv"
)

type (
	Film struct {
		FilmId          string  `json:"film_id"`
		Title           string  `json:"title"`
		Description     string  `json:"description"`
		ReleaseYear     int16   `json:"release_year"`
		LanguageId      int16   `json:"language_id"`
		RentalDuration  int16   `json:"rental_duration"`
		RentalRate      float32 `json:"rental_rate"`
		Length          int32   `json:"length"`
		ReplacementCost float32 `json:"replacement_cost"`
		Rating          string  `json:"rating"`
		LastUpdate      string  `json:"last_update"`
		SpecialFeatures string  `json:"special_features"`
		Fulltext        string  `json:"fulltext"`
	}
	TotalFilmWActor struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Total       int16  `json:"total"`
	}
	TotalFilmWRate struct {
		Rating string `json:"rating"`
		Total  int16  `json:"total"`
	}
	TotalFilmWCat struct {
		Name  string `json:"name"`
		Total int16  `json:"total"`
	}
	AverageFilmDuration struct {
		Name            string  `json:"name"`
		Total           int16   `json:"total"`
		AverageDuration float32 `json:"average_duration"`
	}
	TotalAverageFilmReplacementCost struct {
		Criteria    string  `json:"criteria"`
		Total       int16   `json:"total"`
		AverageCost float32 `json:"average_cost"`
	}
)

func GetAllFilm() (Response, error) {
	var res Response
	db := database.GetDBInstance()

	film := []Film{}

	if err := db.Find(&film).Error; err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Successfully collect " + strconv.Itoa(len(film)) + " data"
	res.Data = film

	return res, nil
}

func GetFilmWMostActor() (Response, error) {
	var res Response
	var obj TotalFilmWActor
	var arrobj []TotalFilmWActor

	db := database.GetDBInstance()

	rows, err := db.Raw("select f.title, f.description, count(1) as total from films f join film_actors fa on fa.film_id = f.film_id group by f.film_id order by 3 desc limit 7").Rows()
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&obj.Title,
			&obj.Description,
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

func GetTotalFilmByRating() (Response, error) {
	var res Response
	var obj TotalFilmWRate
	var arrobj []TotalFilmWRate

	db := database.GetDBInstance()

	rows, err := db.Raw("select rating, count(1) as total from films group by 1 order by 2 desc").Rows()
	if err != nil {
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&obj.Rating,
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

func GetTotalFilmByCat() (Response, error) {
	var res Response
	var obj TotalFilmWCat
	var arrobj []TotalFilmWCat

	db := database.GetDBInstance()

	rows, err := db.Raw("select c.name, count(1) as total from categorys c join film_categorys " +
		"fc on c.category_id = fc.category_id group by c.category_id order by 2 desc").Rows()
	if err != nil {
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&obj.Name,
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

func GetAverageFilmDurationByCat() (Response, error) {
	var res Response
	var obj AverageFilmDuration
	var arrobj []AverageFilmDuration

	db := database.GetDBInstance()

	rows, err := db.Raw("select c.name, count(1) as total, cast(avg(f.length) as decimal(10,2)) as average_duration " +
		" from categorys c join film_categorys fc on c.category_id = fc.category_id join films f on f.film_id = fc.film_id group by c.category_id order by 3 desc").Rows()
	if err != nil {
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&obj.Name,
			&obj.Total,
			&obj.AverageDuration)

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

func GetTotalAverageFilmReplacementCost() (Response, error) {
	var res Response
	var obj TotalAverageFilmReplacementCost
	var arrobj []TotalAverageFilmReplacementCost

	db := database.GetDBInstance()

	rows, err := db.Raw("select case when f.replacement_cost < 16 then 'Cheap' when f.replacement_cost < 23 then 'Moderate' " +
		"else 'Expensive' end as criteria, count(1) as total, cast(avg(f.replacement_cost) as decimal(10,2)) as average_cost " +
		"from films f group by 1 order by 3").Rows()

	if err != nil {
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&obj.Criteria,
			&obj.Total,
			&obj.AverageCost)

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
