package main

import "math"

type Column struct {
	ID                        int
	status                    string
	amountOfFloors            int
	amountOfElevators         int
	amountOfElevatorPerColumn int
	elevatorsList             []Elevator
	callButtonsList           []CallButton
	servedFloors              []int
	isBasement                bool
}

var callButtonID int = 1
var elevatorID int = 1

func NewColumn(_id int, _amountOfFloors int, _amountOfElevators int, _servedFloors []int, _isBasement bool) *Column {
	return &Column{_id, "online", _amountOfFloors, _amountOfElevators, 0, *createElevators(_amountOfFloors, _amountOfElevators), *createCallButtons(_amountOfFloors, _isBasement), _servedFloors, _isBasement}
}

func (c *Column) createCallButtons(_amountOfFloors int, _isBasement bool) {
	if _isBasement {
		buttonFloor := -1
		for i := 0; i < _amountOfFloors; i++ {
			callButton := *NewCallButton(callButtonID, "off", buttonFloor, "up")
			c.callButtonsList = append(c.callButtonsList, callButton)
			buttonFloor--
			callButtonID++
		}
	} else {
		buttonFloor := 1
		for i := 0; i < _amountOfFloors; i++ {
			callButton := *NewCallButton(callButtonID, "off", buttonFloor, "down")
			c.callButtonsList = append(c.callButtonsList, callButton)
			buttonFloor++
			callButtonID++
		}
	}
}

func (c *Column) createElevators(_amountOfFloors int, _amountOfElevators int) *[]Elevator {
	var elevatorsList []Elevator
	for i := 0; i < _amountOfElevators; i++ {
		elevator := *NewElevator(elevatorID, "idle", _amountOfFloors, 1)
		c.elevatorsList = append(c.elevatorsList, elevator)
		elevatorID++
	}
	return elevatorsList
}

//Simulate when a user press a button on a floor to go back to the first floor
func (c *Column) requestElevator(userPosition int, direction string) *Elevator {
	elevator := *findElevator(userPosition, direction)
	elevator.addNewRequest(userPosition)
	elevator.move()

	elevator.addNewRequest(1)
	elevator.move()

	return &elevator
}

/*We use a score system depending on the current elevators state. Since the bestScore and the referenceGap are
  //higher values than what could be possibly calculated, the first elevator will always become the default bestElevator,
  //before being compared with to other elevators. If two elevators get the same score, the nearest one is prioritized. Unlike
  //the classic algorithm, the logic isn't exactly the same depending on if the request is done in the lobby or on a floor.*/

func (c *Column) findElevator(requestedFloor int, requestedDirection string) *Elevator {
	bestElevatorInformations := *BestElevatorInfo()

	if requestedFloor == 1 {
		for _, elevator := range c.elevatorsList {
			//The elevator is at the lobby and already has some requests. It is about to leave but has not yet departed
			if 1 == elevator.currentFloor && elevator.status == "stopped" {
				bestElevatorInformations := *checkIfElevatorIsBetter(1, elevator, bestElevatorInformations, requestedFloor)
				//The elevator is at the lobby and has no requests
			} else if 1 == elevator.currentFloor && elevator.status == "idle" {
				bestElevatorInformations := *checkIfElevatorIsBetter(2, elevator, bestElevatorInformations, requestedFloor)
				//The elevator is lower than me and is coming up. It means that I'm requesting an elevator to go to a basement, and the elevator is on it's way to me.
			} else if 1 > elevator.currentFloor && elevator.direction == "up" {
				bestElevatorInformations := *checkIfElevatorIsBetter(3, elevator, bestElevatorInformations, requestedFloor)
				//The elevator is above me and is coming down. It means that I'm requesting an elevator to go to a floor, and the elevator is on it's way to me
			} else if 1 < elevator.currentFloor && elevator.direction == "down" {
				bestElevatorInformations := *checkIfElevatorIsBetter(3, elevator, bestElevatorInformations, requestedFloor)
				//The elevator is not at the first floor, but doesn't have any request
			} else if elevator.status == "idle" {
				bestElevatorInformations := *checkIfElevatorIsBetter(4, elevator, bestElevatorInformations, requestedFloor)
				//The elevator is not available, but still could take the call if nothing better is found
			} else {
				bestElevatorInformations := *checkIfElevatorIsBetter(5, elevator, bestElevatorInformations, requestedFloor)
			}
		}
		return &bestElevatorInformations.bestElevator
	} else {
		for _, elevator := range c.elevatorsList {
			//The elevator is at the same level as me, and is about to depart to the first floor
			if requestedFloor == elevator.currentFloor && elevator.status == "stopped" && requestedDirection == elevator.direction {
				bestElevatorInformations = *checkIfElevatorIsBetter(1, elevator, bestElevatorInformations, requestedFloor)
				//The elevator is lower than me and is going up. I'm on a basement, and the elevator can pick me up on it's way
			} else if requestedFloor > elevator.currentFloor && elevator.direction == "up" && requestedDirection == "up" {
				bestElevatorInformations := *checkIfElevatorIsBetter(2, elevator, bestElevatorInformations, requestedFloor)
				//The elevator is higher than me and is going down. I'm on a floor, and the elevator can pick me up on it's way
			} else if requestedFloor < elevator.currentFloor && elevator.direction == "down" && requestedDirection == "down" {
				bestElevatorInformations := *checkIfElevatorIsBetter(2, elevator, bestElevatorInformations, requestedFloor)
				//The elevator is idle and has no requests
			} else if elevator.status == "idle" {
				bestElevatorInformations := *checkIfElevatorIsBetter(4, elevator, bestElevatorInformations, requestedFloor)
				//The elevator is not available, but still could take the call if nothing better is found
			} else {
				bestElevatorInformations := *checkIfElevatorIsBetter(5, elevator, bestElevatorInformations, requestedFloor)
			}
		}
		return &bestElevatorInformations.bestElevator
	}
}
func (c *Column) checkIfElevatorIsBetter(scoreToCheck int, newElevator Elevator, bestElevatorInformations BestElevatorInformations, floor int) *BestElevatorInformations {
	if scoreToCheck < bestElevatorInformations.bestScore {
		bestElevatorInformations.bestScore = scoreToCheck
		bestElevatorInformations.bestElevator = newElevator
		bestElevatorInformations.referenceGap = math.Abs(newElevator.currentFloor - floor)
	} else if bestElevatorInformations.bestScore == scoreToCheck {
		gap := math.Abs(newElevator.currentFloor - floor)

		if bestElevatorInformations.referenceGap > gap {
			bestElevatorInformations.bestScore = scoreToCheck
			bestElevatorInformations.bestElevator = newElevator
			bestElevatorInformations.referenceGap = gap
		}
	}
	return &bestElevatorInformations
}
