package main

import "time"

type Employee struct {
	ID        int
	Name      string
	Address   string
	DOB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

var dilbert Employee
