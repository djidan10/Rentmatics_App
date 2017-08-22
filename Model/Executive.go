package Model

type Executive struct {
	Homestruct Home
	images     []string
}

type LoginData struct {
	Id       string
	Username string
	Password string
	Role     string
}

type Login struct {
	Username string
	Password string
}
