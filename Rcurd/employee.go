package main

// Employee is a struct that represents an employee.
type Employee struct {
	ID     uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Dob    string `json:"dob"`
	Mobile string `json:"mobile"`
}
