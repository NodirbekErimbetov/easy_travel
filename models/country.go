package models

type CountryPrimaryKey struct {
	Id string `json:"id"`
}

type Country struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Code      string `json:"code"`
	Continent string `json:"continent"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreateCountry struct {
	Title     string `json:"title"`
	Code      string `json:"code"`
	Continent string `json:"continent"`

}

type UpdateCountry struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Code      string `json:"code"`
	Continent string `json:"continent"`
}

type GetListCountryRequest struct {
	Offset int64  `json:"offset"`
	Limit  int64  `json:"limit"`
	Search string `json:"search"`
	Query  string `json:"query"`
	FromDate string `json:"from_date"`
	Todate string `json:"to_date"`
}

type GetListCounrtyResponse struct {
	Count int64 `json:"count"`
	Countries []*Country `json:"countries"`
}

