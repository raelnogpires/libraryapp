package book

type Book struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CategoryId  int64  `json:"category_id"`
	AuthorId    int64  `json:"author_id"`
	ImgUrl      string `json:"img_url"`
}
