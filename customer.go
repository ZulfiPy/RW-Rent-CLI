package main

import (
	"errors"
	"fmt"
	"net/mail"
	"time"
	"unicode"
)

type Customer struct {
	FirstName    string
	LastName     string
	PersonalID   int
	PhoneNumber  string
	Email        string
	RentedCars   []Vehicle
	CreatedAt    time.Time
	LastEditedAt *time.Time
}

type Customers []Customer

func intLength(number int) int {
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

func (customers *Customers) validateInput(FirstName, LastName, PhoneNumber, Email string, PersonalID int) error {
	if FirstName == "" || len(FirstName) <= 3 {
		return errors.New("invalid input: first name cannot be empty or shorter than 3 characters")
	}

	if LastName == "" || len(LastName) <= 3 {
		return errors.New("invalid input: last name cannot be empty or shorter than 3 characters")
	}

	if PhoneNumber == "" || len(PhoneNumber) < 7 {
		return errors.New("invalid input: phone number cannot be empty or shorter than 7")
	}

	for _, char := range PhoneNumber {
		isChar := unicode.IsLetter(char)

		if isChar {
			return errors.New("invalid input: phone number cannot consist letters")
		}
	}

	if Email == "" || len(Email) < 7 {
		return errors.New("invalid input: email cannot be empty or shorter than 7 characters")
	}

	_, err := mail.ParseAddress(Email)

	if err != nil {
		return fmt.Errorf("invalid %v", err)
	}

	personalIDLen := intLength(PersonalID)

	if personalIDLen != 11 {
		return errors.New("invalid input: personal id of the customer must be exactly 11 digits")
	}

	return nil
}

func (customers *Customers) findCustomerByPersonalID(PersonalID int) *Customer {
	c := *customers

	for _, customer := range c {
		if customer.PersonalID == PersonalID {
			return &customer
		}
	}

	return nil
}

func (customers *Customers) AddCustomer(FirstName, LastName, PhoneNumber, Email string, PersonalID int) error {
	duplicatedCustomer := customers.findCustomerByPersonalID(PersonalID)

	if duplicatedCustomer != nil {
		return errors.New("error: duplicated customer found")
	}

	validatedCustomer := customers.validateInput(FirstName, LastName, PhoneNumber, Email, PersonalID)

	if validatedCustomer != nil {
		return validatedCustomer
	}

	newCustomer := Customer{
		FirstName:    FirstName,
		LastName:     LastName,
		PersonalID:   PersonalID,
		PhoneNumber:  PhoneNumber,
		Email:        Email,
		RentedCars:   []Vehicle{},
		CreatedAt:    time.Now(),
		LastEditedAt: nil,
	}

	*customers = append(*customers, newCustomer)

	return nil

}
