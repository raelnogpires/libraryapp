package models

type Author struct {
	ID   int64  `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
