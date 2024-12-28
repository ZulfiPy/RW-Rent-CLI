package main

import (
	"errors"
	"fmt"
	"net/mail"
	"os"
	"slices"
	"strconv"
	"time"
	"unicode"

	"github.com/aquasecurity/table"
)

type Customer struct {
	FirstName    string
	LastName     string
	PersonalID   int64
	PhoneNumber  string
	Email        string
	RentedCars   []Vehicle
	CreatedAt    time.Time
	LastEditedAt *time.Time
}

type Customers []Customer

func intLength(number int64) int {
	if number == 0 {
		return 1
	}

	length := 0

	for number != 0 {
		number /= 10
		length++
	}

	return length
}

func (customers *Customers) validateInput(firstName, lastName, phoneNumber, email string, personalID int64) error {
	if firstName == "" || len(firstName) < 3 {
		return errors.New("invalid input: first name cannot be empty or shorter than 3 characters")
	}

	if lastName == "" || len(lastName) < 3 {
		return errors.New("invalid input: last name cannot be empty or shorter than 3 characters")
	}

	if phoneNumber == "" || len(phoneNumber) < 7 {
		return errors.New("invalid input: phone number cannot be empty or shorter than 7")
	}

	for _, char := range phoneNumber {
		isChar := unicode.IsLetter(char)

		if isChar {
			return errors.New("invalid input: phone number cannot consist letters")
		}
	}

	if email == "" || len(email) < 7 {
		return errors.New("invalid input: email cannot be empty or shorter than 7 characters")
	}

	_, err := mail.ParseAddress(email)

	if err != nil {
		return fmt.Errorf("invalid %v", err)
	}

	personalIDLen := intLength(personalID)

	if personalIDLen != 11 {
		return errors.New("invalid input: personal id of the customer must be exactly 11 digits")
	}

	return nil
}

func (customers *Customers) validateIdx(idx int) error {
	c := *customers

	if idx < 0 || idx >= len(c) {
		return errors.New("error: idx is out of range")
	}

	return nil
}

func (customers *Customers) FindCustomerByPersonalID(personalID int64) (*Customer, int) {
	c := *customers

	for idx, customer := range c {
		if customer.PersonalID == personalID {
			return &c[idx], idx
		}
	}

	return nil, -1
}

func (customers *Customers) AddCustomer(firstName, lastName, phoneNumber, email string, personalID int64) error {
	duplicatedCustomer, _ := customers.FindCustomerByPersonalID(personalID)

	if duplicatedCustomer != nil {
		return errors.New("error: duplicated customer found")
	}

	validatedCustomer := customers.validateInput(firstName, lastName, phoneNumber, email, personalID)

	if validatedCustomer != nil {
		return validatedCustomer
	}

	newCustomer := Customer{
		FirstName:    firstName,
		LastName:     lastName,
		PersonalID:   int64(personalID),
		PhoneNumber:  phoneNumber,
		Email:        email,
		RentedCars:   []Vehicle{},
		CreatedAt:    time.Now(),
		LastEditedAt: nil,
	}

	*customers = append(*customers, newCustomer)

	return nil
}

func (customers *Customers) DeleteCustomerByIdx(idx int) error {
	c := *customers
	validatedIdx := c.validateIdx(idx)

	if validatedIdx == nil {
		*customers = append(c[:idx], c[idx+1:]...)
		return nil
	}

	return fmt.Errorf("unsuccessful validation of idx %v", idx)
}

func (customers *Customers) DeleteCustomerByPersonalID(personalID int64) error {
	c := *customers
	foundCustomer, idx := c.FindCustomerByPersonalID(personalID)

	if foundCustomer == nil {
		fmt.Println("Customer not found")
		return fmt.Errorf("customer with PersonalID %d not found", personalID)
	}

	*customers = append(c[:idx], c[idx+1:]...)

	return nil
}

func (customers *Customers) EditCustomerContacts(personalID int64, phoneNumber, email string) error {
	c := *customers

	customer, idx := c.FindCustomerByPersonalID(personalID)

	if idx == -1 && customer == nil {
		fmt.Println("Customer not found")
		return fmt.Errorf("Customer not found by PersonalID %d", personalID)
	}

	if phoneNumber != "" && len(phoneNumber) >= 7 {
		customer.PhoneNumber = phoneNumber
	}

	_, err := mail.ParseAddress(email)

	if email != "" && len(email) >= 7 && err == nil {
		customer.Email = email
	}

	editTime := time.Now()
	customer.LastEditedAt = &editTime

	return nil
}

func (customers *Customers) AddVehicleToCustomer(personalID int64, plateNumber string, vehicles Vehicles) error {
	c := *customers

	customer, idx := c.FindCustomerByPersonalID(personalID)

	if customer == nil && idx == -1 {
		return fmt.Errorf("Customer with personalID %v doesn't pesists", personalID)
	}

	vehiclePersists := vehicles.ensureAbsence(plateNumber)

	if vehiclePersists == nil {
		return fmt.Errorf("Vehicle with plate number %v doesn't persist", plateNumber)
	}

	if slices.Contains(customer.RentedCars, vehicles[plateNumber]) {
		return fmt.Errorf("Vehicle with plate number %v already exists in customer's rented vehicles", plateNumber)
	}

	customer.RentedCars = append(customer.RentedCars, vehicles[plateNumber])

	return nil
}

func (customers *Customers) DeleteVehicleFromCustomer(personalID int64, plateNumber string, vehicles Vehicles) error {
	c := *customers

	customer, idx := c.FindCustomerByPersonalID(personalID)

	if customer == nil && idx == -1 {
		return fmt.Errorf("Customer with personalID %v doesn't exist", personalID)
	}

	vehiclePersists := vehicles.ensureAbsence(plateNumber)

	if vehiclePersists == nil {
		return fmt.Errorf("Vehicle with plate number %v doesn't persists", plateNumber)
	}

	if len(customer.RentedCars) == 0 {
		return fmt.Errorf("Customer %v %v doesn't have any cars, deletion cannot be executed", customer.FirstName, customer.LastName)
	}

	if !(slices.Contains(customer.RentedCars, vehicles[plateNumber])) {
		return fmt.Errorf("Vehicle with plate number %v doesn't persist in %v %v rented vehicles lists", plateNumber, customer.FirstName, customer.LastName)
	}

	for idx, vehicle := range customer.RentedCars {
		if vehicle.PlateNumber == plateNumber {
			customer.RentedCars = append(customer.RentedCars[:idx], customer.RentedCars[idx+1:]...)
			return nil
		}
	}

	return nil
}

func (customers *Customers) PrintCustomersTable() {
	t := table.New(os.Stdout)
	c := *customers

	t.SetHeaders("#", "First Name", "Last Name", "Personal ID", "Phone Number", "Email", "Rented Cars", "Created At", "Last Edited")

	for idx, customer := range c {
		lastEdited := "❌"

		if customer.LastEditedAt != nil {
			lastEdited = customer.LastEditedAt.Format(time.RFC1123)
		}

		vehicle := "❌"

		if len(customer.RentedCars) == 1 {
			vehicle = customer.RentedCars[0].PlateNumber
		}

		t.AddRow(strconv.FormatInt(int64(idx+1), 10), customer.FirstName, customer.LastName, strconv.FormatInt(int64(customer.PersonalID), 10), customer.PhoneNumber, customer.Email, vehicle, customer.CreatedAt.Format(time.RFC1123), lastEdited)
	}

	t.Render()
}
