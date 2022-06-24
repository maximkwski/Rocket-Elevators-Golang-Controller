package main

import (
	"math"
)

type Battery struct {
	ID                        int
	status                    string
	amountOfColumns           int
	amountOfFloors            int
	amountOfBasements         int
	amountOfElevatorPerColumn int
	columnsList               []Column
	floorRequestButtonsList   []FloorRequestButton
}

var columnID int = 1
var floorRequestButtonID int = 1

func NewBattery(_id int, _amountOfColumns int, _amountOfFloors int, _amountOfBasements int, _amountOfElevatorPerColumn int) *Battery {
	return &Battery{_id, "online", _amountOfColumns, _amountOfFloors, _amountOfBasements, _amountOfElevatorPerColumn, *createColumns(_amountOfColumns, _amountOfFloors, _amountOfBasements, _amountOfElevatorPerColumn), *createFloorRequestButtons(_amountOfFloors)}
}

func createBasementColumn(_amountOfBasements int, _amountOfElevatorPerColumn int) {

	servedFloors := []int{}
	floor := -1

	for i := 0; i < _amountOfBasements; i++ {
		servedFloors = append(servedFloors, floor)
		floor--
	}
}

func createColumns(_amountOfColumns int, _amountOfFloors int, _amountOfBasements int, _amountOfElevatorPerColumn int) *[]Column {
	var columnsList []Column
	if _amountOfBasements > 0 {
		createBasementFloorRequestButtons(_amountOfBasements)
		createBasementColumn(_amountOfBasements, _amountOfElevatorPerColumn)
		_amountOfColumns -= 1
	} // WHERE THIS GOES
	amountOfFloorsPerColumn := int(math.Ceil(float64(_amountOfFloors) / float64(_amountOfColumns)))
	floor := 1
	for i := 0; i < _amountOfColumns; i++ {
		servedFloors := []int{}
		for j := 0; j < amountOfFloorsPerColumn; j++ {
			if floor <= _amountOfFloors {
				servedFloors = append(servedFloors, floor)
				floor++
			}
		}
		column := *NewColumn(columnID, _amountOfFloors, _amountOfElevatorPerColumn, servedFloors, false)
		columnsList = append(columnsList, column)
		columnID++
	}
	return &columnsList

}

func createFloorRequestButtons(_amountOfFloors int) *[]FloorRequestButton {
	buttonFloor := 1
	var floorRequestButtonsList []FloorRequestButton
	for i := 0; i < _amountOfFloors; i++ {
		floorRequestButton := *NewFloorRequestButton(floorRequestButtonID, "off", buttonFloor, "up")
		floorRequestButtonsList = append(floorRequestButtonsList, floorRequestButton)
		buttonFloor++
		floorRequestButtonID++
	}
	return &floorRequestButtonsList
}

func createBasementFloorRequestButtons(_amountOfBasements int) *[]FloorRequestButton {
	buttonFloor := -1
	var floorRequestButtonsList []FloorRequestButton
	for i := 0; i < _amountOfBasements; i++ {
		floorRequestButton := *NewFloorRequestButton(floorRequestButtonID, "off", buttonFloor, "down")
		floorRequestButtonsList = append(floorRequestButtonsList, floorRequestButton)
		buttonFloor--
		floorRequestButtonID++
	}
	return &floorRequestButtonsList
}

func (b *Battery) findBestColumn(_requestedFloor int) *Column {

	selectedColumn := b.columnsList[0]
	for _, column := range b.columnsList {
		for _, x := range column.servedFloors {
			if x == _requestedFloor {
				// selectedColumn = column
				return &column
			}
		}
	}
	return &selectedColumn
}

//Simulate when a user press a button at the lobby
func (b *Battery) assignElevator(_requestedFloor int, _direction string) (*Column, *Elevator) {
	column := *b.findBestColumn(_requestedFloor)
	elevator := *column.findElevator(1, _direction)
	elevator.addNewRequest(1)
	elevator.move()

	elevator.addNewRequest(_requestedFloor)
	elevator.move()

	return &column, &elevator
}
