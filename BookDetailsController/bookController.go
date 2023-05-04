package BookDetailsControllerPackage

import (
	tp "bookstore/types"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func Addbook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprint(w, "POST Method not allowed..")
			return
		}
		var books tp.Books
		err := json.NewDecoder(r.Body).Decode(&books)
		if err != nil {
			fmt.Println("Error in register json.Decode", err.Error())
			return
		}
		result, err := db.Exec("INSERT INTO books(book_name,author_name,price) VALUES (?,?,?)", books.Book_Name, books.Author_Name, books.Price)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		id, err := result.LastInsertId()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		books.Book_id = int(id)

		json.NewEncoder(w).Encode(books)

	}
}

func GetAllBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprint(w, "GET Method Not Allowed...!")
			return
		}

		rows, err := db.Query("select * from books")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer rows.Close()

		for rows.Next() {
			var data tp.Books
			err := rows.Scan(&data.Book_id, &data.Book_Name, &data.Author_Name, &data.Price)
			if err != nil {
				fmt.Println("err in read ", err)
				return
			}
			json.NewEncoder(w).Encode(data)
		}
	}
}

func Getbookbyid(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprint(w, "GET Method not allowed")
			return
		}
		id := r.URL.Path[len("/getBookById/"):]

		//fmt.Println(id)
		rows := db.QueryRow("select * from books WHERE book_id = ?", id)

		var book tp.Books

		err := rows.Scan(&book.Book_id, &book.Book_Name, &book.Author_Name, &book.Price)
		if err != nil {
			fmt.Fprint(w, err.Error())
			fmt.Fprint(w, "No book Found...!")

		} else {
			json.NewEncoder(w).Encode(book)
		}

	}
}

func Getbookbyname(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprint(w, "Method not allowed")
			return
		}

		name := r.URL.Path[len("/getBookByName/"):]
		row := db.QueryRow("select * from books WHERE book_name = ?", name)

		var book tp.Books
		err := row.Scan(&book.Book_id, &book.Book_Name, &book.Author_Name, &book.Price)

		if err != nil {
			fmt.Fprint(w, err.Error())
		} else {
			json.NewEncoder(w).Encode(book)
		}

	}
}

func UpdateBook(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "PUT" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprint(w, "PUT Method Not Allowed...!")
			return
		}

		id := r.URL.Path[len("/updateBook/"):]
		var book tp.Books
		json.NewDecoder(r.Body).Decode(&book)

		_, err := db.Exec("UPDATE books SET book_name = ? ,author_name =?, price = ? Where book_id = ?", book.Book_Name, book.Author_Name, book.Price, id)
		if err != nil {
			fmt.Fprint(w, "book with given id not available in list")
		} else {
			fmt.Fprint(w, "Book details Updated successfully")
		}
	}
}

func Deletebyid(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprint(w, "Method not allowed...")
			return
		}

		id := r.URL.Path[len("/deleteById/"):]
		_, err := db.Exec("DELETE FROM books WHERE Book_id = ?", id)
		if err != nil {
			fmt.Fprint(w, err.Error())
			return
		}
		fmt.Fprint(w, "deleted book by id :-", id)
	}
}
