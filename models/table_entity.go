package models

var TableEntity = map[string][]string{
	"country": {"guid", "title", "code", "continent"},
	"city":    {"id", "title", "country_id", "city_code", "latitude", "longitude", `"offset"`, "timezone_id", "country_name"},
}
