package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/rs/cors"

	_ "github.com/lib/pq"
)

var db *sql.DB


func connectToDB() {
	var err error
	//My Database Details
	connStr := "user=postgres password=Marykiki dbname=store_db sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging database: ", err)
	}
	fmt.Println("Connected to the database!")
}

// Fetch all products (for the user side)
func fetchProducts(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, description, price FROM products")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []map[string]interface{}

	for rows.Next() {
		var id int
		var name, description string
		var price float64
		if err := rows.Scan(&id, &name, &description, &price); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		product := map[string]interface{} {
			"id":          id,
			"name":        name,
			"description": description,
			"price":       price,
		}
		products = append(products, product)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}


func placeOrder(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        var order struct {
            UserName string `json:"user_name"`
            Comment  string `json:"comment"`
            Items    []struct {
                ProductID int     `json:"product_id"`
                Quantity  int     `json:"quantity"`
                Price     float64 `json:"price"`
            } `json:"items"`
        }

        
        if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        
        fmt.Println("Received Order Data:")
        fmt.Printf("User Name: %s\n", order.UserName)
        fmt.Printf("Comment: %s\n", order.Comment)
        fmt.Println("Items:")

        var totalOrderPrice float64
        // Calculate the total price for the entire order
        for _, item := range order.Items {
            itemTotalPrice := float64(item.Quantity) * item.Price
            fmt.Printf("Item - Product ID: %d, Quantity: %d, Price: %.2f, Item Total Price: %.2f\n", 
                item.ProductID, item.Quantity, item.Price, itemTotalPrice)
            totalOrderPrice += itemTotalPrice
        }

        // Log the calculated total order price
        fmt.Printf("Total Order Price: %.2f\n", totalOrderPrice)

        
        for _, item := range order.Items {
            itemTotalPrice := float64(item.Quantity) * item.Price

            _, err := db.Exec("INSERT INTO orders (product_id, user_name, quantity, total_price, comment) VALUES ($1, $2, $3, $4, $5)",
                item.ProductID, order.UserName, item.Quantity, itemTotalPrice, order.Comment)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }

            
            fmt.Printf("Inserted Order Item - Product ID: %d, Quantity: %d, Price: %.2f, Total Price: %.2f\n", 
                item.ProductID, item.Quantity, item.Price, itemTotalPrice)
        }

        // Send success response
        w.WriteHeader(http.StatusCreated)
        fmt.Fprintf(w, "Order placed successfully")
    }
}


func fetchOrders(w http.ResponseWriter, r *http.Request) {
    
    rows, err := db.Query(`
        SELECT o.id, o.user_name, SUM(o.total_price) as total_price, o.status, o.comment, p.name AS product_name, p.price AS product_price
        FROM orders o
        JOIN products p ON o.product_id = p.id
        GROUP BY o.id, o.user_name, o.status, o.comment, p.name, p.price
    `)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var orders []map[string]interface{}

    
    for rows.Next() {
        var id int
        var userName, status, productName string
        var totalPrice, productPrice float64
        var comment sql.NullString
        if err := rows.Scan(&id, &userName, &totalPrice, &status, &comment, &productName, &productPrice); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        
        commentStr := ""
        if comment.Valid {
            commentStr = comment.String
        }

        
        order := map[string]interface{}{
            "id":            id,
            "user_name":     userName,
            "total_price":   totalPrice,
            "status":        status,
            "comment":       commentStr,
            "product_name":  productName,   
            "product_price": productPrice,  
        }
        orders = append(orders, order)
    }

    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(orders)
}


func main() {
	connectToDB()
	defer db.Close()

	// CORS configuration
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, //CORS 
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	// Routes for fetching products and placing orders
	http.HandleFunc("/api/products", fetchProducts) 
	http.HandleFunc("/api/place-order", placeOrder) 

	// Routes for order management (admin side)
	http.HandleFunc("/api/orders", fetchOrders) 
	

	// Wrap your handler with CORS middleware
	handler := corsHandler.Handler(http.DefaultServeMux)

	fmt.Println("Server is running at port 8080...")
	log.Fatal(http.ListenAndServe(":8080", handler))
}



