# UdacityCRMBackendProject

This is a simple RESTful API for managing customer data.

Prerequisites
~~~~~~~~~~~~~
    Go 1.16 or higher

Installation
~~~~~~~~~~~~
    Clone the repository: 
    git clone https://github.com/example/customer-api.git

Build the application:
    cd customer-api
    // Install mux package in local workspace and then build
    go get github.com/gorilla/mux
    go build

Run the application:
    ./customer-api

Usage
~~~~~
The API provides the following endpoints:

GET /customers 
(Returns a list of all customers.)

GET /customers/{id} 
(Returns a single customer with the specified ID.)

POST /customers 
(Creates a new customer.)

PUT /customers/{id} 
(Updates an existing customer with the specified ID.)

DELETE /customers/{id} (Deletes an existing customer with the specified ID.)

Built With
~~~~~~~~~~
    Go - Programming language used
    Gorilla/mux - Router used for HTTP method-based routing and variables in URL paths
