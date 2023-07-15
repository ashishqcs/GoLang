package parkinglot

type ParkingSlot struct {
	slotsFull bool
	slot      []bool
}

func (ps *ParkingSlot) fill() int {
	availableSlot := ps.getNextAvailableSlot()
	if availableSlot != -1 {
		ps.slot[availableSlot] = true
	}

	return availableSlot
}

func (ps *ParkingSlot) getNextAvailableSlot() int {
	for slotNumber, isFilled := range ps.slot {
		if !isFilled {
			return slotNumber
		}
	}
	return -1
}
