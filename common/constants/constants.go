package constants

import (
	"errors"
)

var (
	CommandSeparator = " "

	Tab = "\t"

	NO_FREE_SLOT = errors.New("Sorry, parking lot is full")

	NO_CARS_PARKED = errors.New("No cars parked")

	NO_CAR_FOUND = errors.New("Not found")

	NO_PARKING_AREA = errors.New("No Parking Area")

	NO_CAR_WITH_REG_NO_FOUND = errors.New("Car with specified registration number not found")

	NO_CAR_WITH_COLOR_FOUND = errors.New("Car with specified color not found")

	CAR_ALREADY_PARKED = errors.New("Car with given reg number is already parked")
)
