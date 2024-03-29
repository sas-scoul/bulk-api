package user

type CreateUserDto struct {
	Name         string `json:"name"`
	TicketNumber string `json:"ticket_number"`
	Password     string `json:"password"`
}
