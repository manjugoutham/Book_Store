package userRegistrationControl

import (
	tp "bookstore/types"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func Register(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprint(w, "POST method not allowed ...!")
			return
		}

		var data tp.User

		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			fmt.Println("err in decoding", err)
			return
		}

		res, err := db.Exec("insert into users (name, email, password) values(?,?,?)", data.Name, data.Email, data.Password)
		if err != nil {
			fmt.Println("err in inserting", err)
			return
		}

		id, err := res.LastInsertId()
		if err != nil {
			fmt.Println("err in last inseret index", err)
			return
		}

		data.Id = int(id)

		json.NewEncoder(w).Encode(data)

	}

}

func AllRecord(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprint(w, "Method not Allowed...!")
			return
		}

		rows, err := db.Query("select * from users")
		if err != nil {
			fmt.Println("err in reading", err)
			return
		}

		defer rows.Close()

		var data []tp.User
		for rows.Next() {
			var val tp.User
			err := rows.Scan(&val.Id, &val.Name, &val.Email, &val.Password)
			if err != nil {
				fmt.Println("err in scan row", err)
				return
			}
			data = append(data, val)
		}

		// Now write the reponse to client
		fmt.Fprint(w, data)
	}

}

func GetUserById(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[len("/userById/"):]

		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprint(w, "Method not Allowed...!")
			return
		}

		row := db.QueryRow("select * from users where id = ?", id)

		var userOfId tp.User

		err := row.Scan(&userOfId.Id, &userOfId.Name, &userOfId.Email, &userOfId.Password)
		if err != nil {
			fmt.Fprint(w, "No user available by this id...!")
			return
		}
		fmt.Fprint(w, userOfId)
	}
}

func UpdateId(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id := r.URL.Path[len("/editUser/"):]

		if r.Method != "PUT" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Println("Method not Allowed...!")
			return
		}
		var u tp.User
		json.NewDecoder(r.Body).Decode(&u)

		//row := db.QueryRow("select * from bookuser where id =? ", id)

		_, err := db.Exec("UPDATE users SET name = ? ,email = ? , password = ?, phone = ? Where id = ?", u.Name, u.Email, u.Password, id)
		if err != nil {
			fmt.Fprint(w, err.Error())
			return
		}
		fmt.Fprint(w, "Updated DB")

	}
}

func DeleteById(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id := r.URL.Path[len("/deleteId/"):]

		if r.Method != "DELETE" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprint(w, "Method not Allowed...!")
			return
		}

		_, err := db.Query("delete from users where id = ?", id)
		if err != nil {
			fmt.Println("err in delete ", err)
			return
		}
		fmt.Fprint(w, "deleted user ", id)

	}
}
