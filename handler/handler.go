package handler

import (
	"fmt"

	"github.com/GO-JEK-CAR/common/constants"
	"github.com/GO-JEK-CAR/models"
)

type ParkingAreaModel struct {
	ParkingArea *models.ParkingArea
}

// this will create a parking area with parking spaces equal to totalParkingSlots
func CreateParkingArea(totalParkingSlots int64) *models.ParkingArea {
	res := models.CreateParkingArea(totalParkingSlots)
	fmt.Println("Parking Area is created")
	return res
}

/* This function use to park the car
1. performs validation whether car with given registration number is already parked or not
2. parks the car in the empty space
*/
func Park(area *ParkingAreaModel, car *models.Car) error {
	// Check if car with same reg no is already in input
	if _, err := models.FindSlotNumberByRegNumber(area.ParkingArea, car.RegistrationNumber); err == nil {
		return constants.CAR_ALREADY_PARKED
	}

	if err := models.Park(area.ParkingArea, car); err != nil {
		return err
	}

	return nil
}

// Remove car from a slot number
func RemoveCarFromSlotNumber(area *ParkingAreaModel, slotNumber int64) error {
	err := models.RemoveBySlotNumber(area.ParkingArea, slotNumber)
	if err != nil {
		return err
	}

	fmt.Printf("Slot number %+v is free", slotNumber)
	return nil

}

// remove car with given registration number
func RemoveCarFromRegNumber(area *ParkingAreaModel, regNumber string) error {
	err := models.RemoveByRegNumber(area.ParkingArea, regNumber)
	if err != nil {
		return err
	}
	fmt.Printf("Slot number %+v is free", regNumber)
	return nil
}

func FindSlotNumberByRegNumber(area *ParkingAreaModel, regNumber string) error {
	parkingSlot, err := models.FindSlotNumberByRegNumber(area.ParkingArea, regNumber)
	if err != nil {
		return err
	}
	fmt.Printf("Slot number is %+v ", parkingSlot.SlotNumber)
	return err

}

func FindAllSlotNumbersByColor(area *ParkingAreaModel, color string) error {

	parkingSlotList, err := models.FindAllSlotNumbersByColor(area.ParkingArea, color)
	if err != nil {
		return err
	}

	for each := range parkingSlotList {
		fmt.Printf("%+v ", parkingSlotList[each].SlotNumber)
	}
	fmt.Print("\n")

	return nil

}

func FindAllRegNoByColor(area *ParkingAreaModel, color string) error {

	parkingSlots, err := models.FindAllSlotNumbersByColor(area.ParkingArea, color)
	if err != nil {
		return err
	}

	for i := 0; i < len(parkingSlots); i++ {
		fmt.Printf("%+v ", parkingSlots[i].Car.RegistrationNumber)
	}
	fmt.Print("\n")
	return nil
}

func GetParkingAreaStatus(area *ParkingAreaModel) error {
	if area == nil {
		return constants.NO_PARKING_AREA
	}

	current := area.ParkingArea.ParkingSlotList
	fmt.Println("Slot Number. Registration No. Color")
	for current != nil {
		car := current.Car
		fmt.Printf("%+v \t\t %+v \t\t %+v\n", current.SlotNumber, car.RegistrationNumber, car.Color)
		current = current.NextParkingSlot
	}
	return nil
}
