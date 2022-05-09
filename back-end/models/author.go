package models

type Author struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" binding:"required"`
	// https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples/
}
