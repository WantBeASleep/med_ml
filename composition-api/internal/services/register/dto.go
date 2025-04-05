package register

import "time"

type RegisterDoctorArg struct {
	Email       string
	Password    string
	FullName    string
	Org         string
	Job         string
	Description *string
}

type RegisterPatientArg struct {
	Email     string
	Password  string
	FullName  string
	Policy    string
	BirthDate time.Time
}
