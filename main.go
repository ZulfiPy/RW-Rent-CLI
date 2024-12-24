package main

import (
	"fmt"
)

func main() {
	fmt.Println("main.go runs")
	customers := Customers{}
	vehicles := Vehicles{}

	// customers.AddCustomer("Jevgeni", "Lebedev", "+37256713141", "eugene.lebedev13@gmail.com", 39011270287)

	customers.AddCustomer("Jevgeni", "Lebedev", "+37256713141", "eugene.lebedev13@gmail.com", 39011270287)

	// CHECK DUPLICATED CUSTOMER ADITION
	// customers.AddCustomer("Jevgeni", "Lebedev", "+37256713141", "eugene.lebedev13@gmail.com", 39011270287)

	fmt.Println("----------------------------------AFTER---------------------------------")
	fmt.Println("customers:", customers)
	fmt.Println("vehicles:", vehicles)

}
