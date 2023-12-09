package models

type AeroportPrimaryKey struct {
	Id string `json:"id"`
}

type Aeroport struct {
	Id           string  `json:"id"`
	Title        string  `json:"title"`
	CountryId    string  `json:"country_id"`
	CityId       string  `json:"city_id"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Radius       string  `json:"radius"`
	Image        string  `json:"image"`
	Address      string  `json:"address"`
	Country      string  `json:"country"`
	City         string  `json:"city"`
	SearchText   string  `json:"search_text"`
	Code         string  `json:"code"`
	ProductCount int64   `json:"product_count"`
	Gmt          string  `json:"gmt"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}

type CreateAeroport struct {
	Title        string  `json:"title"`
	CountryId    string  `json:"country_id"`
	CityId       string  `json:"city_id"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Radius       string  `json:"radius"`
	Image        string  `json:"image"`
	Address      string  `json:"address"`
	Country      string  `json:"country"`
	City         string  `json:"city"`
	SearchText   string  `json:"search_text"`
	Code         string  `json:"code"`
	ProductCount int64   `json:"product_count"`
	Gmt          string  `json:"gmt"`
}

type UpdateAeroport struct {
	Id           string  `json:"id"`
	Title        string  `json:"title"`
	CountryId    string  `json:"country_id"`
	CityId       string  `json:"city_id"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Radius       string  `json:"radius"`
	Image        string  `json:"image"`
	Address      string  `json:"address"`
	Country      string  `json:"country"`
	City         string  `json:"city"`
	SearchText   string  `json:"search_text"`
	Code         string  `json:"code"`
	ProductCount int64   `json:"product_count"`
	Gmt          string  `json:"gmt"`
}

type GetListAeroportRequest struct {
	Offset int64  `json:"offset"`
	Limit  int64  `json:"limit"`
	Search string `json:"search"`
	Query  string `json:"query"`
}

type GetListAeroportResponse struct {
	Count     int64       `json:"count"`
	Aeroports []*Aeroport `json:"aeroports"`
}
