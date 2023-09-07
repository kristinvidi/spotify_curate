package model

type Email string

type User struct {
	DisplayName Name
	Email       Email
	ID          ID
	URI         URI
	Country     CountryCode
}
