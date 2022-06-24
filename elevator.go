package main

import "sort"

type Elevator struct {
	ID                    int
	status                string
	amountOfFloors        int
	currentFloor          int
	direction             string
	floorRequestsList     []int
	completedRequestsList []int
	door                  Door
}

func NewElevator(_id int, _status string, _amountOfFloors int, _currentFLoor int) *Elevator {
	return &Elevator{_id, _status, _amountOfFloors, _currentFLoor, "", []int{}, []int{}, *NewDoor("closed")}
}

func (e *Elevator) move() {

	// screenDisplay := 0
	for len(e.floorRequestsList) > 0 {
		e.sortFloorList()
		destination := e.floorRequestsList[0]
		e.status = "moving"
		if e.currentFloor < destination {
			e.direction = "up"
			for e.currentFloor < destination {
				e.currentFloor++
				// screenDisplay = e.currentFloor
			}
		} else if e.currentFloor > destination {
			e.direction = "down"
			// e.sortFloorList()
			for e.currentFloor > destination {
				e.currentFloor--
				// screenDisplay = e.currentFloor
			}
		}
		e.status = "stopped"
		e.completedRequestsList = append(e.completedRequestsList, e.floorRequestsList[0])
		e.floorRequestsList = e.floorRequestsList[1:]
	}
	e.status = "idle"
}

func (e *Elevator) sortFloorList() {
	if e.direction == "up" {
		sort.Slice(e.floorRequestsList, func(i, j int) bool {
			return e.floorRequestsList[i] < e.floorRequestsList[j]
		})
	} else {
		sort.Slice(e.floorRequestsList, func(i, j int) bool {
			return e.floorRequestsList[i] > e.floorRequestsList[j]
		})
	}
}

func (e *Elevator) operateDoors() {
	e.door.status = "opened"
}

func (e *Elevator) addNewRequest(requestedFloor int) {

	if contains(e.floorRequestsList, requestedFloor) == false {
		e.floorRequestsList = append(e.floorRequestsList, requestedFloor)
	}
	if e.currentFloor < requestedFloor {
		e.direction = "up"
	}
	if e.currentFloor > requestedFloor {
		e.direction = "down"
	}

}
