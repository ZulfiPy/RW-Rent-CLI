package main

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/aquasecurity/table"
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
	var fuelType = []string{"Petrol", "Diesel", "Hybrid", "Electric", "Lpg", "Cng"}
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

	if !(slices.Contains(fuelType, cases.Title(language.English).String(input.FuelType))) {
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

func (vehicles *Vehicles) DeleteVehicle(plateNumber string) error {
	v := *vehicles
	vehiclePersists := v.ensureAbsence(plateNumber)

	if vehiclePersists == nil {
		return fmt.Errorf("Vehicle with plate number %v doesn't persist in vehicles", plateNumber)
	}

	delete(v, plateNumber)

	return nil
}

func (vehicles *Vehicles) PrintVehiclesTable() {
	v := *vehicles
	t := table.New(os.Stdout)

	t.SetHeaders("#", "Plate Number", "Make", "Model", "Year", "Fuel Type", "Gearbox", "Color", "Body")

	idx := 0
	for _, vehicle := range v {
		t.AddRow(strconv.FormatInt(int64(idx+1), 10), vehicle.PlateNumber, vehicle.Make, vehicle.Model, strconv.FormatInt(int64(vehicle.Year), 10), vehicle.FuelType, vehicle.Gearbox, vehicle.Color, vehicle.Body)

		idx++
	}

	t.Render()
}
