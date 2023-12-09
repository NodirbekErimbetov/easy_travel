package models

type CityPrimaryKey struct {
	Id string `json:"id"`
}

type City struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	CountryId   string `json:"country_id"`
	CityCode    string `json:"city_code"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	Offset      string `json:"offset"`
	CountryName string `json:"country_name"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type CreateCity struct {
	Title       string `json:"title"`
	CountryId   string `json:"country_id"`
	CityCode    string `json:"city_code"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	Offset      string `json:"offset"`
	CountryName string `json:"country_name"`
}

type UpdateCity struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	CountryId   string `json:"country_id"`
	CityCode    string `json:"city_code"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	Offset      string `json:"offset"`
	CountryName string `json:"country_name"`
}

type GetListCityRequest struct {
	Offset int64  `json:"offset"`
	Limit  int64  `json:"limit"`
	Search string `json:"search"`
	Query  string `json:"query"`
}

type GetListCityResponse struct {
	Count  int64   `json:"count"`
	Cities []*City `json:"cities"`
}
