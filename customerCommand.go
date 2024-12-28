package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add           string
	AddVehicle    string
	Edit          string
	Delete        int64
	DeleteVehicle string
	List          bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new customer using the following format --> firstName:lastName:phoneNumber:email:personalID")
	flag.StringVar(&cf.AddVehicle, "addVehicle", "", "Add an existing vehicle to a customer using the following format --> personalID:plateNumber")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a customer's contact data using their personal ID using the following format --> personalID:phoneNumber:email")
	flag.Int64Var(&cf.Delete, "delete", 0, "Delete a customer using their personal ID using the following format --> personalID")
	flag.StringVar(&cf.DeleteVehicle, "deleteVehicle", "", "Delete a vehicle from customer's rented cars")
	flag.BoolVar(&cf.List, "list", false, "List all customers")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute(customers *Customers, vehicles Vehicles) {
	switch {
	case cf.List:
		customers.PrintCustomersTable()
	case cf.Add != "":
		parts := strings.SplitN(cf.Add, ":", 5)

		if len(parts) != 5 {
			fmt.Println("Error, invalid format for adding the customer. Please provide the whole personal information.")
			os.Exit(1)
		}

		firstName := parts[0]
		lastName := parts[1]
		phoneNumber := parts[2]
		email := parts[3]
		personalID, _ := strconv.Atoi(parts[4])

		customers.AddCustomer(firstName, lastName, phoneNumber, email, int64(personalID))
	case cf.Edit != "":
		fmt.Println("cf.Edit", cf.Edit)
		parts := strings.SplitN(cf.Edit, ":", 3)

		if len(parts) != 3 {
			fmt.Println("Error, invalid format for editting the customer. Don't forget to sepparate a phone number and email with :")
			os.Exit(1)
		}

		personalID, _ := strconv.Atoi(parts[0])
		phoneNumber := parts[1]
		email := parts[2]

		customers.EditCustomerContacts(int64(personalID), phoneNumber, email)
	case cf.Delete != 0:
		personalID := cf.Delete
		if intLength(personalID) != 11 {
			fmt.Println("Error, invalid format for deleting the customer. The personal ID must be exactly 11 digits long.")
			os.Exit(1)
		}

		customers.DeleteCustomerByPersonalID(personalID)
	case cf.AddVehicle != "":
		parts := strings.SplitN(cf.AddVehicle, ":", 2)

		if len(parts) != 2 {
			fmt.Println("Error, invalid format for adding a vehicle to the customer. Use the following format --> personalID:plateNumber")
			os.Exit(1)
		}

		personalID, _ := strconv.Atoi(parts[0])
		plateNumber := parts[1]

		customers.AddVehicleToCustomer(int64(personalID), plateNumber, vehicles)
	case cf.DeleteVehicle != "":
		parts := strings.SplitN(cf.DeleteVehicle, ":", 2)
		
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
