//package UdacityCRMProject
package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Customer struct {
	ID        int
	Name      string
	Role      string
	Email     string
	Phone     string
	Contacted bool
}

var database []Customer

func main() {
	// Add some customers to the database
	customer1 := Customer{ID: 1, Name: "Abhiram Ved", Role: "Employee", Email: "abhiram.ved@example.com", Phone: "666-4321", Contacted: true}
	customer2 := Customer{ID: 2, Name: "Shan Sa", Role: "Manager", Email: "shan.sa@example.com", Phone: "666-7894", Contacted: false}
	customer3 := Customer{ID: 3, Name: "Hari Sri", Role: "Employee", Email: "hari.sri@example.com", Phone: "666-4561", Contacted: true}
	customer4 := Customer{ID: 4, Name: "Shan Sa", Role: "Manager", Email: "shan.sa@example.com", Phone: "666-7894", Contacted: false}

	database = append(database, customer1, customer2, customer3, customer4)

	// Create a new router
	r := mux.NewRouter()

	// Define routes and handlers
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/customers", getCustomersHandler).Methods("GET")
	r.HandleFunc("/customers/{id}", getCustomerHandler).Methods("GET")
	r.HandleFunc("/customers", addCustomerHandler).Methods("POST")
	r.HandleFunc("/customers/{id}", updateCustomerHandler).Methods("PUT")
	r.HandleFunc("/customers/{id}", deleteCustomerHandler).Methods("DELETE")

	// Start server
	// log.Fatal(http.ListenAndServe(":8080", r))

	http.Handle("/", r)

    fmt.Println("Server listening on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    // Define the available endpoints
    endpoints := []struct {
        Method string
        Path   string
        Desc   string
    }{
        {"GET", "/customers", "Get all customers"},
        {"GET", "/customers/{id}", "Get a single customer"},
        {"POST", "/customers", "Add a new customer"},
        {"PUT", "/customers/{id}", "Update an existing customer"},
        {"DELETE", "/customers/{id}", "Delete a customer"},
    }

    // Generate the HTML response body
    html := "<h1>Welcome to the Customer API</h1>"
    html += "<h2>Available endpoints:</h2>"
    html += "<ul>"
    for _, ep := range endpoints {
        html += fmt.Sprintf("<li><strong>%s</strong> %s - %s</li>", ep.Method, ep.Path, ep.Desc)
    }
    html += "</ul>"

    // Set the response headers and write the response body
    w.Header().Set("Content-Type", "text/html")
    fmt.Fprint(w, html)
}

func getCustomersHandler(w http.ResponseWriter, r *http.Request) {
	// Convert database to JSON
	jsonData, err := json.Marshal(database)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write JSON data to response body
	w.Write(jsonData)
}

func getCustomerHandler(w http.ResponseWriter, r *http.Request) {
	// Get customer ID from URL parameter
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Find customer in database
	var customer Customer
	for _, c := range database {
		if c.ID == id {
			customer = c
			break
		}
	}

	// Return error if customer not found
	if customer.ID == 0 {
		http.NotFound(w, r)
		return
	}

	// Convert customer to JSON
	jsonData, err := json.Marshal(customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write JSON data to response body
	w.Write(jsonData)
}

func addCustomerHandler(w http.ResponseWriter, r *http.Request) {
	// Read request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Parse JSON data into a Customer struct
	var customer Customer
	err = json.Unmarshal(body, &customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generate ID for new customer
	if len(database) > 0 {
		customer.ID = database[len(database)-1].ID + 1
	}
}

func updateCustomerHandler(w http.ResponseWriter, r *http.Request) {
	// Get customer ID from URL parameter
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Find customer in database
	var customerIndex int
	for i, c := range database {
		if c.ID == id {
			customerIndex = i
			break
		}
	}

	// Return error if customer not found
	if customerIndex == 0 && database[customerIndex].ID != id {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Read request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Parse JSON data into a Customer struct
	var customer Customer
	err = json.Unmarshal(body, &customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update customer in database
	customer.ID = id
	database[customerIndex] = customer

	// Return updated customer as JSON
	jsonData, err := json.Marshal(customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write JSON data to response body
	w.Write(jsonData)
}

func deleteCustomerHandler(w http.ResponseWriter, r *http.Request) {
	// Get customer ID from URL parameter
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Find customer in database
	var customerIndex int
	for i, c := range database {
		if c.ID == id {
			customerIndex = i
			break
		}
	}

	// Return error if customer not found
	if customerIndex == 0 && database[customerIndex].ID != id {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Remove customer from database
	database = append(database[:customerIndex], database[customerIndex+1:]...)

	// Set status code to 204 No Content
	w.WriteHeader(http.StatusNoContent)
}