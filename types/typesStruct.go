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

type Orders struct {
	OrderId   int    `json:"orderid"`
	UserId    int    `json:"userid"`
	BookId    int    `json:"bookid"`
	OrderDate string `josn:"orderdate"`
	Quantity  int    `json:"quantity"`
	Status    string `josn:"status"`
}
