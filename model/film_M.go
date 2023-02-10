package model

type (
	Film struct {
		FilmId          string  `json:"film_id"`
		Title           string  `json:"title"`
		Description     string  `json:"description"`
		ReleaseYear     int16   `json:"release_year"`
		LanguageId      int16   `json:"language_id"`
		RentalDuration  int8    `json:"rental_duration"`
		RentalRate      float32 `json:"rental_rate"`
		Length          int8    `json:"length"`
		ReplacementCost float32 `json:"replacement_cost"`
		Rating          string  `json:"rating"`
		LastUpdate      string  `json:"last_update"`
		SpecialFeatures string  `json:"special_features"`
		Fulltext        string  `json:"fulltext"`
	}
	TotalFilmWActor struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Total       int8   `json:"total"`
	}
	TotalFilmWRate struct {
		Rating string `json:"rating"`
		Total  int8   `json:"total"`
	}
	TotalFilmWCat struct {
		Name  string `json:"name"`
		Total int8   `json:"total"`
	}
	AverageFilmDuration struct {
		Name            string  `json:"name"`
		Total           int16   `json:"total"`
		AverageDuration float32 `json:"average_duration"`
	}
)
