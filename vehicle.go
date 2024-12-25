package main

import (
	"errors"
	"fmt"
	"slices"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Vehicle struct {
	PlateNumber string
	Make        string
	Model       string
	Year        int
	FuelType    string
	Gearbox     string
	Color       string
	Body        string
}

type Vehicles map[string]Vehicle

func (vehicles *Vehicles) validateVehicle(input Vehicle) error {
	fmt.Println("validation of vehicle input...", input)
	var fuelType = []string{"Petrol", "Diesel", "Hybrid", "Electric", "LPG", "CNG"}
	var gearbox = []string{"Automatic", "Manual"}
	var colors = []string{"White", "Black", "Red", "Blue", "Green", "Yellow", "Gray", "Silver", "Brown"}
	var bodies = []string{"Sedan", "Touring", "Hatchback", "Minivan", "Coupe", "Cabriolet", "Pickup", "Limousine"}

	caser := cases.Title(language.English)

	if input.PlateNumber == "" {
		return errors.New("invalid input: vehicle plate number may not be empty")
	}

	if input.Make == "" {
		return errors.New("invalid input: vehicle make may not be empty")
	}

	if input.Model == "" {
		return errors.New("invalid input: vehicle model may not be empty")
	}

	if input.Year < 2010 || input.Year > time.Now().Year() {
		return errors.New("invalid input: vehicle year may not be lower than 2010 or greater than the current year")
	}

	if input.FuelType == "" {
		return errors.New("invalid input: vehicle fuel type may not be empty")
	}

	if !(slices.Contains(fuelType, caser.String(input.FuelType))) {
		return errors.New("invalid input: vehicle fuel type may only be (Petrol / Diesel / Hybrid / Electric / LPG / CNG)")
	}

	if input.Gearbox == "" {
		return errors.New("invalid input: vehicle gearbox may not be empty")
	}

	if !(slices.Contains(gearbox, caser.String(input.Gearbox))) {
		return errors.New("invalid input: vehicle gearbox may only be (Automatic or Manual)")
	}

	if input.Color == "" {
		return errors.New("invalid input: vehicle color may not be empty")
	}

	if !(slices.Contains(colors, caser.String(input.Color))) {
		return errors.New("invalid input: wrong vehicle color")
	}

	if input.Body == "" {
		return errors.New("invalid input: vehicle body may not be empty")
	}

	if !(slices.Contains(bodies, caser.String(input.Body))) {
		return errors.New("invalid input: wrong vehicle body")
	}

	return nil
}

func (vehicles *Vehicles) ensureAbsence(plateNumber string) error {
	_, ok := (*vehicles)[plateNumber]

	if ok {
		return fmt.Errorf("vehicle with plate number %v persists in vehicles", plateNumber)
	}

	return nil
}

func (vehicles *Vehicles) AddVehicle(input Vehicle) error {
	fmt.Println("I run in AddVehicle")
	v := *vehicles
	validatedVehicle := v.validateVehicle(input)

	if validatedVehicle != nil {
		return validatedVehicle
	}

	duplicatedVehicle := v.ensureAbsence(input.PlateNumber)

	if duplicatedVehicle != nil {
		return duplicatedVehicle
	}

	v[input.PlateNumber] = input

	return nil
}
