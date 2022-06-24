package main

type Door struct {
	ID     int
	status string
}

func NewDoor(_status string) *Door {
	return &Door{1, _status}
}
