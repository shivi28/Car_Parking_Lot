package models

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

var car1 = &Car{
	RegistrationNumber: "GO_JEK_1111",
	Color:              "White",
}

var car2 = &Car{
	RegistrationNumber: "GO_JEK_2222",
	Color:              "Red",
}

var parkingSlot2 = &ParkingSlot{
	SlotNumber: 2,
	Car:        car2,
}

var parkingSlot1 = &ParkingSlot{
	SlotNumber:      1,
	Car:             car1,
	NextParkingSlot: parkingSlot2,
}

var area = &ParkingArea{
	TotlaParkingSlot: 5,
	ParkingSlotList:  parkingSlot1,
}

func TestCountOccupiedSlots(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(int64(2), area.CountOccupiedSlots())

}

func TestFindParkingSlotBySlotNumber(t *testing.T) {
	assert := assert.New(t)
	slot, err := area.ParkingSlotList.FindParkingSlotBySlotNumber(2)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(parkingSlot2, slot)
}

func TestUpdateSlotNumber(t *testing.T) {
	assert := assert.New(t)
	slot := &ParkingSlot{
		SlotNumber: 1,
	}
	res := slot.UpdateSlotNumber(slot.SlotNumber + 1)
	assert.Equal(int64(2), res.SlotNumber)

}

func TestFindParkingSlotByRegNumber(t *testing.T) {
	assert := assert.New(t)

	slot, err := area.ParkingSlotList.FindParkingSlotByRegNumber("GO_JEK_2222")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(int64(2), slot.SlotNumber)
}

func TestFindAllCarsWithColor(t *testing.T) {
	assert := assert.New(t)

	slot, err := area.ParkingSlotList.FindAllCarsWithColor("White")
	if err != nil {
		t.Error(err)
	}

	assert.Equal([]*ParkingSlot{parkingSlot1}, slot)
}

func TestAddNextSlot(t *testing.T) {
	assert := assert.New(t)

	car := &Car{
		Color:              "Red",
		RegistrationNumber: "GO_JEK_3333",
	}

	slot := &ParkingSlot{
		SlotNumber: 0,
		Car:        car,
	}

	slot = area.ParkingSlotList.AddNextSlot(slot)
	assert.Equal(int64(3), slot.SlotNumber)
}
