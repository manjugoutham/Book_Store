package types

type User struct {
	Id       int    `josn:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type Books struct {
	Book_id     int    `json:"book_id"`
	Book_Name   string `json:"book_name"`
	Author_Name string `json:"author_name"`
	Price       int    `json:"price"`
}
