package Model

type Payment struct {
	Homeid     int
	Loginid    string
	Purpose    string
	Amount     string
	Buyer_name string
	Email      string
	Phone      string
	Send_email bool
	Send_sms   bool
	//Redirect_url string
}

type Paymentinsto struct {
	Purpose      string `json:"purpose"`
	Amount       string `json:"amount"`
	Buyer_name   string `json:"buyer_name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Send_email   bool   `json:"send_email"`
	Send_sms     bool   `json:"send_sms"`
	Redirect_url string `json:"redirect_url"`
}
