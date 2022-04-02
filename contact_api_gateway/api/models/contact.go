package models

type CreateContactModel struct {
	Name  string `json:"name" binding:"required"`
	Phone string `json:"phone" binding:"required"`
}

type ContactModel struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type GetAllContactModel struct {
	Contacts []ContactModel `json:"contacts"`
}

type Contact struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	UserId string `json:"user_id"`
}

type ContactUpdate struct {
	Status string `json:"status"`
}
