package parkinglot

import (
	"testing"
)

func TestShouldParkForRegistrationId(t *testing.T) {
	parking := ParkingLot{maxSlot: 5}

	registrationId := "ABCD"
	isParked, err := parking.Park(registrationId)

	assertNoError(err, t)
	assertTrue(isParked, t)
}

func TestShouldNotParkIfVehicleIsAlreadyParked(t *testing.T) {
	expectedErrorMessage := "Vehicle is already parked"
	registrationId := "ABCD"

	parking := ParkingLot{maxSlot: 5}
	isParked, err := parking.Park(registrationId)
	isParked, err = parking.Park(registrationId)

	assertFalse(isParked, t)
	assertThat(err.Error(), expectedErrorMessage, t)
}
func TestShouldNotParkIfSlotsAreFull(t *testing.T) {
	expectedErrorMessage := "Parking is full"
	registrationId := "PQRS"

	parking := ParkingLot{maxSlot: 2, slot: map[string]bool{
		"ABCD": true,
		"BCDE": true,
	}}

	isParked, err := parking.Park(registrationId)

	assertFalse(isParked, t)
	assertThat(err.Error(), expectedErrorMessage, t)
}

func TestShouldUnParkIfIsParked(t *testing.T) {
	parking := ParkingLot{maxSlot: 5}

	registrationId := "ABCD"
	parking.Park(registrationId)

	isUnParked, err := parking.UnPark(registrationId)

	assertTrue(isUnParked, t)
	assertNoError(err, t)

}

func TestShouldNotUnParkIfVehicleIsNotParked(t *testing.T) {
	parking := ParkingLot{maxSlot: 5}
	expectedErrorMessage := "Vehicle not found to unpark"

	registrationId := "ABCD"
	parking.Park(registrationId)
	parking.UnPark(registrationId)

	isUnParked, err := parking.UnPark(registrationId)

	assertFalse(isUnParked, t)
	assertThat(err.Error(), expectedErrorMessage, t)

}

func assertTrue(actual bool, t *testing.T) {
	assertBool(true, actual, t)
}

func assertFalse(actual bool, t *testing.T) {
	assertBool(false, actual, t)
}

func assertThat(expected string, actual string, t *testing.T) {
	if expected != actual {
		t.Errorf("Failed. Wanted \"%s\" but found \"%s\"", expected, actual)
	}
}

func assertBool(expected bool, actual bool, t *testing.T) {
	if expected != actual {
		t.Errorf("Failed. Wanted %t but Found %t", expected, actual)
	}
}

func assertNoError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("Failed. Error occured \"%s\"", err.Error())
	}
}
