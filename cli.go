package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/GO-JEK-CAR/cmd"
	"github.com/GO-JEK-CAR/common/constants"
	"github.com/GO-JEK-CAR/common/lib"
	"github.com/GO-JEK-CAR/handler"
	"github.com/GO-JEK-CAR/models"
)

func RunFileCommand(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var parkAreaModel *handler.ParkingAreaModel
	firstLine := true
	for scanner.Scan() {
		if firstLine {
			text := scanner.Text()
			command := parseCommand(text)
			if command[0] != cmd.CREATE_PARKING_LOT {
				panic("first command needs to be creating the parking area")
			}
			maxSlots, err := strconv.ParseInt(command[1], 10, 64)
			if err != nil {
				panic(err.Error())
			}

			area := handler.CreateParkingArea(int64(maxSlots))

			parkAreaModel = &handler.ParkingAreaModel{
				ParkingArea: area,
			}

			firstLine = false
			continue
		}
		commands := parseCommand(scanner.Text())
		processCommand(parkAreaModel, commands)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nil
}

func RunCliCommand() error {
	command := "Start"
	fmt.Println("\nStart Entering Input .................")

	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	text = strings.TrimRight(text, "\r\n")

	commands := parseCommand(text)
	if commands[0] != cmd.CREATE_PARKING_LOT {
		panic("first command needs to be creating the parking area")
	}

	maxSlots, err := lib.ConvertStringToInt(commands[1])

	if err != nil {
		panic(err.Error())
	}

	area := handler.CreateParkingArea(int64(maxSlots))

	parkAreaModel := &handler.ParkingAreaModel{
		ParkingArea: area,
	}

	for command != "Exit" {

		fmt.Println("\nInput")

		text, _ := reader.ReadString('\n')
		text = strings.TrimRight(text, "\r\n")

		commands := parseCommand(text)
		processCommand(parkAreaModel, commands)

		command = commands[0]
	}
	return nil
}

func parseCommand(command string) []string {
	parsedCommand := []string{}

	command = strings.Replace(command, constants.Tab, constants.CommandSeparator, -1)

	for _, s := range strings.Split(command, constants.CommandSeparator) {
		if s != "" {
			parsedCommand = append(parsedCommand, s)
		}
	}

	return parsedCommand
}

func processCommand(area *handler.ParkingAreaModel, command []string) {

	switch command[0] {

	case cmd.CREATE_PARKING_LOT:
		maxSlots, err := lib.ConvertStringToInt(command[1])
		if err != nil {
			panic(err.Error())
		}
		handler.CreateParkingArea(int64(maxSlots))

	case cmd.PARK:
		car := &models.Car{
			RegistrationNumber: command[1],
			Color:              command[2],
		}
		if err := handler.Park(area, car); err != nil {
			fmt.Printf("Error: %+v\n", err)
		}

	case cmd.STATUS:
		if err := handler.GetParkingAreaStatus(area); err != nil {
			fmt.Printf("Error: %+v\n", err)
		}

	case cmd.LEAVE:
		slotNumber, err := lib.ConvertStringToInt(command[1])
		if err != nil {
			panic(err.Error())
		}
		if err := handler.RemoveCarFromSlotNumber(area, int64(slotNumber)); err != nil {
			fmt.Printf("Error: %+v\n", err)
		}

	case cmd.REG_NO_OF_ALL_CARS_WITH_COLOR:
		if err := handler.FindAllRegNoByColor(area, command[1]); err != nil {
			fmt.Printf("Error: %+v\n", err)
		}

	case cmd.SLOT_NOS_BY_COLOR:
		if err := handler.FindAllSlotNumbersByColor(area, command[1]); err != nil {
			fmt.Printf("Error: %+v\n", err)
		}

	case cmd.SLOT_NOS_BY_REG_NO:
		if err := handler.FindSlotNumberByRegNumber(area, command[1]); err != nil {
			fmt.Printf("Error: %+v\n", err)
		}
	default:
	}

}
