package models

type Book struct {
	ID          int64  `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CategoryId  int64  `json:"category_id"`
	AuthorId    int64  `json:"author_id"`
	// will be fixing this
	Category Category `gorm:"foreignKey:ID"`
	Author   Author   `gorm:"foreignKey:ID"`
	ImgUrl   string   `json:"img_url"`
}
