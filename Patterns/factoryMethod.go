package Patterns

import "fmt"

type DeviceWrite interface {
	UseDevice()
}

type Device struct {
	Name string
}

func (s *Device) UseDevice() {
	fmt.Println("Use device: ", s.Name)
}

type Pen struct {
	Device
	Ink int
}

func NewPen() DeviceWrite {
	return &Pen{
		Device: Device{Name: "pen"},
		Ink:    10,
	}

}

type Pencil struct {
	Device
	PencilLead int
}

func NewPencil() DeviceWrite {
	return &Pencil{
		Device:     Device{Name: "pen"},
		PencilLead: 10,
	}

}
