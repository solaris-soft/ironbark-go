package serializers

import "github.com/solaris-soft/ironbark-go/db"

type ContactSerializer struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	City      string `json:"city"`
	State     string `json:"state"`
	Zip       string `json:"zip"`
	Country   string `json:"country"`
}

func NewContact(contact *db.Contact) *ContactSerializer {
	return &ContactSerializer{
		FirstName: contact.FirstName,
		LastName:  contact.LastName,
		Email:     contact.Email,
		Phone:     contact.Phone,
		Address:   contact.Address,
		City:      contact.City,
		State:     contact.State,
		Zip:       contact.Zip,
		Country:   contact.Country,
	}
}
