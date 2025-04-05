package patient

import "time"

type CreatePatientArg struct {
	Fullname   string
	Email      string
	Policy     string
	Active     bool
	Malignancy bool
	BirthDate  time.Time
}

type UpdatePatientArg struct {
	Active     *bool
	Malignancy *bool
}
