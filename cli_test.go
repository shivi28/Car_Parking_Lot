package main

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestParseCommand(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]string{"create_parking_lot", "6"}, parseCommand("create_parking_lot   6"))
	assert.Equal([]string{"park", "KA-01-HH-1234", "White"}, parseCommand("park   KA-01-HH-1234	White"))
	assert.Equal([]string{"leave", "4"}, parseCommand("leave 4"))
	assert.Equal([]string{"status"}, parseCommand("status"))
	assert.Equal([]string{"registration_numbers_for_cars_with_colour", "White"}, parseCommand("registration_numbers_for_cars_with_colour White"))
	assert.Equal([]string{"slot_numbers_for_cars_with_colour", "White"}, parseCommand("slot_numbers_for_cars_with_colour White"))
	assert.Equal([]string{"slot_number_for_registration_number", "MH-04-AY-1111"}, parseCommand("slot_number_for_registration_number MH-04-AY-1111"))
	assert.True(true)
}
