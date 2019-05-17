package models

import (
	"fmt"
	"github.com/GO-JEK-CAR/common/constants"
)

func CreateParkingArea(totalParkingSlot int64) *ParkingArea {
	newParkingArea := &ParkingArea{
		TotlaParkingSlot: totalParkingSlot,
	}
	return newParkingArea
}

func Park(area *ParkingArea, car *Car) error {
	maxSlots := area.TotlaParkingSlot

	occupiedSlotsCount := area.CountOccupiedSlots()

	if occupiedSlotsCount >= maxSlots {
		return constants.NO_FREE_SLOT
	}

	var newParkingSlot, currentParkingSlot *ParkingSlot
	if occupiedSlotsCount == 0 {
		newParkingSlot = &ParkingSlot{
			Car:        car,
			SlotNumber: 1,
		}
		area.ParkingSlotList = newParkingSlot
		fmt.Printf("Allocated parking slot: %+v\n", area.ParkingSlotList.SlotNumber)
		return nil
	}

	if area.ParkingSlotList.SlotNumber > 1 {
		currentParkingSlot = area.ParkingSlotList
		newSlot := &ParkingSlot{
			Car:        car,
			SlotNumber: 1,
		}
		area.ParkingSlotList = newSlot
		area.ParkingSlotList.AddNextSlot(currentParkingSlot)
		currentParkingSlot.PrevParkingSlot = area.ParkingSlotList

	}

	newSlot := &ParkingSlot{
		Car:        car,
		SlotNumber: 0,
	}

	area.ParkingSlotList.AddNextSlot(newSlot)
	fmt.Printf("Allocated parking slot: %+v\n", newSlot.SlotNumber)

	return nil
}

func RemoveBySlotNumber(area *ParkingArea, slotNumber int64) error {

	//slot := area.ParkingSlotList
	if area == nil {
		return constants.NO_PARKING_AREA
	}

	if area.ParkingSlotList == nil {
		return constants.NO_CARS_PARKED
	}

	curr := area.ParkingSlotList
	for curr != nil {
		if curr.SlotNumber == slotNumber && curr.PrevParkingSlot != nil {
			temp := curr.PrevParkingSlot
			curr.PrevParkingSlot.NextParkingSlot = curr.NextParkingSlot
			curr.NextParkingSlot.PrevParkingSlot = temp
			break
		} else if curr.SlotNumber == slotNumber && curr.PrevParkingSlot == nil {
			area.ParkingSlotList = curr.NextParkingSlot
			break
		} else {
			curr = curr.NextParkingSlot
		}
	}
	return nil

}

func RemoveByRegNumber(area *ParkingArea, regNumber string) error {
	if area == nil {
		return constants.NO_PARKING_AREA
	}

	if area.ParkingSlotList == nil {
		return constants.NO_CARS_PARKED
	}

	curr := area.ParkingSlotList
	for curr != nil {
		if curr.Car.RegistrationNumber == regNumber && curr.PrevParkingSlot != nil {
			temp := curr.PrevParkingSlot
			curr.PrevParkingSlot.NextParkingSlot = curr.NextParkingSlot
			curr.NextParkingSlot.PrevParkingSlot = temp
			break
		} else if curr.Car.RegistrationNumber == regNumber && curr.PrevParkingSlot == nil {
			area.ParkingSlotList = curr.NextParkingSlot
			break
		} else {
			curr = curr.NextParkingSlot
		}
	}
	return nil
}

func FindSlotNumberByRegNumber(area *ParkingArea, regNumber string) (*ParkingSlot, error) {
	if area.ParkingSlotList == nil {
		return nil, constants.NO_CAR_WITH_REG_NO_FOUND
	}

	return area.ParkingSlotList.FindParkingSlotByRegNumber(regNumber)

}

func FindAllSlotNumbersByColor(area *ParkingArea, color string) ([]ParkingSlot, error) {

	if area.ParkingSlotList == nil {
		return []ParkingSlot{}, constants.NO_CARS_PARKED
	}

	slots, err := area.ParkingSlotList.FindAllCarsWithColor(color)
	if err != nil {
		return []ParkingSlot{}, err
	}

	if len(slots) == 0 {
		return []ParkingSlot{}, constants.NO_CAR_WITH_COLOR_FOUND
	}

	slotsList := []ParkingSlot{}

	for i := range slots {
		slotsList = append(slotsList, *slots[i])
	}
	return slotsList, nil

}
