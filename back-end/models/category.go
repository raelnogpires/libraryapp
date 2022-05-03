package models

type Category struct {
	ID   int64  `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
