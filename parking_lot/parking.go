package parkinglot

import "errors"

var errParkingFull = errors.New("Parking is full")
var errAlreadyParked = errors.New("Vehicle is already parked")
var errInvalidRegistrationId = errors.New("Regration id is invalid")
var errNotFoundToUnPark = errors.New("Vehicle not found to unpark")

type ParkingLot struct {
	maxSlot int
	slot    map[string]bool
}

func (p *ParkingLot) Park(registrationId string) (bool, error) {
	p.initSlots()

	if p.isParked(registrationId) {
		return false, errAlreadyParked
	}

	if len(p.slot) < p.maxSlot {
		p.slot[registrationId] = true
	} else {
		return false, errParkingFull
	}

	return true, nil
}

func (p *ParkingLot) UnPark(registrationId string) (bool, error) {

	if p.isParked(registrationId) {
		delete(p.slot, registrationId)
	} else {
		return false, errNotFoundToUnPark
	}
	return true, nil
}

func (p *ParkingLot) initSlots() {
	if p.slot == nil {
		p.slot = make(map[string]bool, p.maxSlot)
	}
}

func (p *ParkingLot) isParked(registrationId string) bool {
	return p.slot[registrationId] == true
}
