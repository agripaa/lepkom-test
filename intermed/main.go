package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var db *sql.DB

type Order struct {
	ID            int    `json:"id"`
	Kursus        string `json:"kursus"`
	PurchaseTime  string `json:"purchase_time"`
	NamaPraktikan string `json:"nama_praktikan"`
	Status        string `json:"status"`
}

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Database connection
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize the router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/api/orders", getOrders).Methods("GET")
	router.HandleFunc("/api/order/{id}", getOrder).Methods("GET")
	router.HandleFunc("/api/order", createOrder).Methods("POST")
	router.HandleFunc("/api/order/{id}", updateOrder).Methods("PUT")
	router.HandleFunc("/api/order/{id}", deleteOrder).Methods("DELETE")

	// Apply middleware
	router.Use(AuthMiddleware)

	// Start the server
	log.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", router)
}

// Handler to get all orders
func getOrders(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM orders")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	orders := []Order{}
	for rows.Next() {
		var order Order
		if err := rows.Scan(&order.ID, &order.Kursus, &order.PurchaseTime, &order.NamaPraktikan, &order.Status); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		orders = append(orders, order)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

// Handler to get a specific order by ID
func getOrder(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var order Order
	err := db.QueryRow("SELECT * FROM orders WHERE id = ?", id).Scan(&order.ID, &order.Kursus, &order.PurchaseTime, &order.NamaPraktikan, &order.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

// Handler to create a new order
func createOrder(w http.ResponseWriter, r *http.Request) {
	var order Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := db.Exec("INSERT INTO orders (KURSUS, PURCHASE_TIME, NAMA_PRAKTIKAN, STATUS) VALUES (?, ?, ?, ?)", order.Kursus, order.PurchaseTime, order.NamaPraktikan, order.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	order.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

// Handler to update an order by ID
func updateOrder(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var order Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := db.Exec("UPDATE orders SET KURSUS = ?, PURCHASE_TIME = ?, NAMA_PRAKTIKAN = ?, STATUS = ? WHERE id = ?", order.Kursus, order.PurchaseTime, order.NamaPraktikan, order.Status, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	order.ID, _ = strconv.Atoi(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

// Handler to delete an order by ID
func deleteOrder(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	_, err := db.Exec("DELETE FROM orders WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Middleware for authentication
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		internalHeader := r.Header.Get("X-Internal-Request")

		if authHeader != "Bearer "+os.Getenv("AUTH_TOKEN") || internalHeader == "" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
