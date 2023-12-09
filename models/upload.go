package models


type UploadRequest struct {
	TableSlug string `json:"table_slug"`
	Data []map[string]interface{}
}