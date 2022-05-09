package models

type Book struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	CategoryID  uint   `json:"category_id" binding:"required"`
	AuthorID    uint   `json:"author_id" binding:"required"`
	ImgUrl      string `json:"img_url" binding:"required"`
}

// model used in controllers.GetAllBooks and controllers.GetBookById
// returns full information, making front-end work easier
type FullBook struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	CategoryID   uint   `json:"category_id"`
	CategoryName string `json:"category_name"`
	AuthorID     uint   `json:"author_id"`
	AuthorName   string `json:"author_name"`
	ImgUrl       string `json:"img_url"`
}
