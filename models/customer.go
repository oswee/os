package models

import (
	"database/sql"
	"fmt"
	"log"
)

// Customer struct
type Customer struct {
	ID        int    `json:"ID"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
}

// ListCustomers function return list of all customers
func ListCustomers() ([]Customer, error) {

	getCustomers := `SELECT
			id,
			name,
			created_at
		FROM os_customers`

	db := dbLoc()
	rows, err := db.Query(getCustomers)
	// ToDo: Handle error
	checkErr(err)

	defer rows.Close()

	c := Customer{}
	cs := []Customer{}

	for rows.Next() {
		var (
			id        int
			name      string
			createdAt string
		)

		err = rows.Scan(&id, &name, &createdAt)
		// ToDo: Handle error
		checkErr(err)

		c.ID = id
		c.Name = name
		c.CreatedAt = createdAt

		cs = append(cs, c)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	return cs, err
}

// GetCustomer function return list of all customers
func GetCustomer(customerID string) Customer {

	var customer Customer

	getCustomer := `SELECT
			id,
			name,
			created_at
		FROM os_customers
		WHERE os_customers.id = ?`

	db := dbLoc()
	row := db.QueryRow(getCustomer, customerID)
	switch err := row.Scan(&customer.ID, &customer.Name, &customer.CreatedAt); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println("Customer:", customer, "were returned")
	default:
		panic(err)
	}
	defer db.Close()
	return customer
}

// CustomerCreate creates a new customer account
func CustomerCreate(accName string) {
	newCustomer := `INSERT INTO os_customers(name) VALUES(?)`
	db := dbLoc()
	sql, err := db.Prepare(newCustomer)
	if err != nil {
		log.Println(err)
	}
	sql.Exec(accName)
}
