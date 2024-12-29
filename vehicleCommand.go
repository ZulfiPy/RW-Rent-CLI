package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type VehicleCmdFlags struct {
	AddVehicle    string
	DeleteVehicle string
	ListVehicles  bool
}

func VehiclesNewCmdFlags() *VehicleCmdFlags {
	cf := VehicleCmdFlags{}

	flag.BoolVar(&cf.ListVehicles, "listVehicles", false, "List all vehicles")
	flag.StringVar(&cf.AddVehicle, "addVehicle", "", "Add a new vehicles into the vehicles map.")
	flag.StringVar(&cf.DeleteVehicle, "deleteVehicle", "", "Delete a vehicle from the vehicles map.")

	// flag.Parse()

	return &cf
}

func (cf *VehicleCmdFlags) Execute(vehicles *Vehicles) {
	switch {
	case cf.ListVehicles:
		vehicles.PrintVehiclesTable()
	case cf.AddVehicle != "":
		parts := strings.SplitN(cf.AddVehicle, ":", 8)

		if len(parts) != 8 {
			fmt.Println("Error, invalid format for adding a new vehicle. Please the whole personal information. Use the following format: plateNumber:Made:Model:year:gearbox:color:body")
			os.Exit(1)

		}
		year, _ := strconv.Atoi(parts[3])

		newVehicle := Vehicle{
			PlateNumber: parts[0],
			Make:        parts[1],
			Model:       parts[2],
			Year:        year,
			FuelType:    parts[4],
			Gearbox:     parts[5],
			Color:       parts[6],
			Body:        parts[7],
		}

		vehicles.AddVehicle(newVehicle)
	case cf.DeleteVehicle != "":
		plateNumber := strings.SplitN(cf.DeleteVehicle, "", 1)[0]
		vehicles.DeleteVehicle(plateNumber)
	default:
		fmt.Println("Invalid command")
	}
}
