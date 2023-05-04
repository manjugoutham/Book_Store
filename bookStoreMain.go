package main

import (
	bookcontroller "bookstore/BookDetailsController"
	usercontroller "bookstore/userRegistrationControl"
	"database/sql"
	"fmt"
	"net/http"
)

func main() {

	//Database Connection
	db, err := sql.Open("mysql", "root:Tham12@2@tcp(localhost:3306)/bookstore")
	if err != nil {
		fmt.Println("err in sql connection", err)
		return
	}
	defer db.Close()

	//Api
	http.HandleFunc("/register", usercontroller.Register(db))
	http.HandleFunc("/allUsers", usercontroller.AllRecord(db))
	http.HandleFunc("/userById/", usercontroller.GetUserById(db))
	http.HandleFunc("/editUser/", usercontroller.UpdateId(db))
	http.HandleFunc("/deleteId/", usercontroller.DeleteById(db))

	http.HandleFunc("/addBook", bookcontroller.Addbook(db))
	http.HandleFunc("/getAllBooks", bookcontroller.GetAllBooks(db))
	http.HandleFunc("/getBookById/", bookcontroller.Getbookbyid(db))
	http.HandleFunc("/getBookByName/", bookcontroller.Getbookbyname(db))
	http.HandleFunc("/deleteById/", bookcontroller.Deletebyid(db))
	http.HandleFunc("/updateBook/", bookcontroller.UpdateBook(db))

	fmt.Println("Port number : 8000")
	http.ListenAndServe("localhost:8000", nil)

}
