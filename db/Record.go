package db

type Record struct {
	ID        int64 `json:"id" gorm:"primary_key" gorm:"auto_increment"`
	Points    int   `json:"points"`
	Timestamp int64 `json:"date"`
}
