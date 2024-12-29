package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("main.go runs")
	customers := Customers{}
	vehicles := Vehicles{}
	customersCmdFlags := CustomersNewCmdFlags()
	vehiclesCmdFlags := VehiclesNewCmdFlags()

	flag.Parse()

	customersStorage := NewStorage[Customers]("customers.json")
	vehiclesStorage := NewStorage[Vehicles]("vehicles.json")

	customersStorage.Load(&customers)
	vehiclesStorage.Load(&vehicles)

	if customersCmdFlags.ListCustomers || customersCmdFlags.AddCustomer != "" || customersCmdFlags.AddVehicleToCustomer != "" || customersCmdFlags.EditCustomer != "" || customersCmdFlags.DeleteCustomer != 0 || customersCmdFlags.DeleteVehicleFromCustomer != "" {
		customersCmdFlags.Execute(&customers, vehicles)
	}

	if vehiclesCmdFlags.ListVehicles || vehiclesCmdFlags.AddVehicle != "" || vehiclesCmdFlags.DeleteVehicle != "" {
		vehiclesCmdFlags.Execute(&vehicles)
	}

	customersStorage.Save(customers)
	vehiclesStorage.Save(vehicles)

	// customers.PrintCustomersTable()
	// vehicles.PrintVehiclesTable()

	// fmt.Println("----------------------------------AFTER---------------------------------")
	// fmt.Println("customers:", customers)
	// fmt.Println("vehicles:", vehicles)
}
