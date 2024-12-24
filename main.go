package main

import (
	"fmt"
)

func main() {
	fmt.Println("main.go runs")
	customers := Customers{}
	vehicles := Vehicles{}

	// customers.AddCustomer("JEvgen", "liberal", "+37256660000", "evgen.liberal13@gmail.com", 39011212345)

	customers.AddCustomer("JEvgen", "liberal", "+37256660000", "evgen.liberal13@gmail.com", 39011212345)

	// CHECK DUPLICATED CUSTOMER ADITION
	// customers.AddCustomer("JEvgen", "liberal", "+37256660000", "evgen.liberal13@gmail.com", 39011212345)

	fmt.Println("----------------------------------AFTER---------------------------------")
	fmt.Println("customers:", customers)
	fmt.Println("vehicles:", vehicles)

}
