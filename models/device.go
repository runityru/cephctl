package models

type Device struct {
	ID        string
	Daemons   []string
	WearLevel float64
}
