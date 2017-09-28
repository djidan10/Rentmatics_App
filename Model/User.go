package Model

type User struct {
	Id        int
	Username  string
	Password  string
	Loginid   string
	Logintype string
}
type UserResponse struct {
	Username string
	Loginid  string
	Status   string
}
type LoginUser struct {
	Username string
	Password string
}

type StatusResponse struct {
	Status string
}

type Getlogin struct {
	Username string
	Loginid  string
	Password string
}

type Changepass struct {
	Loginid     string
	OldPassword string
	Newpassword string
}

type Subscribeuser struct {
	Sub_City  string
	Sub_Email string
}
