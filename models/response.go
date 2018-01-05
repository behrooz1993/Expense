package models

type Response struct {
	Status bool
	Data   interface{}
	Error  string
}
