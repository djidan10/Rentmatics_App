package Model

type InsertUser struct {
	Id       string
	Username string
	Password string
	Role     string
}

type Admincomp_Req struct {
	PendingComplaints []Complaints
	PendingRequest    []Request
}

type Contactform struct {
	Name        string
	Email       string
	Phonenumber string
	Subject     string
	Message     string
}
