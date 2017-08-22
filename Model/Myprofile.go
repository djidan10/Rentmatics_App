package Model

type Profile struct {
	Id       int
	Loginid  string
	Name     string
	Tittle   string
	Phone    string
	Email    string
	About    string
	Twiter   string
	Facebook string
	Google   string
}

type Getprofile struct {
	Loginid string
}
