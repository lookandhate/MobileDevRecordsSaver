package models

type CreateRecord struct {
	Points    int   `json:"points"`
	Timestamp int64 `json:"date"`
}
