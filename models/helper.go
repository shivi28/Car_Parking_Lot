package models

import (
	"github.com/GO-JEK-CAR/common/constants"
)

//Count the number of occupied parking slots
func (area *ParkingArea) CountOccupiedSlots() int64 {
	if area.ParkingSlotList == nil {
		return 0
	}

	return area.ParkingSlotList.Count()
}

func (slot *ParkingSlot) Count() int64 {
	if slot == nil {
		return 0
	}

	if slot.NextParkingSlot == nil {
		return 1
	}

	return 1 + slot.NextParkingSlot.Count()
}

func (slot *ParkingSlot) RemoveParkingSlot(area *ParkingArea) error {
	if slot.PrevParkingSlot != nil {
		slot.PrevParkingSlot.NextParkingSlot = slot.NextParkingSlot
		return nil
	}

	slot = slot.NextParkingSlot
	slot.PrevParkingSlot = nil
	return nil
}

// find the reference of the parking slot space whose slot number is given
func (slot *ParkingSlot) FindParkingSlotBySlotNumber(slotNumber int64) (*ParkingSlot, error) {
	if slot.SlotNumber == slotNumber {
		return slot, nil
	}

	if slot.NextParkingSlot == nil {
		return &ParkingSlot{}, constants.NO_CAR_FOUND
	}

	return slot.NextParkingSlot.FindParkingSlotBySlotNumber(slotNumber)
}

// find the reference of the parking slot space in which the car with given registration number is parked
func (slot *ParkingSlot) FindParkingSlotByRegNumber(regNumber string) (*ParkingSlot, error) {
	if slot.Car.RegistrationNumber == regNumber {
		return slot, nil
	}

	if slot.NextParkingSlot == nil {
		return &ParkingSlot{}, constants.NO_CAR_FOUND
	}

	return slot.NextParkingSlot.FindParkingSlotByRegNumber(regNumber)
}

func (slot *ParkingSlot) AddNextSlot(sc *ParkingSlot) *ParkingSlot {
	// It takes a parking slot object sc with slot number 0, and update its slot number and add 1 in the slot number of previous
	// parking slot and assign it to sc parking slot
	if slot.NextParkingSlot == nil {
		slot.NextParkingSlot = sc.UpdateSlotNumber(slot.SlotNumber + 1)
		sc.PrevParkingSlot = slot
		return sc
	}

	// this is used when we have empty parking slot in between
	if slot.NextParkingSlot.SlotNumber > (slot.SlotNumber + 1) {
		currentNext := slot.NextParkingSlot
		slot.NextParkingSlot = sc.UpdateSlotNumber(slot.SlotNumber + 1)
		sc.PrevParkingSlot = slot
		sc.NextParkingSlot = currentNext
		currentNext.PrevParkingSlot = sc
		return sc
	}

	return slot.NextParkingSlot.AddNextSlot(sc)

}

func (s *ParkingSlot) UpdateSlotNumber(slotNumber int64) *ParkingSlot {
	s.SlotNumber = slotNumber
	return s
}

// find all parking slots where cars with given color are parked
func (s *ParkingSlot) FindAllCarsWithColor(color string) ([]*ParkingSlot, error) {

	if s.Car.Color == color {
		if s.NextParkingSlot == nil {
			return []*ParkingSlot{
				s,
			}, nil
		}

		slots, err := s.NextParkingSlot.FindAllCarsWithColor(color)
		if err == nil {
			slots = append([]*ParkingSlot{s}, slots...)
		}
		return slots, err
	}

	if s.NextParkingSlot == nil {
		return []*ParkingSlot{}, nil
	}

	return s.NextParkingSlot.FindAllCarsWithColor(color)
}
