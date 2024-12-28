package domain

type Alert struct {
	ID          uint
	Subject     string
	Content     string
	Destination []*Destination
}

func NewAlert(ID uint, Subject string, Content string, Destination []*Destination) *Alert {
	return &Alert{
		ID:          ID,
		Subject:     Subject,
		Content:     Content,
		Destination: Destination,
	}
}

// validations and business rules here
