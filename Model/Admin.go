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
