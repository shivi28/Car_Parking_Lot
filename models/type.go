package models

type ParkingArea struct {
	TotlaParkingSlot int64
	ParkingSlotList  *ParkingSlot
}

type ParkingSlot struct {
	Car             *Car
	PrevParkingSlot *ParkingSlot
	NextParkingSlot *ParkingSlot
	SlotNumber      int64
}

type Car struct {
	RegistrationNumber string
	Color              string
}
