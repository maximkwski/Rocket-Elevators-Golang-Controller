package main

//Button on a floor or basement to go back to lobby
type BestElevatorInformations struct {
	bestElevator *Elevator
	bestScore    int
	referenceGap int
}

func BestElevatorInfo() *BestElevatorInformations {
	return &BestElevatorInformations{nil, 6, 10000000}
}
