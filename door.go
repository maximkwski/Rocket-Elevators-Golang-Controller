package main

type Door struct {
	ID     int
	status string
}

func (d *Door) NewDoor() *Door {
	return &Door{1, "status"}
}
