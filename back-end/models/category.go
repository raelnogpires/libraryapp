package models

type Category struct {
	Id   int64  `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
