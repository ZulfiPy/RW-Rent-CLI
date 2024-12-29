package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CustomerCmdFlags struct {
	AddCustomer               string
	AddVehicleToCustomer      string
	EditCustomer              string
	DeleteCustomer            int64
	DeleteVehicleFromCustomer string
	ListCustomers              bool
}

func CustomersNewCmdFlags() *CustomerCmdFlags {
	cf := CustomerCmdFlags{}

	flag.StringVar(&cf.AddCustomer, "addCustomer", "", "Add a new customer using the following format --> firstName:lastName:phoneNumber:email:personalID")
	flag.StringVar(&cf.AddVehicleToCustomer, "addVehicleToCustomer", "", "Add an existing vehicle to a customer using the following format --> personalID:plateNumber")
	flag.StringVar(&cf.EditCustomer, "editCustomer", "", "Edit a customer's contact data using their personal ID using the following format --> personalID:phoneNumber:email")
	flag.Int64Var(&cf.DeleteCustomer, "deleteCustomer", 0, "Delete a customer using their personal ID using the following format --> personalID")
	flag.StringVar(&cf.DeleteVehicleFromCustomer, "deleteVehicleFromCustomer", "", "Delete a vehicle from customer's rented cars")
	flag.BoolVar(&cf.ListCustomers, "listCustomers", false, "List all customers")

	// flag.Parse()

	return &cf
}

func (cf *CustomerCmdFlags) Execute(customers *Customers, vehicles Vehicles) {
	switch {
	case cf.ListCustomers:
		customers.PrintCustomersTable()
	case cf.AddCustomer != "":
		parts := strings.SplitN(cf.AddCustomer, ":", 5)

		if len(parts) != 5 {
			fmt.Println("Error, invalid format for adding a new customer. Please provide the whole personal information.")
			os.Exit(1)
		}

		firstName := parts[0]
		lastName := parts[1]
		phoneNumber := parts[2]
		email := parts[3]
		personalID, _ := strconv.Atoi(parts[4])

		customers.AddCustomer(firstName, lastName, phoneNumber, email, int64(personalID))
	case cf.EditCustomer != "":
		fmt.Println("cf.EditCustomer", cf.EditCustomer)
		parts := strings.SplitN(cf.EditCustomer, ":", 3)

		if len(parts) != 3 {
			fmt.Println("Error, invalid format for editting the customer. Don't forget to sepparate a phone number and email with :")
			os.Exit(1)
		}

		personalID, _ := strconv.Atoi(parts[0])
		phoneNumber := parts[1]
		email := parts[2]

		customers.EditCustomerContacts(int64(personalID), phoneNumber, email)
	case cf.DeleteCustomer != 0:
		personalID := cf.DeleteCustomer
		if intLength(personalID) != 11 {
			fmt.Println("Error, invalid format for deleting the customer. The personal ID must be exactly 11 digits long.")
			os.Exit(1)
		}

		customers.DeleteCustomerByPersonalID(personalID)
	case cf.AddVehicleToCustomer != "":
		parts := strings.SplitN(cf.AddVehicleToCustomer, ":", 2)

		if len(parts) != 2 {
			fmt.Println("Error, invalid format for adding a vehicle to the customer. Use the following format --> personalID:plateNumber")
			os.Exit(1)
		}

		personalID, _ := strconv.Atoi(parts[0])
		plateNumber := parts[1]

		customers.AddVehicleToCustomer(int64(personalID), plateNumber, vehicles)
	case cf.DeleteVehicleFromCustomer != "":
		parts := strings.SplitN(cf.DeleteVehicleFromCustomer, ":", 2)

		if len(parts) != 2 {
			fmt.Println("Error, invalid format for deleting a vehicle from the customer. Use the following format --> personalID:plateNumber")
			os.Exit(1)
		}

		personalID, _ := strconv.Atoi(parts[0])
		plateNumber := parts[1]

		customers.DeleteVehicleFromCustomer(int64(personalID), plateNumber, vehicles)

	default:
		fmt.Println("invalid command")
	}
}
