package domain

type Destination struct {
	ID    uint
	Email string
	Phone string
}

func NewDestination(ID uint, Email string, Phone string) *Destination {
	return &Destination{
		ID:    ID,
		Email: Email,
		Phone: Phone,
	}
}

// validations and business rules here