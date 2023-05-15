package orderControllerPackage

import (
	tp "bookstore/types"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func NewOrder(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprint(w, "Method Not Allowed...!")
			return
		}
		var order tp.Orders
		json.NewDecoder(r.Body).Decode(&order)
		res, err := db.Exec("insert into orders (userId, BookId, orderDate, quantity, order_status)values(?,?,?,?,?)", order.UserId, order.BookId, order.OrderDate, order.Quantity, order.Status)
		if err != nil {
			fmt.Fprint(w, "error in inserting", err)
			return
		}
		id, err := res.LastInsertId()
		if err != nil {
			fmt.Fprint(w, "error in last inserted index", err)
			return
		}

		order.OrderId = int(id)
		json.NewEncoder(w).Encode(order)
	}
}

func GetAllOrders(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprint(w, "Method Not Allowed...!")
			return
		}

		row, err := db.Query("select * from orders")
		if err != nil {
			fmt.Fprint(w, "error in query", err)
			return
		}

		var rows []tp.Orders

		for row.Next() {
			var val tp.Orders
			row.Scan(&val.OrderId, &val.UserId, &val.BookId, &val.OrderDate, &val.Quantity, &val.Status)
			rows = append(rows, val)
		}
		json.NewEncoder(w).Encode(rows)
	}
}

func GetByOrderId(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprint(w, "Method not allowed...!")
			return
		}

		var order tp.Orders
		id := r.URL.Path[len("/getOrderById/"):]

		rows := db.QueryRow("select * from orders WHERE orderId = ?", id)

		err := rows.Scan(&order.OrderId, &order.BookId, &order.UserId, &order.OrderDate, &order.Quantity, &order.Status)
		if err != nil {
			fmt.Fprint(w, err.Error())
		} else {
			json.NewEncoder(w).Encode(order)

		}

	}
}

func CancelOrder(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "PUT" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprint(w, "Method Not Allowed...!")
		}

		orderId := r.URL.Path[len("/cancelOrder/"):]

		var data tp.Orders
		json.NewDecoder(r.Body).Decode(&data)

		_, err := db.Exec("update  Orders set order_status = ? WHERE orderId = ?", data.Status, orderId)
		if err != nil {
			fmt.Fprint(w, "order with given id not availabe")
		} else {
			fmt.Fprint(w, "order Updated")
		}
	}
}
